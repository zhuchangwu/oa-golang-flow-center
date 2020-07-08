package controllers

import (
	"encoding/json"
	"fmt"
	"google.golang.org/genproto/googleapis/type/date"
	"oa-flow-centor/common"
	"oa-flow-centor/models"
	"time"
)

// todo panic 处理～～～
type FlowController struct {
	BaseController
}

type FlowDetail struct {
	Id        int
	Applicate string
	RecordId  int
	Content   string
}

type FlowTemplate struct {
	Id               int       `json:"id"`
	CreaterUsername  string    `json:"createrUsername"`
	FlowTemplateName string    `json:"flowTemplateName"`
	FlowNodes        string    `json:"flowNodes"`
	CreateTime       date.Date `json:"createTime"`
}

// 分页+根据名称查询～
type QueryInfo struct {
	CurrentPage int    `json:"currentPage"`
	PageSize    int    `json:"limit"`
	Name        string `json:"name"`
}

// 返回的当前页，总页数
type PageInfo struct {
	Count           int
	CurrentPageData *[]models.FlowTemplateForm
}

/**
 *  分页查询+模糊查询
 */
func (c *FlowController) GetFlowTemplateByPage() {
	qi := QueryInfo{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &qi)
	if err != nil {
		fmt.Printf("error : %v", err)
		return
	}
	template := models.NewFlowTemplate()
	templateArr, err := template.FindFlowTemplateByPageAndTemplateName(qi.CurrentPage, qi.PageSize, qi.Name)
	if err != nil {
		c.Data["json"] = common.CreateErrorResultInfo("fail to get flow template by page or name")
		c.ServeJSON()
		fmt.Printf("error : %v", err)
		return
	}
	count, err := template.GetCountFormTable()
	if err != nil {
		fmt.Printf("error : %v", err)
		return
	}
	pageInfo := PageInfo{
		Count:           count,
		CurrentPageData: templateArr,
	}
	c.Data["json"] = common.CreateSuccessfulResultInfo("get flow template by page or name", &pageInfo)
	c.ServeJSON()
}

/**
 *  创建或更新流程模版
 */
func (c *FlowController) CreateOrUpdate() {

	ft := FlowTemplate{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &ft)
	if err != nil {
		fmt.Printf("error : %v", err)
		return
	}

	// 将页面中传递过来的参数struct，转换成From
	template := models.NewFlowTemplate()
	template.Id = ft.Id
	template.CreaterUsername = ft.CreaterUsername
	template.FlowTemplateName = ft.FlowTemplateName
	template.FlowNodes = ft.FlowNodes
	template.CreateTime = time.Now()

	// 插入单挑记录
	insert, err := template.InsertOrUpdate()
	if err != nil {
		fmt.Printf("error : %v", err)
		c.Data["json"] = common.CreateErrorResultInfo("fail to insert flowTemplate " + err.Error())
		c.ServeJSON()
		return
	}
	if insert < 0 {
		fmt.Printf("error : %v", err)
		c.Data["json"] = common.CreateErrorResultInfo("fail to insert flowTemplate" + err.Error())
		return
	}

	c.Data["json"] = common.CreateSuccessfulResultInfo("add flowTemplate successful")
	c.ServeJSON()
}

/**
 * 根据根据 记录id获取记录详情
 */
func (c *FlowController) GetFlowDetailByRecord() {
	recordId, _ := c.GetInt("recordId")
	detail := &FlowDetail{
		Id:        1,
		Applicate: "zhuchangwu",
		RecordId:  recordId,
		Content:   "申请报销发票",
	}
	c.Data["json"] = detail
	c.ServeJSON()
	fmt.Println("准备睡眠")
	time.Sleep(time.Second * 10)
	fmt.Println("结束睡眠")
}

/**
 *  注解路由
 *  获取所有的模版信息
 */
// @router /flowTemplate/getAllFlowTemplateInfo/ [get]
func (c *FlowController) GetAllFlowTemplateInfo() {

	template := models.NewFlowTemplate()
	arr, err := template.GetAllFlowTemplate()
	if err != nil {
		fmt.Printf("error : %v", err)
		info := common.CreateErrorResultInfo("fail to get all flow template info")
		c.Data["json"] = info
		c.ServeJSON()
		return
	}

	c.Data["json"] = common.CreateSuccessfulResultInfo("successful to get all flow template info", arr)
	c.ServeJSON()
}

/**
 * 删除模版信息
 */
func (c *FlowController) DeleteFlowTemplate() {
	id, err := c.GetInt("id")
	if err != nil {
		fmt.Printf("error : %v", err)
		return
	}
	fmt.Println(id)
	c.Data["json"] = "ok"
	c.ServeJSON()
}
