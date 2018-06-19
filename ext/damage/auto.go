package damage

import x "github.com/linuxdeepin/go-x11-client"

// _ns.ext_name: Damage
const MajorVersion = 1
const MinorVersion = 1

var _ext = x.NewExtension("DAMAGE")

func Ext() *x.Extension {
	return _ext
}

// simple ('xcb', 'Damage', 'DAMAGE')
type Damage uint32

// enum ReportLevel
const (
	ReportLevelRawRectangles   = 0
	ReportLevelDeltaRectangles = 1
	ReportLevelBoundingBox     = 2
	ReportLevelNonEmpty        = 3
)

const BadDamageErrorCode = 0
const QueryVersionOpcode = 0

type QueryVersionCookie uint64

const CreateOpcode = 1
const DestroyOpcode = 2
const SubtractOpcode = 3
const AddOpcode = 4
const NotifyEventCode = 0

func NewNotifyEvent(data []byte) (*NotifyEvent, error) {
	var ev NotifyEvent
	r := x.NewReaderFromData(data)
	readNotifyEvent(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

// x_prefix is x.
var readErrorFuncMap = make(map[uint8]func(r *x.Reader) x.Error)

func init() {
	readErrorFuncMap[BadDamageErrorCode] = readBadDamageError
}
