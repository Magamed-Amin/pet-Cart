package database

import (
	"database/sql"
	"fmt"
	"net"
	"net/url"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	if host == "" || port == "" || user == "" ||
		password == "" || dbname == "" || sslmode == "" {
		return nil, fmt.Errorf("не все настройки базы данных заданы")
	}

	u := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(user, password),
		Host:   net.JoinHostPort(host, port),
		Path:   "/" + dbname,
	}

	query := u.Query()
	query.Set("sslmode", sslmode)
	u.RawQuery = query.Encode()

	db, err := sql.Open("pgx", u.String())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
