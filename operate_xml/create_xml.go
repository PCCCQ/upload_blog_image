package operate_xml

import (
	"encoding/xml"
)

type MethodCall struct {
	XMLName    xml.Name `xml:"methodCall"`
	MethodName string   `xml:"methodName"`
	Params     Params   `xml:"params"`
}

type Params struct {
	Param []Param `xml:"param"`
}

type Param struct {
	Value Value `xml:"value"`
}

type Value struct {
	String *StringValue `xml:"string"`
	Struct *StructValue `xml:"struct"`
	Base64 *StringValue `xml:"base64"`
}

type StringValue struct {
	Text string `xml:",chardata"`
}

type StructValue struct {
	Member []Member `xml:"member"`
}

type Member struct {
	Name  string `xml:"name"`
	Value Value  `xml:"value"`
}

func (a *AccountInformation) CreateXml(ImgBase64 string) string {
	methodCall := MethodCall{
		MethodName: "metaWeblog.newMediaObject",
		Params: Params{
			Param: []Param{
				{
					Value: Value{String: &StringValue{Text: a.BlogId}},
				},
				{
					Value: Value{String: &StringValue{Text: a.UserName}},
				},
				{
					Value: Value{String: &StringValue{Text: a.Password}},
				},
				{
					Value: Value{
						Struct: &StructValue{
							Member: []Member{
								{
									Name:  "bits",
									Value: Value{Base64: &StringValue{Text: ImgBase64}},
								},
								{
									Name:  "name",
									Value: Value{String: &StringValue{Text: "111.jpg"}},
								},
								{
									Name:  "type",
									Value: Value{String: &StringValue{Text: "image/jpeg"}},
								},
							},
						},
					},
				},
			},
		},
	}

	xmlData, err := xml.MarshalIndent(methodCall, "", "  ")
	if err != nil {
		panic(err)
	}

	xmlData = []byte(xml.Header + string(xmlData))

	//file, err := os.Create("example.xml")
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()
	//
	//_, err = file.Write(xmlData)
	//if err != nil {
	//	fmt.Println("写入文件失败：", err)
	//}
	//
	//fmt.Println("XML文件已创建：example.xml")
	return string(xmlData)
}
