package controller

import (
	"blc-demo/web/dao"
	"blc-demo/web/model"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

//  数据库测试通过, jdchain未测试
func (app *Application) AddApply(w http.ResponseWriter, r *http.Request) {
	app.ForPreCheck(w, r)
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	addApply := model.TrademarkVO{
		MarkKey:      r.PostFormValue("markname"),
		TranHash:     "default",
		MarkName:     "default",
		MarkPic:      r.PostFormValue("markpic"),
		MarkCate:     r.PostFormValue("markcate"),
		MarkRegister: r.PostFormValue("markregister"),
		ApplyDate:    r.PostFormValue("applydate"),
		ValidityTime: "default",
		Applyer:      r.PostFormValue("applyer"),
		MarkValue:    "default",
		LoanValue:    "default",
		PaybackDate:  "default",
		MoneyNow:     "default",
		Bidder:       "default",
		RestTime:     "default",
	}

	if err = dao.InsertApply(&addApply); err != nil {
		fmt.Println("InsertApply Function Store Data In Mysql Error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// app.TokenVerify2Generate(w, r)
	// if err = app.TrademarkService.AddApply(&addApply); err != nil {
	// 	fmt.Println("addApply Function Store Data In JDChain Error")
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

//  数据库测试通过,返回数据格式未对齐 jdchain未测试
func (app *Application) QueryAllTrademark(w http.ResponseWriter, r *http.Request) {
	app.ForPreCheck(w, r)
	app.TokenVerify2Generate(w, r)
	trademarkSlice, err := app.TrademarkService.QueryAllTrademark()
	if err != nil {
		fmt.Println("QueryAllCompany info from JSChain Error, Retry from mysql")
		trademarkSlice, err = dao.QueryALLTrademark()
		if err != nil {
			fmt.Println("QueryAllTransaction Function Query Data From JDChain Error")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
	fmt.Println("数据库查询出来的结果是: ", trademarkSlice)
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
	fmt.Println("buffer中的数据: ", buffer.String())
	returnData := CommonReturnType{
		Status: "sucess",
		Data:   buffer.String(),
	}
	t, _ := json.Marshal(&returnData)
	w.Write(t)
}

//  数据库测试通过,返回数据格式未对齐 jdchain未测试
func (app *Application) QueryMortmark(w http.ResponseWriter, r *http.Request) {
	app.QueryAllTrademark(w, r)
}

//  数据库测试通过,返回数据格式未对齐 jdchain未测试
func (app *Application) QueryAucmark(w http.ResponseWriter, r *http.Request) {
	app.QueryAllTrademark(w, r)
}

//  逻辑测试通过 jdchain未测试
func (app *Application) AgreeApply(w http.ResponseWriter, r *http.Request) {
	app.ForPreCheck(w, r)
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	agreeApply := model.AgreeApply{
		MortApplyKey: r.PostFormValue("mortapplykey"),
		MarkName:     r.PostFormValue("markname"),
		MarkPic:      r.PostFormValue("markpic"),
		Applyer:      r.PostFormValue("applyer"),
		MarkValue:    r.PostFormValue("markvalue"),
		LoanValue:    r.PostFormValue("loanvalue"),
		PaybackDate:  r.PostFormValue("paybackdate"),
	}

	app.TokenVerify2Generate(w, r)
	if err = app.TrademarkService.AgreeApply(&agreeApply); err != nil {
		fmt.Println("AgreeApply Function Modify Data In JDChain Error")
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

//  逻辑测试通过 jdchain未测试
func (app *Application) RejectApply(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	applyerKey := r.PostFormValue("applyerkey")
	if err = app.TrademarkService.RejectApply(applyerKey); err != nil {
		fmt.Println("RejectApply Function Modify Data In JDChain Error")
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

//  逻辑测试通过 jdchain未测试
func (app *Application) UpdateMortgage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		t, _ := json.Marshal(&CommonReturnType{Status: "fail", Data: nil})
		w.Write(t)
		return
	}
	applyerKey := r.PostFormValue("servertime")
	if err = app.TrademarkService.UpdateMortgage(applyerKey); err != nil {
		fmt.Println("UpdateMortgage Function Modify Data In JDChain Error")
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

//  逻辑测试通过 jdchain未测试
func (app *Application) UpdateAuctioning(w http.ResponseWriter, r *http.Request) {
	if err := app.TrademarkService.UpdateAuctioning(); err != nil {
		fmt.Println("UpdateAuctioning Function Modify Data In JDChain Error")
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

//  逻辑测试通过 jdchain未测试
func (app *Application) Bid(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		t, _ := json.Marshal(&CommonReturnType{Status: "fail", Data: nil})
		w.Write(t)
		return
	}
	bid := &model.Bid{
		AuctioningKey: r.PostFormValue("auctioningkey"),
		Bidder:        r.PostFormValue("bidder"),
		MoneyNow:      r.PostFormValue("moneynow"),
	}

	if err := app.TrademarkService.Bid(bid); err != nil {
		fmt.Println("Bid Function Modify Data In JDChain Error")
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
