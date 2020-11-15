package utils

import (
	"fmt"
	"os"
	"sync"
)

type UT struct {
	Outfile      string
	ReadDistance int
	Process      func([]byte)
	sync.Mutex
}

func NewUT(file string, readDistance int, p func([]byte)) *UT {
	ut := &UT{file, readDistance,
		p, sync.Mutex{}}
	return ut
}

func RR(f *os.File, b []byte) (int64, int, error) {
	n, err := f.Read(b)
	idx := 0
	for i := n - 1; i > 0; i-- {
		if b[i] == 10 {
			idx = i
			break
		}
	}
	if idx >= n {
		return 0, 0, fmt.Errorf("idx bad")
	}
	return int64(idx - n), idx, err
}

func (ut *UT) Squish(file string) {
	ut.Lock()
	defer ut.Unlock()
	f, err := os.Create(ut.Outfile)
	if err != nil {
		return
	}
	defer f.Close()

	f2, err := os.Open(file)
	defer f2.Close()

	var offset, old int64

	b := make([]byte, ut.ReadDistance)
	pt, idx, err := RR(f2, b)
	ut.Process(b[0:idx])

	for {
		offset, err = f2.Seek(pt, 1)
		pt, idx, err = RR(f2, b)
		if err != nil {
			break
		}
		ut.Process(b[0:idx])

		if offset == old {
			break
		}
		old = offset
	}

}
