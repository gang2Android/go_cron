package task

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"sync"
	"time"
)

var fileLocker sync.Mutex

func Run() {
	c := cron.New(cron.WithSeconds())

	//备份数据库
	dbList := make([]TaskDBObj, 0)
	loadTaskDB(&dbList)
	for _, v := range dbList {
		_, _ = c.AddFunc(v.Spec, v.cmd)
	}

	//执行url计划
	lists := make([]TaskObj, 0)
	loadTask(&lists)
	for _, v := range lists {
		_, _ = c.AddFunc(v.Spec, v.cmd)
	}

	//执行系统磁盘监控计划
	sysLists := make([]TaskSystem, 0)
	loadSysDiskTask(&sysLists)
	for _, v := range sysLists {
		_, _ = c.AddFunc(v.Spec, v.getSystemDisk)
	}

	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"), "start")
	c.Start()
	select {}
}
