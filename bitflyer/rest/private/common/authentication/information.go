package authentication

// 認証情報
type Information struct {
	accessKey string
	secretKey string
}

func (i *Information) AccessKey() string {
	return i.accessKey
}

func (i *Information) SecretKey() string {
	return i.secretKey
}

func NewInformation(accessKey string, secretKey string) *Information {
	return &Information{accessKey: accessKey, secretKey: secretKey}
}
