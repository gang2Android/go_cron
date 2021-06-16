package services

import (
	"cronProject/global"
	"cronProject/model"
)

func GetTask(id string) model.Task {
	var task model.Task
	global.DB.Table("planning_tasks").Where("id=?", id).Find(&task)
	return task
}

func GetTasks() []model.Task {
	var task []model.Task
	global.DB.Table("planning_tasks").Where("enable=?", 1).Find(&task)
	return task
}

func SetTaskNo(id, taskNo string) {
	global.DB.Table("planning_tasks").Where("id=?", id).Updates(map[string]interface{}{"task_no": taskNo})
}
