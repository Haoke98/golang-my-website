package video

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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

			db, err := sql.Open("mysql", "root:qwer1234@tcp(139.155.30.83:3306)/izbasar?charset=utf8")
			if err != nil {
				log.Fatal(err)
			} else {
				defer db.Close()
				//_,err = db.Exec()
				var userId uint64
				rows, err := db.Query("SELECT id FROM miniProgram_user where openid=?", openId)
				if err == nil {
					fmt.Println("everything is good until here.")
					var i = 0
					for rows.Next() {
						err = rows.Scan(&userId)
						if err != nil {
							log.Fatal(err)
						} else {
							fmt.Println("the id is :", userId)
							//	TODO:在这里得写更新相关的代码（更新最近一次登录时间）
							i++
						}
					}
					if i == 0 {
						//		TODO：在这里要编写第一次插入相关的代码（第一次登录时间 和 最近一次登录时间，openid等等字段都得插入）
					} else if i > 0 {
						fmt.Println("已经找到了并进行更新了，没必要插入")
					}
				} else {
					log.Fatal(err)
				}
			}

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
