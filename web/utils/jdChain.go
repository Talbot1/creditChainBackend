package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

type JdService struct {
}

const (
	ChainUrl = "http://211.151.11.130:31970"
	Username = "test"
	Password = "Admin@147@!"
)

type JdParameters struct {
	Organization string `json:"organization"`
	Channel      string `json:"channel"`
	Ccname       string `json:"ccname"`
	Function     string `json:"function"`
	Args         string `json:"args"`
}

type LoginResponse struct {
	Code   int    `json:"code"`
	Expire string `json:"expire"`
	Token  string `json:"token"`
}

type JdQueryResponse struct {
	Code int  `json:"code,omitempty"`
	Data Data `json:"data,omitempty"`
}

type Data struct {
	Message       string   `json:"message"`
	TransactionID string   `json:"transactionid"`
	Payload       string   `json:"payload,omitempty"`
	Record        []Record `json:"record"`
}

type Record struct {
	CompanyName         string `json:""`
	LegalRepresentative string //法人代表
	CreditCode          string //统一信用代码
	RegisteredCapital   string //注册资本
	PaidCapital         string
	EstablishedData     string //成立日期
	ApprovedDate        string
	OwnRisk             string
	AssociatedRisk      string
	StaffSize           string //员工规模
	BusinessScope       string //营业范围
	Credit              string //信用分
	Amount              string //贷款金额
	FixedAssets         string //固定资产金额
	ProfitRate          string //利润率
	GuaranteedAmount    string //担保金额
	BrandValue          string //商标价值
	BrandInfo           string //商标信息
	LoanLog             string //贷款日志
}

var Token string = "jdchain"

func (jd *JdService) QueryInfo(url, bearer string, para JdParameters) (*JdQueryResponse, error) {
	jsonByte, _ := json.Marshal(&para)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonByte))
	if err != nil {
		// handle error
		fmt.Println(err)
		return nil, nil
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// handle error
		fmt.Println(err)
		return nil, nil
	}
	code := resp.StatusCode
	if code != 200 {
		return nil, errors.New("string")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	var queryresp JdQueryResponse
	err = json.Unmarshal(body, &queryresp)
	if err != nil {
		// handle error
		fmt.Println(err)
		return nil, nil
	}
	payload := queryresp.Data.Payload
	stringSlice := strings.Split(payload, "},")
	//fmt.Println(stringSlice[0])
	for _, item := range stringSlice {
		ret := regexp.MustCompile(`"Record":{[\S\s]*}`) //参数是一个正则表达式 	``：表示使用原生字符串
		alls := ret.FindAllStringSubmatch(item, 1)
		formatStr := strings.Replace(alls[0][0], `"Record":`, "", 1)
		var record Record
		json.Unmarshal([]byte(formatStr), &record)
		queryresp.Data.Record = append(queryresp.Data.Record, record)
		//fmt.Println(queryresp.Data.Record[index].CompanyName,queryresp.Data.Record ,)
	}
	return &queryresp, nil
}

func (jd *JdService) SetToken() error {
	req, err := http.NewRequest("POST", ChainUrl+"/login", strings.NewReader(""))
	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(Username, Password)

	client := &http.Client{}
	res, err := client.Do(req)

	// res, err := http.Post(ChainUrl+"/login","application/json",bytes.NewBuffer([]byte(`{"Username":"test","Password":"Admin@147!"}`)))
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(res)
	//fmt.Println(string(body))
	loginRes := LoginResponse{}
	err = json.Unmarshal(body, &loginRes)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (jd *JdService) IsValid() error {
	//jd.SetToken()
	url := ChainUrl + "/chaincode/query"
	var bearer = "Bearer " + Token
	fmt.Println(bearer)
	para := JdParameters{
		"user03",
		"testone",
		"fabcarnew",
		"queryAllCompanys",
		"",
	}
	_, err := jd.QueryInfo(url, bearer, para)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func main() {

	var jd *JdService
	if err := jd.IsValid(); err != nil {
		fmt.Println(err)
	}
	//type Payload struct {
	//	Key    string `json: "Key"`
	//	Record string `json: "Record"`
	//}
	//fmt.Println(payload)
	//
	//slice := []string{"peer0.org1.example.com", "peer1.org1.example.com"}
	//str := strings.Replace(strings.Trim(fmt.Sprint(slice), "[]"), " ", ";", -1)
	//fmt.Println("slice:", slice)
	//fmt.Println("string:", str)

}
