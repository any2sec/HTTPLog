package main

import (
	beego "github.com/beego/beego/v2/server/web"
	"httplogs/common"
	_ "httplogs/common"
	_ "httplogs/routers"
)

func main() {
	common.CronInit()
	common.LogsInit()
	beego.Run()
	defer common.DB.Close()
}

