package dbase

import (
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "todo"
)

var (
	once sync.Once

	dbMaster *gorm.DB
)

func DB() *gorm.DB {
	once.Do(func() {
		dbMaster = newConnection()
	})
	return dbMaster
}

func newConnection() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}
	return db
}
