package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type TKContent struct {
	Content string `json:"content"`
}

func content2Options(content string) (resList []string, err error) {
	var tKContentList []*TKContent
	err = json.Unmarshal([]byte(content), &tKContentList)
	if err != nil {
		return
	}
	for _, c := range tKContentList {
		oContent, err := replaceStr(c.Content)
		if err != nil {
			return
		}
		if oContent == "" {
			err = fmt.Errorf("error1")
			return
		}
		resList = append(resList, oContent)
	}
	if len(resList) < 4 {
		err = fmt.Errorf("error2")
	}
	return
}

func replaceStr(str string) (resStr string, err error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(str))
	if err != nil {
		return
	}
	resStr = doc.Text()
	return
}

func main() {
	content2Options(`[{"content":"A","knowledge_id":0,"video_id":0},{"content":"大","knowledge_id":0,"video_id":0},{"content":"12","knowledge_id":0,"video_id":0},{"content":"奥妙","knowledge_id":0,"video_id":0},{"content":"奥秘","knowledge_id":0,"video_id":0}]`)
}
