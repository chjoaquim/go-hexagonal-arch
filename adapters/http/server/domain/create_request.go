package server

type CreateRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
