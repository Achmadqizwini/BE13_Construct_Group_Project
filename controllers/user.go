package controllers

import (
	"be13/account-service-app-project/entities"
	"database/sql"
)

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
