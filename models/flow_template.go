package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/orm" // todo 这里的orm导包得手动搞
	"strconv"
	"time"
)

type FlowTemplateForm struct {
	Id               int       `json:"id" orm:"column(id) pk;auto"`
	CreaterUsername  string    `json:"createrUsername" orm:"column(creater_username) size(64)"`
	FlowTemplateName string    `json:"flowTemplateName" orm:"column(flow_template_name) size(64)"`
	FlowNodes        string    `json:"flowNodes" orm:"column(flow_nodes) size(64)"`
	CreateTime       time.Time `json:"createTime" orm:"column(create_time) type(datetime);auto_now_add"`
	//POrmer           orm.Ormer
}

func (f *FlowTemplateForm) TableName() string {
	return "flowTemplate"
}

func NewFlowTemplate() *FlowTemplateForm {
	return &FlowTemplateForm{}
}

// 获取所有的模版信息
func (f *FlowTemplateForm) GetAllFlowTemplate() (*[]FlowTemplateForm, error) {
	pOrmer := orm.NewOrm()
	var arr []FlowTemplateForm
	_, err := pOrmer.Raw("select * from flowTemplate").QueryRows(&arr)
	if err != nil {
		fmt.Printf("error : %v", err)
		return nil, err
	}
	return &arr, nil
}
// todo 下面直接拼接sql，有sql注入的风险
// 分页查询+模糊查询
func (f *FlowTemplateForm) FindFlowTemplateByPageAndTemplateName(currentPage, pageSize int, name string) (fts *[]FlowTemplateForm, err error) {
	POrmer := orm.NewOrm()
	if currentPage <= 1 {
		currentPage = 1
	}
	if pageSize < 0 {
		pageSize = 10
	}
	var arr []FlowTemplateForm
	var sql string
	if len(name)>0 {
		sql = "select * from flowTemplate where flow_template_name like '%"+name+"%' limit "+ strconv.Itoa((currentPage-1)*pageSize)+","+strconv.Itoa(pageSize)+";"
	}else{
		sql = "select * from flowTemplate limit "+ strconv.Itoa((currentPage-1)*pageSize)+","+strconv.Itoa(pageSize)+";"
	}
	_, err = POrmer.Raw(sql).QueryRows(&arr)
	return &arr, err
}

// 查询库表中的总数
func (f *FlowTemplateForm) GetCountFormTable() (int, error) {
	POrmer := orm.NewOrm()
	var count int
	err := POrmer.Raw("select count(*) from flowTemplate").QueryRow(&count)
	if err != nil {
		fmt.Printf("error : %v", err)
		return -1, err
	}
	return count, nil
}

// 插入一条记录
func (f *FlowTemplateForm) InsertOrUpdate() (id int64, err error) {
	pOrmer := orm.NewOrm()
	id, err = pOrmer.InsertOrUpdate(f)
	if err != nil {
		fmt.Printf("error : %v", err)
		return -1, err
	}
	return
}
