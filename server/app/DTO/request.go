package DTO

//Requestis the request object for Search API and  Insert API
type Request struct {
	SearchString string `json:"word"`
}
