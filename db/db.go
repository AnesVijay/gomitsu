package db

import (
	"fmt"
	"os"

	"github.com/jackc/pgx"
)

func main() {
	// db, err := sql.Open(sql.Drivers()["postgresql"])

	connConf := pgx.ConnConfig{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     5432,
	}

	conn, err := pgx.Connect(connConf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

}
