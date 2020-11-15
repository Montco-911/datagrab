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

func TestUT_Squish(t *testing.T) {
	p := &P{}

	file := "mockfile.txt"
	CreateMock(file)
	ut := NewUT("junk.txt", 50, p.Process)
	ut.Squish(file)

	fmt.Printf("a= %v\n", p.count)
}
