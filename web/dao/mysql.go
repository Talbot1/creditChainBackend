package dao

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

const (
	userName = "root"
	password = "12345678"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "credit"
)

var db *sql.DB

func InitMysql() {
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8&multiStatements=true"}, "")

	if db == nil {
		fmt.Println("---------------------------------------------")
		fmt.Println("Database Connected")
		db, _ = sql.Open("mysql", path)
		DeleteTable()
		CreateTableWithUser()

		CreateTableWithTrademark()
		CreateTableWithTransaction()
		CreateTableWithCompany()
		CreateTableWithLoan()

		CreateUser0InUser()
		CreateUser1InUser()
		CreateStaffInUser()

	}
}

//查询
func QueryRowDB(sqlStr string) *sql.Row {
	return db.QueryRow(sqlStr)
}

func QueryDB(sqlStr string) *sql.Rows {
	rows, _ := db.Query(sqlStr)
	return rows
}

//操作数据库
func Exec(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

func DeleteTable() {
	sqlStr := `SET FOREIGN_KEY_CHECKS = 0;
			DROP TABLE IF EXISTS user;
			SET FOREIGN_KEY_CHECKS = 1;
			DROP TABLE IF EXISTS session;
			DROP TABLE IF EXISTS action;
			DROP TABLE IF EXISTS tea;
			DROP TABLE IF EXISTS trademark;
			DROP TABLE IF EXISTS company;
			DROP TABLE IF EXISTS transaction;
			DROP TABLE IF EXISTS loan`
	fmt.Println("---------------------------------------------")
	fmt.Println("table deleted")
	_, _ = Exec(sqlStr)
}
