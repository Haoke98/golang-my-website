package video

import (
	"log"
	"sadam.com/m/database"
)

func GetVideoById(id int) Video {
	video := Video{Id: id}
	db := database.GetDB()
	rows, err := db.Query("SELECT last_changed_time,showTimes,episodeNum,url,vid,belongTo_id,cover_id FROM miniProgram_video where id=?", id)
	if err == nil {
		for rows.Next() {
			var lastChangedTime string
			err = rows.Scan(&lastChangedTime, &video.ShowTimes, &video.EpisodeNum, &video.Url, &video.Vid, &video.BelongToId, &video.CoverId)
			if err != nil {
				log.Print("An error has been occurred when scanning:", err)
			} else {
				video.SetLastChangedTime(lastChangedTime)
				return video
			}
		}
		log.Println("video:", video, " have not been found on the DataBase.")
		return video
	} else {
		log.Print("An error has been occurred when the querying:", err)
		return video
	}
}

/**
默认是UPDATE模式，INSERT只会在初次创建时使用。
*/
func Save(video Video) {
	db := database.GetDB()
	_, err := db.Exec("UPDATE miniProgram_video SET last_changed_time=?,showTimes=?,episodeNum=?,url=?,vid=?,belongTo_id=?,cover_id=? WHERE id=?", video.LastChangedTime, video.ShowTimes, video.EpisodeNum, video.Url, video.Vid, video.BelongToId, video.CoverId, video.Id)
	if err != nil {
		log.Print("An error when saving video:", video, "to the Database:", err)
	} else {
		/**
		插入更新成功。
		*/
	}
}
