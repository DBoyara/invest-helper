package repository

import (
	"context"
	"fmt"

	"github.com/DBoyara/invest-helper/common"
	"github.com/jackc/pgx/v4/pgxpool"
)

var connPool *pgxpool.Pool

func SetupDB(settings *common.Settings) (err error) {
	urn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		settings.DbUser,
		settings.DbPassword,
		settings.DbHost,
		settings.DbPort,
		settings.DbName,
	)
	connPool, err = pgxpool.Connect(context.Background(), urn)
	return
}

func CloseDB() {
	connPool.Close()
}
