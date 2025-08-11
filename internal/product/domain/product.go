package domain

type ProductPrice struct {
	Source      string
	Price       string
	Link        string
	Description string
}

type Product struct {
	Name string
}

type Suggestions struct {
	Products string `json:"products"`
}
