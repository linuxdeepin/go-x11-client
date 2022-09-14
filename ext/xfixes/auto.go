// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package xfixes

import x "github.com/linuxdeepin/go-x11-client"

// _ns.ext_name: XFixes
const MajorVersion = 5
const MinorVersion = 0

var _ext *x.Extension

func Ext() *x.Extension {
	return _ext
}

const QueryVersionOpcode = 0

type QueryVersionCookie x.SeqNum

// enum SaveSetMode
const (
	SaveSetModeInsert = 0
	SaveSetModeDelete = 1
)

// enum SaveSetTarget
const (
	SaveSetTargetNearest = 0
	SaveSetTargetRoot    = 1
)

// enum SaveSetMapping
const (
	SaveSetMappingMap   = 0
	SaveSetMappingUnmap = 1
)

const ChangeSaveSetOpcode = 1

// enum SelectionEvent
const (
	SelectionEventSetSelectionOwner      = 0
	SelectionEventSelectionWindowDestroy = 1
	SelectionEventSelectionClientClose   = 2
)

// enum SelectionEventMask
const (
	SelectionEventMaskSetSelectionOwner      = 1
	SelectionEventMaskSelectionWindowDestroy = 2
	SelectionEventMaskSelectionClientClose   = 4
)

const SelectionNotifyEventCode = 0

func NewSelectionNotifyEvent(data []byte) (*SelectionNotifyEvent, error) {
	var ev SelectionNotifyEvent
	r := x.NewReaderFromData(data)
	err := readSelectionNotifyEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const SelectSelectionInputOpcode = 2

// enum CursorNotify
const (
	CursorNotifyDisplayCursor = 0
)

// enum CursorNotifyMask
const (
	CursorNotifyMaskDisplayCursor = 1
)

const CursorNotifyEventCode = 1

func NewCursorNotifyEvent(data []byte) (*CursorNotifyEvent, error) {
	var ev CursorNotifyEvent
	r := x.NewReaderFromData(data)
	err := readCursorNotifyEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const SelectCursorInputOpcode = 3
const GetCursorImageOpcode = 4

type GetCursorImageCookie x.SeqNum

// simple ('xcb', 'XFixes', 'REGION')
type Region uint32

const BadRegionErrorCode = 0

// enum Region
const (
	RegionNone = 0
)

const CreateRegionOpcode = 5
const CreateRegionFromBitmapOpcode = 6
const CreateRegionFromWindowOpcode = 7
const CreateRegionFromGCOpcode = 8
const CreateRegionFromPictureOpcode = 9
const DestroyRegionOpcode = 10
const SetRegionOpcode = 11
const CopyRegionOpcode = 12
const UnionRegionOpcode = 13
const IntersectRegionOpcode = 14
const SubtractRegionOpcode = 15
const InvertRegionOpcode = 16
const TranslateRegionOpcode = 17
const RegionExtentsOpcode = 18
const FetchRegionOpcode = 19

type FetchRegionCookie x.SeqNum

const SetGCClipRegionOpcode = 20
const SetWindowShapeRegionOpcode = 21
const SetPictureClipRegionOpcode = 22
const SetCursorNameOpcode = 23
const GetCursorNameOpcode = 24

type GetCursorNameCookie x.SeqNum

const GetCursorImageAndNameOpcode = 25

type GetCursorImageAndNameCookie x.SeqNum

const ChangeCursorOpcode = 26
const ChangeCursorByNameOpcode = 27
const ExpandRegionOpcode = 28
const HideCursorOpcode = 29
const ShowCursorOpcode = 30

// simple ('xcb', 'XFixes', 'BARRIER')
type Barrier uint32

// enum BarrierDirections
const (
	BarrierDirectionsPositiveX = 1
	BarrierDirectionsPositiveY = 2
	BarrierDirectionsNegativeX = 4
	BarrierDirectionsNegativeY = 8
)

const CreatePointerBarrierOpcode = 31
const DeletePointerBarrierOpcode = 32

var errorCodeNameMap = map[uint8]string{
	BadRegionErrorCode: "BadRegion",
}
var requestOpcodeNameMap = map[uint]string{
	QueryVersionOpcode:            "QueryVersion",
	ChangeSaveSetOpcode:           "ChangeSaveSet",
	SelectSelectionInputOpcode:    "SelectSelectionInput",
	SelectCursorInputOpcode:       "SelectCursorInput",
	GetCursorImageOpcode:          "GetCursorImage",
	CreateRegionOpcode:            "CreateRegion",
	CreateRegionFromBitmapOpcode:  "CreateRegionFromBitmap",
	CreateRegionFromWindowOpcode:  "CreateRegionFromWindow",
	CreateRegionFromGCOpcode:      "CreateRegionFromGC",
	CreateRegionFromPictureOpcode: "CreateRegionFromPicture",
	DestroyRegionOpcode:           "DestroyRegion",
	SetRegionOpcode:               "SetRegion",
	CopyRegionOpcode:              "CopyRegion",
	UnionRegionOpcode:             "UnionRegion",
	IntersectRegionOpcode:         "IntersectRegion",
	SubtractRegionOpcode:          "SubtractRegion",
	InvertRegionOpcode:            "InvertRegion",
	TranslateRegionOpcode:         "TranslateRegion",
	RegionExtentsOpcode:           "RegionExtents",
	FetchRegionOpcode:             "FetchRegion",
	SetGCClipRegionOpcode:         "SetGCClipRegion",
	SetWindowShapeRegionOpcode:    "SetWindowShapeRegion",
	SetPictureClipRegionOpcode:    "SetPictureClipRegion",
	SetCursorNameOpcode:           "SetCursorName",
	GetCursorNameOpcode:           "GetCursorName",
	GetCursorImageAndNameOpcode:   "GetCursorImageAndName",
	ChangeCursorOpcode:            "ChangeCursor",
	ChangeCursorByNameOpcode:      "ChangeCursorByName",
	ExpandRegionOpcode:            "ExpandRegion",
	HideCursorOpcode:              "HideCursor",
	ShowCursorOpcode:              "ShowCursor",
	CreatePointerBarrierOpcode:    "CreatePointerBarrier",
	DeletePointerBarrierOpcode:    "DeletePointerBarrier",
}

func init() {
	_ext = x.NewExtension("XFIXES", 0, errorCodeNameMap, requestOpcodeNameMap)
}
