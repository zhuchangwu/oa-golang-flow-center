package routers

import (
	"oa-flow-centor/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/getFlowDetailByRecordId", &controllers.FlowController{},"get:GetFlowDetailByRecord")
}
