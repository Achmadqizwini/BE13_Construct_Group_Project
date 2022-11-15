package main

import (
	"be13/account-service-app-project/config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.ConnectToDB()
	defer db.Close()

	fmt.Println("Menu : \n1. Register \n2. Top Up \n3. Transfer\n ")
	fmt.Println("Masukan pilihan anda :")
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:

	case 2:

	case 3:

	}

}
