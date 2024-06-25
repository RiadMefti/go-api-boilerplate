package db

import (
	"database/sql"
	"fmt"

	"github.com/RiadMefti/go-api-boilerplate/types"
	_ "github.com/lib/pq"
)

type Storage interface {
	test()
}

type PostgresStore struct {
	db *sql.DB
}

func IninDb(config types.Config) (*PostgresStore, error) {

	postgresCOnnectionString := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", config.DbHost, config.DbPort, config.DbUser, config.DbPassword, config.DbName)

	db, err := sql.Open("postgres", postgresCOnnectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil

}

func (store PostgresStore) test() {

}
