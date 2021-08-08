package dao

import (
	"blc-demo/web/model"
	"fmt"
)

func CreateTableWithTransaction() {
	sqlStr := `CREATE TABLE IF NOT EXISTS transaction (
		id BIGINT AUTO_INCREMENT primary key not null,
		name varchar(1024),
		 url varchar(1024)  ,
	 type varchar(1024) );
			`
	Exec(sqlStr)
	fmt.Println("---------------------------------------------")
	fmt.Println("Transaction table created")
}

func InsertTransaction(Transcation *model.TransactionVO) error {
	_, err := Exec("insert into transaction(name,url,type) values (?,?,?);",
		Transcation.Name, Transcation.Url, Transcation.Type)
	return err
}

func QueryAllTransaction() ([]*model.TransactionVO, error) {
	rows := QueryDB(`select id, name, url, type from transaction`)
	var transcationSlice []*model.TransactionVO
	for rows.Next() {
		ts := &model.TransactionVO{}
		err := rows.Scan(&ts.Id, &ts.Name, &ts.Url, &ts.Type)
		if err != nil {
			fmt.Println("读取Transaction全部数据失败", err)
			return nil, err
		}
		transcationSlice = append(transcationSlice, ts)
	}
	return transcationSlice, nil
}

func CreateTableWithCompany() {
	sqlStr := `CREATE TABLE IF NOT EXISTS company (
		companyName varchar(255) primary key not null,
		legalRepresentative varchar(1024),
		 creditCode varchar(1024)  ,
		 registeredCapital varchar(1024)  ,
		 paidCapital varchar(1024)  ,
		 establishedData varchar(1024)  ,
		 approvedDate varchar(1024)  ,
		 ownRisk varchar(1024)  ,
		 associatedRisk varchar(1024)  ,
		 staffSize varchar(1024)  ,
		 businessScope varchar(1024)  ,
		 credit varchar(1024) );
			`
	Exec(sqlStr)
	fmt.Println("---------------------------------------------")
	fmt.Println("Company table created")
}

func InsertCompany(Company *model.CompanyVO) error {
	_, err := Exec("insert into company(companyName,legalRepresentative,creditCode,registeredCapital,paidCapital) values (?,?,?,?,?);",
		Company.CompanyName, Company.LegalRepresentative, Company.CreditCode, Company.RegisteredCapital, Company.PaidCapital)
	return err
}

func QueryAllCompany() ([]*model.CompanyVO, error) {
	rows := QueryDB(`select companyName,legalRepresentative,creditCode,registeredCapital,paidCapital from company`)
	var companySlice []*model.CompanyVO
	for rows.Next() {
		com := &model.CompanyVO{}
		err := rows.Scan(&com.CompanyName, &com.LegalRepresentative, &com.CreditCode, &com.RegisteredCapital, &com.PaidCapital)
		if err != nil {
			fmt.Println("读取CompanyVO全部数据失败", err)
			return nil, err
		}
		companySlice = append(companySlice, com)
	}
	return companySlice, nil
}

func CreateTableWithLoan() {
	sqlStr := `CREATE TABLE IF NOT EXISTS loan (
		id BIGINT primary key not null,
	 value varchar(1024) );
			`
	Exec(sqlStr)
	fmt.Println("---------------------------------------------")
	fmt.Println("loan table created")
}

func InsertLoan(loan *model.Loan) error {
	_, err := Exec("insert into loan(id, value) values (?,?);",
		loan.Id, loan.Value)
	return err
}
