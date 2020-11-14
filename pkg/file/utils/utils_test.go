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

func Process(b []byte) {
	fmt.Printf("%s\n", b)
}

func TestUT_Squish(t *testing.T) {
	file := "mockfile.txt"
	CreateMock(file)
	ut := NewUT("junk.txt", 50, Process)
	ut.Squish(file)
}
