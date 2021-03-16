package conn

import (
	"fmt"
    "log"
    "time"
    "context"
    "github.com/olivere/elastic/v7"
	_ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
    "github.com/go-redis/redis/v8"
    "github.com/beanstalkd/go-beanstalk"
    cpool "github.com/silenceper/pool"
)

var ctx = context.Background()

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

func InitRedisSentinel(dsn []string, psd, name string) *redis.Client {

    reddb := redis.NewFailoverClient(&redis.FailoverOptions{
        MasterName:    name,
        SentinelAddrs: dsn,
        Password:      psd, // no password set
        DB:            0,  // use default DB
        DialTimeout:   10 * time.Second,
        ReadTimeout:   30 * time.Second,
        WriteTimeout:  30 * time.Second,
        PoolSize:      10,
        PoolTimeout:   30 * time.Second,
        MaxRetries:    2,
        IdleTimeout:   5 * time.Minute,
    })
    pong, err := reddb.Ping(ctx).Result()
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
    pong, err := reddb.Ping(ctx).Result()
    if err != nil {
        log.Fatalf("initRedisSlave failed: %s", err.Error())
    }
    fmt.Println(pong, err)

    return reddb
}

func InitES(url []string) *elastic.Client {

	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(url...))
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func InitBeanstalk(beanstalkConn string) cpool.Pool {

    factory    := func() (interface{}, error) { return beanstalk.Dial("tcp", beanstalkConn) }
    closed     := func(v interface{}) error { return v.(*beanstalk.Conn).Close() }
    poolConfig := &cpool.Config{
        InitialCap  : 15,  // 资源池初始连接数
        MaxIdle     : 50,  // 最大空闲连接数
        MaxCap      : 100, // 最大并发连接数
        Factory     : factory,
        Close       : closed,
        IdleTimeout : 15 * time.Second,
    }

    beanPool, err := cpool.NewChannelPool(poolConfig)
    if err != nil {
        log.Fatalln(err)
    }
    return beanPool
}
