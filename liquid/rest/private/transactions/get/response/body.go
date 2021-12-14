package response

type Body struct {
	Models []struct {
		Id                int         `json:"id"`
		CreatedAt         int         `json:"created_at"`
		TransactionType   string      `json:"transaction_type"`
		Notes             interface{} `json:"notes"`
		GrossAmount       float64     `json:"gross_amount"`
		NetAmount         float64     `json:"net_amount"`
		Fee               float64     `json:"fee"`
		FxRate            interface{} `json:"fx_rate"`
		FromFiatAccountId interface{} `json:"from_fiat_account_id"`
		ToFiatAccountId   int         `json:"to_fiat_account_id"`
		FromRole          interface{} `json:"from_role"`
		ToRole            interface{} `json:"to_role"`
		BusinessDate      interface{} `json:"business_date"`
		Currency          string      `json:"currency"`
		ActionId          interface{} `json:"action_id"`
		Loan              interface{} `json:"loan"`
	} `json:"models"`
}
