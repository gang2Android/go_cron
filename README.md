# cron_go

通过go的cron+json配置文件的方式实现可配置的计划任务

根目录的task.json为计划任务的配置文件

```json
[
  {
    "name": "计划任务名称",
    "spec": "任务执行间隔,如每分钟执行一次:0 0/1 * * *",
    "url": "访问的http地址"
  }
]
```

根目录的task_db.json为备份数据库计划任务的配置文件-mysql5.6

```json
[
  {
    "name": "计划任务名称",
    "host": "数据库地址",
    "db": "数据库名称",
    "user": "用户名",
    "pwd": "密码",
    "backPath": "备份到哪个文件夹下",
    "mysqlPath": "mysql安装目录(如/www/server/mysql/bin/)",
    "retain": "数据库保留天数",
    "spec": "任务执行间隔,如每小时执行一次:0 0 0/1 * *"
  }
]
```

根目录的task_disk.json为检查系统磁盘计划任务的配置文件-当低于min就会发生短信到sms_mobile手机上

```json
[
  {
    "name": "数据盘",
    "spec": "0 0 0/1 * * ?",
    "path": "/mnt",
    "min": 5,
    "sms_ak": "阿里云AccessKeyID",
    "sms_as": "阿里云AccessKeySecret",
    "sms_end": "dysmsapi.aliyuncs.com",
    "sms_name": "短信签名",
    "sms_code": "验证码短信模板id",
    "sms_mobile": "接收预警的手机号"
  }
]
```

部署

```shell
set GOOS=linux
set GOARCH=amd64
go build -o cron_app
# 上传到服务器目录下
chmod 777 cron_app
nohup ./cron_app &
```

