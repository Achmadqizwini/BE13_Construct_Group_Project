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

			idAccount := db.QueryRow("select id from users where no_telepon in(?) && password in(?)", account.No_telepon, account.Password)

			var userrow entities.User
			errScan := idAccount.Scan(&userrow.Id)
			if errScan == nil {
				var id_account = userrow.Id
				fmt.Println()
				fmt.Println("login berhasil")
				fmt.Println()
				fmt.Println("Menu : \n1. Profile \n2. Transaksi \n3. Cari pengguna \n4. Log Out")
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
								userrow, err := controllers.LihatProfile(db, account)
								if err != nil {
									fmt.Println("Tidak bisa menampilkan profile")
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

								var query = ("Update users set nama = ?, no_telepon = ?, password = ? where id = ?")
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
								switch pilihanhapus {
								case 1:
									{
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
								case 2:
									{

									}
								}

							}
						}
					}

				case 2:
					{

					}

				case 3:
					{
						caripengguna := entities.User{}
						fmt.Println("Masukkan No Telepon yang dicari :")
						fmt.Scanln(&caripengguna.No_telepon)
						result := db.QueryRow("select nama, gender, no_telepon from users where no_telepon = ?", caripengguna.No_telepon)

						var userrow entities.User
						errScan := result.Scan(&userrow.Nama, &userrow.Gender, &userrow.No_telepon)
						if errScan != nil {
							log.Fatal("error scan", errScan.Error())
						} else {
							fmt.Printf("nama : %s\n gender : %s\n no telepon : %s", userrow.Nama, userrow.Gender, userrow.No_telepon)
						}
					}

				case 4:
					{

					}
				}
			} else if errScan != nil {
				log.Fatal("error scan insert ", errScan.Error())
			}

		}
	case 3:
		{

		}
	}
}
