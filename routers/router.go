package routers

import (
	"github.com/logoocc/tokenservice/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/token", &controllers.TokenController{})
	beego.Router("/validate", &controllers.TokenValidateController{})
}
