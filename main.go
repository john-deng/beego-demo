package main

import (
	_ "github.com/john-deng/beego-demo/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type DataSource struct {
	aliasName string
	driverName string
	host string
	port string
	database string
	username string
	password string
}

func init() {
	dataSource := DataSource{
		"default",
		"mysql",
		"mysql-dev",
		"3306",
		"beego_demo",
		os.Getenv("MYSQL_USERNAME"),
		os.Getenv("MYSQL_PASSWORD"),

	}
	url := dataSource.username + ":" + dataSource.password + "@tcp(" + dataSource.host + ":" + dataSource.port + ")/" + dataSource.database
	orm.RegisterDataBase(dataSource.aliasName, dataSource.driverName, url)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

