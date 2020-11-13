package event

import (
	"fmt"
	"testing"
)

func TestGetEvents(t *testing.T) {
	GetEvents("Event", 30)
}

func TestGetLiveXML(t *testing.T) {
	r := GetLiveXML("LiveXml", 3)
	fmt.Println(r)
}
