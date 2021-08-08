package main

import (
	"blc-demo/web"
	"blc-demo/web/cliInit"
	"blc-demo/web/controller"
	"blc-demo/web/dao"
	"blc-demo/web/service"
	"blc-demo/web/utils"
)

func main() {
	//Web
	dao.InitMysql()

	app := controller.Application{
		cliInit.CliInit(),
		&utils.JdService{},
		&service.CompanyService{},
		&service.TrademarkService{},
	}

	//defer cliInit.SDK.Close()

	web.WebStart(&app)
}
