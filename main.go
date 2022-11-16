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
		{
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

		}

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

				case 2:
				case 3:

				}
			} else if errScan != nil {
				log.Fatal("error scan insert ", errScan.Error())
			}

		}

	case 3:

	}

}
