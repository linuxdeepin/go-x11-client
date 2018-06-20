package composite

import x "github.com/linuxdeepin/go-x11-client"

// _ns.ext_name: Composite
const MajorVersion = 0
const MinorVersion = 4

var _ext *x.Extension

func Ext() *x.Extension {
	return _ext
}

// enum Redirect
const (
	RedirectAutomatic = 0
	RedirectManual    = 1
)

const QueryVersionOpcode = 0

type QueryVersionCookie uint64

const RedirectWindowOpcode = 1
const RedirectSubwindowsOpcode = 2
const UnredirectWindowOpcode = 3
const UnredirectSubwindowsOpcode = 4
const CreateRegionFromBorderClipOpcode = 5
const NameWindowPixmapOpcode = 6
const GetOverlayWindowOpcode = 7

type GetOverlayWindowCookie uint64

const ReleaseOverlayWindowOpcode = 8

func init() {
	_ext = x.NewExtension("Composite", 0, nil)
}
