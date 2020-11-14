package event

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"
)

func TestGetEvents(t *testing.T) {
	GetEvents("Event", 30)
}

func WriteFile(file string, r []Raw) {
	f, err := os.Create(file)
	if err != nil {
		return
	}
	defer f.Close()
	w := bufio.NewWriter(f)

	re := regexp.MustCompile(`[0-9]{4}-[0-9]{2}-[0-9]{2} @ [0-9]{2}:[0-9]{2}:[0-9]{2}`)

	fmt.Fprintf(w, "TimeStamp,Title,Desc,Lng,Lag,Postal,Station,AlertTimeStamp\n")
	for _, iv := range r {
		for _, v := range iv.ActiveAlerts.Events {

			t := re.FindAllString(v.Desc, -1)[0]
			timeStamp := strings.Replace(t, " @ ", " ", -1)
			fmt.Fprintf(w, "%s,%s,%q,%s,%s,%s,%s,%s\n", timeStamp, v.Title, v.Desc, v.Lng, v.Lat, v.Postal, v.Station, iv.TimeStamp)

		}
	}

}

func TestGetLiveXML(t *testing.T) {
	r := GetLiveXML("LiveXml", 3000)
	WriteFile("livexml.csv", r)

}
