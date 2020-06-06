package main

import (
	"google.golang.org/grpc"
	"net"
	_ "oa-flow-centor/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()


}
