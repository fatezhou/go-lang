package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"zoyee-tool/internal/pkg/http"
)

var BuildTime = ""
var Debug = ""

func main(){
	str := "https://12345"
	fmt.Print(strings.IndexAny(str, "https://"))
	c := http.HttpClient{}
	body := make(map[string]interface{})
	bodyData := make(map[string]string)
	bodyData["token"] = "1234567"
	body["data"] = bodyData
	byte, _ := json.Marshal(&body)
	str = c.Post("https://queue1.youyueworld.com/apis/random_url", string(byte), nil)
	fmt.Printf("%s", str)
}
