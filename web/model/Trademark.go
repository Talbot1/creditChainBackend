package model

type TrademarkVO struct {
	MarkKey      string `json:"markKey"`
	TranHash     string `json:"tranHash"`
	MarkName     string `json:"markname"`     //商标名称
	MarkPic      string `json:"markpic"`      //商标图片  url
	MarkCate     string `json:"markcate"`     //商标分类
	MarkRegister string `json:"markregister"` //商标注册号
	ApplyDate    string `json:"applydate"`    //商标申请时间  yyyy-MM-dd hh:mm:ss
	ValidityTime string `json:"validityTime"` //商标有效期
	Applyer      string `json:"applyer"`      //抵押申请人
	MarkValue    string `json:"markValue"`    //商标估值   单位  万元
	LoanValue    string `json:"loanValue"`    //贷款金额  单位 万元
	PaybackDate  string `json:"paybackDate"`  //还款时间  yyyy-MM-dd hh:mm:ss
	MoneyNow     string `json:"moneyNow"`     //当前最高价  单位 万元
	Bidder       string `json:"bidder"`       //出价者
	RestTime     string `json:"restTime"`     //剩余时间   单位 分钟
}

type AgreeApply struct {
	MortApplyKey string `json:"mortApplyKey"`
	MarkName     string `json:"markname"`    //商标名称
	MarkPic      string `json:"markpic"`     //商标图片  url
	Applyer      string `json:"applyer"`     //抵押申请人
	MarkValue    string `json:"markValue"`   //商标估值   单位  万元
	LoanValue    string `json:"loanValue"`   //贷款金额  单位 万元
	PaybackDate  string `json:"paybackDate"` //还款时间  yyyy-MM-dd hh:mm:ss
}

type Bid struct {
	AuctioningKey string `json:"auctioningKey"`
	Bidder        string `json:"bidder"`
	MoneyNow      string `json:"moneyNow"`
}
