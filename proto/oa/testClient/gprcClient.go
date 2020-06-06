package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	flow_center "oa-flow-centor/proto/oa/flow-center"
)

func main() {
	// 创建一个连接
	conn, err := grpc.Dial(":8083",grpc.WithInsecure())
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
	response,err:=client.CreateFlow(context.Background(), f)
	if err != nil {
		fmt.Printf("error : %v", err)
		return
	}

	// 解析响应结果
	fmt.Printf("flowId:[%v]\n",response.GetFlowId())
	fmt.Printf("responseMsg:[%v]\n",response.GetResponseMsg())
}
