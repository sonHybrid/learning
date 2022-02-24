package main

import (
	"todovue/server"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	server.Router()
}
