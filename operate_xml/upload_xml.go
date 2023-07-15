package operate_xml

import (
	"io"
	"net/http"
	"strings"
)

func (a *AccountInformation) UploadImg(xmlData string) string {
	url := a.Address

	response, err := http.Post(url, "application/xml", strings.NewReader(xmlData))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}
