package x

import (
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

var globalConn *Conn
var globalConnMu sync.Mutex

func getConn(t *testing.T) *Conn {
	if os.Getenv("DISPLAY") == "" {
		t.Skip("no set DISPLAY env var")
	}

	globalConnMu.Lock()
	defer globalConnMu.Unlock()

	if globalConn == nil {
		conn, err := NewConn()
		if err != nil {
			t.Fatal(err)
		}

		globalConn = conn
	}

	return globalConn
}

// test CreateWindow + GetGeometry + ChangeWindowAttributes + GetWindowAttributes
// + DestroyWindow
func TestCreateWindow(t *testing.T) {
	c := getConn(t)

	xid, err := c.GenerateID()
	if err != nil {
		t.Fatal(err)
	}

	win := Window(xid)
	t.Log("win:", win)
	root := c.GetDefaultScreen().Root
	err = CreateWindowChecked(c, 0, win, root, 0, 0, 1, 1,
		0, WindowClassInputOutput, 0, 0, nil).Check(c)
	if err != nil {
		t.Fatal(err)
	}

	gem, err := GetGeometry(c, Drawable(win)).Reply(c)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, root, gem.Root)
	assert.Equal(t, int16(0), gem.X)
	assert.Equal(t, int16(0), gem.Y)
	assert.Equal(t, uint16(1), gem.Width)
	assert.Equal(t, uint16(1), gem.Height)
	assert.Equal(t, uint16(0), gem.BorderWidth)

	err = ChangeWindowAttributesChecked(c, win, CWSaveUnder, []uint32{True}).Check(c)
	if err != nil {
		t.Fatal(err)
	}

	attrs, err := GetWindowAttributes(c, win).Reply(c)
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, attrs.SaveUnder)

	err = DestroyWindowChecked(c, win).Check(c)
	if err != nil {
		t.Fatal(err)
	}

}

func TestDestroySubwindows(t *testing.T) {
	c := getConn(t)

	xid, err := c.GenerateID()
	if err != nil {
		t.Fatal(err)
	}

	win := Window(xid)
	t.Log("win:", win)
	root := c.GetDefaultScreen().Root
	err = CreateWindowChecked(c, 0, win, root, 0, 0, 1, 1,
		0, WindowClassInputOutput, 0, 0, nil).Check(c)
	if err != nil {
		t.Fatal(err)
	}

	xid2, err := c.GenerateID()
	if err != nil {
		t.Fatal(err)
	}
	win2 := Window(xid2)
	t.Log("win2:", win2)
	err = CreateWindowChecked(c, 0, win2, win, 0, 0, 1, 1,
		0, WindowClassInputOutput, 0, 0, nil).Check(c)
	if err != nil {
		t.Fatal(err)
	}

	xid3, err := c.GenerateID()
	if err != nil {
		t.Fatal(err)
	}
	win3 := Window(xid3)
	t.Log("win3:", win3)
	err = CreateWindowChecked(c, 0, win3, win, 0, 0, 1, 1,
		0, WindowClassInputOutput, 0, 0, nil).Check(c)
	if err != nil {
		t.Fatal(err)
	}

	err = DestroySubwindowsChecked(c, win).Check(c)
	if err != nil {
		t.Fatal(err)
	}

	_, err = GetWindowAttributes(c, win).Reply(c)
	assert.Nil(t, err)

	_, err = GetWindowAttributes(c, win2).Reply(c)
	assert.NotNil(t, err)

	_, err = GetWindowAttributes(c, win3).Reply(c)
	assert.NotNil(t, err)
}

// test ReparentWindow + QueryTree
func TestReparentWindow(t *testing.T) {
	c := getConn(t)

	xid, err := c.GenerateID()
	if err != nil {
		t.Fatal(err)
	}

	win := Window(xid)
	t.Log("win:", win)
	root := c.GetDefaultScreen().Root
	err = CreateWindowChecked(c, 0, win, root, 0, 0, 1, 1,
		0, WindowClassInputOutput, 0, 0, nil).Check(c)
	if err != nil {
		t.Fatal(err)
	}

	xid2, err := c.GenerateID()
	if err != nil {
		t.Fatal(err)
	}

	win2 := Window(xid2)
	err = CreateWindowChecked(c, 0, win2, root, 0, 0, 1, 1,
		0, WindowClassInputOutput, 0, 0, nil).Check(c)
	if err != nil {
		t.Fatal(err)
	}

	err = ReparentWindowChecked(c, win2, win, 0, 0).Check(c)
	if err != nil {
		t.Fatal(err)
	}

	tree, err := QueryTree(c, win2).Reply(c)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v\n", tree)
	assert.Equal(t, root, tree.Root)
	assert.Equal(t, win, tree.Parent)
	assert.Equal(t, []Window(nil), tree.Children)

	tree, err = QueryTree(c, win).Reply(c)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v\n", tree)
	assert.Equal(t, root, tree.Root)
	assert.Equal(t, root, tree.Parent)
	assert.Equal(t, []Window{win2}, tree.Children)
}

func TestInternAtom(t *testing.T) {
	c := getConn(t)

	ia, err := InternAtom(c, true, "CARDINAL").Reply(c)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, Atom(AtomCardinal), ia.Atom)

	ia, err = InternAtom(c, true, "WM_CLASS").Reply(c)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, Atom(AtomWMClass), ia.Atom)
}

func TestGetAtomName(t *testing.T) {
	c := getConn(t)

	an, err := GetAtomName(c, AtomCardinal).Reply(c)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "CARDINAL", an.Name)

	an, err = GetAtomName(c, AtomPrimary).Reply(c)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "PRIMARY", an.Name)
}

func TestGetProperty(t *testing.T) {
	c := getConn(t)
	root := c.GetDefaultScreen().Root

	ia, err := InternAtom(c, false, "Xorg_Seat").Reply(c)
	if err != nil {
		t.Fatal(err)
	}
	atomXorgSeat := ia.Atom

	p, err := GetProperty(c, false, root, atomXorgSeat, AtomString,
		0, 0xffff).Reply(c)
	if err != nil {
		t.Fatal(err)
	}
	if p.Format != 0 {
		// has property
		assert.Equal(t, uint8(8), p.Format)
		assert.Equal(t, Atom(AtomString), p.Type)
		t.Logf("Xorg_Set: %s", p.Value)
	}

	ia, err = InternAtom(c, false, "_NET_CURRENT_DESKTOP").Reply(c)
	if err != nil {
		t.Fatal(err)
	}
	atomNetCurrentDesktop := ia.Atom
	p, err = GetProperty(c, false, root, atomNetCurrentDesktop, AtomCardinal,
		0, 1).Reply(c)
	if err != nil {
		t.Fatal(err)
	}
	if p.Format != 0 {
		assert.Equal(t, Atom(AtomCardinal), p.Type)
		assert.Equal(t, 4, len(p.Value))
		assert.Equal(t, uint32(1), p.ValueLen)
		t.Log("current desktop:", p.Value)
	}
}

func TestChangeProperty(t *testing.T) {
	c := getConn(t)
	xid, err := c.GenerateID()
	if err != nil {
		t.Fatal(err)
	}

	win := Window(xid)
	t.Log("win:", win)

	root := c.GetDefaultScreen().Root
	err = CreateWindowChecked(c, 0, win, root, 0, 0, 1, 1,
		0, WindowClassInputOutput, 0, 0, nil).Check(c)
	if err != nil {
		t.Fatal(err)
	}

	wmClass := []byte("test\x00test")
	err = ChangePropertyChecked(c, PropModeReplace, win, AtomWMClass, AtomString, 8,
		wmClass).Check(c)
	if err != nil {
		t.Fatal(err)
	}

	prop, err := GetProperty(c, false, win, AtomWMClass, AtomString, 0,
		0xffff).Reply(c)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, Atom(AtomString), prop.Type)
	assert.Equal(t, uint8(8), prop.Format)
	assert.Equal(t, wmClass, prop.Value)

	err = DestroyWindowChecked(c, win).Check(c)
	if err != nil {
		t.Fatal(err)
	}
}

func TestListProperties(t *testing.T) {
	c := getConn(t)
	root := c.GetDefaultScreen().Root
	props, err := ListProperties(c, root).Reply(c)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(props.Atoms)
}

// GrabPointer + UngrabPointer
func TestGrabPointer(t *testing.T) {
	c := getConn(t)
	root := c.GetDefaultScreen().Root
	grab, err := GrabPointer(c, false, root, 0, GrabModeSync, GrabModeSync,
		None, None, CurrentTime).Reply(c)
	if err != nil {
		t.Fatal(err)
	}
	if grab.Status == GrabStatusSuccess {
		err := UngrabPointerChecked(c, CurrentTime).Check(c)
		if err != nil {
			t.Fatal(err)
		}
	}
}

// GrabButton + UngrabButton
func TestGrabButton(t *testing.T) {
	c := getConn(t)
	root := c.GetDefaultScreen().Root
	err := GrabButtonChecked(c, false, root, 0, GrabModeSync, GrabModeSync, None,
		None, ButtonIndex1, 0).Check(c)
	if err != nil {
		t.Fatal(err)
	}

	err = UngrabButtonChecked(c, ButtonIndex1, root, 0).Check(c)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGrabKey(t *testing.T) {
	c := getConn(t)
	root := c.GetDefaultScreen().Root
	key := c.GetSetup().MinKeycode
	err := GrabKeyChecked(c, false, root, 0, key, GrabModeSync, GrabModeSync).Check(c)
	if err != nil {
		t.Fatal(err)
	}

	err = UngrabKeyChecked(c, key, root, 0).Check(c)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGrabKeyboard(t *testing.T) {
	c := getConn(t)
	root := c.GetDefaultScreen().Root
	grab, err := GrabKeyboard(c, false, root, CurrentTime, GrabModeSync,
		GrabModeSync).Reply(c)
	if err != nil {
		t.Fatal(err)
	}
	if grab.Status == GrabStatusSuccess {
		err := UngrabKeyboardChecked(c, CurrentTime).Check(c)
		if err != nil {
			t.Fatal(err)
		}
	}
}

// ListExtension + QueryExtension
func TestListExtensions(t *testing.T) {
	c := getConn(t)
	list, err := ListExtensions(c).Reply(c)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(list.Names)

	for _, extName := range list.Names {
		extInfo, err := QueryExtension(c, extName).Reply(c)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("%s %#v\n", extName, extInfo)
			assert.True(t, extInfo.Present)
			assert.True(t, extInfo.MajorOpcode > 0)
		}

	}
}

// GrabServer + UngrabServer
func TestGrabServer(t *testing.T) {
	c := getConn(t)
	err := GrabServerChecked(c).Check(c)
	if err != nil {
		t.Fatal(err)
	}
	err = UngrabServerChecked(c).Check(c)
	if err != nil {
		t.Fatal(err)
	}
}

func TestQueryPointer(t *testing.T) {
	c := getConn(t)
	root := c.GetDefaultScreen().Root
	p, err := QueryPointer(c, root).Reply(c)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v\n", p)
}

func TestGetMotionEvents(t *testing.T) {
	c := getConn(t)
	root := c.GetDefaultScreen().Root
	events, err := GetMotionEvents(c, root, CurrentTime, CurrentTime).Reply(c)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v\n", events.Events)
}

func TestTranslateCoordinates(t *testing.T) {
	c := getConn(t)
	root := c.GetDefaultScreen().Root
	coord, err := TranslateCoordinates(c, root, root, 10, 20).Reply(c)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v\n", coord)
	assert.Equal(t, int16(10), coord.DstX)
	assert.Equal(t, int16(20), coord.DstY)
	assert.True(t, coord.SameScreen)

	coord, err = TranslateCoordinates(c, root, root, -10, -20).Reply(c)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v\n", coord)
	assert.Equal(t, int16(-10), coord.DstX)
	assert.Equal(t, int16(-20), coord.DstY)
	assert.True(t, coord.SameScreen)
}

func TestWarpPointer(t *testing.T) {
	c := getConn(t)
	root := c.GetDefaultScreen().Root
	err := WarpPointerChecked(c, 0, root, 0, 0,
		0, 0, 10, 20).Check(c)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetInputFocus(t *testing.T) {
	c := getConn(t)
	focus, err := GetInputFocus(c).Reply(c)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v", focus)
}

func TestQueryKeymap(t *testing.T) {
	c := getConn(t)
	keymap, err := QueryKeymap(c).Reply(c)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v\n", keymap)
}

func TestGetKeyboardMapping(t *testing.T) {
	c := getConn(t)
	minKeycode := c.GetSetup().MinKeycode
	maxKeycode := c.GetSetup().MaxKeycode
	// keycodes count
	count := uint8(maxKeycode - minKeycode + 1)
	kbdMapping, err := GetKeyboardMapping(c, minKeycode, count).Reply(c)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v\n", kbdMapping)
	assert.Len(t, kbdMapping.Keysyms, int(count)*int(kbdMapping.KeysymsPerKeycode))
}
