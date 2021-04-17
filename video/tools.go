package video

import (
	"fmt"
	"izbasar.link/web/logger"
	"izbasar.link/web/myUtil/httpHelper"
	"net/http"
)

func GetOfficialAccountVideoPureUlr(vid string) (pureUrl string) {
	pureUrl = ""
	url := fmt.Sprintf("https://mp.weixin.qq.com/mp/videoplayer?action=get_mp_video_play_url&__biz=&mid=&idx=&vid=%s&token=&lang=zh_CN&f=json&ajax=1", vid)
	resp, err := http.Get(url)
	if err != nil {
		logger.Log("An error has occurred when request the mp.weixin.qq.com to get the url of the video", err)
	} else {
		result := httpHelper.ParseBody(resp.Body)
		//myUtil.BeautyConsolePrint(result)
		logger.Log(result)
		urlInfo := result["url_info"]
		i := urlInfo.([]interface{})
		high := i[0]
		j := high.(map[string]interface{})
		k := j["url"].(string)
		pureUrl = k
		return
	}
	return
}
