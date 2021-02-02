package migrations

import (
	"context"
	"testing"

	"github.com/lcslucas/projeto-micro/config"

	database "github.com/lcslucas/projeto-micro/database/mongodb"
)

func TestExecMigrationAlunos(t *testing.T) {
	ctx := context.Background()

	c := config.ConfigDB{
		Driver:   "mongodb",
		Host:     "0.0.0.0",
		User:     "root",
		Password: "root",
		Port:     "27017",
		DBName:   "micro",
	}

	conn, err := database.Connect(nil, c)
	if err != nil {
		t.Fatalf(err.Error())
	}

	err = ExecMigrationAlunos(ctx, c.DBName, conn)

	if err != nil {
		t.Fatalf(err.Error())
	}

}
