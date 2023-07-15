package img_base64

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func ImgToBase64(imagePath string) (encodedImage string, err error) {

	isNetworkImage := strings.HasPrefix(imagePath, "http://") || strings.HasPrefix(imagePath, "https://")

	var imageBytes []byte
	if isNetworkImage {
		// 下载网络图片
		response, err := http.Get(imagePath)
		if err != nil {
			fmt.Println("下载网络图片失败：", err)
			return encodedImage, err
		}
		defer response.Body.Close()

		imageBytes, err = ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("读取网络图片响应失败：", err)
			return encodedImage, err
		}
	} else {
		// 读取本地图片
		imageBytes, err = ioutil.ReadFile(imagePath)
		if err != nil {
			fmt.Println("读取本地图片文件失败：", err)
			return encodedImage, err
		}
	}

	encodedImage = base64.StdEncoding.EncodeToString(imageBytes)
	return encodedImage, err
	//
	//fmt.Println("图片的Base64编码：")
	//fmt.Println(encodedImage)
}
