package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/sagnikc395/comm/cmd/api"
	"github.com/sagnikc395/comm/db"
)

func main() {
	db, err := db.NewMYSQLStorage(mysql.Config{
		User:                 "root",
		Passwd:               "rootpwd",
		Addr:                 "127.0.0.1:8765",
		DBName:               "commDB",
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	server := api.NewAPIServer(":8080", nil)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
