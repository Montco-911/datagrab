package squish

import (
	"fmt"
	"github.com/Montco-911/datagrab/pkg/file/utils"
	"os"
	"strings"
)

func TWP(s string) string {
	twp := []string{
		"UPPER POTTSGROVE", "LOWER MERION", "PLYMOUTH", "ABINGTON",
		"NORRISTOWN", "POTTSTOWN", "LIMERICK", "BUCKS COUNTY",
		"TOWAMENCIN", "FRANCONIA", "LOWER PROVIDENCE", "UPPER MORELAND",
		"LOWER MORELAND", "LANSDALE", "ROYERSFORD", "CHESTER COUNTY",
		"WEST CONSHOHOCKEN", "UPPER MERION", "UPPER DUBLIN", "CHELTENHAM",
		"NARBERTH", "HORSHAM", "LOWER POTTSGROVE", "HATFIELD TOWNSHIP",
		"CONSHOHOCKEN", "MONTGOMERY", "EAST NORRITON", "ROCKLEDGE",
		"HATFIELD BORO", "BRIDGEPORT", "LOWER GWYNEDD", "WHITPAIN",
		"LOWER SALFORD", "UPPER PROVIDENCE", "UPPER GWYNEDD",
		"NORTH WALES", "WORCESTER", "WHITEMARSH", "DOUGLASS", "JENKINTOWN",
		"PERKIOMEN", "WEST NORRITON", "GREEN LANE", "SKIPPACK", "AMBLER",
		"SOUDERTON", "PENNSBURG", "EAST GREENVILLE", "UPPER HANOVER",
		"BERKS COUNTY", "HATBORO", "SPRINGFIELD", "TELFORD",
		"UPPER FREDERICK", "UPPER SALFORD", "LOWER FREDERICK",
		"WEST POTTSGROVE", "COLLEGEVILLE", "SALFORD", "NEW HANOVER",
		"BRYN ATHYN", "DELAWARE COUNTY", "TRAPPE", "LEHIGH COUNTY",
		"SCHWENKSVILLE", "MARLBOROUGH", "PHILA COUNTY", "RED HILL"}

	for _, v := range twp {
		if strings.Contains(s, v) {
			return v
		}

	}

	return ""

}

type P struct {
	count int
	f     *os.File
}

func (p *P) CreateFile(file string) {
	f, err := os.Create(file)
	if err != nil {
		return
	}
	p.f = f
}

func (p *P) Process(b []byte) {
	p.count += 1
	if p.count == 1 {

		// lat,lng,desc,zip,title,timeStamp,twp,addr
		p.f.Write([]byte("timeStamp,title,desc,lng,lat,zip,station,twp\n"))
		return
	}
	lines := strings.Split(string(b), "\n")
	m := map[string]int{}

	for _, v := range lines {
		twp := TWP(v)
		r := strings.Split(v, ",")
		if len(r) > 6 {
			m[r[0]+","+r[1]+","+r[2]+","+r[3]+","+r[4]+","+r[5]+","+r[6]+","+twp] += 1

		}

	}
	for k, _ := range m {
		r := fmt.Sprintf("%s\n", k)
		p.f.Write([]byte(r))
	}

}

func DoSquish() {
	p := &P{}
	p.CreateFile("out.csv")

	//file := "../../../fixtures/alllivexml.csv"
	file := "/Users/rommel/Downloads/alllivexml.csv"

	ut := utils.NewUT("junk2.txt", 5000000, p.Process)
	//ut.SetStop(12000)
	ut.LineGulp(file)
}
