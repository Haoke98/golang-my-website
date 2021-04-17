package video

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"izbasar.link/web/logger"
	"log"
	"net/http"
	"strconv"
)

func VideoHandlerFunc(c *gin.Context) {

}

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

func GetHandleFunc(c *gin.Context) {
	idStr := c.DefaultQuery("id", "0")
	id, err := strconv.Atoi(idStr)
	if err == nil {
		video := GetVideoById(id)
		video.UpdateShowTimes()
		pureUrl := video.GetPureUrl()
		//c.JSON(http.StatusOK,gin.H{"video":video,"pureUrl":pureUrl})
		// http.StatusTemporaryRedirect  临时重定向
		// http.StatusMovedPermanently  永久重定向
		logger.Log(pureUrl)
		//c.Redirect(http.StatusTemporaryRedirect,pureUrl)
		http.Redirect(c.Writer, c.Request, pureUrl, http.StatusTemporaryRedirect)
	} else {
		log.Println(err)
	}
}
