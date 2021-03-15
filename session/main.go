package session

import (
	// "log"
	"errors"
	"fmt"
  	"context"
	"github.com/go-redis/redis/v8"
	"github.com/valyala/fasthttp"
)


var (
	ctx = context.Background()
	client *redis.Client	
)



func New(reddb *redis.Client) {
	client = reddb
}

func Set(value []byte, uid uint64) (string, error) {

	//uuid := fmt.Sprintf("TI%d", uid)
	key  := fmt.Sprintf("%d", Cputicks())
	
	//val  := client.Get(ctx, uuid).Val()
	pipe := client.TxPipeline()

	defer pipe.Close()

	/*
	if len(val) > 0 {
		//同一个用户，一个时间段，只能登录一个
		pipe.Unlink(ctx, val)
	}
	*/
	//pipe.Set(ctx, uuid, key, -1)
	pipe.SetNX(ctx, key, value, defaultGCLifetime)

	_, err := pipe.Exec(ctx)

	return key, err
}

func Update(value []byte, uid uint64) bool {

	uuid := fmt.Sprintf("TI%d", uid)
	
	val  := client.Get(ctx, uuid).Val()
	pipe := client.TxPipeline()
	defer pipe.Close()
  
	if len(val) == 0 {
		return false
	}

	pipe.Unlink(ctx, val)
	pipe.SetNX(ctx, val, value, defaultGCLifetime)

	_, err := pipe.Exec(ctx)
    if err != nil {
    	return false
    }
	return true
}

func Destroy(ctx *fasthttp.RequestCtx) {

	key := string(ctx.Request.Header.Peek("token"))
	if len(key) == 0 {
		return
	}
	client.Unlink(ctx, key)
	//cookie.Delete(ctx, defaultSessionKeyName)
}

func Get(ctx *fasthttp.RequestCtx) ([]byte, error) {

	key := string(ctx.Request.Header.Peek("token"))
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
