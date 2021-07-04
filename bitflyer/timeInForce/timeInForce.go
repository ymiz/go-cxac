package timeInForce

type TimeInForce string

const (
	Gtc = TimeInForce("GTC")
	Ioc = TimeInForce("IOC")
	Fok = TimeInForce("FOK")
)

func (f TimeInForce) IsGtc() bool {
	return f == Gtc
}

func (f TimeInForce) IsIoc() bool {
	return f == Ioc
}

func (f TimeInForce) IsFok() bool {
	return f == Fok
}

func (f TimeInForce) ToString() string {
	return string(f)
}
