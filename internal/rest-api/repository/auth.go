package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lakhan-purohit/net-http/internal/pkg/db"
	"github.com/lakhan-purohit/net-http/internal/pkg/utils"
	"github.com/lakhan-purohit/net-http/internal/rest-api/model"
)

type IAuthRepository interface {
	Login(ctx context.Context, email, password string) (*model.User, error)
	SignUp(username, email, password, avatar string) (*model.User, error)
}

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) Login(
	ctx context.Context,
	email, password string,
) (*model.User, error) {

	query := `
		SELECT uuid, id, username, email, status, password
		FROM users
		WHERE email = ?
		LIMIT 1
	`

	// Internal struct for scanning including the password
	var result struct {
		model.User
		Password string `db:"password"`
	}

	if err := db.FindOne(ctx, query, &result, email); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("invalid credentials")
		}
		return nil, err
	}

	if !utils.ComparePassword(result.Password, password) {
		return nil, errors.New("invalid credentials")
	}

	return &result.User, nil
}

func (r *AuthRepository) SignUp(userName string, email string, password string, avatar string) (*model.User, error) {
	passwordHash, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	uuid := utils.UUID()
	query := `
		INSERT INTO users (uuid, username, email, password, avatar)
		VALUES (?, ?, ?, ?, ?)
	`
	userID, err := db.Insert(context.Background(), query, uuid, userName, email, passwordHash, avatar)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:       int64(userID),
		Username: userName,
		Email:    email,
		UUID:     uuid,
		Status:   1,
		Avatar:   avatar,
	}, nil
}
