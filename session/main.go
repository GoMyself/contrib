package session

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/valyala/fasthttp"
	"time"
)

var (
	ctx    = context.Background()
	client *redis.Client
)

func New(reddb *redis.Client) {
	client = reddb
}

func Set(value []byte, uid string) (string, error) {

	uuid := fmt.Sprintf("TI%s", uid)
	key := fmt.Sprintf("%d", Cputicks())

	val, err  := client.Get(ctx, uuid).Result()

	pipe := client.TxPipeline()
	defer pipe.Close()

	if err != redis.Nil && len(val) > 0 {
		//同一个用户，一个时间段，只能登录一个
		pipe.Unlink(ctx, val)
	}

	pipe.Set(ctx, uuid, key, time.Duration(100) * time.Hour)
	pipe.SetNX(ctx, key, value, defaultExpires)

	_, err = pipe.Exec(ctx)

	return key, err
}

func Update(value []byte, uid uint64) bool {

	uuid := fmt.Sprintf("TI%d", uid)

	val := client.Get(ctx, uuid).Val()
	pipe := client.TxPipeline()
	defer pipe.Close()

	if len(val) == 0 {
		return false
	}

	pipe.Unlink(ctx, val)
	pipe.SetNX(ctx, val, value, renewExpires)

	_, err := pipe.Exec(ctx)
	if err != nil {
		return false
	}

	return true
}

func Offline(sid string) {

	if len(sid) == 0 {
		return
	}

	client.Unlink(ctx, sid)
}

func Destroy(ctx *fasthttp.RequestCtx) {

	key := string(ctx.Request.Header.Peek("t"))
	if len(key) == 0 {
		return
	}

	client.Unlink(ctx, key)
}

func Get(ctx *fasthttp.RequestCtx) ([]byte, error) {

	key := string(ctx.Request.Header.Peek("t"))
	if len(key) == 0 {
		return nil, errors.New("does not exist")
	}

	val, err := client.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return nil, errors.New("does not exist")
	} else if err != nil {
		return nil, err
	} else {
		return val, nil
	}
}

func GetByToken(token string) ([]byte, error) {

	val, err := client.Get(ctx, token).Bytes()
	if err == redis.Nil {
		return nil, errors.New("does not exist")
	} else if err != nil {
		return nil, err
	} else {
		return val, nil
	}
}
