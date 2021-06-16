package initialize

import (
	"cronProject/global"
	"github.com/go-redis/redis"
)

/**
redis-server &	启动redis
redis-cli shutdown	暂停redis
redis-cli	进入redis命令行
set key value	设置 key 的值
get key	获取 key 的值
exists key	查看此 key 是否存在
keys *	查看所有的 key
flushall	消除所有的 key
*/
func GetRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: global.Config.Redis.Host + ":" + global.Config.Redis.Port,
		//Password: "ZcEsd7aPXMwem5RV",
		Password: global.Config.Redis.Pwd,
		DB:       global.Config.Redis.Db,
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	return client
}
