package main

import (
	"fmt"
	"log"
	services "main/internal/services"
	"main/migration"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	conn, err := migration.InitializeDBTable()
	if err != nil {
		log.Fatal("Error creating DB or table: %w", err)
	}
	defer conn.DB.Close()

	//создаем экземпляр структуры, реализующей необходимые методы
	userService := services.NewUserService(conn)

	//Starting transaction
	tx := conn.DB.MustBegin()

	var methodError error

	methodError = userService.CreateUser(tx, "Fanil", "Kamaletdinov", "USD", 1000)
	if methodError != nil {
		log.Fatal(methodError)
	}
	methodError = userService.AlterBalance(tx, "Fanil", "Kamaletdinov", -100)
	if methodError != nil {
		log.Fatal(methodError)
	}

	//Closing transaction
	if err := tx.Commit(); err != nil {
		log.Fatal("Couldn't commit transaction: ", err)
	}

	//Checking the data
	user, err := userService.GetUserInfo("Fanil", "Kamaletdinov")
	if err != nil {
		log.Fatal("Error: %w", err)
	}
	fmt.Printf("User ID: %d\nName: %s\nSurname: %s\nBalance: %.2f%s\n", user.ID, user.Name, user.Surname, user.Balance, user.Currency)

}
