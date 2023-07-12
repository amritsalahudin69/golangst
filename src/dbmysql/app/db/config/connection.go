package config

import (
	"database/sql"
	"fmt"
	"net/url"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func MysqlDBInit() *sql.DB {
	var (
		errMysql error
		dbConn   *sql.DB
	)
	dbHost := "localhost"
	dbPort := "3306"
	dbUser := "root"
	dbPass := ""
	dbName := "salt_school"

	//dbHost := os.Getenv("MYSQL_HOST")
	//dbPort := os.Getenv("MYSQL_PORT")
	//dbUser := os.Getenv("MYSQL_NAME")
	//dbPass := os.Getenv("MYSQL_PASS")
	//dbName := os.Getenv("MYSQL_DB_NAME")

	//fmt.Println("dbhost : ", dbHost)
	//fmt.Println("dbPort : ", dbPort)
	//fmt.Println("dbUser : ", dbUser)
	//fmt.Println("dbPass : ", dbPass)
	//fmt.Println("dbName : ", dbName)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	dbConn, errMysql = sql.Open(`mysql`, dsn)

	if errMysql != nil {
		return nil
	}
	dbConn.SetMaxOpenConns(300)
	dbConn.SetMaxIdleConns(25)
	dbConn.SetConnMaxLifetime(5 * time.Minute)

	return dbConn
}
