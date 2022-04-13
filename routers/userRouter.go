package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"regInfo/controllers"
)

func init() {
	beego.Router("/", &controllers.UsrController{})
	beego.Router("/Check", &controllers.UsrController{}, "get:ShowRegister;post:HandleUsrPost")
}
