package postgres

import (
	"database/sql"
	"my_mod/library"
)

type LibarySorage struct {
	Db *sql.DB
}

func NewLibraryStorage(db *sql.DB) *LibarySorage {
	return &LibarySorage{Db: db}
}

func (l *LibarySorage) LibaryCreate(books *library.AddBookRequest) (*string, error) {
	var id string
	err := l.Db.QueryRow("INSERT INTO BOOK(TITLE,AUTHOR,YEAR) VALUES($1,$2,$3) returns book_id", books.Titile, books.Author, books.Year).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (l *LibarySorage) LIbaryGet(id int32) error {
	_, err := l.Db.Query("selecr * from book WHERE BOOK_ID=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func (l *LibarySorage) LibaryBorrow(id int32) error {
	_, err := l.Db.Query("selecr * from book WHERE BOOK_ID=$1", id)
	if err != nil {
		return err
	}
	return nil
}
