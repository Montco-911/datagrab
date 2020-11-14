package utils

import (
	"fmt"
	"os"
	"testing"
)

func CreateMock() {
	f, _ := os.Create("mockfile.txt")
	for i := 0; i < 1000; i++ {
		result := fmt.Sprintf("......(%d)\n", i)
		f.Write([]byte(result))
	}

	defer f.Close()
}

func TestUT_Squish(t *testing.T) {
	CreateMock()
	ut := NewUT("junk.txt", 50)
	ut.Squish("mockfile.txt")
}
