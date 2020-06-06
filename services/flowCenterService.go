package services

import (
	"context"
	"fmt"
	flow_center "oa-flow-centor/proto/oa/flow-center"
)

// 继承为我们生成的Client
type FlowCentorService struct {
	flow_center.FlowServiceClient
}

// 重写CreateFlow
func (s *FlowCentorService) CreateFlow(ctx context.Context, request *flow_center.Flow) (*flow_center.Response, error) {
	// 解析请求
	fmt.Println("-------------------------接收到请求------------------------")
	fmt.Printf("application：[%v]\n",request.Applicate)
	fmt.Printf("deparmentId：[%v]\n",request.DepartmentId)
	fmt.Printf("recoredId：[%v]\n",request.RecordId)
	fmt.Printf("flowType：[%v]\n",request.FlowType)
	fmt.Printf("roleMap：[%#v]\n",request.RoleMap)
	fmt.Println("-------------------------接收到请求------------------------")


	r := &flow_center.Response{
		FlowId:      1,
		ResponseMsg: "successfully",
	}
	return r, nil
}
