package main

import (
	"encoding/csv"
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
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

	for value := range record { // for i:=0; i<len(record)
		fmt.Println("**", record[value])
	}

}
