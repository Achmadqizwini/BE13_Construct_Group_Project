package main

import (
	"be13/account-service-app-project/config"
	"be13/account-service-app-project/entities"
	"fmt"

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
		newUser := entities.User{}
		fmt.Println("Masukkan Nomor Telepon:")
		fmt.Scanln(&newUser.No_telepon)
		fmt.Println("Masukkan Password:")
		fmt.Scanln(&newUser.Password)

	case 3:

	}

}
