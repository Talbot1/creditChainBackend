package dao

import (
	"blc-demo/web/model"
	"fmt"
)

func CreateTableWithTrademark() {
	sqlStr := `CREATE TABLE IF NOT EXISTS Trademark (
		MarkKey varchar(1024),
	 TranHash varchar(1024),
	 MarkName  varchar(255) PRIMARY KEY NOT NULL,
	 MarkPic  varchar(1024),
	 MarkCate varchar(1024) ,  
	 MarkRegister varchar(1024),  
	 ApplyDate varchar(1024)   ,
	 ValidityTime  varchar(1024), 
	 Applyer  varchar(1024)  ,
	 MarkValue varchar(1024) ,
	 LoanValue varchar(1024) ,
	 PaybackDate varchar(1024),  
	 MoneyNow varchar(1024) ,
	 Bidder varchar(1024)  ,
	 RestTime varchar(1024) 
			);`
	Exec(sqlStr)
	fmt.Println("---------------------------------------------")
	fmt.Println("Trademark table created")
}

func InsertApply(Trademark *model.TrademarkVO) error {
	_, err := Exec("insert into Trademark(MarkName,MarkPic,MarkCate,MarkRegister,ApplyDate,Applyer) values (?,?,?,?,?,?)",
		Trademark.MarkName, Trademark.MarkPic, Trademark.MarkCate, Trademark.MarkRegister, Trademark.ApplyDate, Trademark.Applyer)
	return err
}

func QueryALLTrademark() ([]*model.TrademarkVO, error) {
	rows := QueryDB(`select MarkName, MarkPic, MarkCate, MarkRegister, ApplyDate, Applyer from TradeMark`)
	var trademarkSLice []*model.TrademarkVO
	for rows.Next() {
		td := &model.TrademarkVO{}
		//err := rows.Scan(&td.MarkKey, &td.TranHash, &td.MarkName, &td.MarkPic, &td.MarkCate, &td.MarkRegister, &td.ApplyDate, &td.Applyer, &td.MarkValue, &td.LoanValue, &td.PaybackDate, &td.MoneyNow, &td.Bidder, &td.RestTime)
		err := rows.Scan(&td.MarkName, &td.MarkPic, &td.MarkCate, &td.MarkRegister, &td.ApplyDate, &td.Applyer)
		if err != nil {
			fmt.Println("读取TradeMark全部数据失败", err)
			return nil, err
		}
		trademarkSLice = append(trademarkSLice, td)
	}
	return trademarkSLice, nil
}
