/*
 * @Author: DiheChen
 * @Date: 2021-09-15 00:27:16
 * @LastEditTime: 2021-09-15 00:34:37
 * @LastEditors: DiheChen
 * @Description: None
 * @GitHub: https://github.com/DiheChen
 */
package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var AuthKey = ""

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST("/x/v1/search", handle)
	_ = router.Run(":7777")
}

func handle(context *gin.Context) {
	recv := make(map[string]interface{})
	_ = context.BindJSON(&recv)
	recv["ts"] = time.Now().Unix()
	postJSON, _ := json.Marshal(recv)
	req, _ := http.NewRequest("POST", "https://api.pcrdfans.com/x/v1/search", bytes.NewReader(postJSON))
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.63 Safari/537.36")
	req.Header.Set("Authorization", AuthKey)
	resp, _ := (&http.Client{}).Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	result := make(map[string]interface{})
	_ = json.Unmarshal([]byte(body), &result)
	context.JSON(200, result)
}
