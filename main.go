package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sahith-narahari/otis/alert"
	"github.com/sahith-narahari/otis/config"
	"log"
)

func init() {
	if err := godotenv.Load();err != nil {
		fmt.Println("env file not found")
	}
	config.SetConfig()
}
func main() {
	err := alert.SendAlert()
	if err != nil{
		log.Panicf("Couldn't start alerting bot")
	}
}