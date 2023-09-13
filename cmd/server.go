package cmd

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/zu1k/nali/pkg/entity"
	"os"
	"strconv"
	"time"
)

var (
	APIKEY = "1[P6-CV~HhR?mw(+Z"
)

func server(port int) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/v1/ip/info/:ip", func(c *gin.Context) {
		key := c.GetHeader("API-KEY")
		if key != APIKEY {
			c.JSON(200, resp{Code: -1, Msg: "API-KEY错误", Timestamp: time.Now().Unix()})
			return
		}
		ip := c.Param("ip")
		jsonResult := entity.ParseLine(ip).Json()
		var result ipResp
		err := json.Unmarshal([]byte(jsonResult), &result)
		if err != nil {
			c.JSON(200, resp{Code: -1, Msg: "API-KEY错误", Timestamp: time.Now().Unix()})
			return
		}
		result.Result = result.Text
		c.JSON(200, resp{
			Code:      0,
			Msg:       "success",
			Timestamp: time.Now().Unix(),
			Data:      result,
		})
		return
	})
	os.Setenv("PORT", strconv.Itoa(port))
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

type resp struct {
	Code      int    `json:"code,omitempty"`
	Msg       string `json:"msg,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
	Data      ipResp `json:"data"`
}

type ipResp struct {
	Type   int    `json:"type,omitempty"`
	IP     string `json:"ip,omitempty" `
	Text   string `json:"text"`
	Result string `json:"result,omitempty" `
	Source string `json:"source,omitempty" `
	Info   struct {
		Country string `json:"country,omitempty" `
		Area    string `json:"area,omitempty"`
	} `json:"info" json:"info"`
}
