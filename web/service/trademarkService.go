package service

import (
	"blc-demo/web/model"
	"blc-demo/web/utils"
)

type TrademarkService struct {
	JdService *utils.JdService
}

func (td *TrademarkService) AddApply(vo *model.TrademarkVO) error {
	// 操作数据库
	return errJD
}

func (td *TrademarkService) QueryAllTrademark() ([]*model.TrademarkVO, error) {
	var TDlice []*model.TrademarkVO
	return TDlice, errJD
}

func (td *TrademarkService) AgreeApply(vo *model.AgreeApply) error {
	// 操作数据库
	return errJD
}

func (td *TrademarkService) RejectApply(ak string) error {
	// 操作数据库
	return errJD
}

func (td *TrademarkService) UpdateMortgage(ak string) error {
	// 操作数据库
	return errJD
}

func (td *TrademarkService) UpdateAuctioning() error {
	// 操作数据库
	return errJD
}

func (td *TrademarkService) Bid(bid *model.Bid) error {
	// 操作数据库
	return errJD
}
