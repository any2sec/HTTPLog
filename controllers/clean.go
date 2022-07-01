package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"httplogs/common"
	"httplogs/models"
)

type CleanController struct {
	beego.Controller
}

func (u *CleanController) Get(){
	resp := models.RespData{Msg: "success",HTTPStatusCode: "200"}
	common.DBClean()
	u.Data["json"] = resp
	u.ServeJSON()
}