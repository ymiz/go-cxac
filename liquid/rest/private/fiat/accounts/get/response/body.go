package response

type Body struct {
	Models []struct {
		Id                       int         `json:"id"`
		Currency                 string      `json:"currency"`
		Balance                  float64     `json:"balance"`
		ReservedBalance          float64     `json:"reserved_balance"`
		PusherChannel            string      `json:"pusher_channel"`
		LowestOfferInterestRate  interface{} `json:"lowest_offer_interest_rate"`
		HighestOfferInterestRate interface{} `json:"highest_offer_interest_rate"`
		CurrencySymbol           string      `json:"currency_symbol"`
		SendToBtcAddress         interface{} `json:"send_to_btc_address"`
		ExchangeRate             string      `json:"exchange_rate"`
		CurrencyType             string      `json:"currency_type"`
	} `json:"models"`
}
