package models

import "time"

type VehicleType struct {
	Id        *int
	Name      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeleteAt  *time.Time
}