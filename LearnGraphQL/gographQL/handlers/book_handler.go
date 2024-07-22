package handlers

import (
	"github.com/graphql-go/graphql"
	"gographql/services"
)

func GetBooksField(service services.BookService) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(bookType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return service.GetAllBooks()
		},
	}
}

func GetBookByIDField(service services.BookService) *graphql.Field {
	return &graphql.Field{
		Type: bookType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.String},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id := p.Args["id"].(string)
			return service.GetBookByID(id)
		},
	}
}
