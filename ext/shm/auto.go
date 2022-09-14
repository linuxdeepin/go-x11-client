// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package shm

import x "github.com/linuxdeepin/go-x11-client"

// _ns.ext_name: Shm
const MajorVersion = 1
const MinorVersion = 2

var _ext *x.Extension

func Ext() *x.Extension {
	return _ext
}

// simple ('xcb', 'Shm', 'SEG')
type Seg uint32

const CompletionEventCode = 0

func NewCompletionEvent(data []byte) (*CompletionEvent, error) {
	var ev CompletionEvent
	r := x.NewReaderFromData(data)
	err := readCompletionEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const BadSegErrorCode = 0
const QueryVersionOpcode = 0

type QueryVersionCookie x.SeqNum

const AttachOpcode = 1
const DetachOpcode = 2
const PutImageOpcode = 3
const GetImageOpcode = 4

type GetImageCookie x.SeqNum

const CreatePixmapOpcode = 5
const AttachFdOpcode = 6
const CreateSegmentOpcode = 7

type CreateSegmentCookie x.SeqNum

var errorCodeNameMap = map[uint8]string{
	BadSegErrorCode: "BadSeg",
}
var requestOpcodeNameMap = map[uint]string{
	QueryVersionOpcode:  "QueryVersion",
	AttachOpcode:        "Attach",
	DetachOpcode:        "Detach",
	PutImageOpcode:      "PutImage",
	GetImageOpcode:      "GetImage",
	CreatePixmapOpcode:  "CreatePixmap",
	AttachFdOpcode:      "AttachFd",
	CreateSegmentOpcode: "CreateSegment",
}

func init() {
	_ext = x.NewExtension("MIT-SHM", 0, errorCodeNameMap, requestOpcodeNameMap)
}
