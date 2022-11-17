package entities

import "time"

type Transfer struct {
	Id               int
	User_id_Pengirim int
	User_id_Penerima int
	Nominal          int
	Keterangan       string
	Created_at       time.Time
}
