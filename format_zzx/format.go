package format_zzx

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

// formatZzx 格式化zzx文件函数
func formatZzx(filePath string) map[string]string {
	// 打开文件
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("read file failed, err:", err)
		panic(err)
	}
	// 使用“;”分割开
	contentSlice := strings.Split(string(content), ";")
	// 创建相应的map
	contentMap := make(map[string]string, len(contentSlice))
	for _, v := range contentSlice {
		if strings.Contains(v, "=") {
			tmp := strings.Split(v, "=")
			contentMap[tmp[0]] = tmp[1]
		}
	}
	return contentMap
}

// LoadZzx 加载zzx文件函数，参数为结构体指针，zzx文件地址
func LoadZzx(zzx interface{}, filePath string) {
	contentMap := formatZzx(filePath)
	tmp := reflect.TypeOf(zzx).Elem()
	for i := 0; i < tmp.NumField(); i++ {
		field := tmp.Field(i)
		value, _ := contentMap[field.Tag.Get("zzx")]
		reflect.ValueOf(zzx).Elem().FieldByName(field.Name).SetString(value)
	}
}
