package db

import (
	"database/sql"
	"fmt"

	"github.com/RiadMefti/go-api-boilerplate/types"
	_ "github.com/lib/pq"
)

type Storage interface {
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore(config types.Config) (*PostgresStore, error) {

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

func (s *PostgresStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgresStore) createAccountTable() error {
	query := `create table if not exists users (
		id serial primary key,
		username varchar(100),
		email varchar(100),
		encrypted_password varchar(100)

	)`

	_, err := s.db.Exec(query)
	return err
}
