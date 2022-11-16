package controllers

import (
	"be13/account-service-app-project/entities"
	"database/sql"
)

func Login(db *sql.DB, input entities.User) (int, error) {
	idAccount := db.QueryRow("select id from users where no_telepon in(?) && password in(?)", input.No_telepon, input.Password)
	var userrow entities.User
	errScan := idAccount.Scan(&userrow.Id)
	if errScan != nil {
		return -1, errScan
	} else {
		var id_account = userrow.Id
		return id_account, nil
	}
}

func Get_saldo(db *sql.DB, ent entities.User) (int, error) {
	idSaldo := db.QueryRow("select saldo from users where no_telepon = (?)", ent.No_telepon)
	var userrow entities.User
	errScan := idSaldo.Scan(&userrow.Saldo)
	if errScan != nil {
		return -1, errScan
	} else {
		var id_Saldo = userrow.Saldo
		return id_Saldo, nil
	}
}

// func Get_saldo(db *sql.DB, id_account int) (int, error) {
// 	idSaldo := db.QueryRow("select saldo from users where id = (?)", id_account)
// 	var userrow entities.User
// 	errScan := idSaldo.Scan(&userrow.Saldo)
// 	if errScan != nil {
// 		return -1, errScan
// 	} else {
// 		// var id_Saldo = userrow.Saldo
// 		return userrow.Id, nil
// 	}
// }
