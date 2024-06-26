package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/RiadMefti/go-api-boilerplate/types"
	_ "github.com/lib/pq"
)

type Storage interface {
	CreateUser(ctx context.Context, user *types.User) error
	GetUserByEmail(ctx context.Context, email string) (*types.User, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore(config types.Config) (*PostgresStore, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DbHost, config.DbPort, config.DbUser, config.DbPassword, config.DbName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &PostgresStore{db: db}, nil
}

func (s *PostgresStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgresStore) createAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(100),
		email VARCHAR(100) UNIQUE,
		encrypted_password VARCHAR(100)
	)`

	if _, err := s.db.Exec(query); err != nil {
		return fmt.Errorf("failed to create users table: %w", err)
	}
	return nil
}

func (s *PostgresStore) CreateUser(ctx context.Context, user *types.User) error {
	query := "INSERT INTO users (id, username, email, encrypted_password) VALUES ($1, $2, $3, $4)"
	_, err := s.db.ExecContext(ctx, query, user.ID, user.Username, user.Email, user.EncryptedPassword)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (s *PostgresStore) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {
	query := "SELECT id, username, email, encrypted_password FROM users WHERE email = $1"
	var user types.User
	err := s.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Username, &user.Email, &user.EncryptedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}
	return &user, nil
}
