package response

type Body struct {
	Id                 int           `json:"id"`
	Currency           string        `json:"currency"`
	Balance            float64       `json:"balance"`
	FreeBalance        float64       `json:"free_balance"`
	Pnl                float64       `json:"pnl"`
	Margin             float64       `json:"margin"`
	OrdersMargin       float64       `json:"orders_margin"`
	FreeMargin         float64       `json:"free_margin"`
	Equity             float64       `json:"equity"`
	UserId             int           `json:"user_id"`
	MaintenanceMargin  string        `json:"maintenance_margin"`
	ReservedBalance    string        `json:"reserved_balance"`
	IcoLockingTrackers []interface{} `json:"ico_locking_trackers"`
}
