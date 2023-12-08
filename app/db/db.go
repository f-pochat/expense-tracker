package db

import (
	"database/sql"
	"fmt"
	"github.com/revel/revel"
)

var DB *sql.DB

func InitDB() {
	connstring := fmt.Sprintf("user=%s password='%s' dbname=%s sslmode=disable", "postgres", "password", "expense")

	var err error
	DB, err = sql.Open("postgres", connstring)
	if err != nil {
		revel.AppLog.Info("DB Error", err)
	}
	revel.AppLog.Info("DB Connected")
}
