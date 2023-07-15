package operate_xml

import (
	"strings"
)

func AnalysisUrl(body string) string {
	s1 := strings.Split(body, "<string>")[1]
	s2 := strings.Split(s1, "</string>")[0]
	if !(strings.HasPrefix(s2, "http://") || strings.HasPrefix(s2, "https://")) {
		panic("上传图片失败")
	}
	//fmt.Println(strings.HasPrefix(s2, "http://") || strings.HasPrefix(s2, "https://"))
	return s2
}
