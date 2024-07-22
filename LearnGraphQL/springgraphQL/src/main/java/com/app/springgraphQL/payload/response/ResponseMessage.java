package com.app.springgraphQL.payload.response;

import lombok.Data;

@Data
public class ResponseMessage {
	private String message;
	private String status;
	private String timestamp;
	
	public ResponseMessage(String message, String status, String timestamp) {
		this.message = message;
		this.status = status;
		this.timestamp = timestamp;
	}
}
