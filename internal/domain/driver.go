package domain

import (
	"github.com/google/uuid"
)

type Driver struct {
	ID           uuid.UUID `db:"id"`
	FirstName    string    `db:"first_name"`
	LastName     string    `db:"last_name"`
	Email        string    `db:"email"`
	Password     string    `db:"password"` //hash password
	LicenseID    string    `db:"licenseid"`
	VehicleModel string    `db:"vehicle_model"`
	VehiclePlate string    `db:"vehicle_plate"`
}
