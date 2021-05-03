package option

type Option string

const (
	// TakerOnly is TAKER のデータのみを送信する
	// trades channel 用
	// @see https://api.coin.z.com/docs/#ws-trades
	TakerOnly = Option("TAKER_ONLY")
)
