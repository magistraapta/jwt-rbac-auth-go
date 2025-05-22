package model

type Product struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Description string `json:"description"`
	Image       string `json:"image" gorm:"default:product.png"`
}
