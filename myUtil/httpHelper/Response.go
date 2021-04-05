package httpHelper

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func ParseBody(resp *http.Response) map[string]interface{} {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("An error has occurred when parsing the response body:", err)
	} else {
		result := map[string]interface{}{}
		json.Unmarshal(body, &result)
		//data,err := JSONMarshal(result)
		//if err!=nil{
		//	log.Println("An error has occurred when turn teh interface to the bytes:", err)
		//}else{
		//	json.Unmarshal(data,&result)
		//}
		return result
	}
	return nil
}
func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

func TransHtmlJson(data []byte) []byte {
	data = bytes.Replace(data, []byte("\\u0026"), []byte("&"), -1)
	data = bytes.Replace(data, []byte("\\u003c"), []byte("<"), -1)
	data = bytes.Replace(data, []byte("\\u003e"), []byte(">"), -1)
	return data
}
