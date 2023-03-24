package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
	"os"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

// Entry defines both the CSV layout and database schema
type Entry struct {
	gorm.Model
	Field1 float64 `csv:"field1"`
	Field2 float64 `csv:"field2"`
}

type User struct {
	gorm.Model
	Name     string
	Password string
}

func main() {
	file, err := os.Open("emp.csv")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	record, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error", err)
	}

	var indexArray = make([]string, len(record[0]))
	for _ = range record {
		var arr = record[0]
		for name := range arr {
			fmt.Println(arr[name])
			indexArray = append(indexArray, arr[name])

		}
		break
	}
	fmt.Println(indexArray)
	var sliceRecord = append(record[1:])

	for value := range sliceRecord { // for i:=0; i<len(record)

		fmt.Println("**", record[value])
	}

	connectDB(sliceRecord)

}

func connectDB(data [][]string) {
	postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", postgresqlDbInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Established a successful connection!")

	sqlStr := "INSERT INTO postgres(name,age) VALUES"
	vals := []interface{}{}

	for _, row := range data {
		t1 := row[0]
		t2 := row[1]
		sqlStr += "('1', '2'),"

		fmt.Println(t1, "cgjh", t2)

		//vals = append(vals, row[0], row[1])
	}
	//trim the last ,
	sqlStr = sqlStr[0 : len(sqlStr)-1]
	//prepare the statement
	fmt.Println(sqlStr)
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		panic(err)
	}

	//format all vals at once
	fmt.Println(vals)
	res, err := stmt.Exec(vals...)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

}
