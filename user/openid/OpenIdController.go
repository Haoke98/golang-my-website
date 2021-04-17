package openid

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"izbasar.link/web/myUtil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func OpenIdHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = OpenIdGET(w, r)
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

func OpenIdGET(w http.ResponseWriter, r *http.Request) (err error) {
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
				log.Print(err)
			} else {
				defer db.Close()
				//_,err = db.Exec()
				var userId uint64
				rows, err := db.Query("SELECT id FROM miniProgram_user where openid=?", openId)
				if err == nil {
					fmt.Println("everything is good until here.", time.Now())
					//nowString := time.Now().Format("2006-01-02 15:04:05.5")
					//fmt.Println(nowString)
					//nowString = time.Now().Format("2006-01-02 15:04:05.05")
					//fmt.Println(nowString)
					//nowString = time.Now().Format("2006-01-02 15:04:05.000")
					//fmt.Println(nowString)
					//nowString = time.Now().Format("2006-01-02 15:04:05.999")
					//fmt.Println(nowString)
					//nowString = time.Now().Format("2006-01-02 15:04:05.0")
					//fmt.Println(nowString)
					//nowString = time.Now().Format("2006-01-02 15:04:05.9")
					//fmt.Println(nowString)
					//nowString = time.Now().Format("2006-01-02 15:04:")
					//fmt.Println(nowString)
					nowString := time.Now().String()[0:27]
					fmt.Println(nowString)

					var i = 0
					for rows.Next() {
						err = rows.Scan(&userId)
						if err != nil {
							log.Print(err)
						} else {
							fmt.Println("the id is :", userId)
							//"UPDATE table_name SET field1=new-value1, field2=new-value2\n[WHERE Clause]"
							_, err := db.Exec("UPDATE miniProgram_user SET last_login_time=? WHERE id=?", time.Now(), userId)
							if err == nil {
								fmt.Println("更新一个openid为", openId, "的user成功！")
							} else {
								fmt.Println("更新一个openid为", openId, "的user失败！")
							}
							i++
						}
					}
					if i == 0 {
						_, err := db.Exec("INSERT INTO miniProgram_user(openid,firstTimeLogin,last_login_time,last_changed_time) VALUE(?,?,?,?)", openId, time.Now(), time.Now(), time.Now())
						if err == nil {
							fmt.Println("新建一个openid为", openId, "的user成功！")
						} else {
							fmt.Println("新建一个openid为", openId, "的user成功！")
							log.Print(err)
						}
					} else if i > 0 {
						fmt.Println("已经找到了并进行更新了，没必要插入")
					}
				} else {
					log.Print(err)
				}
			}

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

func GetOpenIdHandleFunc(c *gin.Context) {
	jsCode := c.DefaultQuery("jsCode", "")
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
				log.Print(err)
			} else {
				defer db.Close()
				//_,err = db.Exec()
				var userId uint64
				rows, err := db.Query("SELECT id FROM miniProgram_user where openid=?", openId)
				if err == nil {
					fmt.Println("everything is good until here.", time.Now())
					nowString := time.Now().String()[0:27]
					fmt.Println(nowString)

					var i = 0
					for rows.Next() {
						err = rows.Scan(&userId)
						if err != nil {
							log.Print(err)
						} else {
							fmt.Println("the id is :", userId)
							//"UPDATE table_name SET field1=new-value1, field2=new-value2\n[WHERE Clause]"
							_, err := db.Exec("UPDATE miniProgram_user SET last_login_time=? WHERE id=?", time.Now(), userId)
							if err == nil {
								fmt.Println("更新一个openid为", openId, "的user成功！")
							} else {
								fmt.Println("更新一个openid为", openId, "的user失败！")
							}
							i++
						}
					}
					if i == 0 {
						_, err := db.Exec("INSERT INTO miniProgram_user(openid,firstTimeLogin,last_login_time,last_changed_time) VALUE(?,?,?,?)", openId, time.Now(), time.Now(), time.Now())
						if err == nil {
							fmt.Println("新建一个openid为", openId, "的user成功！")
						} else {
							fmt.Println("新建一个openid为", openId, "的user成功！")
							log.Print(err)
						}
					} else if i > 0 {
						fmt.Println("已经找到了并进行更新了，没必要插入")
					}
				} else {
					log.Print(err)
				}
			}

			b, err := json.Marshal(openId)
			if err == nil {
				c.String(http.StatusOK, "%s", b)
			} else {
				log.Print(err)
			}
		} else {
			log.Print(err)
		}
	} else {
		fmt.Println(err)
	}
}
