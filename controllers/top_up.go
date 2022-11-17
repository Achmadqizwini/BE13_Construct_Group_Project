package controllers

import (
	"be13/account-service-app-project/entities"
	"database/sql"
	"log"
)

func Topup(db *sql.DB, id_user int, nominal int) (int, error) {
	var query = "Insert into Top_up (User_id, nominal) Values (?, ?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return -1, errPrepare
	}
	result, errExec := statement.Exec(id_user, nominal)
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

func TambahSaldo(db *sql.DB, nominal int, saldo int, id_account int) (int, error) {
	var query = ("Update users set saldo = ? where id = ?")
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return -1, errPrepare
	}

	var newsaldo = saldo + nominal
	result, errExec := statement.Exec(newsaldo, id_account)
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

func SelectNom(db *sql.DB, ent entities.User) (int, error) {
	idNom := db.QueryRow("select nominal from top_up where no_telepon = (?)", ent.No_telepon)
	var userrow entities.Top_up
	errScan := idNom.Scan(&userrow.Nominal)
	if errScan != nil {
		return -1, errScan
	} else {
		var id_Nom = userrow.Nominal
		return id_Nom, nil
	}
}

func KurangSaldo(db *sql.DB, nominal int, saldo int, id_account int) (int, error) {
	var query = ("Update users set saldo = ? where id = ?")
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return -1, errPrepare
	}

	var newsaldo = saldo - nominal
	result, errExec := statement.Exec(newsaldo, id_account)
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

func History_topup(db *sql.DB, id_account int) ([]entities.History_topup_respon, error) {
	result, errSelect := db.Query("select top_up.user_id, users.nama, top_up.nominal, top_up.created_at from top_up inner join users on top_up.user_id = users.id where users.id = (?)", id_account)
	if errSelect != nil {
		log.Fatal("error select", errSelect.Error())
	}

	var dataHistory []entities.History_topup_respon
	for result.Next() {
		var userrow2 entities.History_topup_respon
		errScan := result.Scan(&userrow2.User_id, &userrow2.Nama, &userrow2.Nominal, &userrow2.Created_at)
		if errScan != nil {
			log.Fatal("error scan", errScan.Error())
		}
		dataHistory = append(dataHistory, userrow2)

	}

	return dataHistory, nil

}
