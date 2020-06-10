package main

import (
	"github.com/astaxie/beego"
	_ "oa-flow-centor/models"
	_ "oa-flow-centor/routers"
)

func main() {
	beego.Run()
}
