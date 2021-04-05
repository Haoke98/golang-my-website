package myUtil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v6/esapi"
	"io/ioutil"
	"log"
	"os"
	"sadam.com/m/myUtil/httpHelper"
)

func BeautyConsolePrint(s interface{}) {
	fmt.Println(BeautifyString(s))
}
func BeautifyString(s interface{}) string {
	out := beautify(s)
	return out.String()
}
func BeautySaving(s interface{}, fileName string, mode os.FileMode) {
	out := beautify(s)
	err := ioutil.WriteFile(fileName, out.Bytes(), mode)
	if err != nil {
		panic(err)
	}
}
func beautify(s interface{}) bytes.Buffer {
	b, err := httpHelper.JSONMarshal(s)
	if err != nil {
		log.Fatalln(err)
	}

	var out bytes.Buffer

	err = json.Indent(&out, b, "", "\t")

	if err != nil {
		panic(err)
		return out
	} else {
		return out
	}
}
func BeautifyResponse(res *esapi.Response) string {
	var body interface{} = make(map[string]interface{})
	buf, err := ioutil.ReadAll(res.Body)
	if err == nil {
		err := json.Unmarshal(buf, &body)
		if err == nil {
			return BeautifyString(body) + "\n"
		}
	}
	return ""
}
