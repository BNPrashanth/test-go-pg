package config

import (
	"fmt"

	"github.com/go-pg/pg"
	"github.com/spf13/viper"
)

// DB exported
var DB *pg.DB

func init() {
	initializeReadConfig()
	initializeDatabase()
}

func initializeReadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
}

func initializeDatabase() {
	dbusername := viper.GetString("pgdbusername")
	dbpassword := viper.GetString("pgdbpassword")
	database := viper.GetString("pgdatabase")

	db := pg.Connect(&pg.Options{
		User:     dbusername,
		Password: dbpassword,
		Database: database,
	})
	defer db.Close()
	DB = db
}
