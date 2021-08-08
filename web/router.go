package web

import (
	"blc-demo/web/controller"
	"fmt"
	"net/http"
)

func WebStart(app *controller.Application) {

	// fs := http.FileServer(http.Dir("./web/assets"))
	// http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// 查询全部数据
	http.HandleFunc("/queryJdAll", app.QueryJdAll)
	// 商标处理部分
	http.HandleFunc("/mark/addApply", app.AddApply)
	http.HandleFunc("/mark/queryAllTrademark", app.QueryAllTrademark)
	http.HandleFunc("/mark/queryMortmark", app.QueryMortmark)
	http.HandleFunc("/mark/queryAucmark", app.QueryAucmark)
	http.HandleFunc("/mark/agreeApply", app.AgreeApply)
	http.HandleFunc("/mark/rejectApply", app.RejectApply)
	http.HandleFunc("/mark/updateMortgage", app.UpdateMortgage)
	http.HandleFunc("/mark/updateAuctioning", app.UpdateAuctioning)
	http.HandleFunc("/mark/bid", app.Bid)

	// 公司和担保部分
	http.HandleFunc("/admin/addTx", app.AddTransaction)
	http.HandleFunc("/admin/transaction", app.QueryAllTransaction)
	http.HandleFunc("/admin/addcompany", app.AddCompany)
	http.HandleFunc("/admin/company", app.QueryAllCompany)
	http.HandleFunc("/admin/marks", app.ChangeCredit)
	http.HandleFunc("/admin/loan", app.AddLoan)
	http.HandleFunc("/gurantee/promise", app.QueryGuarantee)

	fmt.Println("---------------------------------------------")
	fmt.Println("启动Web服务, 监听端口号: 9000")

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("启动Web服务错误")
	}

}
