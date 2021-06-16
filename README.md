# cron_go

从mysql读取计划任务配置启动计划任务 通过api控制任务

api

- 获取信息 `http://127.0.0.1:9999/api/info`
- 启动所有任务 `http://127.0.0.1:9999/api/start`
- 停止所有任务 `http://127.0.0.1:9999/api/stop`
- 添加任务 `http://127.0.0.1:9999/api/add?id=数据库计划任务表id`
- 停止任务 `http://127.0.0.1:9999/api/remove?id=数据库计划任务表id`

config.yaml说明

|参数名|类型|说明|
|:---- |:----- |----- |
|name |string |服务名 |
|port |int |服务启动端口 |
|mysql |object |mysql数据库配置 |
|mysql-host |string |数据库-地址 |
|mysql-port |int |数据库-端口 |
|mysql-db |string |数据库-名称 |
|mysql-name |string |数据库-用户名 |
|mysql-pwd |string |数据库-密码 |
|redis |object |redis配置 |
|redis-host |string |redis-地址 |
|redis-port |int |redis-端口 |
|redis-db |string |redis-数据库id |
|redis-pwd |string |redis-密码 |
|logs |object |日志 |
|logs-file_path |string |日志-存储路径如：./logs/cron_app.log |
|logs-max_size |int |日志-单个文件大小上限,单位M |
|logs-max_backups |int |日志-保留份数 |
|logs-max_age |int |日志-保存天数 |
|logs-compress |bool |日志-是否压缩 |
|logs-debug |bool |日志-是否debug |

cron_task.sql字段说明

|参数名|类型|说明|
|:---- |:----- |----- |
|id |int |自增id |
|name |string |任务名称 |
|type |int |任务类型,1访问url,2备份mysql |
|spec |string |执行间隔,cron表达式 |
|enable |int |是否启用,1启用2停用 |
|content |string |任务内容1http网址,2mysql对象 |
|task_no |string |cron任务id,用于管理任务 |
|add_time |timestamp |添加时间 |

cron_task.sql - content内容当type=2时

|参数名|类型|说明|
|:---- |:----- |----- |
|name |string |数据库名称 |
|host |string |数据库地址 |
|db |string |数据库名称 |
|user |string |数据库用户名称 |
|pwd |string |数据库密码 |
|backPath |string |备份文件的保存位置,绝对路径 |
|mysqlPath |string |本地mysql安装位置,绝对路径 |
|retain |int |备份文件保存天数 |

部署

```shell
set GOOS=linux
set GOARCH=amd64
go build -o cron_app
# 上传到服务器目录下
# 导入数据库表
mkdir "logs"
chmod 777 cron_app
nohup ./cron_app &
```
