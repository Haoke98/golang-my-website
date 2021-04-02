package video

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sadam.com/m/myUtil"
)

func VideoHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = VideoGET(w, r)
		break
	case "POST":
		break
	case "PUT":
		break
	case "DELETE":
		break
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func VideoGET(w http.ResponseWriter, r *http.Request) (err error) {
	fmt.Println("this is VideoGet, the Request's url is :", r.URL)
	r.ParseForm()
	jsCode := r.Form["jsCode"][0]
	targetUri := "https://api.weixin.qq.com/sns/jscode2session"

	resp, err := http.PostForm(targetUri, url.Values{
		"appid":      {"wx3723124dbb36e3eb"},
		"secret":     {"7b336b4fc0d26313fd848581c5e818af"},
		"grant_type": {"authorization_code"},
		"js_code":    {jsCode},
	})
	if err == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			result := map[string]interface{}{}
			json.Unmarshal(body, &result)
			myUtil.BeautyConsolePrint(result)
			openId := result["openid"]
			sessionKey := result["session_key"]
			fmt.Println(openId, sessionKey)
			b, err := json.Marshal(openId)
			if err == nil {
				w.Write(b)
			} else {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
	} else {
		fmt.Println(err)
	}
	return err
}
