package com.app.springgraphQL.controller;

import com.app.springgraphQL.model.Book;
import com.app.springgraphQL.payload.request.BookPageCountInput;
import com.app.springgraphQL.service.BookService;
import org.springframework.graphql.data.method.annotation.Argument;
import org.springframework.graphql.data.method.annotation.MutationMapping;
import org.springframework.graphql.data.method.annotation.QueryMapping;
import org.springframework.stereotype.Controller;
import org.springframework.beans.factory.annotation.Autowired;

import java.util.List;

@Controller
public class BookController {
	@Autowired
	private BookService bookService;
	
	@QueryMapping
	public List<Book> getAllBooks() {
		return bookService.getAllBooks();
	}
	
	@QueryMapping
	public Book getBookById(@Argument String id) {
		return bookService.getBookById(id);
	}
	
	@MutationMapping
	public Book createBook(@Argument String name, @Argument int pageCount, @Argument String authorId) {
		return bookService.createBook(name, pageCount, authorId);
	}
	
	@MutationMapping
	public Book updateBook(@Argument String id, @Argument String name, @Argument int pageCount, @Argument String authorId) {
		return bookService.updateBook(id, name, pageCount, authorId);
	}
	
	@MutationMapping
	public void deleteBook(@Argument String id) {
		bookService.deleteBook(id);
	}
	
	@QueryMapping
	public List<Book> searchBooks(@Argument String name, @Argument String authorFirstName, @Argument String authorLastName, @Argument Integer minPageCount, @Argument Integer maxPageCount) {
		return bookService.searchBooks(name, authorFirstName, authorLastName, minPageCount, maxPageCount);
	}
	
	@MutationMapping
	public List<Book> updateBooksPageCount(@Argument List<BookPageCountInput> booksPageCountInput) {
		return bookService.updateBooksPageCount(booksPageCountInput);
	}
}
