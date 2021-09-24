package redis

import (
	"context"
	"time"

	"evermos/constants"

	"evermos/config"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb *redis.Client

func Init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "192.168.1.212:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func SaveToken(userId string, token string) {
	rdb.HSet(ctx, constants.AUTH+userId, token, time.Now())
	rdb.PExpire(ctx, userId, time.Duration(config.JWT_EXP)*time.Second)
}

func IsExistToken(userId string, token string) bool {
	return rdb.HExists(ctx, constants.AUTH+userId, token).Val()
}
