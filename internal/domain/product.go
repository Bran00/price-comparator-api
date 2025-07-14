package domain

type ProductPrice struct {
	Source      string `json:"source"`
	Price       string `json:"price"`
	Link        string `json:"link"`
	Description string `json:"description"`
}

type Product struct {
	Name string
}
