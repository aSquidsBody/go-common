package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectionString(user, pass, host, port, db string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, pass, host, port, db)
}

type PgClient struct {
	*gorm.DB
}

func NewPgClient(user, pass, host, port, db string) (*PgClient, error) {
	connStr := connectionString(user, pass, host, port, db)
	client, err := gorm.Open(postgres.Open(connStr))
	if err != nil {
		return nil, err
	}
	return &PgClient{client}, nil
}
