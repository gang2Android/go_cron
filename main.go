package main

import (
	"cronProject/global"
	"cronProject/initialize"
)

func main() {
	initialize.InitBase()

	err := initialize.GetRouter().Run(":" + global.Config.Port)
	if err != nil {
		global.Logger.Error(err.Error())
	} else {
		global.Logger.Info(global.Config.Name + "start success on:" + "127.0.0.1" + global.Config.Port)
	}

	//task.Run()
}
