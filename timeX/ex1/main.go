package main

import (
	"fmt"
	"strconv"
	"time"
)

/**
 * 武鑫宇的时间快速获取工具
 * @Author: 武鑫宇
 * @Date: 2020/6/8 10:15
 */

var YearFormat = "2006"
var MonthFormat = "01"
var DateFormat = "02"

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

// 昨天零点时间
func GiveMeYesterDay() int64 {
	var timeStr = time.Now().Format("2006-01-02")
	var Time, _ = time.ParseInLocation("2006-01-02", timeStr, time.Local)
	return Time.AddDate(0, 0, -1).Unix()
}

// 今日零点时间
func GiveMeZeroPoint() int64 {
	var timeStr = time.Now().Format("2006-01-02")
	var Time, _ = time.ParseInLocation("2006-01-02", timeStr, time.Local)
	return Time.Unix()
}

func main() {
	for {
		now := time.Now()
		desc_time := fmt.Sprintf("%d-%d-%d %d:%d:%d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
		fmt.Println(desc_time)
		fmt.Println(time.Now().Format("2006-1-02 15:4:05"))

		time.Sleep(time.Second)
	}

	//now := time.Now()
	//nextYear, err := time.Parse("2006-01-02","2020-06-23")
	//log.Println(err)
	//log.Println(nextYear)
	//
	//d := nextYear.Sub(now)
	//fmt.Println(d)
}
