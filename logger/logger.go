package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"izbasar.link/web/myUtil"
)

func Log(is ...interface{}) {
	for k, v := range is {
		gin.DefaultWriter.Write([]byte(fmt.Sprintf("		%d%s\n", k, myUtil.BeautifyString(v))))
	}
}
