package task

import (
	"cronProject/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

type TaskDBObj struct {
	Name      string `json:"name"`
	Host      string `json:"host"`
	Db        string `json:"db"`
	User      string `json:"user"`
	Pwd       string `json:"pwd"`
	BackPath  string `json:"backPath"`
	MysqlPath string `json:"mysqlPath"`
	Retain    string `json:"retain"`
	Spec      string `json:"spec"`
}

func (t TaskDBObj) cmd() {
	go func(task TaskDBObj) {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"), "start", task.Name)

		var cmdStr strings.Builder
		cmdStr.WriteString("#!/bin/bash \n")
		cmdStr.WriteString("host=" + t.Host + "; \n")
		cmdStr.WriteString("uName=" + t.User + "; \n")
		cmdStr.WriteString("pwd=" + t.Pwd + "; \n")
		cmdStr.WriteString("dbName=" + t.Db + "; \n")
		cmdStr.WriteString("basePath=" + t.BackPath + "; \n")
		cmdStr.WriteString("dirName=$dbName`date +\\%Y\\%m\\%d`/; \n")
		cmdStr.WriteString("time=`date +\\%Y\\%m\\%d\\%H\\%M`; \n")
		cmdStr.WriteString("cd $basePath; \n")
		cmdStr.WriteString("mkdir $dirName; \n")
		cmdStr.WriteString("echo $basePath$dirName$dbName$time; \n")
		cmdStr.WriteString(t.MysqlPath + "mysqldump -h $host -u$uName -p$pwd --default-character-set=utf8 --skip-extended-insert $dbName > $basePath$dirName$dbName$time.sql; \n")
		cmdStr.WriteString("cd $dirName; \n")
		cmdStr.WriteString("split --verbose -l 500 $dbName$time.sql --additional-suffix=.sql  $dbName$time; \n")
		cmdStr.WriteString("rm -rf $dbName$time.sql; \n")
		cmdStr.WriteString("tar -zcvf $dbName$time.tar.gz --exclude=*.tar.gz ./; \n")
		cmdStr.WriteString("find -name \"*.sql\" -exec rm -Rf {} \\; \n")
		cmdStr.WriteString("cd ..; \n")
		cmdStr.WriteString("rm -rf $dbName`date -d \"" + t.Retain + " days ago\" +%Y%m%d`; \n")
		cmdStr.WriteString("echo $basePath$dirName$dbName$time.tar.gz \n")

		utils.Cmd(cmdStr.String())

		fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"), "end", task.Name)
	}(t)
}

func loadTaskDB(tasks *[]TaskDBObj) {
	fileLocker.Lock()
	data, err := ioutil.ReadFile("./task_db.json")
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
