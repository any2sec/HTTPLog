package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"httplogs/common"
	"httplogs/models"
	"time"
)

type MainController struct {
	beego.Controller
}


func (c *MainController) Get() {
	random :=c.GetString("random")
	if random != ""{
	ip := c.Ctx.Input.IP()
	l,_ := time.LoadLocation("Asia/Shanghai")
	nowTime := time.Now().In(l).Format("2006-01-02 15:04:05")
	record := &models.Record{RequestIP: ip,Random: random,CreatedAt: nowTime}
	err := common.DB.Create(&record)
	if err.Error!=nil{
		common.Logs.Error(err.Error.Error())
	}else {
		c.Data["json"]= record
		c.ServeJSON()
	}

	}else{
		c.TplName = "index.tpl"
	}
}
