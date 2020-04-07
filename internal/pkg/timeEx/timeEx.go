package timeEx

import "time"

func OneDaySeconds()int{
	return 86400
}

func YearDay()int{
	now := time.Now()
	return now.YearDay()
}

func Today()int{
	now := time.Now()
	return now.Year() * 10000 + int(now.Month()) * 100 + now.Day()
}

func TodayTimestamp()int64{
	now := time.Now()
	timeString := now.Format("20060102")
	nowWithoutHMS, _ := time.ParseInLocation("20060102", timeString, now.Location())

	return nowWithoutHMS.Unix()
}

func YestodayTimestamp() int64{
	return TodayTimestamp() - 86400
}

