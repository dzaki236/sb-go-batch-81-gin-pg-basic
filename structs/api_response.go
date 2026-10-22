package structs

type ApiResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}
