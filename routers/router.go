package routers

import (
	"myapp/controllers"
	"myapp/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.AutoRouter(&admin.RoleController{})
}
