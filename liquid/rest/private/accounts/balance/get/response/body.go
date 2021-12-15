package response

type Body struct {
	Models []struct {
		Currency string `json:"currency"`
		Balance  string `json:"balance"`
	} `json:"models"`
}
