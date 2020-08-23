package timer

import "time"


const (
	Nanosecond  = 1
	Microsecond = 1000 * Nanosecond
	Millisecond = 1000 * Microsecond
	Second = 1000 * Millisecond
	Minute = 60 * Second
	Hour = 60 * Minute
)

//当前时间戳
func GetNowTime() time.Time  {
	location,_ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(location)
}

func GetCalculateTime(currentTimer time.Time,d string) (time.Time,error)  {
	//从字符串中解析出来duration
	duration,err := time.ParseDuration(d)

	if err != nil {
		return time.Time{},err
	}

	//加上解析的时间
	return currentTimer.Add(duration),nil
}

