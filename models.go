package main

type requestedFile struct {
	FileName	string	`json:"file_name"`
	FileContent string	`json:"file_content,omitempty"`
}

type httpResponse struct {
	Status		string			`json:"status"`
	Code		int				`json:"code"`
	Message		string			`json:"message"`
	Data		requestedFile	`json:"data,omitempty"`
}
