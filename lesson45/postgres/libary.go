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

	 l.Db.QueryRow("INSERT INTO Book (Title, Author, Year) VALUES ($1, $2, $3) RETURNING book_id", books.Titile, books.Author, books.Year).Scan(&id)
	return &id, nil
}

func (l *LibarySorage) LIbaryGet(query string) ([]*library.Book, error) {

	rows, err := l.Db.Query("select book_id,title,author,year from book WHERE author=$1 or title=$1", query)
	if err != nil {
		return nil, err
	}

	var books []*library.Book

	for rows.Next() {
		book := &library.Book{}

		err = rows.Scan(&book.BookId, &book.Title, &book.Author, &book.YearPublishet)

		if err != nil {
			return nil, err
		}
		books = append(books, book)

	}

	return books, nil
}

func (l *LibarySorage) LibaryBorrow(ids *library.BorrowBookRequest) (*library.BorrowBookRespons, error) {

	secces := false

	err := l.Db.QueryRow("update book set user_Id=$1 , borrow=true where book_id=$2 returning true", ids.UserId, ids.BookId).Scan(&secces)
	
	return &library.BorrowBookRespons{Succes: secces}, err
}
