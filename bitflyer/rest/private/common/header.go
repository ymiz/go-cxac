package common

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

func GenerateHeader(key, method, path, requestBody, secret string) map[string]string {
	ts, sign := makeSign(method, path, requestBody, secret)
	return map[string]string{
		"content-type":     "application/json; charset=UTF-8",
		"ACCESS-KEY":       key,
		"ACCESS-TIMESTAMP": ts,
		"ACCESS-SIGN":      sign,
		"Accept-Encoding":  "gzip",
	}
}

func makeSign(method, path, requestBody, secret string) (string, string) {
	ts := strconv.FormatInt(time.Now().Unix(), 10)
	text := ts + method + path + requestBody

	// API secretでSHA256 署名を生成する
	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write([]byte(text))
	sign := hex.EncodeToString(hash.Sum(nil))
	return ts, sign
}
