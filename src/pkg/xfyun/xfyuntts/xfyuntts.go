package xfyuntts

import (
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"

	"gin-vue-admin/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "tts-api.xfyun.cn:80", "wss service address")

// TTSRequest 方法
func TTSRequest(c *gin.Context, text string) {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "wss", Host: *addr, Path: "/v2/tts"}
	log.Printf("connecting to %s", u.String())

	w, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer w.Close()

	done := make(chan struct{})

	response.SuccessResultWithEmptyData(c)

	go func() {
		defer close(done)
		for {
			_, message, err := w.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()
}
