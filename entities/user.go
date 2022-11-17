package entities

import "time"

type User struct {
	Id         int
	Nama       string
	Gender     int
	No_telepon string
	Password   string
	Saldo      int
	Created_at time.Time
}
