package order

type Order struct {
	ID       string `json:"id"`
	Customer string `json:"customer"`
	Item     string `json:"item"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

type UpdateQuantity struct {
	ID       int `json:"id"`
	Quantity int `json:"quantity"`
}
