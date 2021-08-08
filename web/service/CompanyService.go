package service

import (
	"blc-demo/web/model"
	"blc-demo/web/utils"
	"encoding/json"
	"errors"
	"fmt"
)

type CompanyService struct {
	JdService *utils.JdService
}

var errJD error = errors.New("JdChain service unconnected")

func (c *CompanyService) QueyALl() ([]byte, error) {
	url := utils.ChainUrl + "/chaincode/query"
	var bearer = "Bearer " + utils.Token
	fmt.Println(bearer)
	para := utils.JdParameters{
		Organization: "user03",
		Channel:      "testone",
		Ccname:       "fabcarnew",
		Function:     "queryAllCompanys",
	}
	resp, _ := c.JdService.QueryInfo(url, bearer, para)
	msg, _ := json.Marshal(resp)
	return msg, nil
}

func (c *CompanyService) QueryAllTransaction() ([]*model.TransactionVO, error) {
	var TSlice []*model.TransactionVO
	return TSlice, errJD
}

func (c *CompanyService) AddTransaction(*model.TransactionVO) error {
	return errJD

}
func (c *CompanyService) QueryAllCompany() ([]*model.CompanyVO, error) {
	var comSlice []*model.CompanyVO
	return comSlice, errJD
}

func (c *CompanyService) AddCompany(*model.CompanyVO) error {
	return errJD

}
func (c *CompanyService) ChangeCredit(*model.Loan) error {
	// 与京东链进行操作, 传递参数可以是bytes类型.
	return errJD
}

func (c *CompanyService) AddLoan(*model.Loan) error {
	// 与京东链进行操作, 传递参数可以是bytes类型.
	return errJD
}

func (c *CompanyService) QueryGuarantee() (*model.GuaranteeVO, error) {
	var guaranteeVo *model.GuaranteeVO
	return guaranteeVo, errJD
}
