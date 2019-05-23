package main

import (
	"fmt"

	c "github.com/BNPrashanth/test-go-pg/configs"
	d "github.com/BNPrashanth/test-go-pg/database"
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
}
