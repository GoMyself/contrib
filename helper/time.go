package helper

import "time"

// 月份字符串校验
func CtypeMonth(s string, loc *time.Location) (int64, error) {

	s += "-01 00:00:00"
	t, err := time.ParseInLocation("2006-01-02 15:04:05", s, loc)
	if err != nil {
		return 0, err
	}

	return t.Unix(), nil
}

// 通过时间戳，获取一天的开始时间
// 默认为当天的 00：00：00 时间戳
func DayTST(timestamp int64, loc *time.Location) time.Time {

	t := time.Now().In(loc)
	if timestamp > 0 {
		t = time.Unix(timestamp, 0).In(loc)
	}

	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, loc)
}

// 通过时间戳，获取一天的结束时间
// 默认为当天的 23：59：59 时间戳
func DayTET(timestamp int64, loc *time.Location) time.Time {

	t := time.Now().In(loc)
	if timestamp > 0 {
		t = time.Unix(timestamp, 0).In(loc)
	}

	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, loc)
}

// 通过日期字符串，获取一天的开始时间
// 默认为当天的 00：00：00 时间戳
func DaySST(date string, loc *time.Location) time.Time {

	t := time.Now().In(loc)
	if date != "" {
		t = StrToTime(date, loc)
	}

	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, loc)
}

// 通过时间戳，获取一天的结束时间
// 默认为当天的 23：59：59 时间戳
func DaySET(date string, loc *time.Location) time.Time {

	t := time.Now().In(loc)
	if date != "" {
		t = StrToTime(date, loc)
	}

	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, loc)
}

// 通过时间戳，获取一月的开始时间
// 默认为当前月的第一天 00：00：00 时间戳
func MonthTST(timestamp int64, loc *time.Location) time.Time {

	t := time.Now().In(loc)
	if timestamp > 0 {
		t = time.Unix(timestamp, 0).In(loc)
	}

	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, loc)
}

// 通过时间戳，获取一月的结束时间
// 默认为当前月的最后一天 23：59：59 时间戳
func MonthTET(timestamp int64, loc *time.Location) time.Time {

	t := time.Now().In(loc)
	if timestamp > 0 {
		t = time.Unix(timestamp, 0).In(loc)
	}

	t = time.Date(t.Year(), t.Month(), 1, 23, 59, 59, 999999999, loc)
	return t.AddDate(0, 1, -1)
}

// 通过日期字符串，获取一月的开始时间
// 默认为当前月的第一天 00：00：00 时间戳
func MonthSST(date string, loc *time.Location) time.Time {

	t := time.Now().In(loc)
	if date != "" {
		t = StrToTime(date, loc)
	}

	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, loc)
}

// 通过日期字符串，获取一月的结束时间
// 默认为当前月的最后一天 23：59：59 时间戳
func MonthSET(date string, loc *time.Location) time.Time {

	t := time.Now().In(loc)
	if date != "" {
		t = StrToTime(date, loc)
	}

	t = time.Date(t.Year(), t.Month(), 1, 23, 59, 59, 999999999, loc)
	return t.AddDate(0, 1, -1)
}
