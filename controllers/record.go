package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"httplogs/common"
	"httplogs/models"
	"math/rand"
)

type RecordController struct {
	beego.Controller
}

//获取一个random
func (u *RecordController) Post() {
	localIP,_:= beego.AppConfig.String("local.ip")
	httpPort,_:= beego.AppConfig.String("httpport")
	url:=fmt.Sprintf("http://%s:%s?random=%s",localIP,httpPort,RandomString())
	respData := models.RespData{Msg: url,HTTPStatusCode: "200"}
	u.Data["json"] =respData
	u.ServeJSON()
}

//查询random
func (u *RecordController) Get(){
	var records []models.Record
	query := u.GetString("query")
	if query != ""{
		err1 := common.DB.Where("random=?",query).Order("created_at DESC").Find(&records)
		if err1.Error != nil{
			common.Logs.Error(err1.Error.Error())
		}
	u.Data["json"]=records
	u.ServeJSON()

	}else{
		err1 := common.DB.Order("created_at DESC").Find(&records)
		if err1.Error != nil{
			common.Logs.Error(err1.Error.Error())
		}
		u.Data["json"]=records
		u.ServeJSON()
	}
}

func RandomString() string{
	var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	var letters []rune

	letters = defaultLetters

	b := make([]rune, 16)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}