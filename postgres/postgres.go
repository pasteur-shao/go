package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetConnectionObj(host string, port int, user string, keycode string, dbname string) string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, keycode, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Connect to postgres database successfully !!")

	return db
}
