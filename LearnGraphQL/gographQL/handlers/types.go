package handlers

import (
	"github.com/graphql-go/graphql"
)

var authorType *graphql.Object
var bookType *graphql.Object

func init() {
	// Khởi tạo bookType trước để sử dụng trong authorType
	bookType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Book",
			Fields: graphql.Fields{
				"id":        &graphql.Field{Type: graphql.String},
				"name":      &graphql.Field{Type: graphql.String},
				"pageCount": &graphql.Field{Type: graphql.Int},
				"author":    &graphql.Field{Type: graphql.String}, // Sử dụng authorType sau khi đã khởi tạo
			},
		},
	)

	// Khởi tạo authorType
	authorType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Author",
			Fields: graphql.Fields{
				"id":        &graphql.Field{Type: graphql.String},
				"firstName": &graphql.Field{Type: graphql.String},
				"lastName":  &graphql.Field{Type: graphql.String},
				"books":     &graphql.Field{Type: graphql.NewList(bookType)},
			},
		},
	)

	// Cập nhật lại field author trong bookType
	bookType.AddFieldConfig("author", &graphql.Field{Type: authorType})
}
