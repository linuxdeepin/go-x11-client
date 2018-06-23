package x

import (
	"errors"
	"math/big"
	"sync"
)

func (c *Conn) AllocID() (uint32, error) {
	return c.ridAllocator.alloc()
}

func (c *Conn) FreeID(rid uint32) error {
	return c.ridAllocator.free(rid)
}

type resourceIdAllocator struct {
	mu           sync.Mutex
	base         uint32
	mask         uint32
	last         uint32
	bitmap       *big.Int
	allAllocated bool
}

func (ra *resourceIdAllocator) init(base, mask uint32) {
	ra.base = base
	ra.mask = mask
	ra.bitmap = big.NewInt(0)
}

var errOutOfResourceIds = errors.New("out of resource ids")

func (ra *resourceIdAllocator) alloc() (uint32, error) {
	ra.mu.Lock()
	defer ra.mu.Unlock()

	if ra.allAllocated {
		return 0, errOutOfResourceIds
	}

	i := ra.last
	for ra.bitmap.Bit(int(i)) == 1 {
		i++
		if i > ra.mask {
			i = 0
		}
		if i == ra.last {
			ra.allAllocated = true
			return 0, errOutOfResourceIds
		}
	}
	ra.bitmap.SetBit(ra.bitmap, int(i), 1)
	ra.last = i
	return ra.base | i, nil
}

func (ra *resourceIdAllocator) free(rid uint32) error {
	ra.mu.Lock()
	defer ra.mu.Unlock()

	i := rid & ra.mask
	if rid-i != ra.base {
		return errors.New("resource id outside range")
	}

	if ra.bitmap.Bit(int(i)) == 0 {
		return errors.New("resource id not used")
	}
	ra.bitmap.SetBit(ra.bitmap, int(i), 0)
	ra.allAllocated = false
	return nil
}
