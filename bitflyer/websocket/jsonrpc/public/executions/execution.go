package executions

import "time"

type Execution struct {
	Id                         int       `json:"id"`
	Side                       string    `json:"side"`
	Price                      float64   `json:"price"`
	Size                       float64   `json:"size"`
	ExecDate                   time.Time `json:"exec_date"`
	BuyChildOrderAcceptanceId  string    `json:"buy_child_order_acceptance_id"`
	SellChildOrderAcceptanceId string    `json:"sell_child_order_acceptance_id"`
}
