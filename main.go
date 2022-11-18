package main

import (
	"be13/account-service-app-project/config"
	"be13/account-service-app-project/controllers"
	"be13/account-service-app-project/entities"
	"fmt"
	"os"

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
									fmt.Printf("Nama : %s\nJenis Kelamin : %s\nNo telepon : %s\nSaldo : %d\n", userrow.Nama, userrow.Gender, userrow.No_telepon, userrow.Saldo)
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
							fmt.Println("Masukkan nomor telepon Anda:")
							fmt.Scanln(&top_up.No_telepon)
							fmt.Println("Masukkan nominal top up :")
							var nominal int
							fmt.Scanln(&nominal)
							// nominal2 := &nominal

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

											saldo, err := controllers.Get_saldo(db, id_account)
											if err != nil {
												fmt.Println("error get saldo")
											} else {
												newsaldo, err := controllers.TambahSaldo(db, nominal, saldo, id_account)
												if err != nil {
													fmt.Println("error dalam menambahkan saldo")
												} else {
													fmt.Println("saldo anda saat ini :", newsaldo)
												}
											}

										}
									case 2:
										{
											fmt.Println("Tidak melakukan cek saldo")
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
							fmt.Println("Masukkan keterangan :")
							fmt.Scanln(&transfer.Keterangan)

							saldoUser, err := controllers.Get_saldo(db, id_account)
							if err != nil {
								fmt.Println("error cek saldo", err.Error())
							} else {
								fmt.Println("Saldo Anda saat ini :", saldoUser)
								fmt.Println()
							}

							if saldoUser < nominal {
								fmt.Println("Tidak dapat melakukan transfer, Saldo anda tidak mencukupi")
							} else {
								_, err := controllers.Transfer(db, tujuan, transfer, id_account, nominal)
								if err != nil {
									fmt.Println("Tidak dapat melakukan transfer")
								} else {
									fmt.Println("Transfer Sukses")

									saldoPenerima, err := controllers.Get_saldo_2(db, tujuan)
									if err != nil {
										fmt.Println("error cek saldo", err.Error())
									} else {
										newsaldoReceiver, err := controllers.TambahSaldo_tf(db, nominal, saldoPenerima, tujuan)
										if err != nil {
											fmt.Println("error dalam menambahkan saldo")
										} else {
											fmt.Println("saldo anda saat ini (penerima):", newsaldoReceiver)
										}
									}

									newSaldoSender, err := controllers.KurangSaldo(db, nominal, saldoUser, id_account)
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
							history_topup, err := controllers.History_topup(db, id_account)
							if err != nil {
								fmt.Println("error get the history", err.Error())
							} else {
								for _, value := range history_topup {
									fmt.Printf("Id : %d, Nama : %s, Nominal : %d, Tanggal : %s\n\n", value.User_id, value.Nama, value.Nominal, value.Created_at)
								}
							}

						}
					case 4:
						{
							fmt.Println("History Transfer yang telah anda lakukan :")
							history_tf, err := controllers.Transfer_history(db, id_account)
							if err != nil {
								fmt.Println("error get the history", err.Error())
							} else {
								for _, value := range history_tf {
									fmt.Printf("Nama pengirim : %s, Nama penerima : %s, Nominal Transfer : %d, Keterangan : %s, Tanggal : %s\n\n", value.Nama_Pengirim, value.Nama_Penerima, value.Nominal, value.Keterangan, value.Created_at)
								}
							}
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
		fmt.Println("Terimakasih telah bertransaksi")

		os.Exit(3)
	}

}
