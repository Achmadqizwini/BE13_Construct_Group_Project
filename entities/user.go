package entities

import "time"

type User struct {
	Id         int
	Nama       string
	Gender     string
	No_telepon string
	Password   string
	Saldo      int
	Created_at time.Time
}
