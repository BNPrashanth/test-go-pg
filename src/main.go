package main

import (
    "fmt"

	"github.com/spf13/viper"
    "github.com/go-pg/pg"
    "github.com/go-pg/pg/orm"
)

// User exported
type User struct {
    ID     int64
    Name   string
    Emails []string
}

func (u User) String() string {
    return fmt.Sprintf("User<%d %s %v>", u.ID, u.Name, u.Emails)
}

// Story exported
type Story struct {
    ID       int64
    Title    string
    AuthorID int64
    Author   *User
}

func (s Story) String() string {
    return fmt.Sprintf("Story<%d %s %s>", s.ID, s.Title, s.Author)
}

func main() {
	fmt.Println("Starting...")

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	dbusername := viper.GetString("pgdbusername")
	dbpassword := viper.GetString("pgdbpassword")
	database := viper.GetString("pgdatabase")

	db := pg.Connect(&pg.Options{
		User: dbusername,
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
    for _, model := range []interface{}{(*User)(nil), (*Story)(nil)} {
        err := db.CreateTable(model, &orm.CreateTableOptions{
            Temp: true,
        })
        if err != nil {
            return err
        }
    }
    return nil
}
