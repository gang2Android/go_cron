package task

import (
	"cronProject/global"
	"github.com/robfig/cron/v3"
)

func init() {
	global.CronTask = cron.New(cron.WithSeconds())
}

func Start() {}
