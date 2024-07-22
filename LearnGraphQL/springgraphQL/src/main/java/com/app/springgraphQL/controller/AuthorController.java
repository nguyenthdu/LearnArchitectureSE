package com.app.springgraphQL.controller;

import com.app.springgraphQL.model.Author;
import com.app.springgraphQL.payload.response.ResponseMessage;
import com.app.springgraphQL.service.AuthorService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.graphql.data.method.annotation.Argument;
import org.springframework.graphql.data.method.annotation.MutationMapping;
import org.springframework.graphql.data.method.annotation.QueryMapping;
import org.springframework.stereotype.Controller;

import java.util.List;

@Controller
public class AuthorController {
	@Autowired
	private AuthorService authorService;
	
	@QueryMapping
	public List<Author> getAllAuthors() {
		return authorService.getAllAuthors();
	}
	
	@QueryMapping
	public Author getAuthorById(@Argument String id) {
		return authorService.getAuthorById(id);
	}
	
	@MutationMapping
	public Author createAuthor(@Argument String firstName, @Argument String lastName) {
		return authorService.createAuthor(firstName, lastName);
	}
	
	@MutationMapping
	public ResponseMessage updateAuthor(@Argument String id, @Argument String firstName, @Argument String lastName) {
		return authorService.updateAuthor(id, firstName, lastName);
	}
	
	@MutationMapping
	public ResponseMessage deleteAuthor(@Argument String id) {
		return authorService.deleteAuthor(id);
	}
}