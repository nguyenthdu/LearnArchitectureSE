package com.app.springgraphQL.service;

import com.app.springgraphQL.model.Author;
import com.app.springgraphQL.model.Book;
import com.app.springgraphQL.payload.request.BookPageCountInput;
import com.app.springgraphQL.repository.AuthorRepository;
import com.app.springgraphQL.repository.BookRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.stream.Collectors;

@Service
public class BookService {
	@Autowired
	private BookRepository bookRepository;
	@Autowired
	private AuthorRepository authorRepository;
	
	public List<Book> getAllBooks() {
		return bookRepository.findAll();
	}
	
	public Book getBookById(String id) {
		return bookRepository.findById(id).orElse(null);
	}
	
	public Book createBook(String name, int pageCount, String authorId) {
		Author author = authorRepository.findById(authorId).orElse(null);
		if(author == null) {
			throw new IllegalArgumentException("Author not found");
		}
		Book book = new Book(name, pageCount, author);
		return bookRepository.save(book);
	}
	
	public Book updateBook(String id, String name, int pageCount, String authorId) {
		Book book = bookRepository.findById(id).orElse(null);
		Author author = authorRepository.findById(authorId).orElse(null);
		if(book != null && author != null) {
			book.setName(name);
			book.setPageCount(pageCount);
			book.setAuthor(author);
			bookRepository.save(book);
		}
		return book;
	}
	/*TODO:
	*   chức năng phức tạp hơn thể hiện rõ ưu điểm của GraphQL so với REST, chúng ta sẽ tạo một truy vấn (query) và một mutation phức tạp hơn, bao gồm cả việc lồng các mối quan hệ và sử dụng các tính năng nâng cao của GraphQL như batch loading và data fetching optimization.*/
	//TODO:cho phép tìm kiếm sách dựa trên nhiều tiêu chí và hiển thị thông tin chi tiết về tác giả cùng với sách đó.
	public List<Book> searchBooks(String name, String authorFirstName, String authorLastName, Integer minPageCount, Integer maxPageCount) {
		// Logic tìm kiếm phức tạp sử dụng các tiêu chí đã cung cấp
		return bookRepository.findBooksByCriteria(name, authorFirstName, authorLastName, minPageCount, maxPageCount);
	}
    //TODO:  thêm một tính năng để cập nhật số lượng trang của nhiều sách cùng một lúc (batch update).
	 public List<Book> updateBooksPageCount(List<BookPageCountInput> booksPageCountInput) {
        List<String> bookIds = booksPageCountInput.stream()
                .map(BookPageCountInput::getId)
                .collect(Collectors.toList());

        List<Book> books = bookRepository.findAllById(bookIds);

        for (Book book : books) {
            booksPageCountInput.stream()
                    .filter(input -> input.getId().equals(book.getId()))
                    .findFirst()
                    .ifPresent(input -> book.setPageCount(input.getPageCount()));
        }

        return bookRepository.saveAll(books);
    }
	public void deleteBook(String id) {
		bookRepository.deleteById(id);
	}
}
