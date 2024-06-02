package core

import (
	"CSAMS-Backend/global"
	"context"
	"github.com/go-redis/redis"
	"log"
	"time"
)

// ConnectRedis 默认连接0号库
func ConnectRedis() *redis.Client {
	return ConnectRedisDB(0)
}

func ConnectRedisDB(db int) *redis.Client {
	redisConf := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr(),
		Password: redisConf.Password,
		DB:       db,
		PoolSize: redisConf.PoolSize,
	})
	//限制连接所需要的时长
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	//测试连接
	_, err := rdb.Ping().Result()
	if err != nil {
		log.Printf("redis连接失败 %s", redisConf.Addr())
		return nil
	}
	return rdb
}
