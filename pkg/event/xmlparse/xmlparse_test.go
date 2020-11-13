package xmlparse

import (
	"testing"
)

func MockData() string {
	data := `<activeAlerts timeStamp="2017-08-07 12:05:00" >
<event>
<title>EMS: OVERDOSE</title>
<desc>ORCHARD CT &amp; VAUGHN RD;  UPPER PROVIDENCE; Station 325; 2017-08-07 @ 07:57:47;</desc>
<station> Station 325</station>
<dispatch> 2017-08-07 @ 07:57:47</dispatch>
<lat>40.1839670</lat>
<lng>-75.5228942</lng>
<postal>19468</postal>
<neighborhood></neighborhood>
<address>ORCHARD CT &amp; VAUGHN RD,  UPPER PROVIDENCE</address>
<pubDate>Mon, 7 Aug 2017 12:05:00 GMT</pubDate>
</event>
<event>
<title>EMS: MEDICAL ALERT ALARM</title>
<desc>SHANNONDELL DR &amp; SHANNONDELL BLVD;  LOWER PROVIDENCE; Station 322A; 2017-08-07 @ 07:56:01;</desc>
<station> Station 322A</station>
<dispatch> 2017-08-07 @ 07:56:01</dispatch>
<lat>40.1330371</lat>
<lng>-75.4084631</lng>
<postal>19403</postal>
<neighborhood></neighborhood>
<address>SHANNONDELL DR &amp; SHANNONDELL BLVD,  LOWER PROVIDENCE</address>
<pubDate>Mon, 7 Aug 2017 12:05:00 GMT</pubDate>
</event>
<event>
<title>EMS: RESPIRATORY EMERGENCY</title>
<desc>MAIN ST &amp; OLD SUMNEYTOWN PIKE;  LOWER SALFORD; Station 344; 2017-08-07 @ 07:47:58;</desc>
<station> Station 344</station>
<dispatch> 2017-08-07 @ 07:47:58</dispatch>
<lat>40.2890267</lat>
<lng>-75.3995896</lng>
<postal>19438</postal>
<neighborhood></neighborhood>
<address>MAIN ST &amp; OLD SUMNEYTOWN PIKE,  LOWER SALFORD</address>
<pubDate>Mon, 7 Aug 2017 12:05:00 GMT</pubDate>
</event>
<event>
<title>EMS: GENERAL WEAKNESS</title>
<desc>STONEHAVEN DR &amp; JEFFERSON ST;  RED HILL; Station 369; 2017-08-07 @ 07:42:02;</desc>
<station> Station 369</station>
<dispatch> 2017-08-07 @ 07:42:02</dispatch>
<lat>40.3727342</lat>
<lng>-75.4837547</lng>
<postal>18076</postal>
<neighborhood></neighborhood>
<address>STONEHAVEN DR &amp; JEFFERSON ST,  RED HILL</address>
<pubDate>Mon, 7 Aug 2017 12:05:00 GMT</pubDate>
</event>
<event>
<title>EMS: CARDIAC EMERGENCY</title>
<desc>RIVERFRONT DR &amp; DIAMOND PL;  ROYERSFORD; Station 325; 2017-08-07 @ 07:38:38;</desc>
<station> Station 325</station>
<dispatch> 2017-08-07 @ 07:38:38</dispatch>
<lat>40.1792894</lat>
<lng>-75.5417180</lng>
<postal>19468</postal>
<neighborhood></neighborhood>
<address>RIVERFRONT DR &amp; DIAMOND PL,  ROYERSFORD</address>
<pubDate>Mon, 7 Aug 2017 12:05:00 GMT</pubDate>
</event>
<event>
<title>EMS: RESPIRATORY EMERGENCY</title>
<desc>GLASGOW ST &amp; ELM ST;  POTTSTOWN; Station 329; 2017-08-07 @ 07:31:12;</desc>
<station> Station 329</station>
<dispatch> 2017-08-07 @ 07:31:12</dispatch>
<lat>40.2539246</lat>
<lng>-75.6731501</lng>
<postal>19464</postal>
<neighborhood></neighborhood>
<address>GLASGOW ST &amp; ELM ST,  POTTSTOWN</address>
<pubDate>Mon, 7 Aug 2017 12:05:00 GMT</pubDate>
</event>
<event>
<title>EMS: RESPIRATORY EMERGENCY</title>
<desc>LAUREL AVE &amp; JEFFERSON AVE;  CHELTENHAM; Station 358; 2017-08-07 @ 06:47:22;</desc>
<station> Station 358</station>
<dispatch> 2017-08-07 @ 06:47:22</dispatch>
<lat>40.0676677</lat>
<lng>-75.0965923</lng>
<postal>19012</postal>
<neighborhood>Cheltenham</neighborhood>
<address>LAUREL AVE &amp; JEFFERSON AVE,  CHELTENHAM</address>
<pubDate>Mon, 7 Aug 2017 12:05:00 GMT</pubDate>
</event>
<event>
<title>Traffic: DISABLED VEHICLE -</title>
<desc>DEKALB PIKE &amp; JOLLY RD; WHITPAIN; 2017-08-07 @ 07:59:31;</desc>
<lat>40.1580254</lat>
<lng>-75.2992999</lng>
<postal>19422</postal>
<neighborhood></neighborhood>
<address>DEKALB PIKE &amp; JOLLY RD, WHITPAIN</address>
<pubDate>Mon, 7 Aug 2017 12:05:00 GMT</pubDate>
</event>
<event>
<title>Traffic: VEHICLE ACCIDENT -</title>
<desc>ASHBOURNE RD &amp; JENKINTOWN RD; CHELTENHAM; 2017-08-07 @ 07:56:35;</desc>
<lat>40.0620972</lat>
<lng>-75.1042225</lng>
<postal>19012</postal>
<neighborhood>Cheltenham</neighborhood>
<address>ASHBOURNE RD &amp; JENKINTOWN RD, CHELTENHAM</address>
<pubDate>Mon, 7 Aug 2017 12:05:00 GMT</pubDate>
</event>
<event>
<title>Traffic: DISABLED VEHICLE -</title>
<desc>SUMNEYTOWN PIKE &amp; MAINLAND RD; TOWAMENCIN; 2017-08-07 @ 07:41:13;</desc>
<lat>40.2506263</lat>
<lng>-75.3492741</lng>
<postal>19438</postal>
<neighborhood></neighborhood>
<address>SUMNEYTOWN PIKE &amp; MAINLAND RD, TOWAMENCIN</address>
<pubDate>Mon, 7 Aug 2017 12:05:00 GMT</pubDate>
</event>
<event>
<title>Traffic: VEHICLE ACCIDENT -</title>
<desc>EASTON RD &amp; N YORK RD; UPPER MORELAND; 2017-08-07 @ 07:34:31;</desc>
<lat>40.1445255</lat>
<lng>-75.1160469</lng>
<postal>19090</postal>
<neighborhood></neighborhood>
<address>EASTON RD &amp; N YORK RD, UPPER MORELAND</address>
<pubDate>Mon, 7 Aug 2017 12:05:00 GMT</pubDate>
</event>
<event>
<title>Traffic: VEHICLE ACCIDENT -</title>
<desc>RIVER RD &amp; SWEDELAND RD; UPPER MERION; 2017-08-07 @ 07:34:16;</desc>
<lat>40.0824810</lat>
<lng>-75.3246360</lng>
<postal>19406</postal>
<neighborhood></neighborhood>
<address>RIVER RD &amp; SWEDELAND RD, UPPER MERION</address>
<pubDate>Mon, 7 Aug 2017 12:05:00 GMT</pubDate>
</event>
<event>
<title>Traffic: VEHICLE ACCIDENT -</title>
<desc>RAMP RT422 EB TO EGYPT RD &amp; RT422 EB; UPPER PROVIDENCE; 2017-08-07 @ 07:24:56;</desc>
<lat>40.1723141</lat>
<lng>-75.4927278</lng>
<postal></postal>
<neighborhood></neighborhood>
<address>RAMP RT422 EB TO EGYPT RD &amp; RT422 EB, UPPER PROVIDENCE</address>
<pubDate>Mon, 7 Aug 2017 12:05:00 GMT</pubDate>
</event>
</activeAlerts>`

	return data

}

func TestDecode(t *testing.T) {
	data := []byte(MockData())
	a := Decode(data)
	if len(a.Events) != 13 {
		t.Fatalf(("Didn't get all records"))
	}
	expected := "EMS: GENERAL WEAKNESS"
	if a.Events[3].Title != expected {
		t.Fatalf("Expected: %s, got %s\n", expected, a.Events[3].Title)
	}

}
