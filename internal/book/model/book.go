package model

import "time"

type Book struct {
	BookID    string    `db:"book_id" json:"bookId"`
	BookName  string    `db:"book_name" json:"bookName"`
	Author    string    `db:"author" json:"author"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}

type BookToUpdate struct {
	BookName string `json:"bookName"`
	Author   string `json:"author"`
}
