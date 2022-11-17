package controllers

import (
	"be13/account-service-app-project/entities"
	"database/sql"
	"log"
	"time"
)

func Transfer(db *sql.DB, entity entities.User, entity2 entities.Transfer, id_account int, nominal int) (int, error) {
	entity2.Created_at = time.Now()
	var query = "Insert into Transfer (User_id_Pengirim, User_id_Penerima, nominal, keterangan, created_at) Values (?, ?, ?,?,?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return -1, errPrepare
	}

	idTujuan := db.QueryRow("select id from users where no_telepon in(?)", entity.No_telepon)
	var userrow entities.User
	errScan := idTujuan.Scan(&userrow.Id)
	if errScan != nil {
		return -1, errScan
	}

	result, errExec := statement.Exec(id_account, userrow.Id, nominal, entity2.Keterangan, entity2.Created_at)
	if errExec != nil {
		return -1, errExec
	} else {
		row, _ := result.RowsAffected()
		if row == 0 {
			return -1, errExec
		} else {
			return int(row), nil
		}
	}
}

func Transfer_history(db *sql.DB, id_account int) ([]entities.History_Transfer_Respon, error) {
	result, errSelect := db.Query("select u.nama, v.nama, transfer.nominal, transfer.keterangan, transfer.created_at from transfer inner join users u on transfer.user_id_pengirim = u.id inner join users v on transfer.user_id_penerima = v.id where transfer.user_id_pengirim = (?);", id_account)
	if errSelect != nil {
		log.Fatal("error select", errSelect.Error())
	}

	var dataHistory []entities.History_Transfer_Respon
	for result.Next() {
		var userrow entities.History_Transfer_Respon
		errScan := result.Scan(&userrow.Nama_Pengirim, &userrow.Nama_Penerima, &userrow.Nominal, &userrow.Keterangan, &userrow.Created_at)
		if errScan != nil {
			log.Fatal("error scan", errScan.Error())
		}
		dataHistory = append(dataHistory, userrow)
	}
	return dataHistory, nil
}

func TambahSaldo_tf(db *sql.DB, nominal int, saldo int, ent entities.User) (int, error) {
	var query = ("Update users set saldo = ? where no_telepon = ?")
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return -1, errPrepare
	}

	var newsaldo = saldo + nominal
	result, errExec := statement.Exec(newsaldo, ent.No_telepon)
	if errExec != nil {
		return -1, errExec
	} else {
		_, errRow := result.RowsAffected()
		if errRow != nil {
			return 0, errRow
		}
		return newsaldo, nil
	}
}
