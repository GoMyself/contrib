package conn

import (
	"fmt"
    "log"
    "time"
	_ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
    "github.com/go-redis/redis/v8"
)

func InitDB(dsn string, maxIdleConn, maxOpenConn int) *sqlx.DB {

    db, err := sqlx.Connect("mysql", dsn)
    if err != nil {
        log.Fatalln(err)
    }

    db.SetMaxOpenConns(maxOpenConn)
    db.SetMaxIdleConns(maxIdleConn)
    db.SetConnMaxLifetime(time.Second * 30)
    err = db.Ping()
    if err != nil {
        log.Fatalln(err)
    }

    return db
}

func InitRedisSentinel(dsn []string, psd, name string, db int) *redis.Client {

    reddb := redis.NewFailoverClient(&redis.FailoverOptions{
        MasterName:    name,
        SentinelAddrs: dsn,
        Password:      psd, // no password set
        DB:            db,  // use default DB
        DialTimeout:   10 * time.Second,
        ReadTimeout:   30 * time.Second,
        WriteTimeout:  30 * time.Second,
        PoolSize:      10,
        PoolTimeout:   30 * time.Second,
        MaxRetries:    2,
        IdleTimeout:   5 * time.Minute,
    })
    pong, err := reddb.Ping().Result()
    if err != nil {
        log.Fatalf("initRedisSentinel failed: %s", err.Error())
    }
    fmt.Println(pong, err)

    return reddb
}

func InitRedis(dsn string, psd string) *redis.Client {

    reddb := redis.NewClient(&redis.Options{
        Addr:         dsn,
        Password:     psd, // no password set
        DB:           0,   // use default DB
        DialTimeout:  10 * time.Second,
        ReadTimeout:  30 * time.Second,
        WriteTimeout: 30 * time.Second,
        PoolSize:     10,
        PoolTimeout:  30 * time.Second,
        MaxRetries:   2,
        IdleTimeout:  5 * time.Minute,
    })
    pong, err := reddb.Ping().Result()
    if err != nil {
        log.Fatalf("initRedisSlave failed: %s", err.Error())
    }
    fmt.Println(pong, err)

    return reddb
}