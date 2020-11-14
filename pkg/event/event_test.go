package event

import (
	"testing"
)

func TestGetEvents(t *testing.T) {
	GetEvents("Event", 30)
}



func TestGetLiveXML(t *testing.T) {
	ds := NewDS("livexml.csv")
	ds.GetLiveXML("LiveXml", 900)


}
