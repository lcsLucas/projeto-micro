package postgresql

import (
	"testing"

	"github.com/lcslucas/projeto-micro/config"
)

func TestDB(t *testing.T) {

	c := config.ConfigDB{
		Driver:   "postgres",
		Host:     "0.0.0.0",
		User:     "postgres",
		Password: "postgres123",
		Port:     "5432",
		DBName:   "micro",
	}

	_, err := Connect(c, false)
	if err != nil {
		t.Fatalf(err.Error())
	}

}
