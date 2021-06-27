package request

type Body struct {
	Order Order `json:"order"`
}

type Order struct {
	Quantity string `json:"quantity"`
	Price    string `json:"price"`
}
