package common

import (
	"github.com/astaxie/beego/logs"
)

//var Logs *logs.BeeLogger
var Logs *logs.BeeLogger

func LogsInit(){
	Logs = logs.NewLogger(10000)
	//控制台打印日志
	//Logs.SetLogger(logs.AdapterConsole, `{"level":7,"color":true}`)
	//多文件输出日志
	//Logs.SetLogger(logs.AdapterMultiFile,`{"filename":"HTTPLog.log","level":3,"maxlines":100000000,"maxdays":90,"color":true,"separate":["notice","info","warning","error"]}`)
	//单文件输出日志
	Logs.SetLogger(logs.AdapterFile, `{"filename":"HTTPLog.log","maxlines":100000000,"maxdays":30,"color":true}`)
}
