package entities

import "time"

type Top_up struct {
	Id             int
	User_id        int
	User_id_tujuan int
	Nominal        int
	Created_at     time.Time
}
