package common

import (
	"github.com/beego/beego/v2/adapter/toolbox"
	"httplogs/models"
	"time"
)

func CronInitTask()error{
	return nil
}

func DBCleanTask() error{
	var xMonths int;xMonths=0
	var xDays int;xDays=0
	//获取当前时间
	l,_ := time.LoadLocation("Asia/Shanghai")
	Time := time.Now().In(l)
	Beforex := Time.AddDate(0,xMonths,xDays).Format("2006-01-02 15:04:05")

	Logs.Info("DBCleanExecute Task In " + Time.Format("2006-01-02 15:04:05"))
	//删除result表中xMonths个月前的记录
	err:=DB.Where("created_at < ?", Beforex).Delete(models.Record{})
	if err.Error!=nil{
		Logs.Error(err.Error.Error())
	}
	return nil
	return nil
}

func DBClean() error{
	var xMonths int;xMonths=0
	var xDays int;xDays=0
	//获取当前时间
	l,_ := time.LoadLocation("Asia/Shanghai")
	Time := time.Now().In(l)
	Beforex := Time.AddDate(0,xMonths,xDays).Format("2006-01-02 15:04:05")

	Logs.Info("DBCleanExecute In " + Time.Format("2006-01-02 15:04:05"))
	//删除result表中xMonths个月前的记录
	err:=DB.Where("created_at < ?", Beforex).Delete(models.Record{})
	if err.Error!=nil{
		Logs.Error(err.Error.Error())
	}
	return nil
}

func CronInit() {
	//创建任务             (任务名称， 任务时间描述， 所需执行的函数)
	InitTask := toolbox.NewTask("InitTask", "0 0 21 * * 1", CronInitTask)
	cleanTask := toolbox.NewTask("DBCleanTask", "0 0 */3 * * *", DBCleanTask)

	//任务执行(立即)
	err := InitTask.Run()
	if err != nil{
		Logs.Error(err.Error())
	}
	//err1 :=ProgressNightTask.Run()
	//if err1 != nil{
	//	fmt.Println(err)
	//}
	//添加任务
	toolbox.AddTask("InitTask",InitTask)
	toolbox.AddTask("DBCleanTask", cleanTask)
	//任务执行(会根据添加任务的时间点)
	toolbox.StartTask()
}