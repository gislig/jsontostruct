package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"

	"github.com/gislig/sqltest/middleware"
	"github.com/gislig/sqltest/models/device"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "support_writer"
	password = "helloMyName.123"
	dbname   = "support"
)

// User - Some comment
type User struct {
	UserID uint `json:"user_id"`
	Name   string
}

// Address - Some comment
type Address struct {
	AddressID uint
	Home      string
	UserID    uint
}

func APIReader(api interface{}) {
	APIValue := reflect.ValueOf(api)
	APIType := reflect.TypeOf(api)
	//APIName := APIType.Name()

	for i := 0; i < APIValue.NumField(); i++ {
		fieldName := APIValue.Type().Field(i).Name
		//fieldValue := APIValue.Field(i)
		fieldTag, _ := APIType.FieldByName(fieldName)
		jsonField := fieldTag.Tag.Get("json")
		fmt.Println(fieldName, jsonField)
	}
}

func InsertIntoTable(q interface{}) {
	v := reflect.ValueOf(q)
	T := reflect.TypeOf(q)
	t := T.Name()

	query := fmt.Sprintf("INSERT INTO %s (", t)
	//k := t.Kind()
	//fmt.Println("Type", t)
	//fmt.Println("Kind", k)
	//fmt.Println("Fields", v.NumField())
	//tag = string(field.Tag)
	columns := ""
	values := ""
	for i := 0; i < v.NumField(); i++ {
		////fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
		//PrimaryKey := fmt.Sprintf("%sID", t)
		//if v.Type().Field(i).Name == PrimaryKey {
		//	PrimaryKey = fmt.Sprintf("%s %s PRIMARY KEY NOT NULL,", PrimaryKey, v.Field(i).Kind())
		//	query = fmt.Sprintf("%s %s\n", query, PrimaryKey)
		//}

		fieldName := v.Type().Field(i).Name
		fieldValue := v.Field(i)
		fieldTag, _ := T.FieldByName(fieldName)
		fmt.Println(fieldTag.Tag.Get("json"))
		//fmt.Println(fieldTag.Tag.Get("custom"))

		if i == 0 {
			//query = fmt.Sprintf("%s", query, v.Field(i).Uint())
			columns = fmt.Sprintf("%v%v", columns, fieldName)
			values = fmt.Sprintf("'%v'", fieldValue)
		} else {
			columns = fmt.Sprintf("%v,%v", columns, fieldName)
			values = fmt.Sprintf("%v,'%v'", values, fieldValue)
		}

		/*
			switch v.Field(i).Kind() {
			case reflect.Uint:
				if i == 0 {
					//query = fmt.Sprintf("%s", query, v.Field(i).Uint())
					columns = fmt.Sprintf("%v%v", columns, fieldName)
					values = fmt.Sprintf("'%v'", fieldValue)
				} else {
					columns = fmt.Sprintf("%v,%v", columns, fieldName)
					values = fmt.Sprintf("%v,'%v'", values, fieldValue)
				}
			case reflect.Int:
				if i == 0 {

					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("Unsupported type")
				return
			}*/
		//fmt.Printf("Name: %v\n", v.Type().Field(i).Name)
	}
	//fmt.Println(columns)
	query = fmt.Sprintf("%v%v) VALUES (%v)", query, columns, values)
	fmt.Println(query)
}

func main() {
	conf := middleware.GetConfig()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Password, conf.DBName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	dat, _ := ioutil.ReadFile("sqlvars.sql")
	fileContent := string(dat)

	fileContent = strings.Replace(fileContent, "%tablename%", "tellme", 2)
	fmt.Println(fileContent)

	userItem1 := User{
		UserID: 1,
		Name:   "bob",
	}

	InsertIntoTable(userItem1)
	//CreateQuery(userItem2)

	defer db.Close()

	bios := device.Bios{}
	APIReader(bios)

	//tableName := reflect.TypeOf(User)

	fmt.Println("Successfully connected!")
}
