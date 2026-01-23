package models

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

var DummyProduct = []Product{
	{
		ID:    1,
		Name:  "Indomie",
		Price: 3500,
		Stock: 10,
	},
	{
		ID:    2,
		Name:  "Aqua",
		Price: 3000,
		Stock: 40,
	},
	{
		ID:    3,
		Name:  "Kecap",
		Price: 3000,
		Stock: 40,
	},
}
