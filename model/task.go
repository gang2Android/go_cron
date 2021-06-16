package model

import (
	"cronProject/global"
	"cronProject/utils"
	"encoding/json"
	"strconv"
	"strings"
)

type Task struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Type    int    `json:"type"` // 1url，2数据库
	Spec    string `json:"spec"`
	Enable  int    `json:"enable"`
	AddTime string `json:"add_time"`
	Content string `json:"content"`
	TaskNo  string `json:"task_no"`
}

func (t Task) Cmd() {
	go func(task Task) {
		global.Logger.Info(task.Name + "start")
		switch task.Type {
		case 1:
			{
				if strings.Contains(task.Content, "://") {
					res := utils.Get(task.Content)
					if len(res) > 100 {
						res = res[0:100]
						res += "..."
					}
					global.Logger.Info(task.Name + "=" + res)
				}
			}
		case 2:
			var db DB
			_ = json.Unmarshal([]byte(task.Content), &db)
			utils.Cmd(db.GetCmdStr())
		default:
			global.Logger.Info(task.Name + "default")
		}
		global.Logger.Info(task.Name + "end")
	}(t)
}

func (t Task) String() string {
	return "{\"id\":" + t.Id + ",\"name\":" + t.Name + ",\"type\":" + strconv.Itoa(t.Type) +
		",\"spec\":" + t.Spec + ",\"enable\":" + strconv.Itoa(t.Enable) +
		",\"add_time\":" + t.AddTime + ",\"content\":" + t.Content + ",\"task_no\":" + t.TaskNo + "}"
}
