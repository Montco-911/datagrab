package event

import (
	"bufio"
	"cloud.google.com/go/datastore"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	"log"
	"os"
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
	Raw string
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

		fmt.Printf("%s, %s,%s, %s, %s\n", task.TimeStamp,string(task.Incidentno), task.Lat, task.Lng, task.Title)
		fmt.Fprintf(w, "%s, %s,%s, %s, %s\n", task.TimeStamp,string(task.Incidentno), task.Lat, task.Lng, task.Title)
	}

}

func GetLiveXML(kind string, count int) {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, "mchirico")
	if err != nil {
		fmt.Println(err)
	}

	f, err := os.Create("livexml.csv")
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

		var task LiveXML
		_, err := it.Next(&task)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Error fetching next task: %v", err)
		}

		fmt.Printf("%s, %s\n", task.TimeStamp,task.Raw)
		fmt.Fprintf(w, "%s, %s\n", task.TimeStamp,task.Raw)
	}

}
