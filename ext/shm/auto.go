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
	readCompletionEvent(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

const BadSegErrorCode = 0
const QueryVersionOpcode = 0

type QueryVersionCookie uint64

const AttachOpcode = 1
const DetachOpcode = 2
const PutImageOpcode = 3
const GetImageOpcode = 4

type GetImageCookie uint64

const CreatePixmapOpcode = 5
const AttachFdOpcode = 6
const CreateSegmentOpcode = 7

type CreateSegmentCookie uint64

var readErrorFuncMap = make(map[uint8]x.ReadErrorFunc, 1)

func init() {
	readErrorFuncMap[BadSegErrorCode] = readBadSegError
	_ext = x.NewExtension("MIT-SHM", 0, readErrorFuncMap)
}
