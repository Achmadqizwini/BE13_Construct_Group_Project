package entities

import "time"

type Top_up struct {
	Id         int
	User_id    int
	Nominal    int
	Created_at time.Time
}

type History_topup_respon struct {
	User_id    int
	Nama       string
	Nominal    int
	Created_at time.Time
}
