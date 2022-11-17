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

func AccountRegister(db *sql.DB, newUser entities.User) (int, error) {

	var query = "Insert into users(Nama, Gender, No_telepon, Password) Values (?, ?, ?, ?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return -1, errPrepare
	}

	result, errExec := statement.Exec(newUser.Nama, newUser.Gender, newUser.No_telepon, newUser.Password)
	if errExec != nil {
		return -1, errExec
	} else {
		row, errrow := result.RowsAffected()
		if row > 0 {
			return int(row), nil
		} else {
			return -1, errrow
		}
	}
}

func LihatProfile(db *sql.DB, userrow entities.User) (*entities.User, error) {
	result := db.QueryRow("select nama, gender, no_telepon, saldo from user where id in(?) ", userrow.Id)

	var idData entities.User
	errScan := result.Scan(&idData.Nama, &idData.Gender, &idData.No_telepon, &idData.Saldo)
	if errScan != nil {
		return nil, errScan
	}
	return &idData, nil
}
