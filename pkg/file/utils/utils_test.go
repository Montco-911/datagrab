package utils

import (
	"fmt"
	"os"
	"testing"
)

func CreateMock(file string) {
	f, _ := os.Create("mockfile.txt")
	for i := 0; i < 1000; i++ {
		result := fmt.Sprintf("......(%d)\n", i)
		f.Write([]byte(result))
	}

	defer f.Close()
}

type P struct {
	count int
}

func (p *P) Process(b []byte) {
	p.count += 1
	if p.count < 4 {
		fmt.Printf("%s\n", b)
	}
}

func TestUT_LineGulp(t *testing.T) {
	p := &P{}

	file := "mockfile.txt"
	CreateMock(file)
	ut := NewUT("junk.txt", 50, p.Process)
	ut.LineGulp(file)

	if p.count != 252 {
		t.Fatalf("Expected 252, got: %d\n", p.count)
	}

}

func TestUT_SetStop(t *testing.T) {
	p := &P{}

	file := "../../../fixtures/alllivexml.csv"

	ut := NewUT("junk2.txt", 5000, p.Process)
	ut.SetStop(1200)
	ut.LineGulp(file)
}
