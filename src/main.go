package main

import (
	"fmt"
	"strings"

	c "github.com/BNPrashanth/test-go-pg/src/configs"
	m "github.com/BNPrashanth/test-go-pg/src/models"
	"github.com/spf13/viper"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
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

	err := createSchema(db)
	if err != nil {
		panic(err)
	}
}

func createSchema(db *pg.DB) error {
	tables := []interface{}{(*m.User)(nil), (*m.Story)(nil)}
	for _, model := range tables {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: false,
		})
		if err != nil {
			if strings.Contains(err.Error(), "already exists") {
				fmt.Println(err.Error())
			} else {
				return err
			}
		}
	}
	return nil
}
