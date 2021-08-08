package main

import (
	"blc-demo/web"
	"blc-demo/web/controller"
	"blc-demo/web/dao"
)

func main() {
	//Web
	dao.InitMysql()
	web.WebStart(&controller.Application{})
}
