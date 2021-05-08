package common

type Token struct {
	id     string
	secret string
}

func (t Token) Id() string {
	return t.id
}

func (t Token) Secret() string {
	return t.secret
}

func NewToken(id string, secret string) *Token {
	return &Token{id: id, secret: secret}
}
