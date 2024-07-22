package response

type ResponseMessage struct {
	Message   string `json:"message"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}
