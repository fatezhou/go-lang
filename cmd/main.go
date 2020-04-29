package main

import (
	"context"
	"fmt"
	"github.com/fatezhou/go-lang/pkg/rpc"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

var BuildTime = ""
var Debug = ""

var (
	wg sync.WaitGroup
)

func work(ctx context.Context){
	defer wg.Done()
	for{
		select{
		case <-time.After(1 * time.Second):{
			fmt.Printf("work, timeEx.after\n")
		}
		case <-ctx.Done():{
			fmt.Printf("done")
			return
		}
		}
	}
}

func httpGet(){
	client := &http.Client{}
	postData := map[string]string{
		"topic": "fish_coin_topN",
		"project": "fish",
		"username": "fish_dev",
		"password": "wwNSl^g$w4VG",
		"plat": "weile",
		"limit": "50",
		"begintime": "1576944000",
		"endtime": "1577030400",
		"order": "p200000001 desc",
	}
	data := url.Values{}
	for k, v := range postData{
		data.Set(k, v)
	}

	req, err := http.NewRequest("POST", "http://192.168.67.13:5020/api/getlog",
		strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "*/*")
	httpResp, err := client.Do(req)
	if err != nil{

	}
	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)
	fmt.Printf("%v", string(body))
}

func main(){
	rpc.InitRPC("0.0.0.0:5090")
}
