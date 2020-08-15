package date

/**
时间相关工具
*/

import (
	"cyan/errors"
	"strings"
	"time"
)

const (
	DATE_FORMAT_6                string = "2006-01"             // yyyy-MM
	DATE_FORMAT_10               string = "2006-01-02"          // yyyy-MM-dd
	DATE_FORMAT_14               string = "20060102150405"      // yyyyMMddHHmmss
	DATE_FORMAT_17               string = "20060102 15:04:05"   // yyyyMMdd HH:mm:ss
	DATE_FORMAT_19               string = "2006-01-02 15:04:05" // yyyy-MM-dd HH:mm:ss
	DATE_FORMAT_19_FORWARD_SLASH string = "2006/01/02 15:04:05" // yyyy/MM/dd HH:mm:ss
)

/**
获取当前时间
默认时间格式:2006-01-02 15:04:05
*/
func Now() string {
	return time.Now().Format(DATE_FORMAT_19)
}

/**
根据指定格式获取当前时间
*/
func NowByFormat(format string) string {
	return time.Now().Format(format)
}

/**
根据传入的time获取日期
默认时间格式:2006-01-02 15:04:05
*/
func NowByTime(time *time.Time) string {
	return time.Format(DATE_FORMAT_19)
}

/**
获取当前毫秒值
*/
func GetMillisecond() int64 {
	return time.Now().UnixNano() / 1e6
}

/**
根据传入的time和指定格式获取日期
默认时间格式:2006-01-02 15:04:05
*/
func NowByTimeAndFormat(time *time.Time, format string) string {
	if "" == format {
		format = DATE_FORMAT_19
	}
	return time.Format(format)
}

/**
根据时间戳和指定时间格式转换字符串
*/
func TimestampFormatToString(timestamp int64, format string) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format(format)
}

/**
根据时间戳转换字符串
默认格式"2006-01-02 15:04:05" yyyy-MM-dd HH:mm:ss
*/
func TimestampDefaultFormatToString(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format(DATE_FORMAT_19)
}

//获取第几周
/**
获取当前时间 是一年当中的第几周
*/
func Week() int {
	_, week := time.Now().ISOWeek()
	return week
}

/**
获取年、月、日
*/
func DateYMD() (int, int, int) {
	timeNow := time.Now()
	year, month, day := timeNow.Date()
	return year, int(month), day
}

/**
获取明天日期
默认格式  "2006-01-02" yyyy-MM-dd
*/
func Tomorrow() string {
	return GetSpecifiedDateByYMD(1)
}

/**
获取昨天日期
默认格式  "2006-01-02" yyyy-MM-dd
*/
func Yesterday() string {
	return GetSpecifiedDateByYMD(-1)
}

/**
根据指定天数获取日期
默认日期格式 默认格式  "2006-01-02" yyyy-MM-dd
例：获取明天日期 GetSpecifiedDateByYMD(1)
例：获取昨天日期 GetSpecifiedDateByYMD(-1)
*/
func GetSpecifiedDateByYMD(day int) string {
	nTime := time.Now()
	yesTime := nTime.AddDate(0, 0, day)
	return yesTime.Format(DATE_FORMAT_10)
}

/**
根据指定天数和指定格式 获取日期
默认日期格式 默认格式  "2006-01-02" yyyy-MM-dd
例：获取明天日期 GetSpecifiedDateByYMD(1)
例：获取昨天日期 GetSpecifiedDateByYMD(-1)
*/
func GetSpecifiedDateByYMDAndFormat(day int, format string) string {
	if "" == format {
		format = DATE_FORMAT_10
	}
	nowTime := time.Now()
	dTime := nowTime.AddDate(0, 0, day)
	return dTime.Format(format)
}

/**
获取当天日期是周几
*/
func Weekday() int {
	t := time.Now()
	return int(t.Weekday())
}

/**
根据时间字符串转换成time类
*/
func StringConverToTime(dateString string) (*time.Time, error) {

	format, ok := switchDateFormat(dateString)
	if !ok {
		return nil, errors.New(999, format)
	}
	parseTime, err := time.Parse(format, dateString)
	if err != nil {
		return nil, err
	}
	return &parseTime, nil
}

func switchDateFormat(dateString string) (string, bool) {
	switch len(dateString) {
	case 6:
		return DATE_FORMAT_6, true
	case 10:
		return DATE_FORMAT_10, true
	case 11:
	case 12:
	case 13:
	case 15:
	case 16:
	case 18:
	default:
		return "can not find date format for：" + dateString, false
	case 14:
		return DATE_FORMAT_14, true
	case 17:
		return DATE_FORMAT_17, true
	case 19:
		if strings.Contains(dateString, "-") {
			return DATE_FORMAT_19, true
		}
		return DATE_FORMAT_19_FORWARD_SLASH, true
	}
	return "can not find date format for：" + dateString, false
}
