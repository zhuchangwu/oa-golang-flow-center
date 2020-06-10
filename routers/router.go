package routers

import (
	"github.com/astaxie/beego/plugins/cors"
	"oa-flow-centor/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowOrigins:     []string{"http://192.168.43.52"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "X-TOKEN"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	// 创建或者更新流程模版
	beego.Router("/flowTemplate/CreateOrUpdate", &controllers.FlowController{}, "post:CreateOrUpdate")
	// 分页查询+模糊查询
	beego.Router("/flowTemplate/GetFlowTemplateByPage", &controllers.FlowController{}, "post:GetFlowTemplateByPage")
	//
	beego.Router("/getFlowDetailByRecordId", &controllers.FlowController{}, "get:GetFlowDetailByRecord")

}
