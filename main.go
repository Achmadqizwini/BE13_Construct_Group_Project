package main

import (
	"be13/account-service-app-project/config"
	"be13/account-service-app-project/entities"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.ConnectToDB()
	defer db.Close()

	fmt.Println("Menu : \n1. Register \n2. Log in \n3. Close\n ")
	fmt.Println("Masukan pilihan anda :")
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		// {
		// 	newUser := entities.User{}
		// 	fmt.Println("Masukkan id user:")
		// 	fmt.Scanln(&newUser.Id)
		// 	fmt.Println("Masukkan nama user:")
		// 	fmt.Scanln(&newUser.Nama)
		// 	fmt.Println("Masukkan email user:")
		// 	fmt.Scanln(&newUser.Email)
		// 	fmt.Println("Masukkan password user:")
		// 	fmt.Scanln(&newUser.Password)
		// 	fmt.Println("Masukkan phone user:")
		// 	fmt.Scanln(&newUser.Phone)
		// 	fmt.Println("Masukkan domisili user:")
		// 	fmt.Scanln(&newUser.Domisili)

		// 	rowsAffected, err := controllers.InsertDataToUser(db, newUser)
		// 	if err != nil {
		// 		fmt.Println("error insert data")
		// 	} else {
		// 		if rowsAffected == 0 {
		// 			fmt.Println("gagal insert data")
		// 		} else {
		// 			fmt.Println("insert data berhasil")
		// 		}
		// 	}

		// }

	case 2:
		account := entities.User{}
		fmt.Println("Log in :")
		fmt.Println("Masukkan Nomor Telepon:")
		fmt.Scanln(&account.No_telepon)
		fmt.Println("Masukkan Password:")
		fmt.Scanln(&account.Password)

		{
			idAccount := db.QueryRow("select id from users where no_telepon in(?) && password in(?)", account.No_telepon, account.Password)

			var userrow entities.User
			errScan := idAccount.Scan(&userrow.Id)
			if errScan == nil {
				fmt.Println()
				fmt.Println("login berhasil")
				fmt.Println()
				fmt.Println("Menu : \n1. Profile \n2. Transaksi \n3. Cari pengguna\n ")
				fmt.Println("Masukan pilihan anda :")
				var pilihan int
				fmt.Scanln(&pilihan)

				switch pilihan {
				case 1:
					fmt.Println("Profile :")
				// case 1: read
				// {
				// 	result, errSelect := db.Query("select id, nama, email, password, phone, domisili from user")
				// 	if errSelect != nil {
				// 		log.Fatal("error select", errSelect.Error())
				// 	}

				// 	var dataUser []User
				// 	for result.Next() { // membaca tiap baris/row dari hasil query
				// 		var userrow User                                                                                                         // perbaris                                                                                             // membuat variabel penampung
				// 		errScan := result.Scan(&userrow.Id, &userrow.Nama, &userrow.Email, &userrow.Password, &userrow.Phone, &userrow.Domisili) // melakukan scanning data dari masing" row dan menyimpannya kedalam variabel yang dibuat sebelumnya
				// 		if errScan != nil {                                                                                                      // handling ketika ada error pada saat proses scannign
				// 			log.Fatal("error scan", errScan.Error())
				// 		}
				// 		// fmt.Printf("id %s, nama %s, email %s\n", userrow.Id, userrow.Nama, userrow.Email)
				// 		dataUser = append(dataUser, userrow)
				// 	}

				// 	for _, value := range dataUser {
				// 		fmt.Printf("id %s, nama %s, email %s\n", value.Id, value.Nama, value.Email)
				// 	}
				// }

				// case 2: ubah data

				// {
				// 	updateUser := User{}
				// 	fmt.Println("Masukkan id yang akan di-update: ")
				// 	fmt.Scanln(&updateUser.Id)
				// 	fmt.Println("Update nama : ")
				// 	fmt.Scanln(&updateUser.Nama)
				// 	fmt.Println("Update email : ")
				// 	fmt.Scanln(&updateUser.Email)
				// 	fmt.Println("Update Password : ")
				// 	fmt.Scanln(&updateUser.Password)
				// 	fmt.Println("Update number phone : ")
				// 	fmt.Scanln(&updateUser.Phone)
				// 	fmt.Println("Update domisili : ")
				// 	fmt.Scanln(&updateUser.Domisili)

				// 	var query = ("Update user set nama = ?, email = ?, password = ?, phone = ?, domisili = ? where id = ?")
				// 	statement, errPrepare := db.Prepare(query)
				// 	if errPrepare != nil {
				// 		log.Fatal("error prepare insert ", errPrepare.Error())
				// 	}

				// 	result, errExec := statement.Exec(updateUser.Nama, updateUser.Email, updateUser.Password, updateUser.Phone, updateUser.Domisili, updateUser.Id)
				// 	if errExec != nil {
				// 		log.Fatal("error execution insert ", errExec.Error())
				// 	} else {
				// 		row, _ := result.RowsAffected()
				// 		if row > 0 {
				// 			fmt.Println("update berhasil")
				// 		} else {
				// 			fmt.Println("update gagal")
				// 		}
				// 	}
				// }

				// case 3: hapusakun

				// {
				// 	deleteData := User{}
				// 	fmt.Println("Masukkan id yang akan dihapus :")
				// 	fmt.Scanln(&deleteData.Id)

				// 	var query = "Delete from user where id = ?"
				// 	statement, errPrepare := db.Prepare(query)
				// 	if errPrepare != nil {
				// 		log.Fatal("error prepare insert ", errPrepare.Error())
				// 	}

				// 	result, errExec := statement.Exec(deleteData.Id)
				// 	if errExec != nil {
				// 		log.Fatal("error exec insert", errExec.Error())
				// 	} else {
				// 		row, _ := result.RowsAffected()
				// 		if row > 0 {
				// 			fmt.Println("delete berhasil")
				// 		} else {
				// 			fmt.Println("delete gagal")
				// 		}
				// 	}
				// }

				case 2:
				// transaksi
				case 3:
					// {
					// 	idData := entities.User{}
					// 	fmt.Println("Masukkan id yang ingin dicari : ")
					// 	fmt.Scanln(&idData.Id)

					// 	userrow, errGetId := controllers.GetUserById(db, idData)
					// 	if errGetId != nil {
					// 		fmt.Println("error get data user")
					// 	} else {
					// 		fmt.Printf("id :%s \n nama :%s \n email :%s \n phone :%s \n domisili :%s \n ", userrow.Id, userrow.Nama, userrow.Email, userrow.Phone, userrow.Domisili)
					// 	}
					// }

				}
			} else if errScan != nil {
				log.Fatal("error scan insert ", errScan.Error())
			}

		}

	case 3:
		// var namapengirim string
	}

}
