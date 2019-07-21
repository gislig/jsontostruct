package middleware

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Config - Some comment
type Config struct {
	Host      string `json:"host"`
	Port      int    `json:"port"`
	User      string `json:"user"`
	Password  string `json:"password"`
	DBName    string `json:"dbname"`
	WebServer string `json:"webserver"`
}

func GetConfig() Config {
	var conf Config
	jsonFile, err := os.Open("config.json")

	if err != nil {
		fmt.Println("Could not open file", err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Could not read file", err)
	}

	json.Unmarshal(byteValue, &conf)

	if err := json.Unmarshal(byteValue, &conf); err != nil {
		fmt.Println("Could not unmarshal file", err)
	}

	return conf
}
