package authorization

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"gin-vue-admin/tool"
	"net/http"
	"time"
)

// Signature原始字段由 host，date，request-line三个参数按照格式拼接成，拼接的格式为(\n为换行符,’:’后面有一个空格)
func signature() string {
	APISecret := tool.GetConfig("xfyun.textcorrention.apisecret")
	fmt.Println("secret=", APISecret)
	host := "api.xf-yun.com"
	date := time.Now().UTC().Format(http.TimeFormat)
	requstLine := "POST /v1/private/s9a87e3ec HTTP/1.1"
	strings := "host: " + host + "\ndate: " + date + "\n" + requstLine
	fmt.Println("date=", date)
	fmt.Println("signature strings =", strings)
	return computeHmacSha256(strings, APISecret)
}

// Authorization： 规则api_key="$api_key",algorithm="hmac-sha256",headers="host date request-line",signature="$signature"
func Authorization() string {
	APIKey := tool.GetConfig("xfyun.textcorrention.apikey")
	fmt.Println("APIKey=", APIKey)
	signature := signature()
	fmt.Println("signature=", signature)
	strings := "api_key=\"" + APIKey + "\",algorithm=\"hmac-sha256\"" + ",headers=\"host date request-line\"" + ",signature=\"" + signature + "\""
	fmt.Println("Authorization string: ", strings)
	return base64.StdEncoding.EncodeToString([]byte(strings))
}

func computeHmacSha256(message string, secret string) string {

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))

	sha := h.Sum(nil)

	return base64.StdEncoding.EncodeToString([]byte(sha))
}
