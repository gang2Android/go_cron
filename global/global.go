package global

import (
	"cronProject/config"
	"github.com/go-redis/redis"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config   *config.Config
	DB       *gorm.DB
	Redis    *redis.Client
	CronTask *cron.Cron
	Logger   *zap.Logger
)
