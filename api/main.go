package main

import (
	"api/config"
	"api/config/storage"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Messages struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func main() {
	fmt.Println("Start App v0.01")
	err := godotenv.Load()

	if err != nil {
		fmt.Println("env: ", err.Error())
	}

	storage.InitDB()
	router := config.InitRoute()

	defer func() {
		db := storage.GetDB()
		db.Close()
	}()

	fmt.Println("init")

	err = router.Run(":" + os.Getenv("PORT"))
	if err != nil {
		return
	}
}
