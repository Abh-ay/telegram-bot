package dbConnection

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDB() *sqlx.DB {
	connStr := "postgres://" + os.Getenv("POSTGRES_USERNAME") + ":" + os.Getenv("POSTGRES_PASSWORD") + "@localhost/TeleDB?sslmode=disable"
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		fmt.Println("error connecting to DB", err)
	} else {
		fmt.Println("The database is connected")

	}
	return db
}
