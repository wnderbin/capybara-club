package database

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var (
	rdb *redis.Client
	RDB = initRedis()
	Ctx = context.Background()
)

func CloseRedis() error {
	Log.Info("[redis] closing connection with redis...")
	return RDB.Close()
}

func initRedis() *redis.Client {
	if conf.StartUpStatus == 0 {
		return rdb
	}
	rdb = redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Address,
		Password: conf.Redis.Password,
		DB:       conf.Redis.DB,
	})

	if _, err := rdb.Ping(Ctx).Result(); err != nil {
		Log.Error(fmt.Sprintf("[redis] failed to connect to redis: %e", err))
		return rdb
	}
	Log.Info("[redis] connected to redis")
	return rdb
}
