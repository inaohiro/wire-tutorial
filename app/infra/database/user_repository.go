package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/inaohiro-sandbox/wire-tutorial3/app/domain/domain"
	"log"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) Get(ctx context.Context, id int64) (user *domain.User, err error) {
	query := fmt.Sprintf("SELECT id, name FROM users WHERE id = %d", id)
	row := repo.db.QueryRow(query)

	_user := domain.User{}
	if err := row.Scan(&_user.ID, &_user.Name); err != nil {
		log.Fatalf("user repository failed: %v", err)
	}
	user = &_user

	return
}

func (repo *UserRepository) Save(ctx context.Context, user *domain.User) (err error) {
	query := fmt.Sprintf("INSERT INTO users(id, name) VALUES (%d, %s);", user.ID, user.Name)
	_, err = repo.db.Exec(query)
	if err != nil {
		log.Fatalf("user repository failed: %v", err)
	}
	return
}
