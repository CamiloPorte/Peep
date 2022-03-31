package mysql

import (
	"database/sql"
	"fmt"
	"os"
)

func NewConn() (*sql.DB, error) {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("PASSWORD")
	name := os.Getenv("DBNAME")

	conn := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, pass, name)
	return sql.Open("mysql", conn)
}
