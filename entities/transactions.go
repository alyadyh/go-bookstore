package entities

import "time"

type Transactions struct {
	Id         uint
	CartID     int
	StaffID    int
	TransaDate time.Time
	TotalPrice float64
}
