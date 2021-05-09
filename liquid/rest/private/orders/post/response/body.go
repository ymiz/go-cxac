package response

type Body struct {
	Id             int    `json:"id"`
	Status         string `json:"status"`
	FilledQuantity string `json:"filled_quantity"`
}
