package utils

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func GetCurrTime() int64 {
	return time.Now().UnixNano() / 1e6
}

func RandBetween(min, max int32) int32 {
	if min >= max || max == 0 {
		return max
	}
	return rand.Int31n(max-min+1) + min
}

func Now() int64 {
	return time.Now().Unix()
}

func Today() int32 {
	now := time.Now()
	str := fmt.Sprintf("%04d%02d%02d", now.Year(), now.Month(), now.Day())
	return Str2Int32(str)
}

func NowTimeString() string{
	now := time.Now()
	str := fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	return str
}

func GetWeekDay() int32 {
	now := time.Now()
	return int32(now.Weekday()) //周日到周六, 0到6
}

func GetThisMonday() int32 {
	now := time.Now()
	return int32(now.YearDay()) - int32(now.Weekday()) + 1 //取最近一个周一是这一年内的第几天
}

// GetPublicIP 获取公网IP
func GetPublicIP() string {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	return localAddr[0:idx]
}

// GetInternalIP 获取内网IP
func GetInternalIP() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func GetUrl(url string) ([]byte, error) {
	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

