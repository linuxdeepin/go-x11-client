// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

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

type GetVersionCookie x.SeqNum

// enum Cursor
const (
	CursorNone    = 0
	CursorCurrent = 1
)

const CompareCursorOpcode = 1

type CompareCursorCookie x.SeqNum

const FakeInputOpcode = 2
const GrabControlOpcode = 3

var requestOpcodeNameMap = map[uint]string{
	GetVersionOpcode:    "GetVersion",
	CompareCursorOpcode: "CompareCursor",
	FakeInputOpcode:     "FakeInput",
	GrabControlOpcode:   "GrabControl",
}

func init() {
	_ext = x.NewExtension("XTEST", 0, nil, requestOpcodeNameMap)
}
