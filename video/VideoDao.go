package video

import (
	"database/sql"
	"log"
)

func GetVideoById(id int) Video {
	db, err := sql.Open("mysql", "root:qwer1234@tcp(139.155.30.83:3306)/izbasar?charset=utf8")
	video := Video{Id: id}
	if err != nil {
		log.Print("An error has been occurred when connecting to the database server:", err)
		return video
	} else {
		defer db.Close()
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
}
