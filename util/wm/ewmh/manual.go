// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package ewmh

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/linuxdeepin/go-x11-client"
)

const LENGTH_MAX = math.MaxUint32

type ClientSource uint32

const (
	// No source at all (for clients supporting an older version of EWMH specification)
	ClientSourceNone ClientSource = iota
	// normal application
	ClientSourceNormal
	// pagers and other clients that represent direct user actions
	ClientSourceOther
)

func getAtom(c *x.Conn, name string) x.Atom {
	atom, _ := c.GetAtom(name)
	return atom
}

func getBooleanFromReply(r *x.GetPropertyReply) (bool, error) {
	num, err := getCardinalFromReply(r)
	return num != 0, err
}

func getUTF8StrFromReply(reply *x.GetPropertyReply) (string, error) {
	if reply.Format != 8 {
		return "", errors.New("bad reply")
	}

	return string(reply.Value), nil
}

func getUTF8StrsFromReply(reply *x.GetPropertyReply) ([]string, error) {
	if reply.Format != 8 {
		return nil, errors.New("bad reply")
	}

	data := reply.Value
	var strs []string
	sstart := 0
	for i, c := range data {
		if c == 0 {
			strs = append(strs, string(data[sstart:i]))
			sstart = i + 1
		}
	}
	if sstart < len(data) {
		strs = append(strs, string(data[sstart:]))
	}
	return strs, nil
}

func getWMIconFromReply(reply *x.GetPropertyReply) ([]WMIcon, error) {
	if reply.Format != 32 {
		return nil, errors.New("bad reply")
	}
	return getIcons(reply.Value)
}

type WMIcon struct {
	Width, Height uint32
	Data          []byte
}

func (icon *WMIcon) String() string {
	if icon == nil {
		return "nil"
	}

	return fmt.Sprintf("ewmh.WMIcon{ size: %dx%d, dataLength: %d }", icon.Width, icon.Height, len(icon.Data))
}

func getIcons(p []byte) ([]WMIcon, error) {
	var icons []WMIcon
	for len(p) >= 2*4 {
		width := x.Get32(p)
		height := x.Get32(p[4:])
		area := width * height

		if len(p) >= int(2+area)*4 {
			if area > 0 {
				icon := WMIcon{
					Width:  width,
					Height: height,
					Data:   p[2*4 : (2+area)*4],
				}
				icons = append(icons, icon)
			}

			if len(p) > int(2+area)*4 {
				// next icon
				p = p[(2+area)*4:]
				continue
			} else {
				// end
				break
			}
		} else {
			break
		}
	}
	return icons, nil
}

func sendClientMessage(c *x.Conn, win, dest x.Window, msgType x.Atom, pArray *[5]uint32) x.VoidCookie {
	var data x.ClientMessageData
	data.SetData32(pArray)
	event := x.ClientMessageEvent{
		Format: 32,
		Window: win,
		Type:   msgType,
		Data:   data,
	}
	w := x.NewWriter()
	x.WriteClientMessageEvent(w, &event)
	const evMask = x.EventMaskSubstructureNotify | x.EventMaskSubstructureRedirect
	return x.SendEventChecked(c, false, dest, evMask, w.Bytes())
}

/**
 *    _NET_DESKTOP_GEOMETRY
 */

func RequestChangeDesktopGeometry(c *x.Conn, geo DesktopGeometry) x.VoidCookie {
	array := [5]uint32{
		geo.Width,
		geo.Height,
	}
	root := c.GetDefaultScreen().Root
	atomNetDesktopGeometry, _ := c.GetAtom("_NET_DESKTOP_GEOMETRY")
	return sendClientMessage(c, root, root, atomNetDesktopGeometry, &array)
}

/**
 *    _NET_DESKTOP_VIEWPORT
 */

func RequestChangeDesktopViewport(c *x.Conn, corner ViewportLeftTopCorner) x.VoidCookie {
	array := [5]uint32{
		corner.X,
		corner.Y,
	}
	root := c.GetDefaultScreen().Root
	atomNetDesktopViewport, _ := c.GetAtom("_NET_DESKTOP_VIEWPORT")
	return sendClientMessage(c, root, root, atomNetDesktopViewport, &array)
}

/**
 *    _NET_CURRENT_DESKTOP
 */

func RequestChangeCurrentDesktop(c *x.Conn, desktop uint32, timestamp x.Timestamp) x.VoidCookie {
	array := [5]uint32{
		desktop,
		uint32(timestamp),
	}
	root := c.GetDefaultScreen().Root
	atomNetCurrentDesktop, _ := c.GetAtom("_NET_CURRENT_DESKTOP")
	return sendClientMessage(c, root, root, atomNetCurrentDesktop, &array)
}

/**
 *    _NET_ACTIVE_WINDOW
 */

func RequestChangeActiveWindow(c *x.Conn, windowToActivate x.Window, source ClientSource, timestamp x.Timestamp,
	currentActiveWindow x.Window) x.VoidCookie {
	array := [5]uint32{
		uint32(source),
		uint32(timestamp),
		uint32(currentActiveWindow),
	}
	root := c.GetDefaultScreen().Root
	atomNetActivewindow, _ := c.GetAtom("_NET_ACTIVE_WINDOW")
	return sendClientMessage(c, windowToActivate, root, atomNetActivewindow, &array)
}

/**
 *    _NET_SHOWING_DESKTOP
 */

func SetShowingDesktopChecked(c *x.Conn, show bool) x.VoidCookie {
	val := uint32(0)
	if show {
		val = 1
	}
	w := x.NewWriter()
	w.Write4b(uint32(val))
	root := c.GetDefaultScreen().Root
	atomNetShowingDesktop, _ := c.GetAtom("_NET_SHOWING_DESKTOP")
	return x.ChangePropertyChecked(c, x.PropModeReplace, root, atomNetShowingDesktop,
		x.AtomCardinal, 32, w.Bytes())
}

func SetShowingDesktop(c *x.Conn, show bool) {
	val := uint32(0)
	if show {
		val = 1
	}
	w := x.NewWriter()
	w.Write4b(uint32(val))
	root := c.GetDefaultScreen().Root
	atomNetShowingDesktop, _ := c.GetAtom("_NET_SHOWING_DESKTOP")
	x.ChangeProperty(c, x.PropModeReplace, root, atomNetShowingDesktop, x.AtomCardinal,
		32, w.Bytes())
}

func RequestChangeShowingDesktop(c *x.Conn, show bool) x.VoidCookie {
	val := uint32(0)
	if show {
		val = 1
	}
	array := [5]uint32{
		val,
	}
	root := c.GetDefaultScreen().Root
	atomNetShowingDesktop, _ := c.GetAtom("_NET_SHOWING_DESKTOP")
	return sendClientMessage(c, root, root, atomNetShowingDesktop, &array)
}

/**
 *    _NET_CLOSE_WINDOW
 */

func RequestCloseWindow(c *x.Conn, window x.Window, timestamp x.Timestamp, source ClientSource) x.VoidCookie {
	array := [5]uint32{
		uint32(timestamp),
		uint32(source),
	}
	root := c.GetDefaultScreen().Root
	atomNetCloseWindow, _ := c.GetAtom("_NET_CLOSE_WINDOW")
	return sendClientMessage(c, window, root, atomNetCloseWindow, &array)
}

/**
 *    _NET_MOVERESIZE_WINDOW
 */

type MoveResizeWindowFlags uint32

const (
	MoveResizeWindowX MoveResizeWindowFlags = 1 << (8 + iota)
	MoveResizeWindowY
	MoveResizeWindowWidth
	MoveResizeWindowHeight
)

func RequestMoveResizeWindow(c *x.Conn, window x.Window, gravity uint32, source ClientSource, flags MoveResizeWindowFlags,
	x, y, width, height uint32) x.VoidCookie {
	array := [5]uint32{
		gravity | uint32(flags) | (uint32(source) << 12),
		x, y, width, height,
	}
	root := c.GetDefaultScreen().Root
	atomNetMoveResizeWindow, _ := c.GetAtom("_NET_MOVERESIZE_WINDOW")
	return sendClientMessage(c, window, root, atomNetMoveResizeWindow, &array)
}

/**
 *    _NET_WM_MOVERESIZE
 */

type MoveResizeDirection uint32

const (
	MoveResizeSizeTopLeft MoveResizeDirection = iota
	MoveResizeSizeTop
	MoveResizeSizeTopRight
	MoveResizeSizeRight
	MoveResizeSizeBottomRight
	MoveResizeSizeBottom // 5
	MoveResizeSizeBottomLeft
	MoveResizeSizeLeft
	MoveResizeMove
	MoveResizeSizeKeyboard
	MoveResizeMoveKeyboard // 10
	MoveResizeCancel
)

func RequestWMMoveResize(c *x.Conn, window x.Window, xRoot, yRoot uint32, direction MoveResizeDirection,
	button uint32, source ClientSource) x.VoidCookie {
	array := [5]uint32{
		xRoot, yRoot,
		uint32(direction),
		button,
		uint32(source),
	}
	root := c.GetDefaultScreen().Root
	atomNetWMMoveResize, _ := c.GetAtom("_NET_WM_MOVERESIZE")
	return sendClientMessage(c, window, root, atomNetWMMoveResize, &array)
}

/**
 *    _NET_RESTACK_WINDOW
 */

func RequestRestackWindow(c *x.Conn, window, siblingWindow x.Window, stackMode uint32) x.VoidCookie {
	array := [5]uint32{
		2,
		uint32(siblingWindow),
		stackMode,
	}
	root := c.GetDefaultScreen().Root
	atomNetRestackWindow, _ := c.GetAtom("_NET_RESTACK_WINDOW")
	return sendClientMessage(c, window, root, atomNetRestackWindow, &array)
}

/**
 *    _NET_WM_DESKTOP
 */
func RequestChangeWMDesktop(c *x.Conn, window x.Window, desktop uint32, source ClientSource) x.VoidCookie {
	array := [5]uint32{
		desktop,
		uint32(source),
	}
	root := c.GetDefaultScreen().Root
	atomNetWMDesktop, _ := c.GetAtom("_NET_WM_DESKTOP")
	return sendClientMessage(c, window, root, atomNetWMDesktop, &array)
}

/**
 *    _NET_WM_STATE
 */

type WMStateAction uint32

const (
	WMStateRemove WMStateAction = iota
	WMStateAdd
	WMStateToggle
)

func RequestChangeWMState(c *x.Conn, window x.Window, action WMStateAction, firstProperty, secondProperty x.Atom,
	source ClientSource) x.VoidCookie {
	array := [5]uint32{
		uint32(action),
		uint32(firstProperty),
		uint32(secondProperty),
		uint32(source),
	}
	root := c.GetDefaultScreen().Root
	atomNetWMState, _ := c.GetAtom("_NET_WM_STATE")
	return sendClientMessage(c, window, root, atomNetWMState, &array)
}

/*
 *    _NET_WM_PING
 */

func SendWMPing(c *x.Conn, window x.Window, timestamp x.Timestamp) x.VoidCookie {
	atomNetWMPing, _ := c.GetAtom("_NET_WM_PING")
	array := [5]uint32{
		uint32(atomNetWMPing),
		uint32(timestamp),
		uint32(window),
	}
	atomWMProtocols, _ := c.GetAtom("WM_PROTOCOLS")
	return sendClientMessage(c, window, window, atomWMProtocols, &array)
}

/*
 *    _NET_WM_SYNC_REQUEST
 */

func (counter WMSyncRequestCounter) ToUint64() uint64 {
	return uint64(counter.Low) | (uint64(counter.High) << 32)
}

func (counter WMSyncRequestCounter) String() string {
	return strconv.FormatUint(counter.ToUint64(), 10)
}

func NewWMSyncRequestCounter(val uint64) WMSyncRequestCounter {
	return WMSyncRequestCounter{
		Low:  uint32(val),
		High: uint32(val >> 32),
	}
}

func SendWMSyncRequest(c *x.Conn, window x.Window, timestamp x.Timestamp, counter WMSyncRequestCounter) x.VoidCookie {
	atomNetWMSyncRequest, _ := c.GetAtom("_NET_WM_SYNC_REQUEST")
	array := [5]uint32{
		uint32(atomNetWMSyncRequest),
		uint32(timestamp),
		counter.Low,
		counter.High,
	}
	atomWMProtocols, _ := c.GetAtom("WM_PROTOCOLS")
	return sendClientMessage(c, window, window, atomWMProtocols, &array)
}

/*
 *    _NET_WM_FULLSCREEN_MONITORS
 */
func RequestChangeWMFullscreenMonitors(c *x.Conn, window x.Window, edges WMFullscreenMonitors,
	source ClientSource) x.VoidCookie {
	array := [5]uint32{
		edges.Top,
		edges.Bottom,
		edges.Left,
		edges.Right,
		uint32(source),
	}
	root := c.GetDefaultScreen().Root
	atomNetWMFullscreenMonitors, _ := c.GetAtom("_NET_WM_FULLSCREEN_MONITORS")
	return sendClientMessage(c, window, root, atomNetWMFullscreenMonitors, &array)
}
