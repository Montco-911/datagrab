package utils

import (
	"fmt"
	"os"
	"sync"
)

type UT struct {
	outfile      string
	readDistance int
	process      func([]byte)
	stopOffset   int64
	sync.Mutex
}

func NewUT(file string, readDistance int, p func([]byte)) *UT {
	ut := &UT{file, readDistance,
		p,-1, sync.Mutex{}}
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

func (ut *UT) SetStop(i int64) {
	ut.Lock()
	defer ut.Unlock()
	ut.stopOffset = i
}

func (ut *UT) LineGulp(file string) {
	ut.Lock()
	defer ut.Unlock()
	f, err := os.Create(ut.outfile)
	if err != nil {
		return
	}
	defer f.Close()

	f2, err := os.Open(file)
	defer f2.Close()

	var offset, oldoffset int64

	b := make([]byte, ut.readDistance)
	pt, idx, err := RR(f2, b)
	ut.process(b[0:idx])

	for {
		offset, err = f2.Seek(pt, 1)
		pt, idx, err = RR(f2, b)
		if err != nil {
			break
		}
		ut.process(b[0:idx])

		if ut.stopOffset > 0 && offset >= ut.stopOffset {
			break
		}

		if offset == oldoffset {
			break
		}
		oldoffset = offset
	}

}
