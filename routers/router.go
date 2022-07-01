package routers

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/server/web/context"
	"httplogs/common"
	"httplogs/models"
	"io/ioutil"

	beego "github.com/beego/beego/v2/server/web"
	"httplogs/controllers"
)

func JsonRespData(resp models.RespData) string {
	rs, err := json.Marshal(resp)
	if err != nil {
		common.Logs.Error(err.Error())
	}
	return string(rs)
}

var FilterUser = func(ctx *context.Context) {
	token, _ := beego.AppConfig.String("user.token")
	respData := models.RespData{Msg: "error",HTTPStatusCode: "401"}
	var data map[string]string
	token1, _ := ioutil.ReadAll(ctx.Request.Body)
	err0 := json.Unmarshal(token1,&data)
	header_token := ctx.Request.Header.Get("token")

	if err0 == nil && data["token"] == token{
			respData.Msg="success";respData.HTTPStatusCode="200"
			fmt.Fprintf(ctx.ResponseWriter,JsonRespData(respData))
	}else if header_token != token{
		fmt.Fprintf(ctx.ResponseWriter,JsonRespData(respData))
	}
}


func init() {
	beego.InsertFilter("/record/*", beego.BeforeRouter, FilterUser)
	//beego.InsertFilter("/static/*", beego.BeforeRouter, FilterUser)
	beego.Router("/", &controllers.MainController{})
	beego.Router("/record", &controllers.RecordController{})
	beego.Router("/record/clean", &controllers.CleanController{})

	//错误处理
	beego.ErrorController(&controllers.ErrorController{})
}
