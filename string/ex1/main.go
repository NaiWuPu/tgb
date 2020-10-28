package main

import (
	"fmt"
)

func main() {
	var s = `
#HasPrefix
语法: strings.HasPrefix(s, prefix string) bool    // 判断字符串s是否以prefix开头

#HasSuffix
语法:strings.HasSuffix(s, suffix string) bool    // 判断字符串s是否以suffix结尾

#Index
语法:strings.Index(s, str string) int    // 判断str在s中首次出现的位置, 如果没有, 则返回-1

#LastIndex
语法:strings.LastIndex(s, str string) int    // 判断str在s中最后一次出现的位置, 如果没有,则返回-1

#Replace
语法:strings.Replace(s, old, new string, n int) string    // 字符串替换

#Count
语法:strings.Count(s, substr string) int    // 字符串计数

#Repeat
语法:strings.Repeat(s string, count int) string    // 重复 count 次 s

#ToLower
语法:strings.ToLower(s string) string    // 全部转为小写

#ToUpper
语法:strings.ToUpper(s string) string    // 全部转为大写

#TrimSpace
语法:strings.TrimSpace(s string) string    // 去掉字符串s的首尾空白字符

#Trim
语法:strings.Trim(s string, cutset string) string    // 去掉字符串s的首尾指定cutse字符

#TrimLeft
语法:strings.TrimLeft(s string, cutset string) string    // 去掉字符串s的首部指定的cutset字符

#TrimRight
语法:
strings.TrimRight(s string, cutset string) string    // 去掉字符串s的尾部指定的cutset字符


#Fields
语法:strings.Fields(s string) []string    // 返回以 空格 分隔的所有子串slice


#Split
语法:strings.Split(s, sep string) []string    // 返回以 sep 分隔的所有子串slice

#Join
语法:strings.Join(a []string, sep string) string    // 用sep把a中的所有元素链接起来`
fmt.Println(s)
}
