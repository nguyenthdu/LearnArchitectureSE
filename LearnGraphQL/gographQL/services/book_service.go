package services

import (
	"gographql/database"
	"gographql/models"
	"gographql/payload/request"
	"log"
)

type BookService interface {
	GetAllBooks() ([]*models.Book, error)
	GetBookByID(id string) (*models.Book, error)
	CreateBook(book *models.Book) error
	UpdateBook(book *models.Book) error
	DeleteBook(id string) error
	UpdateBooksPageCount(inputs []*request.BookPageCountInput) ([]*models.Book, error)
	SearchBooks(name, authorFirstName, authorLastName string, minPageCount, maxPageCount int) ([]*models.Book, error)
}

type BookServiceImpl struct{}

func (b BookServiceImpl) GetAllBooks() ([]*models.Book, error) {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, page_count, author_id FROM book")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.Name, &book.PageCount, &book.Author.ID)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}
	return books, nil
}

func (b BookServiceImpl) GetBookByID(id string) (*models.Book, error) {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var book models.Book
	err = db.QueryRow("SELECT id, name, page_count, author_id FROM book WHERE id = ?", id).Scan(&book.ID, &book.Name, &book.PageCount, &book.Author.ID)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (b BookServiceImpl) CreateBook(book *models.Book) error {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO book (id, name, page_count, author_id) VALUES (?, ?, ?, ?)", book.ID, book.Name, book.PageCount, book.Author.ID)
	if err != nil {
		return err
	}
	return nil
}

func (b BookServiceImpl) UpdateBook(book *models.Book) error {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("UPDATE book SET name = ?, page_count = ?, author_id = ? WHERE id = ?", book.Name, book.PageCount, book.Author.ID, book.ID)
	if err != nil {
		return err
	}
	return nil
}

func (b BookServiceImpl) DeleteBook(id string) error {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM book WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (b BookServiceImpl) UpdateBooksPageCount(inputs []*request.BookPageCountInput) ([]*models.Book, error) {
	//db, err := database.Connect()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer db.Close()
	//
	//tx, err := db.Begin()
	//if err != nil {
	//	return nil, err
	//}
	//
	//for _, input := range inputs {
	//	_, err := tx.Exec("UPDATE book SET page_count = ? WHERE id = ?", input.PageCount, input.ID)
	//	if err != nil {
	//		tx.Rollback()
	//		return nil, err
	//	}
	//}
	//
	//err = tx.Commit()
	//if err != nil {
	//	return nil, err
	//}
	//
	//// Fetch updated books
	//var bookIDs []interface{}
	//for _, input := range inputs {
	//	bookIDs = append(bookIDs, input.ID)
	//}
	//
	//query, args, err := sqlx.In("SELECT id, name, page_count, author_id FROM book WHERE id IN (?)", bookIDs)
	//if err != nil {
	//	return nil, err
	//}
	//query = db.Rebind(query)
	//
	//rows, err := db.Query(query, args...)
	//if err != nil {
	//	return nil, err
	//}
	//defer rows.Close()
	//
	//var books []*models.Book
	//for rows.Next() {
	//	var book models.Book
	//	err := rows.Scan(&book.ID, &book.Name, &book.PageCount, &book.Author.ID)
	//	if err != nil {
	//		return nil, err
	//	}
	//	books = append(books, &book)
	//}
	return nil, nil
}

func (b BookServiceImpl) SearchBooks(name, authorFirstName, authorLastName string, minPageCount, maxPageCount int) ([]*models.Book, error) {
	// TODO: Implement search books functionality
	panic("implement me")
}

func NewBookService() BookService {
	return &BookServiceImpl{}
}
