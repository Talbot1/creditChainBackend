package controller

import (
	"blc-demo/web/dao"
	"blc-demo/web/model"
	"blc-demo/web/service"
	"blc-demo/web/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Application struct {
	JdService        *utils.JdService
	CompanyService   *service.CompanyService
	TrademarkService *service.TrademarkService
}

func (app *Application) TokenVerify2Generate(w http.ResponseWriter, r *http.Request) {
	if err := app.JdService.IsValid(); err != nil {
		fmt.Println("当前token已经失效, 重新获取")
		app.JdService.SetToken()
	}
}

func (app *Application) ForPreCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Add("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, X-Custom-Header") //header的类型
	w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")                                                                           //允许请求方法
	if r.Method == "OPTIONS" {
		fmt.Println("r.Method", r.Method, "跨域预检通过")
		w.WriteHeader(http.StatusNoContent)
		return
	}
}

func (app *Application) QueryJdAll(w http.ResponseWriter, r *http.Request) {
	app.TokenVerify2Generate(w, r)
	msg, err := app.CompanyService.QueyALl()
	if err != nil {
		fmt.Println("QueryAll Test Function Store Data In JDChain Error, This fuction is only used in dev mode")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	returnData := CommonReturnType{
		Status: "sucess",
		Data:   msg,
	}
	t, _ := json.Marshal(&returnData)
	w.Write(t)
}

// 数据库连接测试成功, jdchain未测试
func (app *Application) AddTransaction(w http.ResponseWriter, r *http.Request) {
	app.ForPreCheck(w, r)
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	addTransaction := model.TransactionVO{
		Id:   0,
		Name: r.PostFormValue("name"),
		Url:  r.PostFormValue("url"),
		Type: r.PostFormValue("type"),
	}
	if err = dao.InsertTransaction(&addTransaction); err != nil {
		fmt.Println("InsertTransaction Function Store Data In Mysql Error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// app.TokenVerify2Generate(w, r)
	// if err = app.CompanyService.AddTransaction(&addTransaction); err != nil {
	// 	fmt.Println("AddTransaction Function Store Data In JDChain Error")
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
	returnData := CommonReturnType{
		Status: "sucess",
		Data:   nil,
	}
	t, _ := json.Marshal(&returnData)
	w.Write(t)
}

// 数据库连接测试成功, jdchain未测试
func (app *Application) QueryAllTransaction(w http.ResponseWriter, r *http.Request) {
	// 先增添再查询
	app.ForPreCheck(w, r)
	app.TokenVerify2Generate(w, r)
	trademarkSlice, err := app.CompanyService.QueryAllTransaction()
	if err != nil {
		fmt.Println("QueryAllCompany info from JSChain Error, Retry from mysql")
		trademarkSlice, err = dao.QueryAllTransaction()
		if err != nil {
			fmt.Println("QueryAllTransaction Function Query Data From JDChain Error")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	var buffer bytes.Buffer
	bArrayMemberAlreadyWritten := false
	buffer.WriteString("[")
	for _, td := range trademarkSlice {
		if bArrayMemberAlreadyWritten {
			buffer.WriteString(",")
		}
		// Record is a JSON object, so we write as-is
		value, _ := json.Marshal(td)
		buffer.WriteString(string(value))
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")
	returnData := CommonReturnType{
		Status: "sucess",
		Data:   buffer.String(),
	}
	t, _ := json.Marshal(&returnData)
	w.Write(t)
}

// 数据库连接测试成功, jdchain未测试
func (app *Application) AddCompany(w http.ResponseWriter, r *http.Request) {
	app.ForPreCheck(w, r)
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	addCompany := model.CompanyVO{
		CompanyName:         r.PostFormValue("companyName"),
		LegalRepresentative: r.PostFormValue("legalRepresentative"),
		CreditCode:          r.PostFormValue("creditCode"),
		RegisteredCapital:   r.PostFormValue("registeredCapital"),
		PaidCapital:         r.PostFormValue("paidCapital"),
		EstablishedData:     "default",
		ApprovedDate:        "default",
		OwnRisk:             "default",
		AssociatedRisk:      "default",
		StaffSize:           "default",
		BusinessScope:       "default",
		Credit:              "default",
		Amount:              "default",
	}
	if err = dao.InsertCompany(&addCompany); err != nil {
		fmt.Println("InsertCompany Function Store Data In Mysql Error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// app.TokenVerify2Generate(w, r)
	// if err = app.CompanyService.AddCompany(&addCompany); err != nil {
	// 	fmt.Println("AddCompany Function Store Data In JDChain Error")
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
	returnData := CommonReturnType{
		Status: "sucess",
		Data:   nil,
	}
	t, _ := json.Marshal(&returnData)
	w.Write(t)
}

// 数据库连接测试成功, jdchain未测试
func (app *Application) QueryAllCompany(w http.ResponseWriter, r *http.Request) {
	app.TokenVerify2Generate(w, r)
	app.ForPreCheck(w, r)
	com, err := app.CompanyService.QueryAllCompany()
	if err != nil {
		fmt.Println("QueryAllCompany info from JSChain Error, Retry from mysql")
		com, err = dao.QueryAllCompany()
		if err != nil {
			fmt.Println("QueryAllCompany Function Query Data In Mysql Error")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
	var buffer bytes.Buffer
	bArrayMemberAlreadyWritten := false
	buffer.WriteString("[")
	for _, td := range com {
		if bArrayMemberAlreadyWritten {
			buffer.WriteString(",")
		}
		// Record is a JSON object, so we write as-is
		value, _ := json.Marshal(td)
		buffer.WriteString(string(value))
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")
	returnData := CommonReturnType{
		Status: "sucess",
		Data:   buffer.String(),
	}
	t, _ := json.Marshal(&returnData)
	w.Write(t)
}

// 数据库连接测试成功, jdchain未测试
func (app *Application) ChangeCredit(w http.ResponseWriter, r *http.Request) {
	app.ForPreCheck(w, r)
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, _ := strconv.Atoi(r.PostFormValue("id"))
	value, _ := strconv.Atoi(r.PostFormValue("value"))
	temp := model.Loan{
		Id:    id,
		Value: value,
	}
	// 真就完全没必要上数据库啊
	app.TokenVerify2Generate(w, r)
	if err = app.CompanyService.ChangeCredit(&temp); err != nil {
		fmt.Println("ChangeCredit Function Store Data in JdCahin Error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		t, _ := json.Marshal(&CommonReturnType{Status: "fail", Data: nil})
		w.Write(t)
		return
	}
	returnData := CommonReturnType{
		Status: "sucess",
		Data:   nil,
	}
	t, _ := json.Marshal(&returnData)
	w.Write(t)
}

// 数据库连接测试成功, jdchain未测试
func (app *Application) AddLoan(w http.ResponseWriter, r *http.Request) {
	app.ForPreCheck(w, r)
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, _ := strconv.Atoi(r.PostFormValue("id"))
	value, _ := strconv.Atoi(r.PostFormValue("value"))
	temp := model.Loan{
		Id:    id,
		Value: value,
	}
	if err = dao.InsertLoan(&temp); err != nil {
		fmt.Println("InsertLoan Function Store Data In Mysql Error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// app.TokenVerify2Generate(w, r)
	// if err = app.CompanyService.AddLoan(&temp); err != nil {
	// 	fmt.Println("AddLoan Function Store Data in JdCahin Error")
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// t, _ := json.Marshal(&CommonReturnType{Status: "fail", Data: nil})
	// w.Write(t)
	// return
	// }
	returnData := CommonReturnType{
		Status: "sucess",
		Data:   nil,
	}
	t, _ := json.Marshal(&returnData)
	w.Write(t)
}

// 逻辑测试成功, 但是未测试还有好几个key不了解的.
func (app *Application) QueryGuarantee(w http.ResponseWriter, r *http.Request) {
	app.TokenVerify2Generate(w, r)
	app.ForPreCheck(w, r)
	com, err := app.CompanyService.QueryAllCompany()
	if err != nil {
		fmt.Println("QueryGuarantee info from JSChain Error, Retry from mysql")
		com, err = dao.QueryAllCompany()
		if err != nil {
			fmt.Println("QueryGuarantee Function Query Data In Mysql Error")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
	var buffer bytes.Buffer
	bArrayMemberAlreadyWritten := false
	buffer.WriteString("[")
	for _, td := range com {
		if bArrayMemberAlreadyWritten {
			buffer.WriteString(",")
		}
		// Record is a JSON object, so we write as-is
		value, _ := json.Marshal(model.GuaranteeVO{
			CompanyName:      td.CompanyName,
			Credit:           td.Credit,
			FixedAssets:      "default",
			ProfitRate:       "default",
			GuaranteedAmount: "default",
		})
		buffer.WriteString(string(value))
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")
	returnData := CommonReturnType{
		Status: "sucess",
		Data:   buffer.String(),
	}
	t, _ := json.Marshal(&returnData)
	w.Write(t)
}

// // 进入查询页面
// func (app *Application) QueryPage(w http.ResponseWriter, r *http.Request) {
// 	data := utils.CheckLogin(r)
// 	ShowView(w, r, "PublicOption/queryPage.html", data)
// }

// // 根据ID查询信息
// func (app *Application) FindDataByID(w http.ResponseWriter, r *http.Request) {
// 	//data := utils.CheckLogin(r)
// 	//
// 	//ID := r.FormValue("id")
// 	//result, err := app.Setup.FindDataByID(ID)
// 	//if err != nil {
// 	//	log.Println(err)
// 	//}
// 	//var d = service.Company{}
// 	//err = json.Unmarshal(result, &d)
// 	//fmt.Println(d)
// 	//if err != nil {
// 	//	log.Println("unmarshal failed, err:", err)
// 	//}
// 	//var score float64
// 	//data.Data = d
// 	//for i := 0; i < len(d.Score); i++ {
// 	//	score += d.Score[i]
// 	//}
// 	//data.Score, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", score/float64(len(d.Score))), 64)
// 	//if score >= 95.0 {
// 	//	data.Rank = "A+"
// 	//} else if score >= 90.0 {
// 	//	data.Rank = "A"
// 	//} else if score >= 85.0 {
// 	//	data.Rank = "A-"
// 	//} else if score >= 80.0 {
// 	//	data.Rank = "B+"
// 	//} else if score >= 75.0 {
// 	//	data.Rank = "B"
// 	//} else if score >= 75.0 {
// 	//	data.Rank = "B-"
// 	//}
// 	//ShowView(w, r, "PublicOption/queryResult.html", data)
// }

// func SwitchTimeStampToData(timeStamp int64) string {
// 	t := time.Unix(timeStamp, 0)
// 	return t.Format("2006-01-02 15:04:05")
// }

// func (app *Application) AddDataPage(w http.ResponseWriter, r *http.Request) {
// 	data := utils.CheckLogin(r)
// 	if data.IsLogin {
// 		if data.IsStaff {
// 			ShowView(w, r, "StaffOption/addDataPage.html", data)
// 			return
// 		} else {
// 			data.Msg = "无权访问"
// 			ShowView(w, r, "index.html", data)
// 			return
// 		}
// 	} else if !data.IsLogin {
// 		ShowView(w, r, "AccountRelated/login.html", data)
// 		return
// 	}
// }

// func (app *Application) AddData(w http.ResponseWriter, r *http.Request) {
// 	//data := utils.CheckLogin(r)
// 	//
// 	//if data.IsStaff {
// 	//	//r.ParseMultipartForm(32 << 10)
// 	//	//获取表单输入
// 	//	id := r.FormValue("id")
// 	//	name := r.FormValue("company_name")
// 	//	legal := r.FormValue("legal")
// 	//	score, _ := strconv.ParseFloat(r.FormValue("score"), 64)
// 	//	rank := r.FormValue("rank")
// 	//
// 	//	//defer content.Close()
// 	//
// 	//	d := service.Company{
// 	//		id,
// 	//		name,
// 	//		legal,
// 	//		SwitchTimeStampToData(time.Now().Unix()),
// 	//		[]float64{score},
// 	//		rank,
// 	//	}
// 	//
// 	//	_, err := app.Setup.SaveData(d)
// 	//	if err != nil {
// 	//		fmt.Println("err3:", err)
// 	//	}
// 	//	data.Data = d
// 	//
// 	//	ShowView(w, r, "StaffOption/addSuccess.html", data)
// 	//	return
// 	//} else if !data.IsStaff {
// 	//	ShowView(w, r, "index.html", data)
// 	//	return
// 	//}
// }

// func (app *Application) RequestLoanPage(w http.ResponseWriter, r *http.Request) {
// 	data := utils.CheckLogin(r)
// 	if !data.IsLogin {
// 		ShowView(w, r, "AccountRelated/login.html", data)
// 		return
// 	}
// 	if data.IsStaff {
// 		ShowView(w, r, "StaffOption/requestLoanPage.html", data)
// 		return
// 	} else {
// 		data.Msg = "无权访问"
// 		ShowView(w, r, "index.html", data)
// 		return
// 	}

// }

// func (app *Application) RequestLoan(w http.ResponseWriter, r *http.Request) {
// 	data := utils.CheckLogin(r)
// 	if !data.IsLogin {
// 		ShowView(w, r, "AccountRelated/login.html", data)
// 		return
// 	}
// 	if data.IsStaff {
// 		ShowView(w, r, "StaffOption/requestSucess.html", data)
// 		return
// 	} else {

// 		data.Msg = "无权访问"
// 		ShowView(w, r, "index.html", data)
// 		return
// 	}
// }
