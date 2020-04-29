package tools

import "strconv"

func Int32ToBool(number int32)bool{
	if number == 0{
		return false
	}else{
		return true
	}
}

func Str2Int64(str string)int64{
	number, err := strconv.Atoi(str)
	if err != nil{
		return 0
	}
	return int64(number)
}

func Int2Str(number int32)string{
	return strconv.FormatInt(int64(number), 10)
}

func Int64ToStr(number int64)string{
	return strconv.FormatInt(number, 10)
}

func Str2Byte(str string)[]byte{
	return []byte(str)
}

func Byte2Str(bytes []byte)string{
	return string(bytes[:])
}

func Float64ToStrring(num float64) string{
	return  strconv.FormatFloat(num ,'E',-1,64)
}