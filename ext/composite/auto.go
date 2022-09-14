// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

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

type QueryVersionCookie x.SeqNum

const RedirectWindowOpcode = 1
const RedirectSubwindowsOpcode = 2
const UnredirectWindowOpcode = 3
const UnredirectSubwindowsOpcode = 4
const CreateRegionFromBorderClipOpcode = 5
const NameWindowPixmapOpcode = 6
const GetOverlayWindowOpcode = 7

type GetOverlayWindowCookie x.SeqNum

const ReleaseOverlayWindowOpcode = 8

var requestOpcodeNameMap = map[uint]string{
	QueryVersionOpcode:               "QueryVersion",
	RedirectWindowOpcode:             "RedirectWindow",
	RedirectSubwindowsOpcode:         "RedirectSubwindows",
	UnredirectWindowOpcode:           "UnredirectWindow",
	UnredirectSubwindowsOpcode:       "UnredirectSubwindows",
	CreateRegionFromBorderClipOpcode: "CreateRegionFromBorderClip",
	NameWindowPixmapOpcode:           "NameWindowPixmap",
	GetOverlayWindowOpcode:           "GetOverlayWindow",
	ReleaseOverlayWindowOpcode:       "ReleaseOverlayWindow",
}

func init() {
	_ext = x.NewExtension("Composite", 0, nil, requestOpcodeNameMap)
}
