package repository

import (
	"database/sql"
)

type AuthRepository struct {
	DB *sqlx.DB
}
