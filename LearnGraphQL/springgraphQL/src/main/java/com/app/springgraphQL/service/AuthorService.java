package com.app.springgraphQL.service;

import com.app.springgraphQL.model.Author;
import com.app.springgraphQL.payload.response.ResponseMessage;
import com.app.springgraphQL.repository.AuthorRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Service;

import java.time.LocalDateTime;
import java.util.List;
import java.util.Optional;

@Service
public class AuthorService {
    @Autowired
    private AuthorRepository authorRepository;
    public List<Author> getAllAuthors() {
        return authorRepository.findAll();
    }

    public Author getAuthorById(String id) {
        return authorRepository.findById(id).orElse(null);
    }

    public Author createAuthor(String firstName, String lastName) {
        Author author = new Author(firstName, lastName);
        return authorRepository.save(author);
    }

    public ResponseMessage updateAuthor(String id, String firstName, String lastName) {
        Optional<Author> authorOptional = authorRepository.findById(id);
        if (authorOptional.isPresent()) {
            Author author = authorOptional.get();
            author.setFirstName(firstName);
            author.setLastName(lastName);
            authorRepository.save(author);
            return new ResponseMessage("Author updated successfully with id: "+id, HttpStatus.OK.toString(), LocalDateTime.now().toString());
        } else {
            return new ResponseMessage("Author not found with id: "+ id, HttpStatus.NOT_FOUND.toString(), LocalDateTime.now().toString());
        }
    }

    public ResponseMessage deleteAuthor(String id) {
         Optional<Author> authorOptional = authorRepository.findById(id);
    if (authorOptional.isPresent()) {
        authorRepository.deleteById(id);
        return new ResponseMessage("Author deleted successfully with id: "+id, HttpStatus.OK.toString(), LocalDateTime.now().toString());
    } else {
        return new ResponseMessage("Author not found with id: "+ id, HttpStatus.NOT_FOUND.toString(), LocalDateTime.now().toString());
    }
    }
}
