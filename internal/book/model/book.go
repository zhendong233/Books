package model

import "time"

type Book struct {
	BookID    string    `db:"book_id"`
	BookName  string    `db:"book_name"`
	Number    int       `db:"number"`
	Author    string    `db:"author"`
	CreatedAt time.Time `db:"created_at"`
}
