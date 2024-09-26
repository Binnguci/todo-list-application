package global

import (
	"github.com/binnguci/todo-app/pkg/logger"
	"github.com/binnguci/todo-app/pkg/setting"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
	Rdb    *redis.Client
)
