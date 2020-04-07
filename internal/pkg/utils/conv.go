package utils


import (
	"fmt"
	"strconv"
)

func Int2Str(number int)string{
	return fmt.Sprintf("%d", number)
}

func Str2Int(number string)int{
	n, _ := strconv.Atoi(number)
	return n
}

func Int322Str(number int32)string{
	return fmt.Sprintf("%d", number)
}

func Str2Int32(number string)int32{
	n, _ := strconv.ParseInt(number, 10, 32)
	return int32(n)
}

func Int642Str(number int64)string{
	return fmt.Sprintf("%d", number)
}

func Str2Int64(number string)int64{
	n, _ := strconv.ParseInt(number, 10, 64)
	return n
}
