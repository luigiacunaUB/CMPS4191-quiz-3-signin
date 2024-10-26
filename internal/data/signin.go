package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/luigiacunaUB/CMPS4191-quiz-3-signin/internal/validator"
)

type SignINModel struct {
	DB *sql.DB
}

// the database table
type SignIN struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	FullName  string `json: "fullname"`
	LoginDate string `json:"logindate"`
}

func ValidateSignin(v *validator.Validator, signin *SignIN) {
	v.Check(signin.Email != "", "Email", "must be provided")
	v.Check(signin.FullName != "", "FullName", "Must be Provided")
	v.Check(len(signin.Email) <= 100, "Email", "must not be more than 100 bytes")
	v.Check(len(signin.FullName) <= 100, "FullName", "must not be more than 100 bytes")
}

func (s SignINModel) Insert(signin *SignIN) error {
	query := `
		INSERT INTO users (Email,FullName)
		VALUES ($1,$2)
		RETURNING ID
	`

	args := []any{signin.Email, signin.FullName}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return s.DB.QueryRowContext(ctx, query, args...).Scan(&signin.ID)
}

func (s SignINModel) Get(id int64) (*SignIN, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `
		SELECT id,email,fullname,logindate
		FROM users
		WHERE id = $1
	`

	var signin SignIN
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := s.DB.QueryRowContext(ctx, query, id).Scan(&signin.ID, &signin.Email, &signin.FullName, &signin.LoginDate)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &signin, nil
}

func (s SignINModel) Update(signin *SignIN) error {
	query := `
	UPDATE users
	SET email = $1 , fullname = $2
	WHERE id = $3
	RETURNING id
	`
	args := []any{signin.Email, signin.FullName, signin.ID}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return s.DB.QueryRowContext(ctx, query, args...).Scan(&signin.ID)
}

func (s SignINModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `
		DELETE FROM users
		WHERE id = $1
		`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := s.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}
