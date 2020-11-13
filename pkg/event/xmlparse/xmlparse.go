package xmlparse

import "encoding/xml"

type ActiveAlerts struct {
	ActiveAlerts xml.Name `xml:"activeAlerts"`
	TimeStamp    string   `xml:"timeStamp,attr"`
	Events       []Event  `xml:"event"`
}

type Event2 struct {
	Event   xml.Name `xml:"event"`
	Address string   `xml:"address"`
	PubDate string   `xml:"pubDate"`
}

type Event struct {
	Event        xml.Name `xml:"event"`
	Title        string   `xml:"title"`
	Desc         string   `xml:"desc"`
	Station      string   `xml:"station"`
	Dispatch     string   `xml:"dispatch"`
	Lat          string   `xml:"lat"`
	Lng          string   `xml:"lng"`
	Postal       string   `xml:"postal"`
	Neighborhood string   `xml:"neighborhood"`
	Address      string   `xml:"address"`
	PubDate      string   `xml:"pubDate"`
}

func Decode(b []byte) ActiveAlerts {
	var activeAlerts ActiveAlerts
	xml.Unmarshal(b, &activeAlerts)
	return activeAlerts
}
