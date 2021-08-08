package model

type CompanyVO struct {
	CompanyName         string `json:"companyName"`
	LegalRepresentative string `json:"legalRepresentative"`
	CreditCode          string `json:"creditCode"`
	RegisteredCapital   string `json:"registeredCapital"`
	PaidCapital         string `json:"paidCapital"`
	EstablishedData     string `json:"establishedData"`
	ApprovedDate        string `json:"approvedDate"`
	OwnRisk             string `json:"ownRisk"`
	AssociatedRisk      string `json:"associatedRisk"`
	StaffSize           string `json:"staffSize"`
	BusinessScope       string `json:"businessScope"`
	Credit              string `json:"credit"`
	Amount              string `json:"amount"`
}

type GuaranteeVO struct {
	CompanyName string `json:"companyName"`
	//信用分
	Credit string `json:"credit"`
	//固定资产
	FixedAssets string `json:"fixedAssets"`
	//利润率
	ProfitRate string `json:"profitRate"`
	//担保金额
	GuaranteedAmount string `json:"guaranteedAmount"`
}

type TransactionVO struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
	// 记录是什么类型的交易：0：京东链，1：天德链，2：以太坊
	Type string `json:"type"`
}

type Loan struct {
	Id    int `json:"id"`
	Value int `json:"value"`
}
