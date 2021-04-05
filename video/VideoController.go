package video

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = GET(w, r)
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

func GET(w http.ResponseWriter, r *http.Request) (err error) {
	fmt.Println("this is VideoGet, the Request's url is :", r.URL)
	err = r.ParseForm()
	if err != nil {
		log.Println("An error has occurred when the parsing the Form:", err)
	} else {
		idStr := r.Form["id"][0]
		id, err := strconv.Atoi(idStr)
		if err == nil {
			video := GetVideoById(id)
			log.Println(video)
			video.UpdateShowTimes()
			pureUrl := video.GetPureUrl()
			log.Println(pureUrl)
			http.Redirect(w, r, pureUrl, http.StatusTemporaryRedirect)
		} else {
			log.Println(err)
		}
	}
	return err
}
