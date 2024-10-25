package data

import(
	"time"
)

//the database table
type SignIN struct{
	ID int64
	Email string
	FullName string
	LoginDate time.Time
}