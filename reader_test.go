package client

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestReader(t *testing.T) {
	Convey("Reader", t, func() {
		buf := make([]byte, 8)
		buf[0] = 1
		buf[1] = 2
		buf[2] = 3
		buf[3] = 4
		buf[4] = 5
		buf[5] = 6
		buf[6] = 7
		buf[7] = 8

		r := NewReaderFromData(buf)

		Convey("ReadBytes", func() {
			So(r.ReadBytes(4), ShouldResemble, []byte{1, 2, 3, 4})
			So(r.Err(), ShouldBeNil)

			So(r.ReadBytes(3), ShouldResemble, []byte{5, 6, 7})
			So(r.Err(), ShouldBeNil)

			So(r.ReadBytes(2), ShouldEqual, nil)
			So(r.Err(), ShouldNotBeNil)
		})

		Convey("Read1b", func() {
			So(r.Read1b(), ShouldEqual, 1)
			So(r.Read1b(), ShouldEqual, 2)
			So(r.Read1b(), ShouldEqual, 3)
			So(r.Read1b(), ShouldEqual, 4)
			So(r.Err(), ShouldBeNil)
			So(r.Read1b(), ShouldEqual, 5)
			So(r.Err(), ShouldBeNil)
		})

		Convey("Read2b", func() {
			So(r.Read2b(), ShouldEqual, 0x0201)
			So(r.Err(), ShouldBeNil)

			So(r.Read2b(), ShouldEqual, 0x0403)
			So(r.Err(), ShouldBeNil)

			So(r.Read2b(), ShouldEqual, 0x0605)
			So(r.Err(), ShouldBeNil)
		})

		Convey("Read4b", func() {
			So(r.Read4b(), ShouldEqual, 0x04030201)
			So(r.Err(), ShouldBeNil)

			So(r.Read4b(), ShouldEqual, 0x08070605)
			So(r.Err(), ShouldBeNil)

			So(r.Read2b(), ShouldEqual, 0)
			So(r.Err(), ShouldNotBeNil)
		})

		Convey("Read8b", func() {
			So(r.Read8b(), ShouldEqual, 0x0807060504030201)
			So(r.Err(), ShouldBeNil)

			So(r.Read2b(), ShouldEqual, 0)
			So(r.Err(), ShouldNotBeNil)
		})
	})
}
