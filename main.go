package main

import (
	"fmt"
	"log"
	"os"
	db "reku-technical-test/database"
	"reku-technical-test/src/web/server"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	fmt.Println(os.Getenv("MYSQL_CONNECTION"))
	dbMysql := db.InitMysql(os.Getenv("MYSQL_CONNECTION"))
	dbRedis := db.InitRedis()

	server.Run(dbMysql, dbRedis)
}
