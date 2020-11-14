package event

import (
	"testing"
)

func TestGetEvents(t *testing.T) {
	GetEvents("Event", 30)
}



func TestGetLiveXML(t *testing.T) {
	r := GetLiveXML("LiveXml", 900)
	WriteFile("livexml.csv", r)

}
