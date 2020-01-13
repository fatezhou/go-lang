package http

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type HttpClient struct{
}

type httpClient struct{
}

type httpsClient struct{
}

func request(method string, url string, body string, header map[string]string, client *http.Client)string{
	req, err := http.NewRequest("POST", url, strings.NewReader(body))
	if err != nil{
		return ""
	}
	for k, v := range header{
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil{
		fmt.Printf("%s", err.Error())
		return ""
	}else{
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil{
			fmt.Printf("%s", err.Error())
		}else{
			return string(body)
		}
	}
	return ""
}

func (hc *httpsClient)Post(url, body string, header map[string]string)string{
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify:true},
	}
	client := &http.Client{Transport:tr}
	return request("POST", url, body, header, client)
}

func (hc *httpClient)Post(url, body string, header map[string]string)string{
	client := &http.Client{}
	return request("POST", url, body, header, client)
}

func (h *HttpClient)Post(url, body string, header map[string]string)string{
	if strings.IndexAny(url, "http://") == 0{
		client := httpClient{}
		return client.Post(url, body, header)
	}
	if strings.IndexAny(url, "http://") == 0{
		client := httpsClient{}
		return client.Post(url, body, header)
	}
	return ""
}
