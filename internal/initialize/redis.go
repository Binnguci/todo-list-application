package initialize

import (
	"context"
	"fmt"
	"github.com/binnguci/todo-app/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
		Password: r.Password,
		DB:       r.DB,
		PoolSize: 10,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Redis connection error:", zap.Error(err))
	}

	fmt.Printf("Redis connection success\n")
	global.Rdb = rdb
}
