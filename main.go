package main

import (
	"be13/account-service-app-project/config"
	"be13/account-service-app-project/controllers"
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
				var id_account = userrow.Id
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

					fmt.Println("Menu : \n1. Top Up \n2. Transfer \n3. Lihat history top up \n4. Lihat history transfer")
					fmt.Println("Pilih transaksi :")
					var pilihan int
					fmt.Scanln(&pilihan)

					switch pilihan {
					case 1:
						{
							top_up := entities.User{}
							fmt.Println("Masukkan nomor telepon :")
							fmt.Scanln(&top_up.No_telepon)
							fmt.Println("Masukkan nominal top up :")
							var nominal int
							fmt.Scanln(&nominal)

							rowsAffected, err := controllers.Topup(db, id_account)
							if err != nil {
								fmt.Println("error top up")
							} else {
								if rowsAffected == 0 {
									fmt.Println("Top up gagal")
								} else {
									fmt.Println("Top up berhasil")
									fmt.Println()
									fmt.Println("Cek saldo anda :")
									fmt.Println("1. Ya \n2. Tidak ")
									var cek int
									fmt.Scanln(&cek)

									switch cek {
									case 1:
										Saldo, err := controllers.TambahSaldo(db, nominal, id_account)
										if err != nil {
											fmt.Println("error cek saldo")
										} else {
											fmt.Println(Saldo)
										}
									}

								}
							}
						}
					case 2:
						{
							transfer := entities.User{}
							fmt.Println("Masukkan nomor telepon tujuan :")
							fmt.Scanln(&transfer.No_telepon)
							fmt.Println("Masukkan nominal transfer :")
							var nominal int
							fmt.Scanln(&nominal)
							detail := entities.Transfer{}
							fmt.Println("Tambahkan keterangan (opsional) :")
							fmt.Scanln(&detail.Keterangan)

							var query = "Insert into Transfer (User_id, User_id_tujuan, nominal, keterangan, created_at) Values (?, ?, ?,?,?)"
							statement, errPrepare := db.Prepare(query)
							if errPrepare != nil {
								log.Fatal("error prepare top up", errPrepare.Error())
							}
							idTujuan := db.QueryRow("select id from users where no_telepon in(?)", transfer.No_telepon)

							var userrow entities.User
							errScan := idAccount.Scan(&userrow.Id)
							if errScan != nil {
								log.Fatal("error scan no telpon ", errScan.Error())
							}

							if userrow.Saldo > nominal {
								result, errExec := statement.Exec(idAccount, idTujuan, nominal, detail.Keterangan)
								if errExec != nil {
									log.Fatal("error exec transfer", errExec.Error())
								} else {
									row, _ := result.RowsAffected()
									if row > 0 {
										fmt.Println("Transfer berhasil")

									} else {
										fmt.Println("Transfer gagal")
									}
								}
							} else {
								fmt.Println("miskin?")
							}

						}
					case 3:
					case 4:

					}
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
