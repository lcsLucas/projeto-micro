package migrations

import (
	"context"
	"testing"

	"github.com/lcslucas/projeto-micro/config"

	database "github.com/lcslucas/projeto-micro/database/postgresql"
)

func TestExecMigrationProva(t *testing.T) {
	ctx := context.Background()

	c := config.ConfigDB{
		Driver:   "postgres",
		Host:     "0.0.0.0",
		User:     "postgres",
		Password: "postgres123",
		Port:     "5432",
		DBName:   "micro",
	}

	conn, err := database.Connect(c, false)
	if err != nil {
		t.Fatalf(err.Error())
	}

	err = ExecMigrationProva(ctx, c.DBName, conn)

	if err != nil {
		t.Fatalf(err.Error())
	}

}
