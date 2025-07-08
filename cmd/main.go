package main

import (
	"log"

	"database/sql"

	"github.com/absakran01/ecom/cmd/api"
	"github.com/absakran01/ecom/config"
	"github.com/absakran01/ecom/db"
	"github.com/go-sql-driver/mysql"
)

func main(){
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                 "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}
	initStorage(db)

	
	server := api.NewAPIServer(":8080", nil)


	if err := server.Run(); err != nil {
		log.Fatal(err)
	}


}


func initStorage(db *sql.DB){
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}else {
		log.Println("Database connection established successfully")
	}
}