package db

import (
	"database/sql"
	"fmt"

	"github.com/RiadMefti/go-api-boilerplate/types"
	_ "github.com/lib/pq"
)

type Storage interface {
	CreateUser(*types.User) error
	GetUserById(string) (*types.User, error)
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

func (s *PostgresStore) CreateUser(user *types.User) error {

	stmt, err := s.db.Prepare("INSERT INTO users(id, username, email, encrypted_password) VALUES ($1, $2, $3, $4)")

	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.ID, user.Username, user.Email, user.EncryptedPassword)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) GetUserById(id string) (*types.User, error) {
	var user types.User

	stmt, err := s.db.Prepare("SELECT id, username, email, encrypted_password FROM users WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&user.ID, &user.Username, &user.Email, &user.EncryptedPassword)
	if err != nil {
		if err == sql.ErrNoRows {

			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
