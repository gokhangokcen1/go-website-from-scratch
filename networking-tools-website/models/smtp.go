package models

type SMTPMessageRequest struct {
	Host    string `json:"host"`
	Port    string `json:"port"`
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type SMTPMessageResponse struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Trace   []string `json:"trace"`
}
