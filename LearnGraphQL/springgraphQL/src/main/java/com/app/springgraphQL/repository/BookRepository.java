package com.app.springgraphQL.repository;

import com.app.springgraphQL.model.Book;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface BookRepository extends JpaRepository<Book, String> {
	
 @Query("SELECT b FROM Book b JOIN b.author a WHERE " +
            "(b.name LIKE %:name%) AND " +
            "(a.firstName LIKE %:authorFirstName%) AND " +
            "(a.lastName LIKE %:authorLastName%) AND " +
            "(b.pageCount >= :minPageCount) AND " +
            "(b.pageCount <= :maxPageCount)")
    List<Book> findBooksByCriteria(
            @Param("name") String name,
            @Param("authorFirstName") String authorFirstName,
            @Param("authorLastName") String authorLastName,
            @Param("minPageCount") Integer minPageCount,
            @Param("maxPageCount") Integer maxPageCount);}
