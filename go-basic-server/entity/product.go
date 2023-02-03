package entity

type Product struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Category    *Category `json:"category,omitempty"`
	Quantity    int       `json:"quantity"`
}

type Category struct {
	Name string `json:"name,omitempty"`
}
