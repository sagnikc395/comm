package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/sagnikc395/comm/cmd/api"
	"github.com/sagnikc395/comm/config"
	"github.com/sagnikc395/comm/db"
)

func main() {
	db, err := db.NewMYSQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatalf("Failed to intialize a new DB %v", err)
	}

	server := api.NewAPIServer(":8080", db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
