package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"strings"
)

//数据库配置
const (
	userName = "root"
	password = "123456"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "posts_article"
)

var DB *sql.DB

func InitDB() {
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	DB, _ = sql.Open("mysql", path)
	//最大连接数
	DB.SetConnMaxLifetime(100)
	//最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("opon database fail")
	}
	fmt.Println("connnect success")
}

func GetItem(table string, oid int, column_name string, column string) (map[string]interface{}, error) {
	var sql string = "SELECT " + column_name + " FROM " + table + " WHERE " + column + " = " + strconv.Itoa(oid)
	rows, _ := DB.Query(sql)
	columns, _ := rows.Columns()
	count := len(columns)
	tableData := make(map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = entry
	}
	return tableData, nil
}

func GetItems(table string, column_name string, condition string) ([]map[string]interface{}, error) {
	var sql string
	if condition != "" {
		sql = "SELECT " + column_name + " FROM " + table + " WHERE " + condition
	} else {
		sql = "SELECT " + column_name + " FROM " + table
	}
	rows, _ := DB.Query(sql)
	columns, _ := rows.Columns()
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	return tableData, nil
}
func ModifyItem(table string, oid string, condition map[string]string, column string) int64 {
	var condition_str string
	for k, v := range condition {
		condition_str += k + "='" + v + "',"
	}
	var sql string = "UPDATE " + table + " SET " + condition_str + " WHERE " + column + " = " + oid
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
	}
	//准备sql语句
	stmt, err := tx.Prepare(sql)
	res, err := stmt.Exec()
	num, err := res.RowsAffected()
	tx.Commit()
	return num
}

func CreateItem(table string, condition map[string]string) int64 {
	var cloumn_list []string
	var cloumn_str string
	var values_list []string
	var values_str string
	for k, v := range condition {
		cloumn_list = append(cloumn_list, []string{k}...)
		cloumn_str = strings.Join(cloumn_list, ",")
		values_list = append(values_list, []string{v}...)
		values_str = strings.Join(values_list, "','")
	}
	var sql string = "INSERT INTO " + table + "(" + cloumn_str + ") VALUES(" + "'" + values_str + "'" + ")"
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
	}
	stmt, err := tx.Prepare(sql)
	res, err := stmt.Exec()
	num, err := res.RowsAffected()
	tx.Commit()
	return num
}

func DeleteItem(table string, oid string, column string) int64 {
	var sql string = "DELETE FROM " + table + " WHERE " + column + " = " + oid
	fmt.Println(sql)
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
	}
	stmt, err := tx.Prepare(sql)
	res, err := stmt.Exec()
	num, err := res.RowsAffected()
	tx.Commit()
	return num
}

func Query(sql string) ([]map[string]interface{}, error) {
	rows, _ := DB.Query(sql)
	columns, _ := rows.Columns()
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	return tableData, nil
}
