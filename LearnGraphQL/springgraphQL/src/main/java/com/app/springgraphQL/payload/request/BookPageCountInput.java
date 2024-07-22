package com.app.springgraphQL.payload.request;

import lombok.Data;

@Data
public class BookPageCountInput {
	String id;
	int pageCount;
}
