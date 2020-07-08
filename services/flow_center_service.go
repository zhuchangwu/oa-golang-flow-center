package services

import (
	"context"
	"fmt"
	"io"
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
 * server接受client obj，返回给client obj
 */
func (s *FlowCenterService) CreateFlow(ctx context.Context, request *flow_center.Flow) (*flow_center.Response, error) {
	// 解析请求
	printFlowInfo(request)

	r := &flow_center.Response{
		FlowId:      1,
		ResponseMsg: "successfully",
	}
	return r, nil
}

/**
 * server接受client obj，返回给client stream
 */
func (s *FlowCenterService) CreateFlowResponseByServerStream(request *flow_center.Flow, stream flow_center.FlowService_CreateFlowResponseByServerStreamServer) error {

	// 解析请求
	printFlowInfo(request)

	for i := 0; i < 20; i++ {
		// 假设是循环读取某数据，并将每次读取到的结果发送到客户端中

		// 这里我们每次简单处理，每次返回一个response
		r := &flow_center.Response{
			FlowId:      1,
			ResponseMsg: "successfully",
		}
		err := stream.Send(r)
		if err != nil {
			fmt.Printf("error : %v", err)
			return err
		}

	}

	return nil
}

/**
 * server接受client obj，返回给client stream
 */
func (s *FlowCenterService) CreateFlowRequestByClientStream(clientStream flow_center.FlowService_CreateFlowRequestByClientStreamServer) (err error) {

	// client会发送stream过来，故我们在这里循环处理
	for {
		flow, err := clientStream.Recv()
		if err == io.EOF {
			fmt.Println("接受完毕，关闭stream")
			// 处理完成，向客户端写回响应
			r := &flow_center.Response{
				FlowId:      1,
				ResponseMsg: "successfully",
			}
			err = clientStream.SendAndClose(r)
			break
		}
		if err != nil {
			fmt.Printf("error : %v", err)
			return err
		}
		printFlowInfo(flow)
	}

	return
}

/**
 * client ，server 双向流通信
 */
func (s *FlowCenterService) CreateFlowByBidirectionalStream(stream flow_center.FlowService_CreateFlowByBidirectionalStreamServer) (err error) {
	// 接受循环处理client 消息
	for {
		flow, err := stream.Recv()
		if err == io.EOF {
			fmt.Printf("client ，server 双向流通信 err:[%v]",err)
			break
		}
		if err != nil {
			fmt.Printf("client ，server 双向流通信 err:[%v]",err)
			return err
		}

		// 处理数据
		printFlowInfo(flow)

		// 往client写会stream
		r := &flow_center.Response{
			FlowId:      1,
			ResponseMsg: "successfully",
		}
		err = stream.Send(r)
		if err != nil {
			fmt.Printf("error : %v", err)
			return err
		}
	}

	return err
}

// 打印
func printFlowInfo(flow *flow_center.Flow) {
	// 解析请求
	fmt.Println("-------------------------接收到请求------------------------")
	fmt.Printf("application：[%v]\n", flow.Applicate)
	fmt.Printf("deparmentId：[%v]\n", flow.DepartmentId)
	fmt.Printf("recoredId：[%v]\n", flow.RecordId)
	fmt.Printf("flowType：[%v]\n", flow.FlowType)
	fmt.Printf("roleMap：[%#v]\n", flow.RoleMap)
	fmt.Println("-------------------------接收到请求------------------------")
}
