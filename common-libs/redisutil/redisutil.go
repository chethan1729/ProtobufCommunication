package redisutil

import (
	"context"
	"log"
	"time"
)

type Handler interface {
	GetHandle() *RedisConn
}

type RedisConn struct {
	IsInited bool
	PoolSize int
	MinIdleConns int
	Conn *redis.Client
}

type RedisConf struct {
	Ip string
	Port string
	Pass string
	Db int
	PoolSize int
	MinIdleConns int
}

func GetHandle(conf RedisConf) *RedisConn {
	var rhdl = new(RedisConn)
	if !rhdl.IsInited {
		log.Println("Redis Util : Initialising Redis conn handle")
		rhdl.InitConn(conf)
	}
	return rhdl
}

// InitConn Initializes and returns a Redis Client
func (rh *RedisConn) InitConn(conf RedisConf) {
	log.Println("Redis Util : Redis IP config => ",conf.Ip,":",conf.Port)
	rh.IsInited = true
	rh.PoolSize = conf.PoolSize
	rh.MinIdleConns = conf.MinIdleConns
	rh.Conn = redis.NewClient(&redis.Options{
		Addr: conf.Ip+ ":" + conf.Port,
		Password: conf.Pass,
		DB: conf.Db,
		PoolSize: conf.PoolSize,
		MinIdleConns: conf.MinIdleConns,
	})
}

// Ping returns PONG and this command is often used to test
// if a connection is still alive, or to messure latency
func (rh *RedisConn) Ping(ctx context.Context) string {
	res := rh.Conn.Ping(ctx)
	return res.Val()
}

// PushToQueue pushes data into redis queue
func (rh *RedisConn) PushToQueue(ctx context.Context, key string, values ...interface{}) int64 {
	args := make([]interface{}, 0, len(values))
	args = appendArgs(args, values)
	res := rh.Conn.LPush(ctx, key, args...)
	return res.Val()
}

// PollRedisForResp returns a struct with key and the popped byte
// array value from the redis queue
func (rh *RedisConn) PollRedisForResp(ctx context.Context, timeout time.Duration, keys ...string) []string {
	res := rh.Conn.BRPop(ctx, timeout, keys...)
	return res.Val()
}