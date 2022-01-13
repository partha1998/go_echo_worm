package main

import (
	"go_echo_gorm/db"
	"go_echo_gorm/routers"
)

func main() {
	db.ConnectDb()
	routers.Start()
}
