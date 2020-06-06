package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	flow_center "oa-flow-centor/proto/oa/flow-center"
	"time"
)

func main() {
	//	client()
		client2()
}

/**
 * 服务端返回给客户端stream
 */
func client2() {

	// 创建client
	conn, err := grpc.Dial(":8083", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("error : %v", err)
		return
	}
	defer conn.Close()
	client := flow_center.NewFlowServiceClient(conn)

	f := &flow_center.Flow{
		FlowType:  "reibersement",
		RecordId:  1,
		Applicate: "changwu",
		RoleMap: map[string]string{
			"manager": "tom",
			"boss":    "jerry",
		},
		DepartmentId: 2,
	}

	// 发送请求
	streamClient, err := client.CreateFlowResponseByServerStream(context.Background(), f)
	if err != nil {
		fmt.Printf("error : %v", err)
		return
	}

	for  {
		res, err := streamClient.Recv()
		if err==io.EOF{
			fmt.Printf("flowId:[%v]\n", res.GetFlowId())
			fmt.Printf("responseMsg:[%v]\n", res.GetResponseMsg())
			return
		}
		if err != nil {
			fmt.Printf("error : %v", err)
			return
		}
		// 解析响应结果
		fmt.Printf("flowId:[%v]\n", res.GetFlowId())
		fmt.Printf("responseMsg:[%v]\n", res.GetResponseMsg())
	}

	time.Sleep(time.Second*5)
}

/**
 * client和server交互 obj
 */
func client() {
	// 创建一个连接
	conn, err := grpc.Dial(":8083", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("error : %v", err)
		return
	}
	defer conn.Close()

	// 基于第一个连接，创建客户端
	client := flow_center.NewFlowServiceClient(conn)

	// 构建请求参数
	f := &flow_center.Flow{
		FlowType:  "reibersement",
		RecordId:  1,
		Applicate: "changwu",
		RoleMap: map[string]string{
			"manager": "tom",
			"boss":    "jerry",
		},
		DepartmentId: 2,
	}

	// 发送grpc请求
	response, err := client.CreateFlow(context.Background(), f)
	if err != nil {
		fmt.Printf("error : %v", err)
		return
	}

	// 解析响应结果
	fmt.Printf("flowId:[%v]\n", response.GetFlowId())
	fmt.Printf("responseMsg:[%v]\n", response.GetResponseMsg())
}
