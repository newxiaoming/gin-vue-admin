package tool

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpPost(url string, contentType string, reqBody string) string {
	client := http.Client{}
	resp, err := client.Post(url, contentType, strings.NewReader(reqBody))

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	return string(body)
}
