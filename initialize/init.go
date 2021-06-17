package initialize

import (
	"cronProject/config"
	"cronProject/global"
	"cronProject/routers"
	"cronProject/task"
	"github.com/gin-gonic/gin"
)

func InitBase() {
	global.Config = config.LoadConfig("./config.yaml")
	global.DB = GetDB()
	global.Redis = GetRedis()
	global.Logger = GetLogger()

	task.Start()
}

func GetRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	group := r.Group("task")
	{
		routers.Cron(group)
	}
	return r
}
