package main

import (
	"cnblogs_upload/format_zzx"
	"cnblogs_upload/img_base64"
	"cnblogs_upload/operate_xml"
	"fmt"
	"os"
	"strings"
)

func main() {
	var account operate_xml.AccountInformation
	//format_zzx.LoadZzx(&account, "config.zzx")
	//fmt.Println("请输入上传图片地址：")
	//var img string
	//_, err := fmt.Scan(&img)
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			if index == 0 {
				i := strings.LastIndex(arg, string(os.PathSeparator))
				arg = arg[:i]
				format_zzx.LoadZzx(&account, arg+"\\config.zzx")
				//fmt.Println()
			} else {
				imgBase64, err := img_base64.ImgToBase64(arg)
				if err != nil {
					return
				}
				requestXml := account.CreateXml(imgBase64)
				if err != nil {
					return
				}
				//fmt.Println(requestXml)
				//fmt.Println(string(resXml))
				responseXml := account.UploadImg(requestXml)
				//fmt.Println(responseXml)
				ImgUrl := operate_xml.AnalysisUrl(responseXml)
				fmt.Println(ImgUrl)
				//fmt.Println(ImgUrl)
				//time.Sleep(time.Second * 60)
				//fmt.Println(string(resXml))
			}
		}
	}

}
