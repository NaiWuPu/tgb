package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// 反射读取数据
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
}

type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(fileName string, data interface{}) (err error) {
	t := reflect.TypeOf(data)
	//fmt.Println(t, t.Kind())
	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		err = errors.New("data param should be a pointer")
		return
	}

	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	// 将文件内容转化为租房因此

	lineSlice := strings.Split(string(b), "\r\n")
	//fmt.Printf("%#v\n", lineSlice)

	var structName string
	// 一行一行读数据
	for idx, line := range lineSlice {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		if strings.HasPrefix(line, "[") {
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if (line[0] == '[' && line[len(line)-1] != ']') || len(sectionName) == 0 {
				err = fmt.Errorf("line :%d sybtax error", idx+1)
				return
			}
			// 根据字符串sectionName 去data里面根据反射找到对应的结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					structName = field.Name
					//fmt.Printf("找到%s对应的嵌套结构体%s", structName, structName)
					break
				}
			}
		} else {
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line :%d sybtax error", idx+1)
				return
			}

			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])

			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName)
			sType := sValue.Type()

			//structObj := v.Elem().FieldByName(structName)

			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("data 中的 %s 字段应该是一个结构体", structName)
			}
			var fieldName string
			var fileType reflect.StructField

			for i := 0; i < sValue.NumField(); i++ {
				field := sType.Field(i)
				fileType = field
				if field.Tag.Get("ini") == key {
					fieldName = field.Name
					break
				}
			}
			if len(fieldName) == 0 {
				continue
			}

			fileObj := sValue.FieldByName(fieldName)
			//fmt.Println(fieldName, fileType.Type.Kind())

			switch fileType.Type.Kind() {
			case reflect.String:
				fileObj.SetString(value)
			case reflect.Int, reflect.Int64, reflect.Int8, reflect.Int32, reflect.Int16:
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					return
				}
				fileObj.SetInt(valueInt)
			default:

			}
		}
	}

	return
}

func main() {
	var cfg Config
	err := loadIni("./conf.ini", &cfg)
	if err != nil {
		fmt.Printf("load ini failed, err:%v\n", err)
		return
	}
	fmt.Println(cfg)
}
