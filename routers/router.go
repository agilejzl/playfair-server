package routers

import (
	"playfair-server/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/api/service", &controllers.ServiceController{})
}
