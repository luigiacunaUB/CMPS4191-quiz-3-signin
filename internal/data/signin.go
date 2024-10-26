package data

import (
	"context"
	"database/sql"
	"time"

	"github.com/luigiacunaUB/CMPS4191-quiz-3-signin/internal/validator"
)

type SignINModel struct {
	DB *sql.DB
}

// the database table
type SignIN struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	FullName  string    `json: "fullname"`
	LoginDate time.Time `json: "logindate"`
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
		RETURNING ID,LoginDate
	`

	args := []any{signin.Email, signin.FullName}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return s.DB.QueryRowContext(ctx, query, args...).Scan(&signin.ID, &signin.LoginDate)
}
