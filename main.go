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
		// {
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

		// }

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
					fmt.Println("Profile :")
					{
						fmt.Println("Menu Profile : \n1. Lihat Profile \n2. Update Profile \n3. Hapus Profile")
						fmt.Println("Masukkan Pilihan Anda")
						var pilihan int
						fmt.Scanln(&pilihan)
						switch pilihan {
						case 1:
							{
								result := db.QueryRow("select nama, gender, no_telepon, saldo from users where id in(?)", id_account)

								var userrow entities.User // membuat variabel penampung
								errScan := result.Scan(&userrow.Nama, &userrow.Gender, &userrow.No_telepon, &userrow.Saldo)
								if errScan != nil {
									log.Fatal("error scan", errScan.Error())
								}
								fmt.Printf("nama : %s\n, gender : %s\n, no_telepon : %s\n, saldo : %d\n", userrow.Nama, userrow.Gender, userrow.No_telepon, userrow.Saldo)
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

								var query = ("Update users set nama = ?, no_telepon = ?, password = ? where id = (?)")
								statement, errPrepare := db.Prepare(query)
								if errPrepare != nil {
									log.Fatal("error prepare insert ", errPrepare.Error())
								}

								result, errExec := statement.Exec(updateUser.Nama, updateUser.No_telepon, updateUser.Password, id_account)
								if errExec != nil {
									log.Fatal("error execution insert ", errExec.Error())
								} else {
									row, _ := result.RowsAffected()
									if row > 0 {
										fmt.Println("update berhasil")
									} else {
										fmt.Println("update gagal")
									}
								}
							}

						case 3:
							{
								fmt.Println("Apa anda yakin akan menghapus profil?")
								fmt.Println("1. Ya 2. Tidak")
								var pilihanhapus int
								fmt.Scanln(&pilihanhapus)
								// deleteData := entities.User{} tambahkan if

								var query = ("Delete from users where id = ?")
								statement, errPrepare := db.Prepare(query)
								if errPrepare != nil {
									log.Fatal("error prepare insert ", errPrepare.Error())
								}

								result, errExec := statement.Exec(id_account)
								if errExec != nil {
									log.Fatal("error exec insert", errExec.Error())
								} else {
									row, _ := result.RowsAffected()
									if row > 0 {
										fmt.Println("delete berhasil")
									} else {
										fmt.Println("delete gagal")
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
										// case 2:
										// 	{

										// 	}
									}
								}
							}
						}
					case 2:
						{
							transfer := entities.Transfer{}
							tujuan := entities.User{}
							fmt.Println("Masukkan nomor telepon tujuan :")
							fmt.Scanln(&tujuan.No_telepon)
							fmt.Println("Tambahkan nominal transfer :")
							var nominal int
							fmt.Scanln(&nominal)
							var nominal2 = &nominal
							fmt.Println("Masukkan keterangan :")
							fmt.Scanln(&transfer.Keterangan)

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
							//history transfer
						}

					}
				case 3:
					{
						caripengguna := entities.User{}
						fmt.Println("Masukkan No Telepon yang dicari :")
						fmt.Scanln(&caripengguna.No_telepon)
						result := db.QueryRow("select nama, gender, no_telepon, from users where no_telepon in(?)", caripengguna.No_telepon)

						var userrow entities.User // membuat variabel penampung
						errScan := result.Scan(&userrow.Nama, &userrow.Gender, &userrow.No_telepon)
						if errScan != nil {
							log.Fatal("error scan", errScan.Error())
						} else {
							fmt.Printf("nama : %s\n gender : %s\n no telepon : %s", userrow.Nama, userrow.Gender, userrow.No_telepon)
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
