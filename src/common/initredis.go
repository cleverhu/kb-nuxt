package common

import (
	"github.com/go-redis/redis"
)

var Rds *redis.Client

func init() {
	//初始化redis
	Rds = redisClient()
}

func redisClient() *redis.Client {
	rds := redis.NewClient(&redis.Options{
		Addr:     "101.132.107.3:6379",
		Password: "", //
		DB:       0,  //
	})
	//fmt.Println(rds)
	return rds
}