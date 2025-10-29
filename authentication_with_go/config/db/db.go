package config

import (
	env "AuthInGo/config/env"
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func SetupDb() (*sql.DB, error) {
	cfg  := mysql.NewConfig()

	cfg.User = env.GetString("DB_USER", "root")
	cfg.Passwd = env.GetString("DB_PASSWORD", "no-pass")
	cfg.Net = env.GetString("DB_NET", "tcp")
	cfg.Addr = env.GetString("DB_ADDR", "localhost:3306")
	cfg.DBName = env.GetString("DBName", "airbnb")

	fmt.Println("Connecting to database: ", cfg.DBName, cfg.FormatDSN()+"?parseTime=true")

	db, error := sql.Open("mysql", cfg.FormatDSN()+"?parseTime=true")

	if error!=nil {
		fmt.Println("Error connecting to db: ", error)
		return nil, error
	}

	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println("Error pinging db: ", pingErr)
		return nil, pingErr
	}

	fmt.Println("Connected to database successfully!")
	return db, nil
}