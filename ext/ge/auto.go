package ge

import x "github.com/linuxdeepin/go-x11-client"

// _ns.ext_name: GenericEvent
const MajorVersion = 1
const MinorVersion = 0

var _ext *x.Extension

func Ext() *x.Extension {
	return _ext
}

const QueryVersionOpcode = 0

type QueryVersionCookie uint64

func init() {
	_ext = x.NewExtension("Generic Event Extension", 0, nil)
}
