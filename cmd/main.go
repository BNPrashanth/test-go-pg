package main

import (
	"fmt"
	"strings"

	c "github.com/BNPrashanth/test-go-pg/configs"
	d "github.com/BNPrashanth/test-go-pg/database"
	m "github.com/BNPrashanth/test-go-pg/models"
	"github.com/spf13/viper"

	"github.com/go-pg/pg"
)

func main() {
	fmt.Println("Starting...")
	c.InitializeReadConfig()

	dbusername := viper.GetString("pgdbusername")
	dbpassword := viper.GetString("pgdbpassword")
	database := viper.GetString("pgdatabase")

	db := pg.Connect(&pg.Options{
		User:     dbusername,
		Password: dbpassword,
		Database: database,
	})
	defer db.Close()

	err := d.CreateSchema(db)
	if err != nil {
		panic(err)
	}

	user := m.User{
		ID:     4,
		Name:   "3rd User",
		Emails: []string{"email@gmail.com", "email2@gmail.com"},
	}
	err = d.AddNewUser(db, &user)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			panic(err.Error())
		}
		fmt.Println(err.Error())
	}
	err = d.SelectOneUser(db, 1)
	err = d.SelectAllUsers(db)

	modifiedUser := m.User{
		ID:     4,
		Name:   "4th User",
		Emails: []string{"email@gmail.com", "email2@gmail.com"},
	}
	err = d.UpdateUser(db, &modifiedUser)
	err = d.SelectOneUser(db, 4)
}
