package main

import (
	_ "beelike/routers"
	"github.com/astaxie/beego"
	"beelike/models"
	"fmt"
	"beelike/util/mongo"
	"beelike/util/helper"
	"github.com/goinggo/tracelog"
	"os"
)

//应用配置文件
var configFile = "conf/conf.yml"

func init() {
	models.ParseConfigFile(configFile)
}

func main() {

	tracelog.Start(tracelog.LevelTrace)
	err := mongo.Startup(helper.MainGoRoutine)
	if err != nil {
		tracelog.CompletedError(err, helper.MainGoRoutine, "initApp")
		os.Exit(1)
	}else {
		fmt.Println("mongodb启动正常")
	}
	beego.Run()

	tracelog.Completed(helper.MainGoRoutine, "Website Shutdown")
	tracelog.Stop()
}

