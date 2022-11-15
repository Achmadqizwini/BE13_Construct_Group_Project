package controllers

import (
	"database/sql"
	"log"
)

func Topup(db *sql.DB, id_user int) (int, error) {
	var nominal int
	var query = "Insert into Top_up (User_id, nominal) Values (?, ?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		log.Fatal("error prepare top up", errPrepare.Error())
	}

	result, errExec := statement.Exec(id_user, &nominal)
	if errExec != nil {
		return -1, errExec
	} else {
		row, errRow := result.RowsAffected()
		if row > 0 {
			return int(row), nil
		} else {
			return -1, errRow
		}
	}
}

func GetSaldo(db *sql.DB, id_user int) (int, error) {
	Saldo := db.QueryRow("select saldo from users where id in = (?)", id_user)
	var saldo int
	errScan := Saldo.Scan(&saldo)
	if errScan != nil {
		return -1, errScan
	}
	return saldo, nil
}

// var userrow entities.User
// var nominal int
// errScan := Saldo.Scan(&userrow.Saldo)
// if errScan != nil {
// 	log.Fatal("error scan saldo ", errScan.Error())
// } else {
// 	totalsaldo := userrow.Saldo + nominal
// 	var newSaldo = "update users set saldo = ?"
// 	statement, errPrepare := db.Prepare(newSaldo)
// 	if errPrepare != nil {
// 		log.Fatal("error prepare insert ", errPrepare.Error())
// 	}

// 	result, errExec := statement.Exec(totalsaldo)
// 	if errExec != nil {
// 		log.Fatal("error execution update ", errExec.Error())
// 	} else {
// 		row, _ := result.RowsAffected()
// 		if row > 0 {
// 			fmt.Println("update berhasil")
// 		} else {
// 			fmt.Println("update gagal")
// 		}
// 	}
// }

func TambahSaldo(db *sql.DB, nominal int, id_user int) (int, error) {
	var query = ("Update users set saldo = ? where id = ?")
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return -1, errPrepare
	}
	saldo, err := GetSaldo(db, id_user)
	if err != nil {
		return -1, err
	}

	var newsaldo = saldo + nominal
	result, errExec := statement.Exec(&newsaldo, id_user)
	if errExec != nil {
		return -1, errExec
	} else {
		row, errRow := result.RowsAffected()
		if errRow != nil {
			return 0, errRow
		}
		return int(row), nil
	}
}
