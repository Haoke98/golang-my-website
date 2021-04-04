package video

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sadam.com/m/myUtil"
	"strconv"
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
	idStr := r.Form["id"][0]
	id, err := strconv.Atoi(idStr)
	if err == nil {
		video := GetVideoById(id)
		log.Println(video)
	} else {
		log.Println(err)
	}
	targetUri := "https://api.weixin.qq.com/sns/jscode2session"

	resp, err := http.PostForm(targetUri, url.Values{
		"appid":      {"wx3723124dbb36e3eb"},
		"secret":     {"7b336b4fc0d26313fd848581c5e818af"},
		"grant_type": {"authorization_code"},
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
				log.Print(err)
			}
		} else {
			log.Print(err)
		}
	} else {
		fmt.Println(err)
	}
	return err
}
