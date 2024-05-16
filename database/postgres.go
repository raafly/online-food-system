package database

import (
	"fmt"
	"log"
	"time"

	"github.com/raafly/online-food-service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(conf *config.AppConfig) *gorm.DB {
	host := conf.Postgres.Host
	port := conf.Postgres.Port
	user := conf.Postgres.User
	pass := conf.Postgres.Pass
	name := conf.Postgres.Name

	url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, user, pass, name)
	sqlDb, err := gorm.Open(postgres.Open(url))
	if err != nil {
		log.Fatal(err)
	}

	db, err := sqlDb.DB()
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100) 
	db.SetConnMaxIdleTime(5 *time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return sqlDb
}