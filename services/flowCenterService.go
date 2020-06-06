package services

import (
	"context"
	"fmt"
	flow_center "oa-flow-centor/proto/oa/flow-center"
)

/**
 * 编译器为我们生成了 FlowServiceServer 接口，里面定义了我们的在 proto文件中定义的接口信息
 * 我们需要实现实现这个接口中的抽象方法～
 */
type FlowCenterService struct {
	flow_center.FlowServiceServer
}

/**
 * 服务端接受client obj，返回给client obj
 */
func (s *FlowCenterService) CreateFlow(ctx context.Context, request *flow_center.Flow) (*flow_center.Response, error) {
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


/**
 * 服务端接受client obj，返回给client stream
 */
func (s *FlowCenterService)CreateFlowResponseByServerStream(request *flow_center.Flow,stream flow_center.FlowService_CreateFlowResponseByServerStreamServer)(error){

	// 解析请求
	fmt.Println("-------------------------接收到请求------------------------")
	fmt.Printf("application：[%v]\n",request.Applicate)
	fmt.Printf("deparmentId：[%v]\n",request.DepartmentId)
	fmt.Printf("recoredId：[%v]\n",request.RecordId)
	fmt.Printf("flowType：[%v]\n",request.FlowType)
	fmt.Printf("roleMap：[%#v]\n",request.RoleMap)
	fmt.Println("-------------------------接收到请求------------------------")

	for i := 0; i < 20; i++ {
		// 假设是循环读取某数据，并将每次读取到的结果发送到客户端中

		// 这里我们每次简单处理，每次返回一个response
		r := &flow_center.Response{
			FlowId:      1,
			ResponseMsg: "successfully",
		}
		err:=stream.Send(r)
		if err != nil {
			fmt.Printf("error : %v", err)
			return err
		}

	}

	return nil
}
