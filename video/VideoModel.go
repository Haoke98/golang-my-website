package video

import (
	"log"
	"strings"
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
func (v *Video) updateChangedTime() {
	v.LastChangedTime = time.Now()
}
func (v *Video) Save() {
	v.updateChangedTime()
	Save(v)
}

func (v *Video) UpdateShowTimes() {
	v.ShowTimes++
	v.Save()
}
func (v *Video) GetPureUrl() string {
	if strings.Contains(v.Vid, "wxv_") {
		log.Println("this vide comes from official account space.")
		return GetOfficialAccountVideoPureUlr(v.Vid)
	} else {
		log.Println("this vide comes from Tencent TV space.")
		v.Vid = strings.Replace(v.Vid, "'", "", 1)
		return v.Url
	}
}
