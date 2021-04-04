package video

import (
	"log"
	"time"
)

type Video struct {
	Id              int
	Name            string
	Url             string
	LastChangedTime time.Time
	ShowTimes       int
	EpisodeNum      int
	Vid             string
	BelongToId      int
	CoverId         int
}

func (v *Video) SetLastChangedTime(timeStr string) {
	DefaultTimeLoc := time.Local
	_time, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, DefaultTimeLoc)
	if err == nil {
		v.LastChangedTime = _time
	} else {
		log.Print("An err occurred when set last changed time to the Video Object:", err)
	}
}

func (v *Video) String() string { return "" }
