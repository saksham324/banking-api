package migrations

import (
	"saksham324/go-bank/helpers"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
}

type Account struct {
	gorm.Model
	Type    string
	Name    string
	Balance uint
	UserId  uint
}

func connectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=sakshamarora dbname=bankapp")
	helpers.HandleError(err)

	return db
}

func createAccounts() {
	db := connectDB()

	users := [2]User{
		{Username: "saksham324", Email: "saksham.arora.23@dartmouth.edu"},
		{Username: "devpnn", Email: "dev.punaini.22@dartmouth.edu"},
	}

	for i := 0; i < len(users); i++ {
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		db.Create(&user)

		account := Account{Type: "checking", Name: string(users[i].Username + "'s" + " account"), Balance: uint(10000 * int(i+1)), UserId: user.ID}
		db.Create(&account)
	}
	defer db.Close()
}

func Migrate() {
	db := connectDB()
	db.AutoMigrate(&User{}, &Account{})

	defer db.Close()
	createAccounts()
}
