package global

import (
	"my-admin/config"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var Logger *zap.SugaredLogger
var LogLevel = zap.NewAtomicLevel()
var DBClient *gorm.DB
var RedisClient *redis.Client

var Config *config.Conf

// var PluginManager Plugins
