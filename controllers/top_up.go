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

func TambahSaldo(db *sql.DB, nominal int, saldo int, ent entities.User) (int, error) {
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

func KurangSaldo(db *sql.DB, nominal int, saldo int, ent entities.User) (int, error) {
	var query = ("Update users set saldo = ? where no_telepon = ?")
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return -1, errPrepare
	}

	var newsaldo = saldo - nominal
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
