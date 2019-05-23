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
