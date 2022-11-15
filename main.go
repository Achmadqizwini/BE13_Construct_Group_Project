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
