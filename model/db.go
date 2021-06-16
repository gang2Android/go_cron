package model

import "strings"

type DB struct {
	Name      string `json:"name"`
	Host      string `json:"host"`
	Db        string `json:"db"`
	User      string `json:"user"`
	Pwd       string `json:"pwd"`
	BackPath  string `json:"backPath"`
	MysqlPath string `json:"mysqlPath"`
	Retain    string `json:"retain"`
}

func (db DB) String() string {
	return "{\"Name\":" + db.Name + ",\"Host\":" + db.Host + ",\"Db\":" + db.Db +
		",\"User\":" + db.User + ",\"Pwd\":" + db.Pwd +
		",\"BackPath\":" + db.BackPath + ",\"MysqlPath\":" + db.MysqlPath + ",\"Retain\":" + db.Retain + "}"
}

func (db DB) GetCmdStr() string {
	var cmdStr strings.Builder
	cmdStr.WriteString("#!/bin/bash \n")
	cmdStr.WriteString("host=" + db.Host + "; \n")
	cmdStr.WriteString("uName=" + db.User + "; \n")
	cmdStr.WriteString("pwd=" + db.Pwd + "; \n")
	cmdStr.WriteString("dbName=" + db.Db + "; \n")
	cmdStr.WriteString("basePath=" + db.BackPath + "; \n")
	cmdStr.WriteString("dirName=$dbName`date +\\%Y\\%m\\%d`/; \n")
	cmdStr.WriteString("time=`date +\\%Y\\%m\\%d\\%H\\%M`; \n")
	cmdStr.WriteString("cd $basePath; \n")
	cmdStr.WriteString("mkdir $dirName; \n")
	cmdStr.WriteString("echo $basePath$dirName$dbName$time; \n")
	cmdStr.WriteString(db.MysqlPath + "mysqldump -h $host -u$uName -p$pwd --default-character-set=utf8 --skip-extended-insert $dbName > $basePath$dirName$dbName$time.sql; \n")
	cmdStr.WriteString("cd $dirName; \n")
	cmdStr.WriteString("split --verbose -l 500 $dbName$time.sql --additional-suffix=.sql  $dbName$time; \n")
	cmdStr.WriteString("rm -rf $dbName$time.sql; \n")
	cmdStr.WriteString("tar -zcvf $dbName$time.tar.gz --exclude=*.tar.gz ./; \n")
	cmdStr.WriteString("find -name \"*.sql\" -exec rm -Rf {} \\; \n")
	cmdStr.WriteString("cd ..; \n")
	cmdStr.WriteString("rm -rf $dbName`date -d \"" + db.Retain + " days ago\" +%Y%m%d`; \n")
	cmdStr.WriteString("echo $basePath$dirName$dbName$time.tar.gz \n")
	return cmdStr.String()
}
