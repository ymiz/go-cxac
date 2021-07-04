package method

type Method string

const (
	Get    = Method("GET")
	Post   = Method("POST")
	Put    = Method("PUT")
	Delete = Method("DELETE")
)

func (m Method) IsGet() bool {
	return m == Get
}

func (m Method) IsPost() bool {
	return m == Post
}

func (m Method) IsPut() bool {
	return m == Put
}

func (m Method) IsDelete() bool {
	return m == Delete
}

func (m Method) ToString() string {
	return string(m)
}
