package database

import (
	"fmt"
	"strings"

	m "github.com/BNPrashanth/test-go-pg/models"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// CreateSchema exported
func CreateSchema(db *pg.DB) error {
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

// AddNewUser exported
func AddNewUser(db *pg.DB, user *m.User) error {
	if err := db.Insert(user); err != nil {
		fmt.Println(err.Error())
	}
	return nil
}

// SelectOneUser exported
func SelectOneUser(db *pg.DB, id int64) error {
	user := &m.User{ID: id}
	err := db.Select(user)
	if err != nil {
		fmt.Println("Select error: " + err.Error())
	}

	fmt.Println("User: ", user.String())
	return nil
}

// SelectAllUsers exported
func SelectAllUsers(db *pg.DB) error {
	var users []m.User
	err := db.Model().Table("users").Select(&users)
	if err != nil {
		fmt.Println("Select error: " + err.Error())
	}

	for _, u := range users {
		fmt.Println(u.String())
	}
	return nil
}

// UpdateUser exported
func UpdateUser(db *pg.DB, modifiedUser *m.User) error {
	err := db.Update(modifiedUser)
	if err != nil {
		fmt.Println("Update error: " + err.Error())
	}
	return nil
}
