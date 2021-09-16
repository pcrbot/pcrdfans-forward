/*
 * @Author: DiheChen
 * @Date: 2021-09-15 00:27:16
 * @LastEditTime: 2021-09-16 23:50:08
 * @LastEditors: DiheChen
 * @Description: None
 * @GitHub: https://github.com/DiheChen
 */
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var AuthKey = ""

func main() {
	http.HandleFunc("/x/v1/search", Handle)
	_ = http.ListenAndServe("0.0.0.0:7777", nil)
}

func Handle(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Host, request.Method, request.Proto)
	if request.Method != "POST" {
		return
	}
	recv := make(map[string]interface{})
	_ = json.NewDecoder(request.Body).Decode(&recv)
	recv["ts"] = time.Now().Unix()
	postJSON, _ := json.Marshal(&recv)
	req, _ := http.NewRequest("POST", "https://api.pcrdfans.com/x/v1/search", bytes.NewReader(postJSON))
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.63 Safari/537.36")
	req.Header.Set("Authorization", AuthKey)
	resp, _ := (&http.Client{}).Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(body)
}
