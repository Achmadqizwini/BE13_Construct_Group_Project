package controllers

import (
	"fmt"
	"log"
)

func AccountRegister() {
	var query = "Insert into user(Nama, Gender, No_telepon, Password) Values (?, ?, ?, ?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		log.Fatal("error prepare Register ", errPrepare.Error())
	}

	result, errExec := statement.Exec(newUser.Nama, newUser.Gender, newUser.No_telepon, newUser.Password)
	if errExec != nil {
		log.Fatal("error exec Register", errExec.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("Register berhasil")
		} else {
			fmt.Println("Register gagal")
		}
	}
}
