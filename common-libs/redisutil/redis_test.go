package redisutil

import (
	"context"
	"log"
	"testing"
)

func TestConnectivity (t *testing.T) {
	conf := RedisConf {
		Ip: "0.0.0.0",
		Port: "6379",
		Pass: "",
		Db: 0,
		PoolSize: 20,
		MinIdleConns: 10,
	}

	conn := GetHandle(conf)
	res := conn.Ping(context.Background)
	log.Println("Reponse : Ping => ",res)
	if res == "" {
		t.Errorf("Redis connection error \n")
	}
}
