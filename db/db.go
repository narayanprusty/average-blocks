package db

import (
	"log"

	pg "github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/narayanprusty/average-blocks/config"
)

var DB *pg.DB

func init() {
	db := pg.Connect(&pg.Options{
		User:     config.Config.DatabaseUsername,
		Password: config.Config.DatabasePassword,
		Database: config.Config.DatabaseName,
		Addr:     config.Config.DatabaseHost + ":" + config.Config.DatabasePort,
	})

	models := []interface{}{
		(*User)(nil),
		(*APIKey)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			log.Fatal(err)
		}
	}

	DB = db
}
