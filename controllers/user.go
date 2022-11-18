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

		return userrow.Id, nil
	}
}

func Get_saldo(db *sql.DB, id_account int) (int, error) {
	idSaldo := db.QueryRow("select saldo from users where id = (?)", id_account)
	var userrow entities.User
	errScan := idSaldo.Scan(&userrow.Saldo)
	if errScan != nil {
		return -1, errScan
	} else {
		return userrow.Saldo, nil
	}
}

func Get_saldo_2(db *sql.DB, ent entities.User) (int, error) {
	idSaldo := db.QueryRow("select saldo from users where no_telepon = (?)", ent.No_telepon)
	var userrow entities.User
	errScan := idSaldo.Scan(&userrow.Saldo)
	if errScan != nil {
		return -1, errScan
	} else {
		return userrow.Saldo, nil
	}
}

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

func LihatProfile(db *sql.DB, id_account int) (*entities.User, error) {

	result := db.QueryRow("select nama, gender, no_telepon, saldo from users where id in(?) ", id_account)


	var idData entities.User
	errScan := result.Scan(&idData.Nama, &idData.Gender, &idData.No_telepon, &idData.Saldo)
	if errScan != nil {
		return nil, errScan
	}
	return &idData, nil
}

func UpdateProfile(db *sql.DB, updateUser entities.User, id_account int) (int, error) {
	var query = ("Update users set nama = ?, no_telepon = ?,password = ? where id = ?")
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return -1, errPrepare
	}
	result, errExec := statement.Exec(updateUser.Nama, updateUser.No_telepon, updateUser.Password, id_account)
	if errExec != nil {
		return 0, errExec
	} else {
		row, error := result.RowsAffected()
		if row > 0 {
			return int(row), nil
		} else {
			return 0, error
		}
	}
}

func CariPengguna(db *sql.DB, carinomor entities.User) (entities.User, error) {
	result := db.QueryRow("select nama, gender, no_telepon from users where no_telepon = ?", carinomor.No_telepon)

	var penggunalain entities.User
	errScan := result.Scan(&penggunalain.Nama, &penggunalain.Gender, &penggunalain.No_telepon)
	if errScan != nil {
		return penggunalain, errScan
	} else {
		return penggunalain, nil
	}

}

func HapusProfile(db *sql.DB, id_account int) (int, error) {
	query := ("Delete from users where id in (?)")
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return -1, errPrepare
	}

	result, errExec := statement.Exec(id_account)
	if errExec != nil {
		return -1, errExec

	} else {
		row, error := result.RowsAffected()
		if error != nil {
			return -1, error
		} else {
			return int(row), nil
		}
	}
}
