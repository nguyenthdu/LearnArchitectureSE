package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"gographql/handlers"
	"gographql/services"
)

func main() {
	authorService := services.NewAuthorService()
	bookService := services.NewBookService()

	fields := graphql.Fields{
		"getAllAuthors": handlers.GetAuthorsField(authorService),
		"getAuthorById": handlers.GetAuthorByIDField(authorService),
		"getAllBooks":   handlers.GetBooksField(bookService),
		"getBookById":   handlers.GetBookByIDField(bookService),
	}

	rootQuery := graphql.NewObject(graphql.ObjectConfig{Name: "Query", Fields: fields})

	schemaConfig := graphql.SchemaConfig{Query: rootQuery}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})

	http.Handle("/graphql", h)
	log.Println("Now server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
