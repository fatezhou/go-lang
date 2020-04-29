package tools

import (
	"errors"
	"git.jiaxianghudong.com/go/utils"
	"net/url"
)

func DecodeFromDatabase(strEncode string, strKey string) (strDecode string, err error) {
	if strKey == "" {
		strKey = "ea65f1e12D13348E4825168585b6966d"
	}
	strDecode, strErr := utils.AesDecryptStr(strEncode, strKey)
	if strErr != "ok" {
		return "", errors.New(strErr)
	} else {
		return strDecode, nil
	}
}

func UrlDecode(strUrlDecode string) string {
	strText, _ := url.QueryUnescape(strUrlDecode)
	return strText
}
