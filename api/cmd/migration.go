package main

import (
	"api/config"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("env: ", err.Error())
	}

	config.InitMigrationDB()

	fmt.Println("миграции отработали")
}
