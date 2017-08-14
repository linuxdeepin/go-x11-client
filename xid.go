package client

import (
	"errors"
	"sync"
)

func (c *Conn) GenerateID() (uint32, error) {
	return c.xg.generate()
}

type xidGenerator struct {
	mu   sync.Mutex
	last uint32
	base uint32
	max  uint32
	inc  uint32
}

func newXidGenerator(base, mask uint32) *xidGenerator {
	xg := new(xidGenerator)
	xg.base = base
	xg.inc = mask & -mask
	xg.max = mask
	return xg
}

func (xg *xidGenerator) generate() (ret uint32, err error) {
	xg.mu.Lock()
	defer xg.mu.Unlock()

	if xg.last >= xg.max-xg.inc+1 {
		// TODO
		if xg.last != xg.max {
			panic("assert xg.last == xg.max failed")
		}
		// inc >= 1
		return 0, errors.New("no more available id")
	} else {
		xg.last += xg.inc
	}
	ret = xg.last | xg.base
	return ret, nil
}
