package textcorrention

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"gin-vue-admin/pkg/response"
	"gin-vue-admin/tool"
	"net/url"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
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

type textData struct {
	Char   interface{} `json:"char"`
	Word   interface{} `json:"word"`
	Redund interface{} `json:"redund"`
	Miss   interface{} `json:"miss"`
	Order  interface{} `json:"order"`
	Dapei  interface{} `json:"dapei"`
	Punc   interface{} `json:"punc"`
	Idm    interface{} `json:"idm"`
	Org    interface{} `json:"org"`
	Leader interface{} `json:"leader"`
	Number interface{} `json:"number"`
}

type responseData struct {
	Header struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		SID     string `json:"sid"`
	} `json:"header"`
	Payload struct {
		Result struct {
			Compress string `json:"compress"`
			Encoding string `json:"encoding"`
			Format   string `json:"format"`
			Seq      string `json:"seq"`
			Status   string `json:"status"`
			Text     string `json:"text"`
		} `json:"result"`
	} `json:"payload"`
}

var (
	respData     responseData
	textOrigData textData
)

// PostData 提交数据
func PostData(c *gin.Context, text string) response.TextCorrentionResponseItem {
	var (
		data response.TextCorrentionResponseItem
	)
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
	reqBody.Payload.Input.Text = base64.StdEncoding.EncodeToString([]byte(text))

	r, err := json.Marshal(reqBody)
	if err != nil {
		response.FailResult(500, err.Error(), c)
	}

	result := tool.HttpPost(reqURL, contentType, string(r))
	if data, err = formatData([]byte(result)); err != nil {
		fmt.Println("PostData:", err)
	}
	data.Text = text

	return data
}

// formatData 格式化返回的数据
func formatData(data []byte) (response.TextCorrentionResponseItem, error) {
	var items response.TextCorrentionResponseItem

	if err := json.Unmarshal(data, &respData); err != nil {
		return items, errors.New("responseData is no json")
	}

	if respData.Header.Code != 0 {
		return items, errors.New("request error")
	}

	textBase64Data, err := base64.StdEncoding.DecodeString(respData.Payload.Result.Text)
	fmt.Println("textBase64Data:", string(data))
	if err != nil {
		return items, errors.New("responseData text is no base64")
	}

	if err := json.Unmarshal(textBase64Data, &textOrigData); err != nil {
		fmt.Println("textOrigData error :", err)
		return items, errors.New("textOrigData is no json")
	}

	k := reflect.ValueOf(textOrigData)
	count := k.NumField()
	var lists []map[string]interface{}
	for i := 0; i < count; i++ {
		if k.Field(i).CanInterface() && reflect.ValueOf(k.Field(i).Interface()).Len() > 0 { //判断是否为可导出字段
			for _, value := range k.Field(i).Interface().([]interface{}) {
				item := map[string]interface{}{
					"OriFrag":     tool.Strval(value.([]interface{})[1]),
					"BeginPos":    tool.Strval(value.([]interface{})[0]),
					"CorrectFrag": tool.Strval(value.([]interface{})[2]),
					"EndPos":      "0",
				}
				fmt.Println("item:", item)
				lists = append(lists, item)
				err := mapstructure.Decode(lists, &items.VecFragment)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
	fmt.Println(items)

	return items, nil
}
