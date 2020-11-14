package event

import (
	"bufio"
	"cloud.google.com/go/datastore"
	"context"
	"fmt"
	"github.com/Montco-911/datagrab/pkg/event/xmlparse"
	"google.golang.org/api/iterator"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

type Event struct {
	TimeStamp    time.Time
	Address      []byte
	Boundry      []byte
	Dispatch     []byte
	Incidentno   []byte
	station      []byte
	Desc         string
	Lat          string
	Lng          string
	NeLat        string
	NeLng        string
	Neighborhood string
	Postal       string
	Station      string
	SwLat        string
	SwLng        string
	Title        string
}

type LiveXML struct {
	TimeStamp time.Time
	Raw       string
}

type Raw struct {
	TimeStamp    time.Time
	ActiveAlerts xmlparse.ActiveAlerts
}

func GetEvents(kind string, count int) {

	ctx := context.Background()
	client, err := datastore.NewClient(ctx, "mchirico")
	if err != nil {
		fmt.Println(err)
	}

	f, err := os.Create("events.csv")
	if err != nil {
		return
	}
	defer f.Close()

	w := bufio.NewWriter(f)

	query := datastore.NewQuery(kind)
	mcount := 0

	it := client.Run(ctx, query)
	for {
		mcount += 1
		if mcount >= count {
			break
		}

		var task Event
		_, err := it.Next(&task)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Error fetching next task: %v", err)
		}

		fmt.Printf("%s, %s,%s, %s, %s\n", task.TimeStamp, string(task.Incidentno), task.Lat, task.Lng, task.Title)
		fmt.Fprintf(w, "%s, %s,%s, %s, %s\n", task.TimeStamp, string(task.Incidentno), task.Lat, task.Lng, task.Title)
	}

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

			t := re.FindAllString(v.Desc, -1)
			if len(t) != 0 {
				timeStamp := strings.Replace(t[0], " @ ", " ", -1)
				fmt.Fprintf(w, "%s,%s,%q,%s,%s,%s,%s,%s\n", timeStamp, v.Title, v.Desc, v.Lng, v.Lat, v.Postal, v.Station, iv.TimeStamp)
			}
		}
	}

}

type DS struct {
	File    string
	re      *regexp.Regexp
	Create  func(string) (*os.File, *bufio.Writer, error)
	Heading func(*bufio.Writer) error
	Write   func(w *bufio.Writer, re *regexp.Regexp, r Raw)
}

func Write(w *bufio.Writer, re *regexp.Regexp, r Raw) {
	for _, v := range r.ActiveAlerts.Events {

		t := re.FindAllString(v.Desc, -1)
		if len(t) != 0 {
			timeStamp := strings.Replace(t[0], " @ ", " ", -1)
			fmt.Fprintf(w, "%s,%s,%q,%s,%s,%s,%s,%s\n", timeStamp, v.Title, v.Desc, v.Lng, v.Lat, v.Postal, v.Station, r.TimeStamp)
		}
	}
}

func Create(file string) (*os.File, *bufio.Writer, error) {
	f, err := os.Create(file)
	if err != nil {
		return nil, nil, err
	}

	w := bufio.NewWriter(f)
	return f, w, err
}

func Heading(w *bufio.Writer) error {
	_, err := fmt.Fprintf(w, "TimeStamp,Title,Desc,Lng,Lag,Postal,Station,AlertTimeStamp\n")
	return err
}

func NewDS(file string) DS {
	re := regexp.MustCompile(`[0-9]{4}-[0-9]{2}-[0-9]{2} @ [0-9]{2}:[0-9]{2}:[0-9]{2}`)
	ds := DS{file, re, Create, Heading, Write}
	return ds
}

func (ds DS) GetLiveXML(kind string, count int) {

	ctx := context.Background()
	client, err := datastore.NewClient(ctx, "mchirico")
	if err != nil {
		fmt.Println(err)
	}

	query := datastore.NewQuery(kind).Order("-timeStamp")

	mcount := 0

	f, w, err := ds.Create(ds.File)
	defer f.Close()
	ds.Heading(w)

	it := client.Run(ctx, query)
	for {
		mcount += 1
		if mcount >= count {
			break
		}

		if mcount%1000 == 0 {
			fmt.Printf("count: %v\n", mcount)
			f.Sync()
		}

		var task LiveXML
		_, err := it.Next(&task)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Error fetching next task: %v", err)
		}

		raw := Raw{}
		raw.TimeStamp = task.TimeStamp
		raw.ActiveAlerts = xmlparse.Decode([]byte(task.Raw))
		ds.Write(w, ds.re, raw)

	}

}
