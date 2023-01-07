package routers

import (
	"cppm/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/insert", &controllers.MainController{}, "post:Insert")
	beego.Router("/cmds", &controllers.MainController{}, "post:Cmds")
    beego.Router("/policy", &controllers.PolicyController{})
	beego.Router("/nodes", &controllers.NodeController{})
	beego.Router("/nodes/gather", &controllers.NodeController{}, "post:Gather")
	beego.Router("/pods/proclist", &controllers.PodController{})
	beego.Router("/nodes/?:name", &controllers.PodController{})
	//beego.Router("/nodes/pods", &controllers.PodController{}, "post:RegisterPod")
}
