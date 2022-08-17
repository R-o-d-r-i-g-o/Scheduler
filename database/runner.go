package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	db            *gorm.DB
	configuration string
}

func NewMysql(config string) *Mysql {
	return &Mysql{
		db:            &gorm.DB{},
		configuration: config,
	}
}

func (m *Mysql) connect() error {
	var err error
	m.db, err = gorm.Open(mysql.Open(m.configuration), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database")
	}

	return nil
}
