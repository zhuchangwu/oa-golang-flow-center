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
	//	client1()
	// client2()
	// client3()
	 client4()
}

/**
 * 客户端服务端使用双向stream
 * 这种场景中，客户端可以将数据分批发送到服务端
 */
func client4() {
	// 创建客户端
	conn, err := grpc.Dial(":8083", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("error : %v", err)
		return
	}
	defer conn.Close()
	client := flow_center.NewFlowServiceClient(conn)

	// 调用方法发送请求
	stream, err := client.CreateFlowByBidirectionalStream(context.Background())
	if err != nil {
		fmt.Printf("error : %v", err)
		return
	}

	// 模拟有好多批数据
	for i := 0; i < 10; i++ {
		// 通过流发送一段数据
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
		err = stream.Send(f)

		res, err := stream.Recv()
		if err==io.EOF{
			break
		}
		if err != nil {
			fmt.Printf("error : %v", err)
			return
		}
		fmt.Printf("flowId:[%v]\n", res.GetFlowId())
		fmt.Printf("responseMsg:[%v]\n", res.GetResponseMsg())
	}

}

/**
 * 客户端发送stream
 * 这种场景中，客户端可以将数据分批发送到服务端
 */
func client3() {
	// 创建客户端
	conn, err := grpc.Dial(":8083", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("error : %v", err)
		return
	}
	defer conn.Close()
	client := flow_center.NewFlowServiceClient(conn)

	stream, err := client.CreateFlowRequestByClientStream(context.Background())
	if err != nil {
		fmt.Printf("error : %v", err)
		return
	}

	// 模拟客户端将大数据拆分成小分片，通过stream发往server
	for i := 0; i < 10; i++ {
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
		err = stream.Send(f)
		if err != nil {
			fmt.Printf("error : %v", err)
			return
		}
	}

	// 数据全部发送完成后，通过这个方法得到服务端的响应
	res, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Printf("error : %v", err)
		return
	}

	// 解析响应结果
	fmt.Printf("flowId:[%v]\n", res.GetFlowId())
	fmt.Printf("responseMsg:[%v]\n", res.GetResponseMsg())

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

	for {
		res, err := streamClient.Recv()
		if err == io.EOF {
			fmt.Println("********************************")
			// 这里面没有数据了～
			// fmt.Printf("flowId:[%v]\n", res.GetFlowId())
			// fmt.Printf("responseMsg:[%v]\n", res.GetResponseMsg())
			fmt.Println("********************************")
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

	time.Sleep(time.Second * 5)
}

/**
 * client和server交互 obj
 */
func client1() {
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
