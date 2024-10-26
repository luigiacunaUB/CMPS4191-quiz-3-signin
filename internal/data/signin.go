package data

import (
	"time"

	"github.com/luigiacunaUB/CMPS4191-quiz-3-signin/internal/validator"
)

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
