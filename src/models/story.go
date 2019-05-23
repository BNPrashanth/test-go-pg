package models

import "fmt"

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
