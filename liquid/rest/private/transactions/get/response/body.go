package response

type Body struct {
	Models []struct {
		Id                int         `json:"id"`
		CreatedAt         int         `json:"created_at"`
		GrossAmount       string      `json:"gross_amount"`
		TransactionType   string      `json:"transaction_type"`
		NetAmount         string      `json:"net_amount"`
		FromRole          interface{} `json:"from_role"`
		ToRole            interface{} `json:"to_role"`
		Currency          string      `json:"currency"`
		ActionId          interface{} `json:"action_id"`
		Notes             interface{} `json:"notes"`
		Fee               string      `json:"fee"`
		FxRate            interface{} `json:"fx_rate"`
		FromFiatAccountId int         `json:"from_fiat_account_id"`
		ToFiatAccountId   int         `json:"to_fiat_account_id"`
		BusinessDate      interface{} `json:"business_date"`
		Loan              interface{} `json:"loan"`
	} `json:"models"`
}
