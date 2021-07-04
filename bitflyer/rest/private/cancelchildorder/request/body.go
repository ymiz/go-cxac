package request

type Body struct {
	ProductCode            string `json:"product_code"`
	ChildOrderAcceptanceId string `json:"child_order_acceptance_id"`
}
