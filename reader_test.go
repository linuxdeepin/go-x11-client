package x

import (
	"testing"

	. "gopkg.in/check.v1"
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
	v, err := s.r.ReadBytes(4)
	c.Assert(v, DeepEquals, []byte{1, 2, 3, 4})
	c.Assert(err, IsNil)

	v, err = s.r.ReadBytes(3)
	c.Assert(v, DeepEquals, []byte{5, 6, 7})
	c.Assert(err, IsNil)
}

func (s *MySuite) TestReaderReadBytesWithPad(c *C) {
	_, v, err := s.r.ReadBytesWithPad(3)
	c.Assert(v, DeepEquals, []byte{1, 2, 3})
	c.Assert(s.r.Pos(), Equals, 4)
	c.Assert(err, IsNil)
}

func (s *MySuite) TestReaderRead1b(c *C) {
	for i := 1; i <= 8; i++ {
		c.Assert(s.r.Read1b(), Equals, uint8(i))
	}
}

func (s *MySuite) TestReaderRead2b(c *C) {
	results := []uint16{0x0201, 0x0403, 0x0605, 0x0807}
	for _, result := range results {
		c.Assert(s.r.Read2b(), Equals, result)
	}
}

func (s *MySuite) TestReaderRead4b(c *C) {
	results := []uint32{0x04030201, 0x08070605}
	for _, result := range results {
		c.Assert(s.r.Read4b(), Equals, result)
	}
}

func (s *MySuite) TestReaderReadNulTermStr(c *C) {
	data := []byte{'h', 'e', 'l', 'l', 'o'}
	r := NewReaderFromData(data)
	str := r.ReadNulTermStr()
	c.Assert(str, Equals, "hello")
	c.Assert(r.Pos(), Equals, 5)

	data = []byte{'h', 'e', 'l', 'l', 'o', 0, 'w', 'o', 'r', 'l', 'd'}
	r = NewReaderFromData(data)
	str = r.ReadNulTermStr()
	c.Assert(str, Equals, "hello")
	c.Assert(r.Pos(), Equals, 5)
}

func (s *MySuite) TestReaderRemainAtLeast(c *C) {
	s.r.ReadPad(3)
	c.Assert(s.r.RemainAtLeast(5), Equals, true)
	c.Assert(s.r.RemainAtLeast(6), Equals, false)

	s.r.Reset()
	c.Assert(s.r.RemainAtLeast4b(0), Equals, true)
	c.Assert(s.r.RemainAtLeast4b(1), Equals, true)
	c.Assert(s.r.RemainAtLeast4b(2), Equals, true)
	c.Assert(s.r.RemainAtLeast4b(3), Equals, false)
}
