package com.app.springgraphQL.model;

import jakarta.persistence.*;
import lombok.*;
	
	@Getter
	@Setter
	@NoArgsConstructor
	@AllArgsConstructor
	@Entity
	@Table(name = "book")
	public class Book {
		@Id
		@GeneratedValue(strategy = GenerationType.UUID)
		private String id;
		private String name;
		private int pageCount;
		@ManyToOne(fetch = FetchType.LAZY)
		@JoinColumn(name = "author_id")
		private Author author;
		
		public Book(String name, int pageCount, Author author) {
			this.name = name;
			this.pageCount = pageCount;
			this.author = author;
		}
	}
