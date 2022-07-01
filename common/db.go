package common
import (
	"fmt"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"httplogs/models"

)

var DB *gorm.DB
var err error

func init() {
	dbhost, _ := web.AppConfig.String("db.host")
	dbport, _ := web.AppConfig.String("db.port")
	dbuser, _ := web.AppConfig.String("db.user")
	dbpassword, _ := web.AppConfig.String("db.passwd")
	dbname, _ := web.AppConfig.String("db.name")
	//timezone, _ := web.AppConfig.String("db.timezone")
	pre, _ := web.AppConfig.String("db.pre")
	if dbport == "" {
		dbport = "3306"
	}
	args := dbuser + ":" + dbpassword + "@(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&parseTime=true&loc=Local" //&serverTimezone="+timezone
	DB, err = gorm.Open("mysql", args)
	if err != nil {
		fmt.Println(err)
		return
	}
	sqlDB := DB.DB()
	//设置数据库连接池参数
	sqlDB.SetMaxOpenConns(100) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。

	// 全局禁用表名复数
	//DB.SingularTable(true)
	//更改默认表名，添加prefix
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return pre + defaultTableName
	}

	//数据库初始化
	//初始record表
	// 检查模型`record`表是否存在
	HasTableModel := DB.HasTable(&models.Record{})
	// 检查表`record`是否存在
	HasTable := DB.HasTable("httplog_record")
	// 为模型`record`创建表
	if !HasTable && !HasTableModel {
		DB.CreateTable(&models.Record{})
		// 创建表`record'时将“ENGINE = InnoDB”附加到SQL语句
		DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&models.Record{})
	}

}