package postgresql

import (
	"fmt"

	"github.com/lcslucas/projeto-micro/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBURL string

func Connect(c config.ConfigDB, flagDB bool) (*gorm.DB, error) {
	DBURL = fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo", c.Host, c.User, c.Password, c.Port)

	if flagDB {
		DBURL = fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable TimeZone=America/Sao_Paulo", c.Host, c.User, c.Password, c.Port, c.DBName)
	}

	fmt.Println(DBURL)

	db, err := gorm.Open(postgres.Open(DBURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	fmt.Printf("Database connected: %s", DBURL)
	return db, nil
}
