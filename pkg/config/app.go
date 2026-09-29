package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("postgres", "postgresql://themoeaegon:moee5055@localhost:5432/go-bookstore?sslmode=disable") // i dont want to use env this is just getting familar with go
	if err != nil {
		panic(err)
	}
	if err := d.DB().Ping(); err != nil {
		panic(err)
	}
	db = d
}

func GetDb() *gorm.DB {
	return db
}
