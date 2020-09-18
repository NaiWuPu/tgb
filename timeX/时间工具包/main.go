package main

import (
	"fmt"
	"strconv"
	"time"
)

var str = "10:00:00"

/**
 * 武鑫宇的时间快速获取工具
 * @Author: 武鑫宇
 * @Date: 2020/6/8 10:15
 */
const (
	YearFormat  = "2006"
	MonthFormat = "01"
	DateFormat  = "02"
	Format      = "2006-01-02 15:04:05"
	FormatY_M_D = "2006-01-02"
	FormatYMD   = "20060102"
	FormatShort = "2006-1-2 15:4:5"
)

var weekForm = map[string]string{
	"Sunday":    "星期日",
	"Monday":    "星期一",
	"Tuesday":   "星期二",
	"Wednesday": "星期三",
	"Thursday":  "星期四",
	"Friday":    "星期五",
	"Saturday":  "星期六",
}

// 获取指定时间星期几
func TimeToWeek(timeUnix int64) string {
	t := time.Unix(timeUnix, 0)
	return weekForm[t.Weekday().String()]
}

// 获取现在星期几
func TodayToWeek() string {
	t := time.Now()
	return weekForm[t.Weekday().String()]
}

// 获取指定年
func Time2Year(timeUnix int64) int64 {
	i, _ := strconv.Atoi(time.Unix(timeUnix, 0).Format(YearFormat))
	return int64(i)
}

// 获取指定月
func Time2Month(timeUnix int64) int64 {
	i, _ := strconv.Atoi(time.Unix(timeUnix, 0).Format(MonthFormat))
	return int64(i)
}

// 获取指定日
func Time2Date(timeUnix int64) int64 {
	i, _ := strconv.Atoi(time.Unix(timeUnix, 0).Format(DateFormat))
	return int64(i)
}

func Time2Ymd(timeUnix int64) string {
	return time.Unix(timeUnix, 0).Format(FormatY_M_D)
}

// 昨天零点时间
func GiveMeYesterDay() int64 {
	var timeStr = time.Now().Format(FormatYMD)
	var Time, _ = time.ParseInLocation(FormatYMD, timeStr, time.Local)
	return Time.AddDate(0, 0, -1).Unix()
}

// 今日零点时间
func GiveMeZeroPoint() int64 {
	var timeStr = time.Now().Format(FormatYMD)
	var Time, _ = time.ParseInLocation(FormatYMD, timeStr, time.Local)
	return Time.Unix()
}

func main() {
	var FeatureLibraryUpTimeRange int64
	var str = "14:50:00"
	var Time, _ = time.ParseInLocation(Format, Time2Ymd(time.Now().Unix())+" "+str, time.Local)
	fmt.Println(Time.Unix())
	fmt.Println(Time.Unix() - time.Now().Unix())
	if Time.Unix() < time.Now().Unix() {
		FeatureLibraryUpTimeRange = Time.Unix() + 86400
	} else {
		FeatureLibraryUpTimeRange = Time.Unix()
	}
	fmt.Println(FeatureLibraryUpTimeRange)
}
