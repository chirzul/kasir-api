package models

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var DummyCategories = []Category{
	{
		ID:          1,
		Name:        "Food & Beverages",
		Description: "Category for food and drinks",
	},
	{
		ID:          2,
		Name:        "Electronics",
		Description: "Category for electronic items",
	},
}
