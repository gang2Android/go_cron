package routers

import (
	"cronProject/global"
	"cronProject/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"runtime"
	"strconv"
)

func Cron(groupRouter *gin.RouterGroup) (R gin.IRoutes) {
	groupRouter.Handle("GET", "info", func(ctx *gin.Context) {
		goroutine := runtime.NumGoroutine()
		ctx.JSON(200, gin.H{"goroutine": goroutine})
	})

	groupRouter.Handle("GET", "start", func(ctx *gin.Context) {
		tasks := services.GetTasks()
		for i := 0; i < len(tasks); i++ {
			task := tasks[i]
			if task.TaskNo != "" {
				no, _ := strconv.Atoi(task.TaskNo)
				global.CronTask.Remove(cron.EntryID(no))
				services.SetTaskNo(task.Id, "")
			}
			taskNo, _ := global.CronTask.AddFunc(task.Spec, task.Cmd)
			services.SetTaskNo(task.Id, strconv.Itoa(int(taskNo)))
		}
		global.CronTask.Stop()
		global.CronTask.Start()
		ctx.JSON(200, gin.H{"status": 0})
	})

	groupRouter.Handle("GET", "stop", func(ctx *gin.Context) {
		global.CronTask.Stop()
		ctx.JSON(200, gin.H{"status": 0})
	})

	groupRouter.Handle("GET", "add", func(ctx *gin.Context) {
		id := ctx.DefaultQuery("id", "")
		if id == "" {
			ctx.JSON(200, gin.H{"status": -1, "msg": "id not empty"})
			return
		}

		task := services.GetTask(id)
		if task.Id == "" {
			ctx.JSON(200, gin.H{"status": -2, "msg": "id is error"})
			return
		}
		if task.TaskNo != "" {
			no, _ := strconv.Atoi(task.TaskNo)
			global.CronTask.Remove(cron.EntryID(no))
			services.SetTaskNo(task.Id, "")
		}
		global.Logger.Info(task.String())

		if global.CronTask == nil {
			fmt.Println("err")
			return
		}

		taskNo, err := global.CronTask.AddFunc(task.Spec, task.Cmd)
		if err != nil {
			fmt.Println(err)
			return
		}
		global.CronTask.Stop()
		global.CronTask.Start()
		services.SetTaskNo(task.Id, strconv.Itoa(int(taskNo)))

		ctx.JSON(200, gin.H{"status": 1, "msg": "success"})
	})

	groupRouter.Handle("GET", "remove", func(ctx *gin.Context) {
		id := ctx.DefaultQuery("id", "")
		if id == "" {
			ctx.JSON(200, gin.H{"status": -1, "msg": "id not empty"})
			return
		}

		task := services.GetTask(id)
		if task.Id == "" {
			ctx.JSON(200, gin.H{"status": -2, "msg": "id is error"})
			return
		}
		if task.TaskNo != "" {
			no, _ := strconv.Atoi(task.TaskNo)
			global.CronTask.Remove(cron.EntryID(no))
			services.SetTaskNo(task.Id, "")
		}
		global.CronTask.Stop()
		global.CronTask.Start()
		ctx.JSON(200, gin.H{"status": 0})
	})

	groupRouter.Handle("GET", "run", func(ctx *gin.Context) {
		id := ctx.DefaultQuery("id", "")
		if id == "" {
			ctx.JSON(200, gin.H{"status": -1, "msg": "id not empty"})
			return
		}

		task := services.GetTask(id)
		if task.Id == "" {
			ctx.JSON(200, gin.H{"status": -2, "msg": "id is error"})
			return
		}
		global.Logger.Info(task.String())
		if task.Enable == 1 {
			if task.TaskNo == "" {
				if global.CronTask == nil {
					fmt.Println("err")
					ctx.JSON(200, gin.H{"status": -3, "msg": "global.CronTask == nil"})
					return
				}
				taskNo, err := global.CronTask.AddFunc(task.Spec, task.Cmd)
				if err != nil {
					fmt.Println(err)
					ctx.JSON(200, gin.H{"status": -4, "msg": err.Error()})
					return
				}
				global.CronTask.Stop()
				global.CronTask.Start()
				services.SetTaskNo(task.Id, strconv.Itoa(int(taskNo)))
			}
		}
		task.Cmd()
		ctx.JSON(200, gin.H{"status": 1, "msg": "success"})
	})

	return groupRouter
}
