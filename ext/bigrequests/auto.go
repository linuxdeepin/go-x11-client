// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package bigrequests

import x "github.com/linuxdeepin/go-x11-client"

// _ns.ext_name: BigRequests
const MajorVersion = 0
const MinorVersion = 0

var _ext *x.Extension

func Ext() *x.Extension {
	return _ext
}

const EnableOpcode = 0

type EnableCookie x.SeqNum

var requestOpcodeNameMap = map[uint]string{
	EnableOpcode: "Enable",
}

func init() {
	_ext = x.NewExtension("BIG-REQUESTS", 0, nil, requestOpcodeNameMap)
}
