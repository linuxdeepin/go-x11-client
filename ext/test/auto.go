package test

import x "github.com/linuxdeepin/go-x11-client"

// _ns.ext_name: Test
const MajorVersion = 2
const MinorVersion = 2

var _ext *x.Extension

func Ext() *x.Extension {
	return _ext
}

const GetVersionOpcode = 0

type GetVersionCookie uint64

// enum Cursor
const (
	CursorNone    = 0
	CursorCurrent = 1
)

const CompareCursorOpcode = 1

type CompareCursorCookie uint64

const FakeInputOpcode = 2
const GrabControlOpcode = 3

func init() {
	_ext = x.NewExtension("XTEST", 0, nil)
}
