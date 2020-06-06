package controllers

type FlowController struct {
	BaseController
}

type FlowDetail struct {
	Id        int
	Applicate string
	RecordId  int
	Content   string
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
}
