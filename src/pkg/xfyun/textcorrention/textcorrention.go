package textcorrention

import (
	"encoding/json"
	"fmt"
	"gin-vue-admin/pkg/response"
	"gin-vue-admin/tool"
	"net/url"

	"github.com/gin-gonic/gin"
)

type requestBody struct {
	Header struct {
		AppID  string `json:"app_id"`
		Status int    `json:"status"`
	} `json:"header"`
	Parameter struct {
		S9a87e3ec struct {
			Result struct {
				Encoding string `json:"encoding"`
				Compress string `json:"compress"`
				Format   string `json:"format"`
			} `json:"result"`
		} `json:"s9a87e3ec"`
	} `json:"parameter"`
	Payload struct {
		Input struct {
			Encoding string `json:"encoding"`
			Compress string `json:"compress"`
			Format   string `json:"format"`
			Status   int    `json:"status"`
			Text     string `json:"text"`
		} `json:"input"`
	} `json:"payload"`
}

func PostData(c *gin.Context) string {
	uri := "https://api.xf-yun.com/v1/private/s9a87e3ec?authorization=%s&host=%s&date=%s"
	contentType := "application/json"
	reqURL := fmt.Sprintf(uri, c.GetString("authorization"), c.GetString("host"), url.QueryEscape(c.GetString("date")))
	fmt.Println(reqURL)

	reqBody := &requestBody{}
	reqBody.Header.AppID = tool.GetConfig("xfyun.textcorrention.appid")
	reqBody.Header.Status = 3
	reqBody.Parameter.S9a87e3ec.Result.Encoding = "utf8"
	reqBody.Parameter.S9a87e3ec.Result.Compress = "raw"
	reqBody.Parameter.S9a87e3ec.Result.Format = "json"
	reqBody.Payload.Input.Encoding = "utf8"
	reqBody.Payload.Input.Compress = "raw"
	reqBody.Payload.Input.Format = "json"
	reqBody.Payload.Input.Status = 3
	reqBody.Payload.Input.Text = "5a+55b6F5q+P5LiA6aG55bel5L2c6YO96KaB5LiA5Lid5LiN5aSf77yM5Lu75L2V5LqL5oOF5Y+q6KaB55So5b+D5Y675YGa77yM5oC75Lya5pyJ5omA5pS255uK77yM5L2c5Li65L6b55S15omA5omA6ZW/6Z2i5a+555qE5LqL5oOF5b6I5aSaLOS5n+W+iOe5geeQkO+8jOS9huaIkeS7rOWPquimgeS7pemrmOW6pui0n+i0o+eahOaAgeW6puadpeWBmuWlveavj+S4gOS7tuS6i+aDhe+8jOaIkeS7rOWwseS8muWcqOWwj+S6i+S5i+S4reS9k+aCn+WIsOactOWunuiAjOa3seWIu+eahOmBk+eQhuOAgg=="

	r, err := json.Marshal(reqBody)
	if err != nil {
		response.FailResult(500, err.Error(), c)
	}

	return tool.HttpPost(reqURL, contentType, string(r))
}
