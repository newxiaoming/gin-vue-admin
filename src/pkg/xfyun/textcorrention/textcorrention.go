package textcorrention

import (
	"encoding/base64"
	"encoding/json"
	"errors"
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

type textData struct {
	Char   [][]interface{} `json:"char"`
	Word   [][]interface{} `json:"word"`
	Redund [][]interface{} `json:"redund"`
	Miss   [][]interface{} `json:"miss"`
	Order  [][]interface{} `json:"order"`
	Dapei  [][]interface{} `json:"dapei"`
	Punc   [][]interface{} `json:"punc"`
	Idm    [][]interface{} `json:"idm"`
	Org    [][]interface{} `json:"org"`
	Leader [][]interface{} `json:"leader"`
	Number [][]interface{} `json:"number"`
}

type textDataItem struct {
	Pos         int
	Cur         string
	Corrent     string
	Description string
}

type textDataItemJSON struct {
	*textDataItem
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
	fmt.Println(result)
	if data, err = formatData([]byte(result)); err != nil {
		fmt.Println("PostData:", err)
	}
	data.Text = text

	return data
}

// formatData 格式化返回的数据
func formatData(data []byte) (response.TextCorrentionResponseItem, error) {
	var items response.TextCorrentionResponseItem
	var lists []map[string]interface{}

	if err := json.Unmarshal(data, &respData); err != nil {
		return items, errors.New("responseData is no json")
	}

	if respData.Header.Code != 0 {
		return items, errors.New("request error")
	}

	textBase64Data, err := base64.StdEncoding.DecodeString(respData.Payload.Result.Text)
	fmt.Println("textBase64Data:", string(textBase64Data))
	if err != nil {
		return items, errors.New("responseData text is no base64")
	}
	fmt.Println("====")
	fmt.Println(string(textBase64Data))
	fmt.Println("====")
	if err := json.Unmarshal(textBase64Data, &textOrigData); err != nil {
		fmt.Println("textOrigData error :", err)
		return items, errors.New("textOrigData is no json")
	}

	fmt.Println("------------------")
	textDataItems := &textData{}
	err = json.Unmarshal(textBase64Data, textDataItems)
	fmt.Println(textDataItems)
	// if len(textDataItems.Char) > 0 {
	// 	for _, value := range textDataItems.Char {
	// 		item := map[string]interface{}{
	// 			"OriFrag":     tool.Strval(value[1]),
	// 			"BeginPos":    tool.Strval(value[0]),
	// 			"CorrectFrag": tool.Strval(value[2]),
	// 			"EndPos":      "0",
	// 		}
	// 		lists = append(lists, item)
	// 		// err := mapstructure.Decode(lists, &items.VecFragment)
	// 		// if err != nil {
	// 		// 	fmt.Println(err)
	// 		// }
	// 	}
	// }

	itemsChar := getItem(textDataItems.Char)
	itemsWord := getItem(textDataItems.Word)
	itemsRedund := getItem(textDataItems.Redund)
	itemsMiss := getItem(textDataItems.Miss)
	itemsOrder := getItem(textDataItems.Order)
	itemsDapei := getItem(textDataItems.Dapei)
	itemsPunc := getItem(textDataItems.Punc)
	itemsIdm := getItem(textDataItems.Idm)
	itemsOrg := getItem(textDataItems.Org)
	itemsLeader := getItem(textDataItems.Leader)
	itemsNumber := getItem(textDataItems.Number)

	lists = append(lists, itemsChar, itemsWord, itemsRedund, itemsMiss, itemsOrder, itemsDapei, itemsPunc, itemsIdm, itemsOrg, itemsLeader, itemsNumber)

	fmt.Println(lists)
	d, _ := json.Marshal(items)
	fmt.Println(string(d))
	fmt.Println("------------------")

	fmt.Println(items)

	return items, nil
}

func getItem(inputItem [][]interface{}) []map[string]interface{} {
	var lists []map[string]interface{}
	item := map[string]interface{}{}
	// items := response.TextCorrentionResponseItem{}
	if len(inputItem) > 0 {
		for _, value := range inputItem {
			item = map[string]interface{}{
				"OriFrag":     tool.Strval(value[1]),
				"BeginPos":    tool.Strval(value[0]),
				"CorrectFrag": tool.Strval(value[2]),
				"EndPos":      "0",
			}

			lists = append(lists, item)

			// err := mapstructure.Decode(lists, &items.VecFragment)
			// if err != nil {
			// 	fmt.Println(err)
			// }
		}
	}
	// else {
	// 	items.VecFragment = make([]map[string]interface{}, 0)
	// }
	fmt.Println("==")
	fmt.Println(item)
	// fmt.Println(items)
	fmt.Println("==")
	return lists
}
