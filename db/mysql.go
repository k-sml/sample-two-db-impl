package db

import (
	"github.com/k-sml/sample-two-db-impl/config/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

func BuildDBConnection(config env.Config) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(config.WriterDBURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	var replicas []gorm.Dialector
	replicas = append(replicas, mysql.Open(config.ReaderDBURL))

	db.Use(dbresolver.Register(dbresolver.Config{
		Replicas: replicas,
		Policy:   dbresolver.RandomPolicy{},
	}))

	return db, nil
}
