package tools

import (
	"fmt"
	"github.com/fatezhou/go-lang/pkg/logs"
	"os"
	"strconv"
	"time"
)

func LoadFile(path string) []byte {
	logs.Debugf("%s", path)
	fp, err := os.Open(path) // 获取文件指针
	if err != nil {
		panic(fmt.Sprintf("no found databin error:%+v", err))
		return nil
	}
	defer fp.Close()

	fileInfo, err := fp.Stat()
	if err != nil {
		panic(fmt.Sprintf("Stat databin error:%+v", err))
		return nil
	}
	buffer := make([]byte, fileInfo.Size())
	_, err = fp.Read(buffer) // 文件内容读取到buffer中
	return buffer
}

func Str2Int32(str string) int32 {
	n, err := strconv.Atoi(str)
	if err != nil {
		return 0
	} else {
		return int32(n)
	}
}

func ValueToIP(databaseValue int64) string {
	//3232239295 = 192.168.14.191
	strHex := fmt.Sprintf("%08x", databaseValue)
	var ip1, ip2, ip3, ip4 uint64
	ip1, _ = strconv.ParseUint(strHex[0:2], 16, 32)
	ip2, _ = strconv.ParseUint(strHex[2:4], 16, 32)
	ip3, _ = strconv.ParseUint(strHex[4:6], 16, 32)
	ip4, _ = strconv.ParseUint(strHex[6:8], 16, 32)
	return fmt.Sprintf("%d.%d.%d.%d", ip1, ip2, ip3, ip4)
}

func Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func StrToTimestamp(strTime string) (int64, string) {
	now, _ := time.ParseInLocation("2006-01-02 15:04:05", strTime, time.Local)
	return now.Unix(), Int64ToStr(now.Unix())
}
