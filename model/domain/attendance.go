package domain

import (
	"time"

	"github.com/google/uuid"
)

type Attendance struct {
	ID        string    `json:"id"`
	IPAddress string    `json:"ip_address"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	ClockIn   time.Time `json:"clock_in"`
	ClockOut  time.Time `json:"clock_out"`
	UserId    string    `json:"user_id"`
}

func (attendance *Attendance) GenerateID(id string) {
	uuid := uuid.New().String()
	attendance.ID = uuid
}
