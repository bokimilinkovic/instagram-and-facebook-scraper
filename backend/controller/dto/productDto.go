package dto

// ProductDto represent a request for creating product
type ProductDto struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Sponsor     string  `json:"sponsor"`
	Image       string  `json:"omitempty"`
}
