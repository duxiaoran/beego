package main

import (
	_ "myapp/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"myapp/models"
)

func init() {
	models.Init()
	beego.BConfig.WebConfig.Session.SessionOn = true
}

func main() {
	beego.Run()
}

