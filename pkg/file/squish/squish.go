package squish

import (
	"fmt"
	"github.com/Montco-911/datagrab/pkg/file/utils"
	"os"
	"strings"
)

type P struct {
	count int
	f *os.File
}

func (p *P) CreateFile(file string){
	f, err := os.Create(file)
	if err != nil {
		return
	}
	p.f = f
}


func (p *P) Process(b []byte) {
	p.count += 1
    if p.count == 1 {

		p.f.Write([]byte("TimeStamp,Title,Desc,Lng,Lag,Postal,Station\n"))
		return
	}
	lines := strings.Split(string(b), "\n")
	m := map[string]int{}
	for _, v := range lines {
		r := strings.Split(v, ",")
		if len(r) > 6 {
			m[r[0]+","+r[1]+","+r[2]+","+r[3]+","+r[4]+","+r[5]+","+r[6]] += 1

		}

	}
   for k,_ := range m {
	   r := fmt.Sprintf("%s\n", k)
	   p.f.Write([]byte(r))
   }


}

func DoSquish() {
	p := &P{}
	p.CreateFile("out.csv")

	file := "../../../fixtures/alllivexml.csv"

	ut := utils.NewUT("junk2.txt", 5000, p.Process)
	ut.SetStop(12000)
	ut.LineGulp(file)
}
