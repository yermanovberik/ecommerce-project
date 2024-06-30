package cmd

import (
	"database/sql"
	"ecommerce-project/cmd/api"
	"ecommerce-project/config"
	db2 "ecommerce-project/db"
	"github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := db2.NewMYSQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPass,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)
	server := api.NewApiServer(":8080", db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	log.Fatal("Succesfully connected to database")

}
