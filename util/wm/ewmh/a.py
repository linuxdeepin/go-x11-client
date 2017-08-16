#!/usr/bin/env python3

_lines = []
CHECKED = ['Checked', '']
UNCHECKED = ['', 'Unchecked']


def l(fmt, *arg):
    _lines.append(fmt % arg)


def DO_GET_PROPERTY(sth, _property, atype, length, cookie_type):
    return_type = 'Get' + cookie_type + 'Cookie'
    for unchecked in UNCHECKED:
        l('\nfunc (c *Conn) Get%s%s(window x.Window) %s{', sth, unchecked,
          return_type)
        l('    cookie := x.GetProperty%s(c.conn, x.False, window,\nc.GetAtom("%s"), %s, 0, %s)',
          unchecked, _property, atype, length)
        l('    return %s(cookie)', return_type)
        l('}')


def DO_GET_ROOT_PROPERTY(sth, _property, atype, length, cookie_type):
    return_type = 'Get' + cookie_type + 'Cookie'
    for unchecked in UNCHECKED:
        l('\nfunc (c *Conn) Get%s%s() %s{', sth, unchecked, return_type)
        l('    cookie := x.GetProperty%s(c.conn, x.False, c.GetRootWin(),\nc.GetAtom("%s"),%s, 0, %s)',
          unchecked, _property, atype, length)
        l('    return %s(cookie)', return_type)
        l('}')


def DO_REPLY_SINGLE_VALUE(sth, atype, go_type, cookie_type):
    # getXXXFromReply
    l('\nfunc get%sFromReply(r *x.GetPropertyReply) (%s, error) {', sth,
      go_type)
    l('    if r.Format != 32 || len(r.Value) != 4 {')
    l('        return 0, errors.New("bad reply")')
    l('    }')  # end if
    l('    return %s(x.Get32(r.Value)), nil', go_type)
    l('}')  # end func

    # define type GetYYYCookie
    DO_COOKIE_REPLY(cookie_type, atype, go_type, 0)


def DO_SINGLE_VALUE(sth, _property, atype, go_type, cookie_type):
    DO_GET_PROPERTY(sth, _property, atype, '1', cookie_type)

    for checked in CHECKED:
        l('\nfunc (c *Conn) Set%s%s(window x.Window, val %s) x.VoidCookie {',
          sth, checked, go_type)
        l('    w := x.NewWriter()')
        l('    w.Write4b(uint32(val))')
        l('    return x.ChangeProperty%s(c.conn, x.PropModeReplace, window,\nc.GetAtom("%s"), %s, 32, 1, w.Bytes())',
          checked, _property, atype)
        l('}')


def DO_ROOT_SINGLE_VALUE(sth, _property, atype, go_type, cookie_type):
    DO_GET_ROOT_PROPERTY(sth, _property, atype, '1', cookie_type)

    for checked in CHECKED:
        l('\nfunc (c *Conn) Set%s%s(val %s) x.VoidCookie {', sth, checked,
          go_type)
        l('    w := x.NewWriter()')
        l('    w.Write4b(uint32(val))')
        l('    return x.ChangeProperty%s(c.conn, x.PropModeReplace, c.GetRootWin(),\nc.GetAtom("%s"), %s, 32, 1, w.Bytes())',
          checked, _property, atype)
        l('}')


def DO_REPLY_LIST_VALUES(sth, atype, go_type):
    get_sth_from_reply = 'get' + sth + 'FromReply'
    l('\nfunc %s(r *x.GetPropertyReply) ([]%s, error) {', get_sth_from_reply,
      go_type)
    l('    if r.Format != 32 {')
    l('        return nil, errors.New("bad reply")')
    l('    }')  # end if

    l('    count := len(r.Value) / 4')
    l('    ret := make([]%s, count)', go_type)
    l('    rdr := x.NewReaderFromData(r.Value)')
    l('    for i := 0; i < count; i++ {')
    l('        ret[i] = %s(rdr.Read4b())', go_type)
    l('    }')  # end for
    l('    return ret, nil')
    l('}')  # end func

    DO_COOKIE_REPLY(sth, atype, '[]' + go_type, 'nil')


def DO_ROOT_LIST_VALUES(sth, _property, atype, go_type, cookie_type):
    DO_GET_ROOT_PROPERTY(sth, _property, atype, 'LENGTH_MAX', cookie_type)

    for checked in CHECKED:
        l('\nfunc (c *Conn) Set%s%s(vals []%s) x.VoidCookie {', sth, checked,
          go_type)
        l('    w := x.NewWriter()')
        l('    for _, val := range vals {')
        l('        w.Write4b(uint32(val))')
        l('    }')  # end for

        l('    return x.ChangeProperty%s(c.conn, x.PropModeReplace, c.GetRootWin(),\n'
          +
          'c.GetAtom("_NET_SUPPORTED"), %s, 32, uint32(len(vals)), w.Bytes())',
          checked, atype)
        l('}')  # end func


def DO_LIST_VALUES(sth, _property, atype, go_type, cookie_type):
    DO_GET_PROPERTY(sth, _property, atype, 'LENGTH_MAX', cookie_type)

    for checked in CHECKED:
        l('\nfunc (c *Conn) Set%s%s(window x.Window,vals []%s) x.VoidCookie {',
          sth, checked, go_type)
        l('    w := x.NewWriter()')
        l('    for _, val := range vals {')
        l('        w.Write4b(uint32(val))')
        l('    }')  # end for

        l('    return x.ChangeProperty%s(c.conn, x.PropModeReplace, window,\n'
          +
          'c.GetAtom("_NET_SUPPORTED"), %s, 32, uint32(len(vals)), w.Bytes())',
          checked, atype)
        l('}')  # end func


def DO_REPLY_STRUCTURE(sth, fields):
    go_type = sth
    bad_go_value = go_type + '{}'
    DO_COOKIE_REPLY(sth, 'x.AtomCARDINAL', go_type, bad_go_value)

    # fields is str list
    # all field type is uint32
    # property type is CARDINAL
    l('type %s struct {', go_type)
    l('    %s uint32', ', '.join(fields))
    l('}')

    l('\nfunc get%sFromReply(reply *x.GetPropertyReply) (%s, error) {', sth,
      go_type)
    l('    list, err := getCardinalsFromReply(reply)')
    l('    if err != nil {')
    l('        return %s, err', bad_go_value)
    l('    }')  # end if

    l('\n    if len(list) != %d {', len(fields))
    l('        return %s, errors.New("length of list is not %d")',
      bad_go_value, len(fields))
    l('    }')  # end if

    # return normal value
    l('    return %s{', go_type)
    for idx, field in enumerate(fields):
        l('        %s: list[%d],', field, idx)
    l('    }, nil')  # end return

    l('}')  # end func


def DO_REPLY_STRUCTURES(sth, go_type, fields):
    n_fields = len(fields)
    DO_COOKIE_REPLY(sth, 'x.AtomCARDINAL', '[]' + go_type, 'nil')

    l('type %s struct {', go_type)
    l('    %s uint32', ', '.join(fields))
    l('}')

    l('\nfunc get%sFromReply(reply *x.GetPropertyReply) ([]%s, error) {', sth,
      go_type)
    l('    list, err := getCardinalsFromReply(reply)')
    l('    if err != nil {')
    l('        return nil, err')
    l('    }')  # end if

    l('    if len(list) %% %d != 0 {', n_fields)
    l('        return nil, errors.New("length of list is not a multiple of %d")',
      n_fields)
    l('    }')  # end if

    l('    ret := make([]%s, len(list)/%d)', go_type, n_fields)
    l('    idx := 0')

    l('    for i := 0; i < len(list); i+=%d {', n_fields)
    l('        ret[idx] = %s{', go_type)
    for field_idx, field in enumerate(fields):
        if field_idx == 0:
            l('            %s: list[i],', field)
        else:
            l('            %s: list[i+%d],', field, field_idx)
    l('        }')  # end struct new
    l('        idx++')
    l('    }')  # end for

    l('    return ret, nil')
    l('}')  # end func


def DO_UTF8_STRING(sth, _property):
    DO_GET_PROPERTY(sth, _property, 'c.GetAtom("UTF8_STRING")', 'LENGTH_MAX',
                    'UTF8Str')

    for checked in CHECKED:
        l('\nfunc (c *Conn) Set%s%s(window x.Window, val string) x.VoidCookie {',
          sth, checked)
        l('    return x.ChangeProperty%s(c.conn, x.PropModeReplace, window,\nc.GetAtom("%s"),'
          + 'c.GetAtom("UTF8_STRING"), 8, uint32(len(val)), []byte(val))',
          checked, _property)
        l('}')  # end func


def DO_ROOT_UTF8_STRING(sth, _property):
    DO_GET_ROOT_PROPERTY(sth, _property, 'c.GetAtom("UTF8_STRING")',
                         'LENGTH_MAX', 'UTF8Str')

    for checked in CHECKED:
        l('\nfunc (c *Conn) Set%s%s(val string) x.VoidCookie {', sth, checked)
        l('    return x.ChangeProperty%s(c.conn, x.PropModeReplace, c.GetRootWin(),\nc.GetAtom("%s"),'
          + 'c.GetAtom("UTF8_STRING"), 8, uint32(len(val)), []byte(val))',
          checked, _property)
        l('}')  # end func


def DO_ROOT_UTF8_STRINGS(sth, _property):
    DO_GET_ROOT_PROPERTY(sth, _property, 'c.GetAtom("UTF8_STRING")',
                         'LENGTH_MAX', 'UTF8Strs')

    for checked in CHECKED:
        l('\nfunc (c *Conn) Set%s%s(vals []string) x.VoidCookie {', sth,
          checked)
        l('    w := x.NewWriter()')
        l('    length := 0')
        l('    for _, val := range vals {')
        l('        w.WriteString(val)')
        l('        w.Write1b(0)')
        l('        length += len(val) + 1')
        l('    }')  # end for
        l('    return x.ChangeProperty%s(c.conn, x.PropModeReplace, c.GetRootWin(),\nc.GetAtom("%s"),'
          + 'c.GetAtom("UTF8_STRING"), 8, uint32(length), w.Bytes())', checked,
          _property)
        l('}')  # end func


def DO_SET_ROOT_STRUCTURES(sth, _property, go_type, fields):
    n_fields = len(fields)
    for checked in CHECKED:
        l('\nfunc (c *Conn) Set%s%s(vals []%s) x.VoidCookie {', sth, checked,
          go_type)
        l('    w := x.NewWriter()')
        l('    for _, val := range vals {')
        for field in fields:
            l('        w.Write4b(val.%s)', field)
        l('    }')  # end for
        l('    length := uint32(len(vals) * %d)', n_fields)
        l('    return x.ChangeProperty%s(c.conn, x.PropModeReplace, c.GetRootWin(),\n'
          + 'c.GetAtom("%s"), x.AtomCARDINAL, 32, length, w.Bytes())', checked,
          _property)
        l('}')  # end func


# TODO
# def DO_SET_STRUCTURES


def DO_SET_ROOT_STRUCTURE(sth, _property, go_type, fields):
    for checked in CHECKED:
        l('\nfunc (c *Conn) Set%s%s(val %s) x.VoidCookie {', sth, checked,
          go_type)
        l('    w := x.NewWriter()')
        for field in fields:
            l('    w.Write4b(val.%s)', field)

        l('    return x.ChangeProperty%s(c.conn, x.PropModeReplace, c.GetRootWin(),\nc.GetAtom("%s"),'
          + 'x.AtomCARDINAL, 32, %d, w.Bytes())', checked, _property,
          len(fields))
        l('}')  # end func


def DO_SET_STRUCTURE(sth, _property, go_type, fields):
    for checked in CHECKED:
        l('\nfunc (c *Conn) Set%s%s(window x.Window, val %s) x.VoidCookie {',
          sth, checked, go_type)
        l('    w := x.NewWriter()')
        for field in fields:
            l('    w.Write4b(val.%s)', field)

        l('    return x.ChangeProperty%s(c.conn, x.PropModeReplace, window,\nc.GetAtom("%s"),'
          + 'x.AtomCARDINAL, 32, %d, w.Bytes())', checked, _property,
          len(fields))
        l('}')  # end func


def DO_COOKIE_REPLY(sth, atype, go_type, bad_go_value):
    cookie_type = 'Get' + sth + 'Cookie'
    get_sth_from_reply = 'get' + sth + 'FromReply'
    l('type %s x.GetPropertyCookie', cookie_type)

    l('\nfunc (cookie %s) Reply(c *Conn) (%s, error) {', cookie_type, go_type)
    l('    reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)')
    l('    if err != nil {')
    l('        return %s, err', bad_go_value)
    l('    }')  # end if

    # check reply type
    l('    if reply.Type != %s {', atype)
    l('        return %s, errors.New("bad reply")', bad_go_value)
    l('    }')  # end if

    l('    return %s(reply)', get_sth_from_reply)  # return good value
    l('}')  # end func


def begin():
    l('package ewmh\n')
    l('import x "github.com/linuxdeepin/go-x11-client"')
    l('import "errors"')


def end():
    for l in _lines:
        print(l)


def do_header(name):
    l('/**')
    l(' *    %s', name)
    l('*/')


def main():
    begin()

    DO_COOKIE_REPLY('UTF8Str', 'c.GetAtom("UTF8_STRING")', 'string', '""')
    DO_COOKIE_REPLY('UTF8Strs', 'c.GetAtom("UTF8_STRING")', '[]string', 'nil')
    DO_COOKIE_REPLY('Boolean', 'x.AtomCARDINAL', 'bool', 'false')

    DO_REPLY_SINGLE_VALUE('Window', 'x.AtomWINDOW', 'x.Window', 'Window')
    DO_REPLY_SINGLE_VALUE('Cardinal', 'x.AtomCARDINAL', 'uint32', 'Cardinal')

    DO_REPLY_LIST_VALUES('Windows', 'x.AtomWINDOW', 'x.Window')
    DO_REPLY_LIST_VALUES('Atoms', 'x.AtomATOM', 'x.Atom')
    DO_REPLY_LIST_VALUES('Cardinals', 'x.AtomCARDINAL', 'uint32')

    #> _NET_SUPPORTED
    do_header('_NET_SUPPORTED')
    DO_ROOT_LIST_VALUES('Supported', '_NET_SUPPORTED', 'x.AtomATOM', 'x.Atom',
                        'Atoms')

    #> _NET_CLIENT_LIST
    do_header('_NET_CLIENT_LIST')
    DO_ROOT_LIST_VALUES('ClientList', '_NET_CLIENT_LIST', 'x.AtomWINDOW',
                        'x.Window', 'Windows')

    #> _NET_CLIENT_LIST_STACKING
    do_header('_NET_CLIENT_LIST_STACKING')
    DO_ROOT_LIST_VALUES('ClientListStacking', '_NET_CLIENT_LIST_STACKING',
                        'x.AtomWINDOW', 'x.Window', 'Windows')

    #> _NET_NUMBER_OF_DESKTOPS
    do_header('_NET_NUMBER_OF_DESKTOPS')
    DO_ROOT_SINGLE_VALUE('NumberOfDesktop', '_NET_NUMBER_OF_DESKTOPS',
                         'x.AtomCARDINAL', 'uint32', 'Cardinal')

    #> _NET_DESKTOP_GEOMETRY
    do_header('_NET_DESKTOP_GEOMETRY')
    DO_GET_ROOT_PROPERTY('DesktopGeometry', '_NET_DESKTOP_GEOMETRY',
                         'x.AtomCARDINAL', '2', 'DesktopGeometry')
    desktop_geometry_fields = ['Width', 'Height']
    DO_REPLY_STRUCTURE('DesktopGeometry', desktop_geometry_fields)
    DO_SET_ROOT_STRUCTURE('DesktopGeometry', '_NET_DESKTOP_GEOMETRY',
                          'DesktopGeometry', desktop_geometry_fields)

    #> _NET_DESKTOP_VIEWPORT
    do_header('_NET_DESKTOP_VIEWPORT')
    DO_GET_ROOT_PROPERTY('DesktopViewport', '_NET_DESKTOP_VIEWPORT',
                         'x.AtomCARDINAL', 'LENGTH_MAX', 'DesktopViewport')
    viewport_fields = ['X', 'Y']
    DO_REPLY_STRUCTURES('DesktopViewport', 'ViewportLeftTopCorner',
                        viewport_fields)
    DO_SET_ROOT_STRUCTURES('DesktopViewport', '_NET_DESKTOP_VIEWPORT',
                           'ViewportLeftTopCorner', viewport_fields)

    #> _NET_CURRENT_DESKTOP
    do_header('_NET_CURRENT_DESKTOP')
    DO_ROOT_SINGLE_VALUE('CurrentDesktop', '_NET_CURRENT_DESKTOP',
                         'x.AtomCARDINAL', 'uint32', 'Cardinal')

    #> _NET_DESKTOP_NAMES
    do_header('_NET_DESKTOP_NAMES')
    DO_ROOT_UTF8_STRINGS('DesktopNames', '_NET_DESKTOP_NAMES')

    #> _NET_ACTIVE_WINDOW
    do_header('_NET_ACTIVE_WINDOW')
    DO_ROOT_SINGLE_VALUE('ActiveWindow', '_NET_ACTIVE_WINDOW', 'x.AtomWINDOW',
                         'x.Window', 'Window')

    #> _NET_WORKAREA
    do_header('_NET_WORKAREA')
    DO_GET_ROOT_PROPERTY('Workarea', '_NET_WORKAREA', 'x.AtomCARDINAL',
                         'LENGTH_MAX', 'Workarea')
    workarea_geo_fields = ['X', 'Y', 'Width', 'Height']
    DO_REPLY_STRUCTURES('Workarea', 'WorkareaGeometry', workarea_geo_fields)
    DO_SET_ROOT_STRUCTURES('Workarea', '_NET_WORKAREA', 'WorkareaGeometry',
                           workarea_geo_fields)

    #> _NET_SUPPORTING_WM_CHECK
    do_header('_NET_SUPPORTING_WM_CHECK')
    DO_SINGLE_VALUE('SupportingWmCheck', '_NET_SUPPORTING_WM_CHECK',
                    'x.AtomWINDOW', 'x.Window', 'Window')

    #>  _NET_VIRTUAL_ROOTS
    do_header('_NET_VIRTUAL_ROOTS')
    DO_ROOT_LIST_VALUES('VirtualRoots', '_NET_VIRTUAL_ROOTS', 'x.AtomWINDOW',
                        'x.Window', 'Windows')

    #> _NET_DESKTOP_LAYOUT
    do_header('_NET_DESKTOP_LAYOUT')
    DO_GET_ROOT_PROPERTY('DesktopLayout', '_NET_DESKTOP_LAYOUT',
                         'x.AtomCARDINAL', '4', 'DesktopLayout')
    desktop_layout_fields = [
        'Orientation', 'Columns', 'Rows', 'StartingCorner'
    ]
    DO_REPLY_STRUCTURE('DesktopLayout', desktop_layout_fields)
    DO_SET_ROOT_STRUCTURE('DesktopLayout', '_NET_DESKTOP_LAYOUT',
                          'DesktopLayout', desktop_layout_fields)

    #> _NET_SHOWING_DESKTOP
    do_header('_NET_SHOWING_DESKTOP')
    DO_GET_ROOT_PROPERTY('ShowingDesktop', '_NET_SHOWING_DESKTOP',
                         'x.AtomCARDINAL', '1', 'Boolean')

    #> _NET_WM_NAME
    do_header('_NET_WM_NAME')
    DO_UTF8_STRING('WmName', '_NET_WM_NAME')

    #> _NET_WM_VISIBLE_NAME
    do_header('_NET_WM_VISIBLE_NAME')
    DO_UTF8_STRING('WmVisibleName', '_NET_WM_VISIBLE_NAME')

    #> _NET_WM_ICON_NAME
    do_header('_NET_WM_ICON_NAME')
    DO_UTF8_STRING('WmIconName', '_NET_WM_ICON_NAME')

    #> _NET_WM_VISIBLE_ICON_NAME
    do_header('_NET_WM_VISIBLE_ICON_NAME')
    DO_UTF8_STRING('WmVisibleIconName', '_NET_WM_VISIBLE_ICON_NAME')

    #> _NET_WM_DESKTOP
    do_header('_NET_WM_DESKTOP')
    DO_SINGLE_VALUE('WmDesktop', '_NET_WM_DESKTOP', 'x.AtomCARDINAL', 'uint32',
                    'Cardinal')

    #> _NET_WM_WINDOW_TYPE
    do_header('_NET_WM_WINDOW_TYPE')
    DO_LIST_VALUES('WmWindowType', '_NET_WM_WINDOW_TYPE', 'x.AtomATOM',
                   'x.Atom', 'Atoms')

    #> _NET_WM_STATE
    do_header('_NET_WM_STATE')
    DO_LIST_VALUES('WmState', '_NET_WM_STATE', 'x.AtomATOM', 'x.Atom', 'Atoms')

    #> _NET_WM_ALLOWED_ACTIONS
    do_header('_NET_WM_ALLOWED_ACTIONS')
    DO_LIST_VALUES('WmAllowedActions', '_NET_WM_ALLOWED_ACTIONS', 'x.AtomATOM',
                   'x.Atom', 'Atoms')

    #> _NET_WM_STRUT
    do_header('_NET_WM_STRUT')
    DO_GET_PROPERTY('WmStrut', '_NET_WM_STRUT', 'x.AtomCARDINAL', '4',
                    'WmStrut')
    strut_fields = ['Left', 'Right', 'Top', 'Bottom']
    DO_REPLY_STRUCTURE('WmStrut', strut_fields)
    DO_SET_STRUCTURE('WmStrut', '_NET_WM_STRUT', 'WmStrut', strut_fields)

    #> _NET_WM_STRUT_PARTIAL
    do_header('_NET_WM_STRUT_PARTIAL')
    DO_GET_PROPERTY('WmStrutPartial', '_NET_WM_STRUT_PARTIAL',
                    'x.AtomCARDINAL', '12', 'WmStrutPartial')
    strut_partial_fields = strut_fields + [
        'LeftStartY', 'LeftEndY', 'RightStartY', 'RightEndY', 'TopStartX',
        'TopEndX', 'BottomStartX', 'BottomEndX'
    ]
    DO_REPLY_STRUCTURE('WmStrutPartial', strut_partial_fields)
    DO_SET_STRUCTURE('WmStrutPartial', '_NET_WM_STRUT_PARTIAL',
                     'WmStrutPartial', strut_partial_fields)

    #> _NET_WM_ICON_GEOMETRY
    do_header('_NET_WM_ICON_GEOMETRY')
    DO_GET_PROPERTY('WmIconGeometry', '_NET_WM_ICON_GEOMETRY',
                    'x.AtomCARDINAL', '4', 'WmIconGeometry')
    icon_geometry = ['X', 'Y', 'Width', 'Height']
    DO_REPLY_STRUCTURE('WmIconGeometry', icon_geometry)
    DO_SET_STRUCTURE('WmIconGeometry', '_NET_WM_ICON_GEOMETRY',
                     'WmIconGeometry', icon_geometry)

    #> _NET_WM_ICON
    do_header('_NET_WM_ICON')
    DO_GET_PROPERTY('WmIcon', '_NET_WM_ICON', 'x.AtomCARDINAL', 'LENGTH_MAX',
                    'WmIcon')
    DO_COOKIE_REPLY('WmIcon', 'x.AtomCARDINAL', '[]WmIcon', 'nil')

    #> _NET_WM_PID
    do_header('_NET_WM_PID')
    DO_SINGLE_VALUE('WmPid', '_NET_WM_PID', 'x.AtomCARDINAL', 'uint32',
                    'Cardinal')

    #> _NET_WM_HANDLED_ICONS
    do_header('_NET_WM_HANDLED_ICONS')
    DO_SINGLE_VALUE('WmHandledIcons', '_NET_WM_HANDLED_ICONS',
                    'x.AtomCARDINAL', 'uint32', 'Cardinal')

    #> _NET_WM_USER_TIME
    do_header('_NET_WM_USER_TIME')
    DO_SINGLE_VALUE('WmUserTime', '_NET_WM_USER_TIME', 'x.AtomCARDINAL',
                    'uint32', 'Cardinal')

    #> _NET_WM_USER_TIME_WINDOW
    do_header('_NET_WM_USER_TIME_WINDOW')
    DO_SINGLE_VALUE('WmUserTimeWindow', '_NET_WM_USER_TIME_WINDOW',
                    'x.AtomWINDOW', 'x.Window', 'Window')

    #> _NET_FRAME_EXTENTS
    do_header('_NET_FRAME_EXTENTS')
    DO_GET_PROPERTY('FrameExtents', '_NET_FRAME_EXTENTS', 'x.AtomCARDINAL',
                    '4', 'FrameExtents')
    frame_extents_fields = ['Left', 'Right', 'Top', 'Bottom']
    DO_REPLY_STRUCTURE('FrameExtents', frame_extents_fields)
    DO_SET_STRUCTURE('FrameExtents', '_NET_FRAME_EXTENTS', 'FrameExtents',
                     frame_extents_fields)

    #> _NET_WM_SYNC_REQUEST_COUNTER
    do_header('_NET_WM_SYNC_REQUEST_COUNTER')
    DO_GET_PROPERTY('WmSyncRequestCounter', '_NET_WM_SYNC_REQUEST_COUNTER',
                    'x.AtomCARDINAL', '2', 'WmSyncRequestCounter')
    request_counter_fields = ['Low', 'High']
    DO_REPLY_STRUCTURE('WmSyncRequestCounter', request_counter_fields)
    DO_SET_STRUCTURE('WmSyncRequestCounter', '_NET_WM_SYNC_REQUEST_COUNTER',
                     'WmSyncRequestCounter', request_counter_fields)

    #> _NET_WM_FULLSCREEN_MONITORS
    do_header('_NET_WM_FULLSCREEN_MONITORS')
    DO_GET_PROPERTY('WmFullscreenMonitors', '_NET_WM_FULLSCREEN_MONITORS',
                    'x.AtomCARDINAL', '4', 'WmFullscreenMonitors')
    fullscreen_monitors_fields = ['Top', 'Bottom', 'Left', 'Right']
    DO_REPLY_STRUCTURE('WmFullscreenMonitors', fullscreen_monitors_fields)
    DO_SET_STRUCTURE('WmFullscreenMonitors', '_NET_WM_FULLSCREEN_MONITORS',
                     'WmFullscreenMonitors', fullscreen_monitors_fields)

    end()


if __name__ == '__main__':
    main()
