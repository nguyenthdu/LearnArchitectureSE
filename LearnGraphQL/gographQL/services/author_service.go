package services

import (
	"gographql/database"
	"gographql/models"
	"log"
)

// services/authorService.go

type AuthorService interface {
	GetAllAuthors() ([]*models.Author, error)
	GetAuthorByID(id string) (*models.Author, error)
	CreateAuthor(author *models.Author) error
	UpdateAuthor(author *models.Author) error
	DeleteAuthor(id string) error
}

type AuthorServiceImpl struct{}

func (a AuthorServiceImpl) GetAllAuthors() ([]*models.Author, error) {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, first_name, last_name FROM author")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []*models.Author
	for rows.Next() {
		var author models.Author
		err := rows.Scan(&author.ID, &author.FirstName, &author.LastName)
		if err != nil {
			return nil, err
		}
		authors = append(authors, &author)
	}
	return authors, nil
}

func (a AuthorServiceImpl) GetAuthorByID(id string) (*models.Author, error) {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var author models.Author
	err = db.QueryRow("SELECT id, first_name, last_name FROM author WHERE id = ?", id).Scan(&author.ID, &author.FirstName, &author.LastName)
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func (a AuthorServiceImpl) CreateAuthor(author *models.Author) error {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO author (id, first_name, last_name) VALUES (?, ?, ?)", author.ID, author.FirstName, author.LastName)
	if err != nil {
		return err
	}
	return nil
}

func (a AuthorServiceImpl) UpdateAuthor(author *models.Author) error {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("UPDATE author SET first_name = ?, last_name = ? WHERE id = ?", author.FirstName, author.LastName, author.ID)
	if err != nil {
		return err
	}
	return nil
}

func (a AuthorServiceImpl) DeleteAuthor(id string) error {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM author WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func NewAuthorService() AuthorService {
	return &AuthorServiceImpl{}
}
