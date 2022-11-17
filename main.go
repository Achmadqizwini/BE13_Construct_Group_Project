package main

import (
	"be13/account-service-app-project/config"
	"be13/account-service-app-project/controllers"
	"be13/account-service-app-project/entities"
	"fmt"
	"log"
	"time"

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

		newUser := entities.User{}
		var jeniskel int
		fmt.Println("Masukkan nama user:")
		fmt.Scanln(&newUser.Nama)
		fmt.Println("Pilih Nomor Jenis Kelamin: \n1. Male \n2. Female")
		fmt.Scanln(&jeniskel)
		switch jeniskel {
		case 1:
			{
				newUser.Gender = "male"
			}
		case 2:
			{
				newUser.Gender = "female"
			}
		}
		fmt.Println("Masukkan Nomor Telephone:")
		fmt.Scanln(&newUser.No_telepon)
		fmt.Println("Masukkan Password:")
		fmt.Scanln(&newUser.Password)

		rowsAffected, err := controllers.AccountRegister(db, newUser)
		if err != nil {
			fmt.Println("error insert data")
			// fmt.Println(err)
		} else {
			if rowsAffected == 0 {
				fmt.Println("gagal insert data")
			} else {
				fmt.Println("insert data berhasil")
			}
		}

	case 2:
		{
			account := entities.User{}

			fmt.Println("Log in :")
			fmt.Println("Masukkan Nomor Telepon:")
			fmt.Scanln(&account.No_telepon)
			fmt.Println("Masukkan Password:")
			fmt.Scanln(&account.Password)

			id_account, err := controllers.Login(db, account)
			if err == nil {
				fmt.Println()
				fmt.Println("login berhasil")
				fmt.Println()
				fmt.Println("Menu : \n1. Profile \n2. Transaksi \n3. Cari pengguna\n ")
				fmt.Println("Masukan pilihan anda :")
				var pilihan int
				fmt.Scanln(&pilihan)

				switch pilihan {
				case 1:
					{
						fmt.Println("Menu Profile : \n1. Lihat Profile \n2. Update Profile \n3. Hapus Profile")
						fmt.Println("Masukkan Pilihan Anda")
						var pilihan int
						fmt.Scanln(&pilihan)
						switch pilihan {
						case 1:
							{
								userrow, err := controllers.LihatProfile(db, id_account)
								if err != nil {
									fmt.Println("Tidak bisa menampilkan profile", err.Error())
								} else {
									fmt.Printf("nama : %s\n, gender : %s\n, no_telepon : %s\n, saldo : %d\n", userrow.Nama, userrow.Gender, userrow.No_telepon, userrow.Saldo)
								}
							}

						case 2:
							{
								fmt.Println("Update Profile")
								updateUser := entities.User{}
								fmt.Println("Update nama : ")
								fmt.Scanln(&updateUser.Nama)
								fmt.Println("Update Nomor Telepon : ")
								fmt.Scanln(&updateUser.No_telepon)
								fmt.Println("Update Password : ")
								fmt.Scanln(&updateUser.Password)

								rows, error := controllers.UpdateProfile(db, updateUser, id_account)
								if error != nil {
									fmt.Println("Eksekusi Gagal")
								} else {
									if rows > 0 {
										fmt.Println("Update Berhasil")
									} else {
										fmt.Println("Update Gagal")
									}
								}
							}

						case 3:
							{
								fmt.Println("Apa anda yakin akan menghapus profil?")
								fmt.Println("1. Ya 2. Tidak")
								var pilihanhapus int

								fmt.Scanln(&pilihanhapus)
								switch pilihanhapus {
								case 1:
									{
										row, error := controllers.HapusProfile(db, id_account)
										if error != nil {
											fmt.Println("Eksekusi Gagal", error.Error())
										} else {
											if row > 0 {
												fmt.Println("Hapus Berhasil")
											} else {
												fmt.Println("Hapus Gagal")
											}
										}
									}

								case 2:
									{
										fmt.Println("Tidak Jadi Hapus")
									}
								}

							}
						}
					}

				case 2:

					fmt.Println("Menu : \n1. Top Up \n2. Transfer \n3. Lihat history top up \n4. Lihat history transfer")
					fmt.Println("Pilih transaksi :")
					var pilihan int
					fmt.Scanln(&pilihan)

					switch pilihan {
					case 1:
						{
							top_up := entities.User{}
							fmt.Println()
							fmt.Println("Masukkan nomor telepon :")
							fmt.Scanln(&top_up.No_telepon)
							fmt.Println("Masukkan nominal top up :")
							var nominal int
							fmt.Scanln(&nominal)
							nominal2 := &nominal

							rowsAffected, err := controllers.Topup(db, id_account, nominal)
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
										{

											saldo, err := controllers.Get_saldo(db, top_up)
											if err != nil {
												fmt.Println("error get saldo")
											} else {
												newsaldo, err := controllers.TambahSaldo(db, *nominal2, saldo, top_up)
												if err != nil {
													fmt.Println("error dalam menambahkan saldo, kasian")
												} else {
													fmt.Println("saldo anda saat ini :", newsaldo)
												}
											}

										}
									case 2:
										{
											// idNom := db.QueryRow("select nominal from top_up where user_id = (?)", id_account)
											// var userrow entities.Top_up
											// errScan := idNom.Scan(&userrow.Nominal)
											// if errScan != nil {
											// 	log.Fatal("error", errScan.Error())
											// } else {
											// 	var id_Nom = userrow.Nominal
											// 	fmt.Println(id_Nom)
											// }
										}
									}
								}
							}
						}
					case 2:
						{
							transfer := entities.Transfer{}
							tujuan := entities.User{}
							fmt.Println("Masukkan nomor telepon penerima :")
							fmt.Scanln(&tujuan.No_telepon)
							fmt.Println("Tambahkan nominal transfer :")
							var nominal int
							fmt.Scanln(&nominal)
							var nominal2 = &nominal
							fmt.Println("Masukkan keterangan :")
							fmt.Scanln(&transfer.Keterangan)
							// var nom = &transfer.Nominal

							saldoUser, err := controllers.Get_saldo(db, account)
							if err != nil {
								fmt.Println("error cek saldo", err.Error())
							} else {
								fmt.Println(saldoUser)
							}
							if saldoUser < transfer.Nominal {
								fmt.Println("Tidak dapat melakukan transfer, Saldo anda tidak mencukupi")
							} else {
								_, err := controllers.Transfer(db, tujuan, transfer, id_account, nominal2)
								if err != nil {
									fmt.Println("Tidak dapat melakukan transfer")
								} else {
									fmt.Println("Transfer Sukses")
									transfer.Created_at = time.Now()
									saldoPenerima, err := controllers.Get_saldo(db, tujuan)
									if err != nil {
										fmt.Println("error cek saldo", err.Error())
									} else {
										newsaldoReceiver, err := controllers.TambahSaldo(db, *nominal2, saldoPenerima, tujuan)
										if err != nil {
											fmt.Println("error dalam menambahkan saldo, kasian")
										} else {
											fmt.Println("saldo anda saat ini (penerima):", newsaldoReceiver)
										}
									}

									newSaldoSender, err := controllers.KurangSaldo(db, *nominal2, saldoUser, account)
									if err != nil {
										fmt.Println("Saldo anda tetap sama")
									} else {
										fmt.Println("Saldo anda saat ini (pengirim): ", newSaldoSender)
									}
								}

							}

						}
					case 3:
						{
							fmt.Println("History Top up yang dilakukan Anda :")
							result, errSelect := db.Query("select top_up.user_id, users.nama, top_up.nominal, top_up.created_at from top_up inner join users on top_up.user_id = users.id where users.id = (?)", id_account)
							if errSelect != nil {
								log.Fatal("error select", errSelect.Error())
							}

							idAccount := db.QueryRow("select nama from users where id = (?)", id_account)
							var userrow entities.User
							errScan := idAccount.Scan(&userrow.Nama)
							if errScan == nil {
								var nama_user = userrow.Nama
								var dataHistory []entities.Top_up
								for result.Next() {
									var userrow entities.Top_up
									userrow.Created_at = time.Now()
									errScan := result.Scan(&userrow.User_id, &nama_user, &userrow.Nominal, &userrow.Created_at)
									if errScan != nil {
										log.Fatal("error scan", errScan.Error())
									}
									dataHistory = append(dataHistory, userrow)

								}

								for _, value := range dataHistory {
									fmt.Printf("User id : %d, Nama : %s, nominal : %d, created_at: %s\n", value.User_id, nama_user, value.Nominal, value.Created_at)

								}
							} else {

								fmt.Println("error", errScan.Error())
							}

						}
					case 4:
						{
							// fmt.Println("History Transfer yang Anda Lakukan:")
							// result, errSelect := db.Query("select u.nama, v.nama, transfer.nominal, transfer.keterangan, transfer.created_at from transfer inner join users u on transfer.user_id_pengirim = u.id inner join users v on transfer.user_id_penerima = v.id where transfer.user_id_pengirim = 1;", id_account)
							// if errSelect != nil {
							// 	log.Fatal("error select", errSelect.Error())
							// }

							// idAccount := db.QueryRow("select nama from users where id = (?)", id_account)
							// var user entities.User
							// errScan := idAccount.Scan(&user.Nama)
							// if errScan == nil {
							// 	var nama_user = user.Nama
							// 	var dataHistory []entities.Transfer
							// 	for result.Next() {
							// 		var userrow entities.Transfer
							// 		userrow.Created_at = time.Now()
							// 		errScan := result.Scan(&nama_user, &userrow.Nominal, &userrow.Keterangan, &userrow.Created_at)
							// 		if errScan != nil {
							// 			log.Fatal("error scan", errScan.Error())
							// 		}
							// 		dataHistory = append(dataHistory, userrow)

							// 	}

							// 	for _, value := range dataHistory {
							// 		fmt.Printf("Nama Pengirim : %s, Nama Penerima : %s, Nominal : %d, Keterangan : %s, created_at: %s\n", nama_user, namapenerima, value.Nominal, value.keterangan, value.Created_at)

							// 	}
							// } else {

							// 	fmt.Println("error", errScan.Error())
							// }
						}

					}
				case 3:
					{
						caripengguna := entities.User{}
						fmt.Println("Masukkan No Telepon yang dicari :")
						fmt.Scanln(&caripengguna.No_telepon)

						caripengguna2, err := controllers.CariPengguna(db, caripengguna)
						if err != nil {
							fmt.Println("Profil Tidak Ditemukan")
						} else {
							fmt.Printf("nama : %s\ngender : %s\nno_telepon : %s\n", caripengguna2.Nama, caripengguna2.Gender, caripengguna2.No_telepon)
						}
					}

				}
			} else {
				fmt.Println("Login gagal, data yang anda masukkan salah")
			}

		}

	case 3:
		// var namapengirim string
	}

}
