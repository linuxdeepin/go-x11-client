package x

func CreateWindow(conn *Conn, depth uint8, wid, parent Window, x, y int16, width, height, borderWidth uint16, class uint16, visual VisualID, valueMask uint32, valueList []uint32) {
	w := NewWriter()
	writeCreateWindow(w, depth, wid, parent, x, y, width, height, borderWidth, class, visual, valueMask, valueList)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  CreateWindowOpcode,
	}
	conn.SendRequest(0, d, req)
}

func CreateWindowChecked(conn *Conn, depth uint8, wid, parent Window, x, y int16, width, height, borderWidth uint16, class uint16, visual VisualID, valueMask uint32, valueList []uint32) VoidCookie {
	w := NewWriter()
	writeCreateWindow(w, depth, wid, parent, x, y, width, height, borderWidth, class, visual, valueMask, valueList)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  CreateWindowOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func ChangeWindowAttributes(conn *Conn, window Window, valueMask uint32, valueList []uint32) {
	w := NewWriter()
	writeChangeWindowAttributes(w, window, valueMask, valueList)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  ChangeWindowAttributesOpcode,
	}
	conn.SendRequest(0, d, req)
}

func ChangeWindowAttributesChecked(conn *Conn, window Window, valueMask uint32, valueList []uint32) VoidCookie {
	w := NewWriter()
	writeChangeWindowAttributes(w, window, valueMask, valueList)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  ChangeWindowAttributesOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func GetWindowAttributes(conn *Conn, window Window) GetWindowAttributesCookie {
	w := NewWriter()
	writeGetWindowAttributes(w, window)
	d := w.Bytes()
	req := &ProtocolRequest{
		Opcode: GetWindowAttributesOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return GetWindowAttributesCookie(seq)
}

func (cookie GetWindowAttributesCookie) Reply(conn *Conn) (*GetWindowAttributesReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := NewReaderFromData(replyBuf)
	var reply GetWindowAttributesReply
	err = readGetWindowAttributesReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func DestroyWindow(conn *Conn, window Window) {
	w := NewWriter()
	writeDestroyWindow(w, window)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  DestroyWindowOpcode,
	}
	conn.SendRequest(0, d, req)
}

func DestroyWindowChecked(conn *Conn, window Window) VoidCookie {
	w := NewWriter()
	writeDestroyWindow(w, window)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  DestroyWindowOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func DestroySubwindows(conn *Conn, window Window) {
	w := NewWriter()
	writeDestroySubwindows(w, window)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  DestroySubwindowsOpcode,
	}
	conn.SendRequest(0, d, req)
}

func DestroySubwindowsChecked(conn *Conn, window Window) VoidCookie {
	w := NewWriter()
	writeDestroySubwindows(w, window)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  DestroySubwindowsOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func ChangeSaveSet(conn *Conn, mode uint8, window Window) {
	w := NewWriter()
	writeChangeSaveSet(w, mode, window)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  ChangeSaveSetOpcode,
	}
	conn.SendRequest(0, d, req)
}

func ChangeSaveSetChecked(conn *Conn, mode uint8, window Window) VoidCookie {
	w := NewWriter()
	writeChangeSaveSet(w, mode, window)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  ChangeSaveSetOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func ReparentWindow(conn *Conn, window, parent Window, x, y int16) {
	w := NewWriter()
	writeReparentWindow(w, window, parent, x, y)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  ReparentWindowOpcode,
	}
	conn.SendRequest(0, d, req)
}

func ReparentWindowChecked(conn *Conn, window, parent Window, x, y int16) VoidCookie {
	w := NewWriter()
	writeReparentWindow(w, window, parent, x, y)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  ReparentWindowOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func MapWindow(conn *Conn, window Window) {
	w := NewWriter()
	writeMapWindow(w, window)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  MapWindowOpcode,
	}
	conn.SendRequest(0, d, req)
}

func MapWindowChecked(conn *Conn, window Window) VoidCookie {
	w := NewWriter()
	writeMapWindow(w, window)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  MapWindowOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func MapSubwindows(conn *Conn, window Window) {
	w := NewWriter()
	writeMapSubwindows(w, window)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  MapSubwindowsOpcode,
	}
	conn.SendRequest(0, d, req)
}

func MapSubwindowsChecked(conn *Conn, window Window) VoidCookie {
	w := NewWriter()
	writeMapSubwindows(w, window)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  MapSubwindowsOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func UnmapWindow(conn *Conn, window Window) {
	w := NewWriter()
	writeUnmapWindow(w, window)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  UnmapWindowOpcode,
	}
	conn.SendRequest(0, d, req)
}

func UnmapWindowChecked(conn *Conn, window Window) VoidCookie {
	w := NewWriter()
	writeUnmapWindow(w, window)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  UnmapWindowOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func UnmapSubwindows(conn *Conn, window Window) {
	w := NewWriter()
	writeUnmapSubwindows(w, window)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  UnmapSubwindowsOpcode,
	}
	conn.SendRequest(0, d, req)
}

func UnmapSubwindowsChecked(conn *Conn, window Window) VoidCookie {
	w := NewWriter()
	writeUnmapSubwindows(w, window)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  UnmapSubwindowsOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func ConfigureWindow(conn *Conn, window Window, valueMask uint16, valueList []uint32) {
	w := NewWriter()
	writeConfigureWindow(w, window, valueMask, valueList)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  ConfigureWindowOpcode,
	}
	conn.SendRequest(0, d, req)
}

func ConfigureWindowChecked(conn *Conn, window Window, valueMask uint16, valueList []uint32) VoidCookie {
	w := NewWriter()
	writeConfigureWindow(w, window, valueMask, valueList)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  ConfigureWindowOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func CirculateWindow(conn *Conn, direction uint8, window Window) {
	w := NewWriter()
	writeCirculateWindow(w, direction, window)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  CirculateWindowOpcode,
	}
	conn.SendRequest(0, d, req)
}

func CirculateWindowChecked(conn *Conn, direction uint8, window Window) VoidCookie {
	w := NewWriter()
	writeCirculateWindow(w, direction, window)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  CirculateWindowOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func GetGeometry(conn *Conn, drawable Drawable) GetGeometryCookie {
	w := NewWriter()
	writeGetGeometry(w, drawable)
	d := w.Bytes()
	req := &ProtocolRequest{
		Opcode: GetGeometryOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return GetGeometryCookie(seq)
}

func (cookie GetGeometryCookie) Reply(conn *Conn) (*GetGeometryReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := NewReaderFromData(replyBuf)
	var reply GetGeometryReply
	err = readGetGeometryReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func QueryTree(conn *Conn, window Window) QueryTreeCookie {
	w := NewWriter()
	writeQueryTree(w, window)
	d := w.Bytes()
	req := &ProtocolRequest{
		Opcode: QueryTreeOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return QueryTreeCookie(seq)
}

func (cookie QueryTreeCookie) Reply(conn *Conn) (*QueryTreeReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := NewReaderFromData(replyBuf)
	var reply QueryTreeReply
	err = readQueryTreeReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func InternAtom(conn *Conn, onlyIfExists bool, name string) InternAtomCookie {
	w := NewWriter()
	writeInternAtom(w, onlyIfExists, name)
	d := w.Bytes()
	req := &ProtocolRequest{
		Opcode: InternAtomOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return InternAtomCookie(seq)
}

func (cookie InternAtomCookie) Reply(conn *Conn) (*InternAtomReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := NewReaderFromData(replyBuf)
	var reply InternAtomReply
	err = readInternAtomReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func GetAtomName(conn *Conn, atom Atom) GetAtomNameCookie {
	w := NewWriter()
	writeGetAtomName(w, atom)
	d := w.Bytes()
	req := &ProtocolRequest{
		Opcode: GetAtomNameOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return GetAtomNameCookie(seq)
}

func (cookie GetAtomNameCookie) Reply(conn *Conn) (*GetAtomNameReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := NewReaderFromData(replyBuf)
	var reply GetAtomNameReply
	err = readGetAtomNameReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func ChangeProperty(conn *Conn, mode uint8, window Window, property, type0 Atom, format uint8, data []byte) {
	w := NewWriter()
	writeChangeProperty(w, mode, window, property, type0, format, data)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  ChangePropertyOpcode,
	}
	conn.SendRequest(0, d, req)
}

func ChangePropertyChecked(conn *Conn, mode uint8, window Window, property, type0 Atom, format uint8, data []byte) VoidCookie {
	w := NewWriter()
	writeChangeProperty(w, mode, window, property, type0, format, data)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  ChangePropertyOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func DeleteProperty(conn *Conn, window Window, property Atom) {
	w := NewWriter()
	writeDeleteProperty(w, window, property)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  DeletePropertyOpcode,
	}
	conn.SendRequest(0, d, req)
}

func DeletePropertyChecked(conn *Conn, window Window, property Atom) VoidCookie {
	w := NewWriter()
	writeDeleteProperty(w, window, property)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  DeletePropertyOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func GetProperty(conn *Conn, delete bool, window Window, property, type0 Atom, longOffset, longLength uint32) GetPropertyCookie {
	w := NewWriter()
	writeGetProperty(w, delete, window, property, type0, longOffset, longLength)
	d := w.Bytes()
	req := &ProtocolRequest{
		Opcode: GetPropertyOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return GetPropertyCookie(seq)
}

func (cookie GetPropertyCookie) Reply(conn *Conn) (*GetPropertyReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := NewReaderFromData(replyBuf)
	var reply GetPropertyReply
	err = readGetPropertyReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func ListProperties(conn *Conn, window Window) ListPropertiesCookie {
	w := NewWriter()
	writeListProperties(w, window)
	d := w.Bytes()
	req := &ProtocolRequest{
		Opcode: ListPropertiesOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return ListPropertiesCookie(seq)
}

func (cookie ListPropertiesCookie) Reply(conn *Conn) (*ListPropertiesReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := NewReaderFromData(replyBuf)
	var reply ListPropertiesReply
	err = readListPropertiesReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func SetSelectionOwner(conn *Conn, owner Window, selection Atom, time Timestamp) {
	w := NewWriter()
	writeSetSelectionOwner(w, owner, selection, time)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  SetSelectionOwnerOpcode,
	}
	conn.SendRequest(0, d, req)
}

func SetSelectionOwnerChecked(conn *Conn, owner Window, selection Atom, time Timestamp) VoidCookie {
	w := NewWriter()
	writeSetSelectionOwner(w, owner, selection, time)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  SetSelectionOwnerOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func GetSelectionOwner(conn *Conn, selection Atom) GetSelectionOwnerCookie {
	w := NewWriter()
	writeGetSelectionOwner(w, selection)
	d := w.Bytes()
	req := &ProtocolRequest{
		Opcode: GetSelectionOwnerOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return GetSelectionOwnerCookie(seq)
}

func (cookie GetSelectionOwnerCookie) Reply(conn *Conn) (*GetSelectionOwnerReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := NewReaderFromData(replyBuf)
	var reply GetSelectionOwnerReply
	err = readGetSelectionOwnerReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func ConvertSelection(conn *Conn, requestor Window, selection, target, property Atom, time Timestamp) {
	w := NewWriter()
	writeConvertSelection(w, requestor, selection, target, property, time)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  ConvertSelectionOpcode,
	}
	conn.SendRequest(0, d, req)
}

func ConvertSelectionChecked(conn *Conn, requestor Window, selection, target, property Atom, time Timestamp) VoidCookie {
	w := NewWriter()
	writeConvertSelection(w, requestor, selection, target, property, time)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  ConvertSelectionOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func SendEvent(conn *Conn, propagate bool, destination Window, eventMask uint32, event []byte) {
	w := NewWriter()
	writeSendEvent(w, propagate, destination, eventMask, event)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  SendEventOpcode,
	}
	conn.SendRequest(0, d, req)
}

func SendEventChecked(conn *Conn, propagate bool, destination Window, eventMask uint32, event []byte) VoidCookie {
	w := NewWriter()
	writeSendEvent(w, propagate, destination, eventMask, event)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  SendEventOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func GrabPointer(conn *Conn, ownerEvents bool, grabWindow Window, eventMask uint16, pointerMode, keyboardMode uint8, confineTo Window, cursor Cursor, time Timestamp) GrabPointerCookie {
	w := NewWriter()
	writeGrabPointer(w, ownerEvents, grabWindow, eventMask, pointerMode, keyboardMode, confineTo, cursor, time)
	d := w.Bytes()
	req := &ProtocolRequest{
		Opcode: GrabPointerOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return GrabPointerCookie(seq)
}

func (cookie GrabPointerCookie) Reply(conn *Conn) (*GrabPointerReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := NewReaderFromData(replyBuf)
	var reply GrabPointerReply
	err = readGrabPointerReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func UngrabPointer(conn *Conn, time Timestamp) {
	w := NewWriter()
	writeUngrabPointer(w, time)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  UngrabPointerOpcode,
	}
	conn.SendRequest(0, d, req)
}

func UngrabPointerChecked(conn *Conn, time Timestamp) VoidCookie {
	w := NewWriter()
	writeUngrabPointer(w, time)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  UngrabPointerOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func GrabButton(conn *Conn, ownerEvents bool, grabWindow Window, eventMask uint16, pointerMode, keyboardMode uint8, confineTo Window, cursor Cursor, button uint8, modifiers uint16) {
	w := NewWriter()
	writeGrabButton(w, ownerEvents, grabWindow, eventMask, pointerMode, keyboardMode, confineTo, cursor, button, modifiers)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  GrabButtonOpcode,
	}
	conn.SendRequest(0, d, req)
}

func GrabButtonChecked(conn *Conn, ownerEvents bool, grabWindow Window, eventMask uint16, pointerMode, keyboardMode uint8, confineTo Window, cursor Cursor, button uint8, modifiers uint16) VoidCookie {
	w := NewWriter()
	writeGrabButton(w, ownerEvents, grabWindow, eventMask, pointerMode, keyboardMode, confineTo, cursor, button, modifiers)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  GrabButtonOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func UngrabButton(conn *Conn, button uint8, grabWindow Window, modifiers uint16) {
	w := NewWriter()
	writeUngrabButton(w, button, grabWindow, modifiers)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  UngrabButtonOpcode,
	}
	conn.SendRequest(0, d, req)
}

func UngrabButtonChecked(conn *Conn, button uint8, grabWindow Window, modifiers uint16) VoidCookie {
	w := NewWriter()
	writeUngrabButton(w, button, grabWindow, modifiers)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  UngrabButtonOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func ChangeActivePointerGrab(conn *Conn, cursor Cursor, time Timestamp, eventMask uint16) {
	w := NewWriter()
	writeChangeActivePointerGrab(w, cursor, time, eventMask)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  ChangeActivePointerGrabOpcode,
	}
	conn.SendRequest(0, d, req)
}

func ChangeActivePointerGrabChecked(conn *Conn, cursor Cursor, time Timestamp, eventMask uint16) VoidCookie {
	w := NewWriter()
	writeChangeActivePointerGrab(w, cursor, time, eventMask)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  ChangeActivePointerGrabOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func GrabKeyboard(conn *Conn, ownerEvents bool, grabWindow Window, time Timestamp, pointerMode, keyboardMode uint8) GrabKeyboardCookie {
	w := NewWriter()
	writeGrabKeyboard(w, ownerEvents, grabWindow, time, pointerMode, keyboardMode)
	d := w.Bytes()
	req := &ProtocolRequest{
		Opcode: GrabKeyboardOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return GrabKeyboardCookie(seq)
}

func (cookie GrabKeyboardCookie) Reply(conn *Conn) (*GrabKeyboardReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := NewReaderFromData(replyBuf)
	var reply GrabKeyboardReply
	err = readGrabKeyboardReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func UngrabKeyboard(conn *Conn, time Timestamp) {
	w := NewWriter()
	writeUngrabKeyboard(w, time)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  UngrabKeyboardOpcode,
	}
	conn.SendRequest(0, d, req)
}

func UngrabKeyboardChecked(conn *Conn, time Timestamp) VoidCookie {
	w := NewWriter()
	writeUngrabKeyboard(w, time)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  UngrabKeyboardOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func GrabKey(conn *Conn, ownerEvents bool, grabWindow Window, modifiers uint16, key Keycode, pointerMode, keyboardMode uint8) {
	w := NewWriter()
	writeGrabKey(w, ownerEvents, grabWindow, modifiers, key, pointerMode, keyboardMode)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  GrabKeyOpcode,
	}
	conn.SendRequest(0, d, req)
}

func GrabKeyChecked(conn *Conn, ownerEvents bool, grabWindow Window, modifiers uint16, key Keycode, pointerMode, keyboardMode uint8) VoidCookie {
	w := NewWriter()
	writeGrabKey(w, ownerEvents, grabWindow, modifiers, key, pointerMode, keyboardMode)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  GrabKeyOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func UngrabKey(conn *Conn, key Keycode, grabWindow Window, modifiers uint16) {
	w := NewWriter()
	writeUngrabKey(w, key, grabWindow, modifiers)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  UngrabKeyOpcode,
	}
	conn.SendRequest(0, d, req)
}

func UngrabKeyChecked(conn *Conn, key Keycode, grabWindow Window, modifiers uint16) VoidCookie {
	w := NewWriter()
	writeUngrabKey(w, key, grabWindow, modifiers)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  UngrabKeyOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func AllowEvents(conn *Conn, mode uint8, time Timestamp) {
	w := NewWriter()
	writeAllowEvents(w, mode, time)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  AllowEventsOpcode,
	}
	conn.SendRequest(0, d, req)
}

func AllowEventsChecked(conn *Conn, mode uint8, time Timestamp) VoidCookie {
	w := NewWriter()
	writeAllowEvents(w, mode, time)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  AllowEventsOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func GrabServer(conn *Conn) {
	w := NewWriter()
	writeGrabServer(w)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  GrabServerOpcode,
	}
	conn.SendRequest(0, d, req)
}

func GrabServerChecked(conn *Conn) VoidCookie {
	w := NewWriter()
	writeGrabServer(w)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  GrabServerOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func UngrabServer(conn *Conn) {
	w := NewWriter()
	writeUngrabServer(w)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  UngrabServerOpcode,
	}
	conn.SendRequest(0, d, req)
}

func UngrabServerChecked(conn *Conn) VoidCookie {
	w := NewWriter()
	writeUngrabServer(w)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  UngrabServerOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func QueryPointer(conn *Conn, window Window) QueryPointerCookie {
	w := NewWriter()
	writeQueryPointer(w, window)
	d := w.Bytes()
	req := &ProtocolRequest{
		Opcode: QueryPointerOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return QueryPointerCookie(seq)
}

func (cookie QueryPointerCookie) Reply(conn *Conn) (*QueryPointerReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := NewReaderFromData(replyBuf)
	var reply QueryPointerReply
	err = readQueryPointerReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func GetMotionEvents(conn *Conn, window Window, start, stop Timestamp) GetMotionEventsCookie {
	w := NewWriter()
	writeGetMotionEvents(w, window, start, stop)
	d := w.Bytes()
	req := &ProtocolRequest{
		Opcode: GetMotionEventsOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return GetMotionEventsCookie(seq)
}

func (cookie GetMotionEventsCookie) Reply(conn *Conn) (*GetMotionEventsReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := NewReaderFromData(replyBuf)
	var reply GetMotionEventsReply
	err = readGetMotionEventsReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func TranslateCoordinates(conn *Conn, srcWindow, dstWindow Window, srcX, srcY int16) TranslateCoordinatesCookie {
	w := NewWriter()
	writeTranslateCoordinates(w, srcWindow, dstWindow, srcX, srcY)
	d := w.Bytes()
	req := &ProtocolRequest{
		Opcode: TranslateCoordinatesOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return TranslateCoordinatesCookie(seq)
}

func (cookie TranslateCoordinatesCookie) Reply(conn *Conn) (*TranslateCoordinatesReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := NewReaderFromData(replyBuf)
	var reply TranslateCoordinatesReply
	err = readTranslateCoordinatesReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func WarpPointer(conn *Conn, srcWindow, dstWindow Window, srcX, srcY int16, srcWidth, srcHeight uint16, dstX, dstY int16) {
	w := NewWriter()
	writeWarpPointer(w, srcWindow, dstWindow, srcX, srcY, srcWidth, srcHeight, dstX, dstY)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  WarpPointerOpcode,
	}
	conn.SendRequest(0, d, req)
}

func WarpPointerChecked(conn *Conn, srcWindow, dstWindow Window, srcX, srcY int16, srcWidth, srcHeight uint16, dstX, dstY int16) VoidCookie {
	w := NewWriter()
	writeWarpPointer(w, srcWindow, dstWindow, srcX, srcY, srcWidth, srcHeight, dstX, dstY)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  WarpPointerOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func SetInputFocus(conn *Conn, revertTo uint8, focus Window, time Timestamp) {
	w := NewWriter()
	writeSetInputFocus(w, revertTo, focus, time)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  SetInputFocusOpcode,
	}
	conn.SendRequest(0, d, req)
}

func SetInputFocusChecked(conn *Conn, revertTo uint8, focus Window, time Timestamp) VoidCookie {
	w := NewWriter()
	writeSetInputFocus(w, revertTo, focus, time)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  SetInputFocusOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func GetInputFocus(conn *Conn) GetInputFocusCookie {
	w := NewWriter()
	writeGetInputFocus(w)
	d := w.Bytes()
	req := &ProtocolRequest{
		Opcode: GetInputFocusOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return GetInputFocusCookie(seq)
}

func (cookie GetInputFocusCookie) Reply(conn *Conn) (*GetInputFocusReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := NewReaderFromData(replyBuf)
	var reply GetInputFocusReply
	err = readGetInputFocusReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func QueryKeymap(conn *Conn) QueryKeymapCookie {
	w := NewWriter()
	writeQueryKeymap(w)
	d := w.Bytes()
	req := &ProtocolRequest{
		Opcode: QueryKeymapOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return QueryKeymapCookie(seq)
}

func (cookie QueryKeymapCookie) Reply(conn *Conn) (*QueryKeymapReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := NewReaderFromData(replyBuf)
	var reply QueryKeymapReply
	err = readQueryKeymapReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func OpenFont(conn *Conn, fid Font, name string) {
	w := NewWriter()
	writeOpenFont(w, fid, name)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  OpenFontOpcode,
	}
	conn.SendRequest(0, d, req)
}

func OpenFontChecked(conn *Conn, fid Font, name string) VoidCookie {
	w := NewWriter()
	writeOpenFont(w, fid, name)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  OpenFontOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func CloseFont(conn *Conn, font Font) {
	w := NewWriter()
	writeCloseFont(w, font)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  CloseFontOpcode,
	}
	conn.SendRequest(0, d, req)
}

func CloseFontChecked(conn *Conn, font Font) VoidCookie {
	w := NewWriter()
	writeCloseFont(w, font)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  CloseFontOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func QueryFont(conn *Conn, font Fontable) QueryFontCookie {
	w := NewWriter()
	writeQueryFont(w, font)
	d := w.Bytes()
	req := &ProtocolRequest{
		Opcode: QueryFontOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return QueryFontCookie(seq)
}

func (cookie QueryFontCookie) Reply(conn *Conn) (*QueryFontReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := NewReaderFromData(replyBuf)
	var reply QueryFontReply
	err = readQueryFontReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func ListFonts(conn *Conn, maxNames uint16, pattern string) ListFontsCookie {
	w := NewWriter()
	writeListFonts(w, maxNames, pattern)
	d := w.Bytes()
	req := &ProtocolRequest{
		Opcode: ListFontsOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return ListFontsCookie(seq)
}

func (cookie ListFontsCookie) Reply(conn *Conn) (*ListFontsReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := NewReaderFromData(replyBuf)
	var reply ListFontsReply
	err = readListFontsReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func ListFontsWithInfo(conn *Conn, maxNames uint16, pattern string) ListFontsWithInfoCookie {
	w := NewWriter()
	writeListFontsWithInfo(w, maxNames, pattern)
	d := w.Bytes()
	req := &ProtocolRequest{
		Opcode: ListFontsWithInfoOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return ListFontsWithInfoCookie(seq)
}

func (cookie ListFontsWithInfoCookie) Reply(conn *Conn) (*ListFontsWithInfoReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := NewReaderFromData(replyBuf)
	var reply ListFontsWithInfoReply
	err = readListFontsWithInfoReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func SetFontPath(conn *Conn, paths []string) {
	w := NewWriter()
	writeSetFontPath(w, paths)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  SetFontPathOpcode,
	}
	conn.SendRequest(0, d, req)
}

func SetFontPathChecked(conn *Conn, paths []string) VoidCookie {
	w := NewWriter()
	writeSetFontPath(w, paths)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  SetFontPathOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func GetFontPath(conn *Conn) GetFontPathCookie {
	w := NewWriter()
	writeGetFontPath(w)
	d := w.Bytes()
	req := &ProtocolRequest{
		Opcode: GetFontPathOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return GetFontPathCookie(seq)
}

func (cookie GetFontPathCookie) Reply(conn *Conn) (*GetFontPathReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := NewReaderFromData(replyBuf)
	var reply GetFontPathReply
	err = readGetFontPathReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func CreatePixmap(conn *Conn, depth uint8, pid Pixmap, drawable Drawable, width, height uint16) {
	w := NewWriter()
	writeCreatePixmap(w, depth, pid, drawable, width, height)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  CreatePixmapOpcode,
	}
	conn.SendRequest(0, d, req)
}

func CreatePixmapChecked(conn *Conn, depth uint8, pid Pixmap, drawable Drawable, width, height uint16) VoidCookie {
	w := NewWriter()
	writeCreatePixmap(w, depth, pid, drawable, width, height)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  CreatePixmapOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func FreePixmap(conn *Conn, pixmap Pixmap) {
	w := NewWriter()
	writeFreePixmap(w, pixmap)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  FreePixmapOpcode,
	}
	conn.SendRequest(0, d, req)
}

func FreePixmapChecked(conn *Conn, pixmap Pixmap) VoidCookie {
	w := NewWriter()
	writeFreePixmap(w, pixmap)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  FreePixmapOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func CreateGC(conn *Conn, cid GContext, drawable Drawable, valueMask uint32, valueList []uint32) {
	w := NewWriter()
	writeCreateGC(w, cid, drawable, valueMask, valueList)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  CreateGCOpcode,
	}
	conn.SendRequest(0, d, req)
}

func CreateGCChecked(conn *Conn, cid GContext, drawable Drawable, valueMask uint32, valueList []uint32) VoidCookie {
	w := NewWriter()
	writeCreateGC(w, cid, drawable, valueMask, valueList)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  CreateGCOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func ChangeGC(conn *Conn, gc GContext, valueMask uint32, valueList []uint32) {
	w := NewWriter()
	writeChangeGC(w, gc, valueMask, valueList)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  ChangeGCOpcode,
	}
	conn.SendRequest(0, d, req)
}

func ChangeGCChecked(conn *Conn, gc GContext, valueMask uint32, valueList []uint32) VoidCookie {
	w := NewWriter()
	writeChangeGC(w, gc, valueMask, valueList)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  ChangeGCOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func CopyGC(conn *Conn, srcGC, dstGC GContext, valueMask uint32) {
	w := NewWriter()
	writeCopyGC(w, srcGC, dstGC, valueMask)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  CopyGCOpcode,
	}
	conn.SendRequest(0, d, req)
}

func CopyGCChecked(conn *Conn, srcGC, dstGC GContext, valueMask uint32) VoidCookie {
	w := NewWriter()
	writeCopyGC(w, srcGC, dstGC, valueMask)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  CopyGCOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func SetDashes(conn *Conn, gc GContext, dashOffset uint16, dashes []uint8) {
	w := NewWriter()
	writeSetDashes(w, gc, dashOffset, dashes)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  SetDashesOpcode,
	}
	conn.SendRequest(0, d, req)
}

func SetDashesChecked(conn *Conn, gc GContext, dashOffset uint16, dashes []uint8) VoidCookie {
	w := NewWriter()
	writeSetDashes(w, gc, dashOffset, dashes)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  SetDashesOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func SetClipRectangles(conn *Conn, ordering uint8, gc GContext, clipXOrigin, clipYOrigin int16, rectangles []Rectangle) {
	w := NewWriter()
	writeSetClipRectangles(w, ordering, gc, clipXOrigin, clipYOrigin, rectangles)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  SetClipRectanglesOpcode,
	}
	conn.SendRequest(0, d, req)
}

func SetClipRectanglesChecked(conn *Conn, ordering uint8, gc GContext, clipXOrigin, clipYOrigin int16, rectangles []Rectangle) VoidCookie {
	w := NewWriter()
	writeSetClipRectangles(w, ordering, gc, clipXOrigin, clipYOrigin, rectangles)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  SetClipRectanglesOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func FreeGC(conn *Conn, gc GContext) {
	w := NewWriter()
	writeFreeGC(w, gc)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  FreeGCOpcode,
	}
	conn.SendRequest(0, d, req)
}

func FreeGCChecked(conn *Conn, gc GContext) VoidCookie {
	w := NewWriter()
	writeFreeGC(w, gc)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  FreeGCOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func ClearArea(conn *Conn, exposures bool, window Window, x, y int16, width, height uint16) {
	w := NewWriter()
	writeClearArea(w, exposures, window, x, y, width, height)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  ClearAreaOpcode,
	}
	conn.SendRequest(0, d, req)
}

func ClearAreaChecked(conn *Conn, exposures bool, window Window, x, y int16, width, height uint16) VoidCookie {
	w := NewWriter()
	writeClearArea(w, exposures, window, x, y, width, height)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  ClearAreaOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func CopyArea(conn *Conn, srcDrawable, dstDrawable Drawable, gc GContext, srcX, srcY, dstX, dstY int16, width, height uint16) {
	w := NewWriter()
	writeCopyArea(w, srcDrawable, dstDrawable, gc, srcX, srcY, dstX, dstY, width, height)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  CopyAreaOpcode,
	}
	conn.SendRequest(0, d, req)
}

func CopyAreaChecked(conn *Conn, srcDrawable, dstDrawable Drawable, gc GContext, srcX, srcY, dstX, dstY int16, width, height uint16) VoidCookie {
	w := NewWriter()
	writeCopyArea(w, srcDrawable, dstDrawable, gc, srcX, srcY, dstX, dstY, width, height)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  CopyAreaOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func CopyPlane(conn *Conn, srcDrawable, dstDrawable Drawable, gc GContext, srcX, srcY, dstX, dstY int16, width, height uint16, bitPlane uint32) {
	w := NewWriter()
	writeCopyPlane(w, srcDrawable, dstDrawable, gc, srcX, srcY, dstX, dstY, width, height, bitPlane)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  CopyPlaneOpcode,
	}
	conn.SendRequest(0, d, req)
}

func CopyPlaneChecked(conn *Conn, srcDrawable, dstDrawable Drawable, gc GContext, srcX, srcY, dstX, dstY int16, width, height uint16, bitPlane uint32) VoidCookie {
	w := NewWriter()
	writeCopyPlane(w, srcDrawable, dstDrawable, gc, srcX, srcY, dstX, dstY, width, height, bitPlane)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  CopyPlaneOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func PolyPoint(conn *Conn, coordinateMode uint8, drawable Drawable, gc GContext, points []Point) {
	w := NewWriter()
	writePolyPoint(w, coordinateMode, drawable, gc, points)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  PolyPointOpcode,
	}
	conn.SendRequest(0, d, req)
}

func PolyPointChecked(conn *Conn, coordinateMode uint8, drawable Drawable, gc GContext, points []Point) VoidCookie {
	w := NewWriter()
	writePolyPoint(w, coordinateMode, drawable, gc, points)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  PolyPointOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func PolyLine(conn *Conn, coordinateMode uint8, drawable Drawable, gc GContext, points []Point) {
	w := NewWriter()
	writePolyLine(w, coordinateMode, drawable, gc, points)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  PolyLineOpcode,
	}
	conn.SendRequest(0, d, req)
}

func PolyLineChecked(conn *Conn, coordinateMode uint8, drawable Drawable, gc GContext, points []Point) VoidCookie {
	w := NewWriter()
	writePolyLine(w, coordinateMode, drawable, gc, points)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  PolyLineOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func PolySegment(conn *Conn, drawable Drawable, gc GContext, segments []Segment) {
	w := NewWriter()
	writePolySegment(w, drawable, gc, segments)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  PolySegmentOpcode,
	}
	conn.SendRequest(0, d, req)
}

func PolySegmentChecked(conn *Conn, drawable Drawable, gc GContext, segments []Segment) VoidCookie {
	w := NewWriter()
	writePolySegment(w, drawable, gc, segments)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  PolySegmentOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func PolyRectangle(conn *Conn, drawable Drawable, gc GContext, rectangles []Rectangle) {
	w := NewWriter()
	writePolyRectangle(w, drawable, gc, rectangles)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  PolyRectangleOpcode,
	}
	conn.SendRequest(0, d, req)
}

func PolyRectangleChecked(conn *Conn, drawable Drawable, gc GContext, rectangles []Rectangle) VoidCookie {
	w := NewWriter()
	writePolyRectangle(w, drawable, gc, rectangles)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  PolyRectangleOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func PolyArc(conn *Conn, drawable Drawable, gc GContext, arcs []Arc) {
	w := NewWriter()
	writePolyArc(w, drawable, gc, arcs)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  PolyArcOpcode,
	}
	conn.SendRequest(0, d, req)
}

func PolyArcChecked(conn *Conn, drawable Drawable, gc GContext, arcs []Arc) VoidCookie {
	w := NewWriter()
	writePolyArc(w, drawable, gc, arcs)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  PolyArcOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func FillPoly(conn *Conn, drawable Drawable, gc GContext, shape uint8, coordinateMode uint8, points []Point) {
	w := NewWriter()
	writeFillPoly(w, drawable, gc, shape, coordinateMode, points)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  FillPolyOpcode,
	}
	conn.SendRequest(0, d, req)
}

func FillPolyChecked(conn *Conn, drawable Drawable, gc GContext, shape uint8, coordinateMode uint8, points []Point) VoidCookie {
	w := NewWriter()
	writeFillPoly(w, drawable, gc, shape, coordinateMode, points)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  FillPolyOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func PolyFillRectangle(conn *Conn, drawable Drawable, gc GContext, rectangles []Rectangle) {
	w := NewWriter()
	writePolyFillRectangle(w, drawable, gc, rectangles)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  PolyFillRectangleOpcode,
	}
	conn.SendRequest(0, d, req)
}

func PolyFillRectangleChecked(conn *Conn, drawable Drawable, gc GContext, rectangles []Rectangle) VoidCookie {
	w := NewWriter()
	writePolyFillRectangle(w, drawable, gc, rectangles)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  PolyFillRectangleOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func PolyFillArc(conn *Conn, drawable Drawable, gc GContext, arcs []Arc) {
	w := NewWriter()
	writePolyFillArc(w, drawable, gc, arcs)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  PolyFillArcOpcode,
	}
	conn.SendRequest(0, d, req)
}

func PolyFillArcChecked(conn *Conn, drawable Drawable, gc GContext, arcs []Arc) VoidCookie {
	w := NewWriter()
	writePolyFillArc(w, drawable, gc, arcs)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  PolyFillArcOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func PutImage(conn *Conn, format uint8, drawable Drawable, gc GContext, width, height uint16, dstX, dstY int16, leftPad, depth uint8, data []byte) {
	w := NewWriter()
	writePutImage(w, format, drawable, gc, width, height, dstX, dstY, leftPad, depth, data)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  PutImageOpcode,
	}
	conn.SendRequest(0, d, req)
}

func PutImageChecked(conn *Conn, format uint8, drawable Drawable, gc GContext, width, height uint16, dstX, dstY int16, leftPad, depth uint8, data []byte) VoidCookie {
	w := NewWriter()
	writePutImage(w, format, drawable, gc, width, height, dstX, dstY, leftPad, depth, data)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  PutImageOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func GetImage(conn *Conn, format uint8, drawable Drawable, x, y int16, width, height uint16, planeMask uint32) GetImageCookie {
	w := NewWriter()
	writeGetImage(w, format, drawable, x, y, width, height, planeMask)
	d := w.Bytes()
	req := &ProtocolRequest{
		Opcode: GetImageOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return GetImageCookie(seq)
}

func (cookie GetImageCookie) Reply(conn *Conn) (*GetImageReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := NewReaderFromData(replyBuf)
	var reply GetImageReply
	err = readGetImageReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func QueryExtension(conn *Conn, name string) QueryExtensionCookie {
	w := NewWriter()
	writeQueryExtension(w, name)
	d := w.Bytes()
	req := &ProtocolRequest{
		Opcode: QueryExtensionOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return QueryExtensionCookie(seq)
}

func (cookie QueryExtensionCookie) Reply(conn *Conn) (*QueryExtensionReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := NewReaderFromData(replyBuf)
	var reply QueryExtensionReply
	err = readQueryExtensionReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func ListExtensions(conn *Conn) ListExtensionsCookie {
	w := NewWriter()
	writeListExtensions(w)
	d := w.Bytes()
	req := &ProtocolRequest{
		Opcode: ListExtensionsOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return ListExtensionsCookie(seq)
}

func (cookie ListExtensionsCookie) Reply(conn *Conn) (*ListExtensionsReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := NewReaderFromData(replyBuf)
	var reply ListExtensionsReply
	err = readListExtensionsReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func KillClient(conn *Conn, resource uint32) {
	w := NewWriter()
	writeKillClient(w, resource)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  KillClientOpcode,
	}
	conn.SendRequest(0, d, req)
}

func KillClientChecked(conn *Conn, resource uint32) VoidCookie {
	w := NewWriter()
	writeKillClient(w, resource)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  KillClientOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func GetKeyboardMapping(conn *Conn, firstKeycode Keycode, count uint8) GetKeyboardMappingCookie {
	w := NewWriter()
	writeGetKeyboardMapping(w, firstKeycode, count)
	d := w.Bytes()
	req := &ProtocolRequest{
		Opcode: GetKeyboardMappingOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return GetKeyboardMappingCookie(seq)
}

func (cookie GetKeyboardMappingCookie) Reply(conn *Conn) (*GetKeyboardMappingReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := NewReaderFromData(replyBuf)
	var reply GetKeyboardMappingReply
	err = readGetKeyboardMappingReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func SetScreenSaver(conn *Conn, timeout, interval int16, preferBlanking, allowExposures uint8) {
	w := NewWriter()
	writeSetScreenSaver(w, timeout, interval, preferBlanking, allowExposures)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  SetScreenSaverOpcode,
	}
	conn.SendRequest(0, d, req)
}

func SetScreenSaverChecked(conn *Conn, timeout, interval int16, preferBlanking, allowExposures uint8) VoidCookie {
	w := NewWriter()
	writeSetScreenSaver(w, timeout, interval, preferBlanking, allowExposures)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  SetScreenSaverOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func GetScreenSaver(conn *Conn) GetScreenSaverCookie {
	w := NewWriter()
	writeGetScreenSaver(w)
	d := w.Bytes()
	req := &ProtocolRequest{
		Opcode: GetScreenSaverOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return GetScreenSaverCookie(seq)
}

func (cookie GetScreenSaverCookie) Reply(conn *Conn) (*GetScreenSaverReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := NewReaderFromData(replyBuf)
	var reply GetScreenSaverReply
	err = readGetScreenSaverReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func ForceScreenSaver(conn *Conn, mode uint8) {
	w := NewWriter()
	writeForceScreenSaver(w, mode)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  ForceScreenSaverOpcode,
	}
	conn.SendRequest(0, d, req)
}

func ForceScreenSaverChecked(conn *Conn, mode uint8) VoidCookie {
	w := NewWriter()
	writeForceScreenSaver(w, mode)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  ForceScreenSaverOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}

func NoOperation(conn *Conn, n int) {
	w := NewWriter()
	writeNoOperation(w, n)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  NoOperationOpcode,
	}
	conn.SendRequest(0, d, req)
}

func NoOperationChecked(conn *Conn, n int) VoidCookie {
	w := NewWriter()
	writeNoOperation(w, n)
	d := w.Bytes()
	req := &ProtocolRequest{
		NoReply: true,
		Opcode:  NoOperationOpcode,
	}
	seq := conn.SendRequest(RequestChecked, d, req)
	return VoidCookie(seq)
}
