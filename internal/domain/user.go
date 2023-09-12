package domain

import (
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `db:"id"`
	FirstName string    `db:"first_name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"` //hash password
}
