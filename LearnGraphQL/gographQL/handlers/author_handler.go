package handlers

import (
	"github.com/graphql-go/graphql"
	"gographql/services"
)

func GetAuthorsField(service services.AuthorService) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(authorType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return service.GetAllAuthors()
		},
	}
}

func GetAuthorByIDField(service services.AuthorService) *graphql.Field {
	return &graphql.Field{
		Type: authorType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.String},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id := p.Args["id"].(string)
			return service.GetAuthorByID(id)
		},
	}
}
