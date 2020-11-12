package main

import (
	"fmt"
	"postgres/postgres.go"
)

const (
	host    = "192.168.63.160"
	port    = 5432
	user    = "postgres"
	keycode = "postgres"
	dbname  = "api"
)

func main() {
	fmt.Println("hiii")
	postgres.GetConnectionObj(host, port, user, keycode, dbname)
}
