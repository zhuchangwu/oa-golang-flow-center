package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	flow_center "oa-flow-centor/proto/oa/flow-center"
	"oa-flow-centor/services"
)

func main() {
	// 构建server
	grpcSewrver := grpc.NewServer()
	// 注册server
	flow_center.RegisterFlowServiceServer(grpcSewrver,new(services.FlowCenterService))
	// 监听，协议
	lis, _ :=net.Listen("tcp", ":8083")
	fmt.Println("Server started")
	// 开启服务
	grpcSewrver.Serve(lis)

}
