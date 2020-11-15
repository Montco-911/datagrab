package squish

import (
	"fmt"
	"testing"
)

func TestDoSquish(t *testing.T) {
	DoSquish()
}

func TestTWP(t *testing.T) {
	v := "a;b ;CHELTENHAM"
	r := TWP(v)
	fmt.Println(r)
}
