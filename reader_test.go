package client

import (
	. "gopkg.in/check.v1"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct {
	r *Reader
}

var _ = Suite(&MySuite{})

func (s *MySuite) SetUpTest(c *C) {
	buf := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	s.r = NewReaderFromData(buf)
}

func (s *MySuite) TestReaderReadBytes(c *C) {
	c.Assert(s.r.ReadBytes(4), DeepEquals, []byte{1, 2, 3, 4})
	c.Assert(s.r.Err(), IsNil)

	c.Assert(s.r.ReadBytes(3), DeepEquals, []byte{5, 6, 7})
	c.Assert(s.r.Err(), IsNil)

	c.Assert(s.r.ReadBytes(2), IsNil)
	c.Assert(s.r.Err(), NotNil)
}

func (s *MySuite) TestReaderRead1b(c *C) {
	for i := 1; i <= 8; i++ {
		c.Assert(s.r.Read1b(), Equals, uint8(i))
		c.Assert(s.r.Err(), IsNil)
	}
	c.Assert(s.r.Read1b(), Equals, uint8(0))
	c.Assert(s.r.Err(), NotNil)
}

func (s *MySuite) TestReaderRead2b(c *C) {
	results := []uint16{0x0201, 0x0403, 0x0605, 0x0807}
	for _, result := range results {
		c.Assert(s.r.Read2b(), Equals, result)
		c.Assert(s.r.Err(), IsNil)
	}
	c.Assert(s.r.Read2b(), Equals, uint16(0))
	c.Assert(s.r.Err(), NotNil)
}

func (s *MySuite) TestReaderRead4b(c *C) {
	results := []uint32{0x04030201, 0x08070605}
	for _, result := range results {
		c.Assert(s.r.Read4b(), Equals, result)
		c.Assert(s.r.Err(), IsNil)
	}
	c.Assert(s.r.Read4b(), Equals, uint32(0))
	c.Assert(s.r.Err(), NotNil)
}

func (s *MySuite) TestReaderRead8b(c *C) {
	c.Assert(s.r.Read8b(), Equals, uint64(0x0807060504030201))
	c.Assert(s.r.Err(), IsNil)

	c.Assert(s.r.Read8b(), Equals, uint64(0))
	c.Assert(s.r.Err(), NotNil)
}
