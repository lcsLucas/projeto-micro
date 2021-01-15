package mongodb

import (
	"testing"

	"github.com/lcslucas/projeto-micro/config"
)

func TestConnectDB(t *testing.T) {

	c := config.ConfigDB{
		Driver:   "mongodb",
		Host:     "0.0.0.0",
		User:     "root",
		Password: "root",
		Port:     "27017",
		DBName:   "micro",
	}

	_, err := Connect(nil, c)
	if err != nil {
		t.Fatalf(err.Error())
	}

}
