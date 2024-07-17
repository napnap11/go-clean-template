package dto

type ExampleRequest struct {
    Query string `json:"query" binding:"required"`
}

type ExampleResponse struct {
    Message string `json:"message"`
    Data    string `json:"data"`
}