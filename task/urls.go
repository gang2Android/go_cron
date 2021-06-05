package task

import (
	"cronProject/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type TaskObj struct {
	Name string `json:"name"`
	Spec string `json:"spec"`
	Url  string `json:"url"`
}

func (t TaskObj) cmd() {
	go func(task TaskObj) {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"), "start", task.Name)
		res := utils.Get(task.Url)
		if len(res) > 100 {
			res = res[0:100]
			res += "..."
		}
		fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"), "end", task.Name, "结果=["+res+"]")
	}(t)
}

func loadTask(tasks *[]TaskObj) {
	fileLocker.Lock()
	data, err := ioutil.ReadFile("./task.json")
	fileLocker.Unlock()
	if err != nil {
		fmt.Println("read json file error")
		return
	}
	dataJson := []byte(data)
	err = json.Unmarshal(dataJson, tasks)
	if err != nil {
		fmt.Println("unmarshal json file error")
		return
	}
	return
}
