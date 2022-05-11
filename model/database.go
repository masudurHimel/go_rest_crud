package model

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Database() (*gorm.DB, error) {
	dsn := "host=0.0.0.0 user=masud password=12345 dbname=test_pg_go port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	if err = db.AutoMigrate(&Grocery{}); err != nil {
		log.Println(err)
	}

	return db, err
}
