package controllers

import (
	"be13/account-service-app-project/entities"
	"database/sql"
	"log"
	"time"
)

// func Transfer(db *sql.DB, transfer entities.User, id_account int, saldo_user int) (int, error) {
// 	var nominal int
// 	var detail entities.Transfer
// 	var query = "Insert into Transfer (User_id, User_id_tujuan, nominal, keterangan) Values (?, ?, ?,?,?)"
// 	statement, errPrepare := db.Prepare(query)
// 	if errPrepare != nil {
// 		log.Fatal("error prepare top up", errPrepare.Error())
// 	}
// 	idTujuan := db.QueryRow("select id from users where no_telepon in(?)", transfer.No_telepon)

// 	var userrow entities.User
// 	errScan := idTujuan.Scan(&userrow.Id)
// 	if errScan != nil {
// 		log.Fatal("error scan no telpon ", errScan.Error())
// 	}
// 	if saldo_user > nominal {
// 		result, errExec := statement.Exec(id_account, userrow.Id, nominal, detail.Keterangan, detail.Created_at)
// 		if errExec != nil {
// 			return -1, errExec
// 		} else {
// 			row, errRow := result.RowsAffected()
// 			if row > 0 {
// 				return int(row), nil
// 			} else {
// 				return -1, errRow
// 			}
// 		}
// 	} else {
// 		return -1, sql.ErrNoRows
// 	}
// }

func Transfer(db *sql.DB, entity entities.User, entity2 entities.Transfer, id_account int, nominal *int) (int, error) {
	entity2.Created_at = time.Now()
	var query = "Insert into Transfer (User_id_Pengirim, User_id_Penerima, nominal, keterangan, created_at) Values (?, ?, ?,?,?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		log.Fatal("error prepare top up", errPrepare.Error())
	}
	idTujuan := db.QueryRow("select id from users where no_telepon in(?)", entity.No_telepon)

	var userrow entities.User
	errScan := idTujuan.Scan(&userrow.Id)
	if errScan != nil {
		log.Fatal("error scan no telpon ", errScan.Error())
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
