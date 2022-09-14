#!/usr/bin/env python3
# SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
#
# SPDX-License-Identifier: GPL-3.0-or-later

_lines = []
CHECKED = ['Checked', '']
UNCHECKED = ['', 'Unchecked']


def l(fmt, *arg):
    _lines.append(fmt % arg)


def DO_GET_PROPERTY(sth, _property, atype, length, cookie_type):
    return_type = 'Get' + cookie_type + 'Cookie'
    l('\nfunc Get%s(c *x.Conn, window x.Window) %s{', sth, return_type)
    l('    atom, _ := c.GetAtom("%s")', _property)
    l('    cookie := x.GetProperty(c, false, window,\natom, %s, 0, %s)',
        atype, length)
    l('    return %s(cookie)', return_type)
    l('}')


def DO_GET_ROOT_PROPERTY(sth, _property, atype, length, cookie_type):
    return_type = 'Get' + cookie_type + 'Cookie'
    l('\nfunc Get%s(c *x.Conn) %s{', sth, return_type)
    l('    rootWin := c.GetDefaultScreen().Root')
    l('    atom, _ := c.GetAtom("%s")', _property)
    l('    cookie := x.GetProperty(c, false, rootWin,\natom,%s, 0, %s)',
        atype, length)
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
        if checked == 'Checked':
            returnType = 'x.VoidCookie'
            returnStr = 'return'
        else:
            returnType = ''
            returnStr = ''
        l('\nfunc Set%s%s(c *x.Conn, window x.Window, val %s) %s {',
          sth, checked, go_type, returnType)
        l('    w := x.NewWriter()')
        l('    w.Write4b(uint32(val))')
        l('    atom, _ := c.GetAtom("%s")', _property)
        l('    %s x.ChangeProperty%s(c, x.PropModeReplace, window,\natom, %s, 32, w.Bytes())',
          returnStr, checked, atype)
        l('}')


def DO_ROOT_SINGLE_VALUE(sth, _property, atype, go_type, cookie_type):
    DO_GET_ROOT_PROPERTY(sth, _property, atype, '1', cookie_type)

    for checked in CHECKED:
        if checked == 'Checked':
            returnType = 'x.VoidCookie'
            returnStr = 'return'
        else:
            returnType = ''
            returnStr = ''
        l('\nfunc Set%s%s(c *x.Conn, val %s) %s {', sth, checked,
          go_type, returnType)
        l('    w := x.NewWriter()')
        l('    w.Write4b(uint32(val))')
        l('    rootWin := c.GetDefaultScreen().Root')
        l('    atom, _ := c.GetAtom("%s")', _property)
        l('    %s x.ChangeProperty%s(c, x.PropModeReplace, rootWin,\natom, %s, 32, w.Bytes())',
          returnStr, checked, atype)
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
        if checked == 'Checked':
            returnType = "x.VoidCookie"
            returnStr = 'return'
        else:
            returnType = ''
            returnStr = ''

        l('\nfunc Set%s%s(c *x.Conn, vals []%s) %s {', sth, checked,
          go_type, returnType)
        l('    w := x.NewWriter()')
        l('    for _, val := range vals {')
        l('        w.Write4b(uint32(val))')
        l('    }')  # end for

        l('    rootWin := c.GetDefaultScreen().Root')
        l('    atom, _ := c.GetAtom("%s")', _property)
        l('    %s x.ChangeProperty%s(c, x.PropModeReplace, rootWin,\n'
          +
          'atom, %s, 32, w.Bytes())',
          returnStr, checked, atype)
        l('}')  # end func


def DO_LIST_VALUES(sth, _property, atype, go_type, cookie_type):
    DO_GET_PROPERTY(sth, _property, atype, 'LENGTH_MAX', cookie_type)

    for checked in CHECKED:
        if checked == 'Checked':
            returnType = "x.VoidCookie"
            returnStr = 'return'
        else:
            returnType = ''
            returnStr = ''

        l('\nfunc Set%s%s(c *x.Conn, window x.Window,vals []%s) %s {',
          sth, checked, go_type, returnType)
        l('    w := x.NewWriter()')
        l('    for _, val := range vals {')
        l('        w.Write4b(uint32(val))')
        l('    }')  # end for

        l('    atom, _ := c.GetAtom("%s")', _property)
        l('    %s x.ChangeProperty%s(c, x.PropModeReplace, window,\n'
          +
          'atom, %s, 32, w.Bytes())',
          returnStr, checked, atype)
        l('}')  # end func


def DO_REPLY_STRUCTURE(sth, fields):
    go_type = sth
    bad_go_value = go_type + '{}'
    DO_COOKIE_REPLY(sth, 'x.AtomCardinal', go_type, bad_go_value)

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
    DO_COOKIE_REPLY(sth, 'x.AtomCardinal', '[]' + go_type, 'nil')

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
    DO_GET_PROPERTY(sth, _property, 'getAtom(c,"UTF8_STRING")', 'LENGTH_MAX',
                    'UTF8Str')

    for checked in CHECKED:
        if checked == 'Checked':
            returnType = 'x.VoidCookie'
            returnStr = 'return'
        else:
            returnType = ''
            returnStr = ''

        l('\nfunc Set%s%s(c *x.Conn, window x.Window, val string) %s {',
          sth, checked, returnType)
        l('    atom, _ := c.GetAtom("%s")', _property)
        l('    atomUTF8String, _ := c.GetAtom("UTF8_STRING")')
        l('    %s x.ChangeProperty%s(c, x.PropModeReplace, window,\natom,'
          + 'atomUTF8String, 8, []byte(val))',
          returnStr, checked)
        l('}')  # end func


def DO_ROOT_UTF8_STRING(sth, _property):
    DO_GET_ROOT_PROPERTY(sth, _property, 'getAtom(c,"UTF8_STRING")',
                         'LENGTH_MAX', 'UTF8Str')

    for checked in CHECKED:
        if checked == 'Checked':
            returnType = 'x.VoidCookie'
            returnStr = 'return'
        else:
            returnType = ''
            returnStr = ''
        l('\nfunc Set%s%s(c *x.Conn,val string) %s {', sth, checked, returnType)
        l('    rootWin := c.GetDefaultScreen().Root')
        l('    atom, _ := c.GetAtom("%s")', _property)
        l('    atomUTF8String, _ := c.GetAtom("UTF8_STRING")')
        l('    %s x.ChangeProperty%s(c, x.PropModeReplace, rootWin,\natom,'
          + 'atomUTF8String, 8, uint32(len(val)), []byte(val))',
          returnStr, checked)
        l('}')  # end func


def DO_ROOT_UTF8_STRINGS(sth, _property):
    DO_GET_ROOT_PROPERTY(sth, _property, 'getAtom(c,"UTF8_STRING")',
                         'LENGTH_MAX', 'UTF8Strs')

    for checked in CHECKED:
        if checked == 'Checked':
            returnType = 'x.VoidCookie'
            returnStr = 'return'
        else:
            returnType = ''
            returnStr = ''
        l('\nfunc Set%s%s(c *x.Conn, vals []string) %s {', sth,
          checked, returnType)
        l('    w := x.NewWriter()')
        l('    for _, val := range vals {')
        l('        w.WriteString(val)')
        l('        w.Write1b(0)')
        l('    }')  # end for
        l('    rootWin := c.GetDefaultScreen().Root')
        l('    atom, _ := c.GetAtom("%s")', _property)
        l('    atomUTF8String, _ := c.GetAtom("UTF8_STRING")')
        l('    %s x.ChangeProperty%s(c, x.PropModeReplace, rootWin,\natom,'
          + 'atomUTF8String, 8, w.Bytes())', returnStr, checked)
        l('}')  # end func


def DO_SET_ROOT_STRUCTURES(sth, _property, go_type, fields):
    n_fields = len(fields)
    for checked in CHECKED:
        if checked == 'Checked':
            returnType = 'x.VoidCookie'
            returnStr = 'return'
        else:
            returnType = ''
            returnStr = ''
        l('\nfunc Set%s%s(c *x.Conn, vals []%s) %s {', sth, checked,
          go_type, returnType)
        l('    w := x.NewWriter()')
        l('    for _, val := range vals {')
        for field in fields:
            l('        w.Write4b(val.%s)', field)
        l('    }')  # end for
        l('    rootWin := c.GetDefaultScreen().Root')
        l('    atom, _ := c.GetAtom("%s")', _property)
        l('    %s x.ChangeProperty%s(c, x.PropModeReplace, rootWin,\n'
          + 'atom, x.AtomCardinal, 32, w.Bytes())', returnStr, checked)
        l('}')  # end func


# TODO
# def DO_SET_STRUCTURES


def DO_SET_ROOT_STRUCTURE(sth, _property, go_type, fields):
    for checked in CHECKED:
        if checked == 'Checked':
            returnType = 'x.VoidCookie'
            returnStr = 'return'
        else:
            returnType = ''
            returnStr = ''

        l('\nfunc Set%s%s(c *x.Conn, val %s) %s {', sth, checked,
          go_type, returnType)
        l('    w := x.NewWriter()')
        for field in fields:
            l('    w.Write4b(val.%s)', field)

        l('    rootWin := c.GetDefaultScreen().Root')
        l('    atom, _ := c.GetAtom("%s")', _property)
        l('    %s x.ChangeProperty%s(c, x.PropModeReplace, rootWin,\natom,'
          + 'x.AtomCardinal, 32, w.Bytes())', returnStr, checked)
        l('}')  # end func


def DO_SET_STRUCTURE(sth, _property, go_type, fields):
    for checked in CHECKED:
        if checked == 'Checked':
            returnType = 'x.VoidCookie'
            returnStr = 'return'
        else:
            returnType = ''
            returnStr = ''
        l('\nfunc Set%s%s(c *x.Conn, window x.Window, val %s) %s {',
          sth, checked, go_type, returnType)
        l('    w := x.NewWriter()')
        for field in fields:
            l('    w.Write4b(val.%s)', field)

        l('    atom, _ := c.GetAtom("%s")', _property)
        l('    %s x.ChangeProperty%s(c, x.PropModeReplace, window,\natom,'
          + 'x.AtomCardinal, 32, w.Bytes())', returnStr, checked)
        l('}')  # end func


def DO_COOKIE_REPLY(sth, atype, go_type, bad_go_value):
    cookie_type = 'Get' + sth + 'Cookie'
    get_sth_from_reply = 'get' + sth + 'FromReply'
    l('type %s x.GetPropertyCookie', cookie_type)

    l('\nfunc (cookie %s) Reply(c *x.Conn) (%s, error) {', cookie_type, go_type)
    l('    reply, err := x.GetPropertyCookie(cookie).Reply(c)')
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

    DO_COOKIE_REPLY('UTF8Str', 'getAtom(c,"UTF8_STRING")', 'string', '""')
    DO_COOKIE_REPLY('UTF8Strs', 'getAtom(c,"UTF8_STRING")', '[]string', 'nil')
    DO_COOKIE_REPLY('Boolean', 'x.AtomCardinal', 'bool', 'false')

    DO_REPLY_SINGLE_VALUE('Window', 'x.AtomWindow', 'x.Window', 'Window')
    DO_REPLY_SINGLE_VALUE('Cardinal', 'x.AtomCardinal', 'uint32', 'Cardinal')

    DO_REPLY_LIST_VALUES('Windows', 'x.AtomWindow', 'x.Window')
    DO_REPLY_LIST_VALUES('Atoms', 'x.AtomAtom', 'x.Atom')
    DO_REPLY_LIST_VALUES('Cardinals', 'x.AtomCardinal', 'uint32')

    #> _NET_SUPPORTED
    do_header('_NET_SUPPORTED')
    DO_ROOT_LIST_VALUES('Supported', '_NET_SUPPORTED', 'x.AtomAtom', 'x.Atom',
                        'Atoms')

    #> _NET_CLIENT_LIST
    do_header('_NET_CLIENT_LIST')
    DO_ROOT_LIST_VALUES('ClientList', '_NET_CLIENT_LIST', 'x.AtomWindow',
                        'x.Window', 'Windows')

    #> _NET_CLIENT_LIST_STACKING
    do_header('_NET_CLIENT_LIST_STACKING')
    DO_ROOT_LIST_VALUES('ClientListStacking', '_NET_CLIENT_LIST_STACKING',
                        'x.AtomWindow', 'x.Window', 'Windows')

    #> _NET_NUMBER_OF_DESKTOPS
    do_header('_NET_NUMBER_OF_DESKTOPS')
    DO_ROOT_SINGLE_VALUE('NumberOfDesktop', '_NET_NUMBER_OF_DESKTOPS',
                         'x.AtomCardinal', 'uint32', 'Cardinal')

    #> _NET_DESKTOP_GEOMETRY
    do_header('_NET_DESKTOP_GEOMETRY')
    DO_GET_ROOT_PROPERTY('DesktopGeometry', '_NET_DESKTOP_GEOMETRY',
                         'x.AtomCardinal', '2', 'DesktopGeometry')
    desktop_geometry_fields = ['Width', 'Height']
    DO_REPLY_STRUCTURE('DesktopGeometry', desktop_geometry_fields)
    DO_SET_ROOT_STRUCTURE('DesktopGeometry', '_NET_DESKTOP_GEOMETRY',
                          'DesktopGeometry', desktop_geometry_fields)

    #> _NET_DESKTOP_VIEWPORT
    do_header('_NET_DESKTOP_VIEWPORT')
    DO_GET_ROOT_PROPERTY('DesktopViewport', '_NET_DESKTOP_VIEWPORT',
                         'x.AtomCardinal', 'LENGTH_MAX', 'DesktopViewport')
    viewport_fields = ['X', 'Y']
    DO_REPLY_STRUCTURES('DesktopViewport', 'ViewportLeftTopCorner',
                        viewport_fields)
    DO_SET_ROOT_STRUCTURES('DesktopViewport', '_NET_DESKTOP_VIEWPORT',
                           'ViewportLeftTopCorner', viewport_fields)

    #> _NET_CURRENT_DESKTOP
    do_header('_NET_CURRENT_DESKTOP')
    DO_ROOT_SINGLE_VALUE('CurrentDesktop', '_NET_CURRENT_DESKTOP',
                         'x.AtomCardinal', 'uint32', 'Cardinal')

    #> _NET_DESKTOP_NAMES
    do_header('_NET_DESKTOP_NAMES')
    DO_ROOT_UTF8_STRINGS('DesktopNames', '_NET_DESKTOP_NAMES')

    #> _NET_ACTIVE_WINDOW
    do_header('_NET_ACTIVE_WINDOW')
    DO_ROOT_SINGLE_VALUE('ActiveWindow', '_NET_ACTIVE_WINDOW', 'x.AtomWindow',
                         'x.Window', 'Window')

    #> _NET_WORKAREA
    do_header('_NET_WORKAREA')
    DO_GET_ROOT_PROPERTY('Workarea', '_NET_WORKAREA', 'x.AtomCardinal',
                         'LENGTH_MAX', 'Workarea')
    workarea_geo_fields = ['X', 'Y', 'Width', 'Height']
    DO_REPLY_STRUCTURES('Workarea', 'WorkareaGeometry', workarea_geo_fields)
    DO_SET_ROOT_STRUCTURES('Workarea', '_NET_WORKAREA', 'WorkareaGeometry',
                           workarea_geo_fields)

    #> _NET_SUPPORTING_WM_CHECK
    do_header('_NET_SUPPORTING_WM_CHECK')
    DO_ROOT_SINGLE_VALUE('SupportingWMCheck', '_NET_SUPPORTING_WM_CHECK',
                    'x.AtomWindow', 'x.Window', 'Window')

    #>  _NET_VIRTUAL_ROOTS
    do_header('_NET_VIRTUAL_ROOTS')
    DO_ROOT_LIST_VALUES('VirtualRoots', '_NET_VIRTUAL_ROOTS', 'x.AtomWindow',
                        'x.Window', 'Windows')

    #> _NET_DESKTOP_LAYOUT
    do_header('_NET_DESKTOP_LAYOUT')
    DO_GET_ROOT_PROPERTY('DesktopLayout', '_NET_DESKTOP_LAYOUT',
                         'x.AtomCardinal', '4', 'DesktopLayout')
    desktop_layout_fields = [
        'Orientation', 'Columns', 'Rows', 'StartingCorner'
    ]
    DO_REPLY_STRUCTURE('DesktopLayout', desktop_layout_fields)
    DO_SET_ROOT_STRUCTURE('DesktopLayout', '_NET_DESKTOP_LAYOUT',
                          'DesktopLayout', desktop_layout_fields)

    #> _NET_SHOWING_DESKTOP
    do_header('_NET_SHOWING_DESKTOP')
    DO_GET_ROOT_PROPERTY('ShowingDesktop', '_NET_SHOWING_DESKTOP',
                         'x.AtomCardinal', '1', 'Boolean')

    #> _NET_WM_NAME
    do_header('_NET_WM_NAME')
    DO_UTF8_STRING('WMName', '_NET_WM_NAME')

    #> _NET_WM_VISIBLE_NAME
    do_header('_NET_WM_VISIBLE_NAME')
    DO_UTF8_STRING('WMVisibleName', '_NET_WM_VISIBLE_NAME')

    #> _NET_WM_ICON_NAME
    do_header('_NET_WM_ICON_NAME')
    DO_UTF8_STRING('WMIconName', '_NET_WM_ICON_NAME')

    #> _NET_WM_VISIBLE_ICON_NAME
    do_header('_NET_WM_VISIBLE_ICON_NAME')
    DO_UTF8_STRING('WMVisibleIconName', '_NET_WM_VISIBLE_ICON_NAME')

    #> _NET_WM_DESKTOP
    do_header('_NET_WM_DESKTOP')
    DO_SINGLE_VALUE('WMDesktop', '_NET_WM_DESKTOP', 'x.AtomCardinal', 'uint32',
                    'Cardinal')

    #> _NET_WM_WINDOW_TYPE
    do_header('_NET_WM_WINDOW_TYPE')
    DO_LIST_VALUES('WMWindowType', '_NET_WM_WINDOW_TYPE', 'x.AtomAtom',
                   'x.Atom', 'Atoms')

    #> _NET_WM_STATE
    do_header('_NET_WM_STATE')
    DO_LIST_VALUES('WMState', '_NET_WM_STATE', 'x.AtomAtom', 'x.Atom', 'Atoms')

    #> _NET_WM_ALLOWED_ACTIONS
    do_header('_NET_WM_ALLOWED_ACTIONS')
    DO_LIST_VALUES('WMAllowedActions', '_NET_WM_ALLOWED_ACTIONS', 'x.AtomAtom',
                   'x.Atom', 'Atoms')

    #> _NET_WM_STRUT
    do_header('_NET_WM_STRUT')
    DO_GET_PROPERTY('WMStrut', '_NET_WM_STRUT', 'x.AtomCardinal', '4',
                    'WMStrut')
    strut_fields = ['Left', 'Right', 'Top', 'Bottom']
    DO_REPLY_STRUCTURE('WMStrut', strut_fields)
    DO_SET_STRUCTURE('WMStrut', '_NET_WM_STRUT', 'WMStrut', strut_fields)

    #> _NET_WM_STRUT_PARTIAL
    do_header('_NET_WM_STRUT_PARTIAL')
    DO_GET_PROPERTY('WMStrutPartial', '_NET_WM_STRUT_PARTIAL',
                    'x.AtomCardinal', '12', 'WMStrutPartial')
    strut_partial_fields = strut_fields + [
        'LeftStartY', 'LeftEndY', 'RightStartY', 'RightEndY', 'TopStartX',
        'TopEndX', 'BottomStartX', 'BottomEndX'
    ]
    DO_REPLY_STRUCTURE('WMStrutPartial', strut_partial_fields)
    DO_SET_STRUCTURE('WMStrutPartial', '_NET_WM_STRUT_PARTIAL',
                     'WMStrutPartial', strut_partial_fields)

    #> _NET_WM_ICON_GEOMETRY
    do_header('_NET_WM_ICON_GEOMETRY')
    DO_GET_PROPERTY('WMIconGeometry', '_NET_WM_ICON_GEOMETRY',
                    'x.AtomCardinal', '4', 'WMIconGeometry')
    icon_geometry = ['X', 'Y', 'Width', 'Height']
    DO_REPLY_STRUCTURE('WMIconGeometry', icon_geometry)
    DO_SET_STRUCTURE('WMIconGeometry', '_NET_WM_ICON_GEOMETRY',
                     'WMIconGeometry', icon_geometry)

    #> _NET_WM_ICON
    do_header('_NET_WM_ICON')
    DO_GET_PROPERTY('WMIcon', '_NET_WM_ICON', 'x.AtomCardinal', 'LENGTH_MAX',
                    'WMIcon')
    DO_COOKIE_REPLY('WMIcon', 'x.AtomCardinal', '[]WMIcon', 'nil')

    #> _NET_WM_PID
    do_header('_NET_WM_PID')
    DO_SINGLE_VALUE('WMPid', '_NET_WM_PID', 'x.AtomCardinal', 'uint32',
                    'Cardinal')

    #> _NET_WM_HANDLED_ICONS
    do_header('_NET_WM_HANDLED_ICONS')
    DO_SINGLE_VALUE('WMHandledIcons', '_NET_WM_HANDLED_ICONS',
                    'x.AtomCardinal', 'uint32', 'Cardinal')

    #> _NET_WM_USER_TIME
    do_header('_NET_WM_USER_TIME')
    DO_SINGLE_VALUE('WMUserTime', '_NET_WM_USER_TIME', 'x.AtomCardinal',
                    'uint32', 'Cardinal')

    #> _NET_WM_USER_TIME_WINDOW
    do_header('_NET_WM_USER_TIME_WINDOW')
    DO_SINGLE_VALUE('WMUserTimeWindow', '_NET_WM_USER_TIME_WINDOW',
                    'x.AtomWindow', 'x.Window', 'Window')

    #> _NET_FRAME_EXTENTS
    do_header('_NET_FRAME_EXTENTS')
    DO_GET_PROPERTY('FrameExtents', '_NET_FRAME_EXTENTS', 'x.AtomCardinal',
                    '4', 'FrameExtents')
    frame_extents_fields = ['Left', 'Right', 'Top', 'Bottom']
    DO_REPLY_STRUCTURE('FrameExtents', frame_extents_fields)
    DO_SET_STRUCTURE('FrameExtents', '_NET_FRAME_EXTENTS', 'FrameExtents',
                     frame_extents_fields)

    #> _NET_WM_SYNC_REQUEST_COUNTER
    do_header('_NET_WM_SYNC_REQUEST_COUNTER')
    DO_GET_PROPERTY('WMSyncRequestCounter', '_NET_WM_SYNC_REQUEST_COUNTER',
                    'x.AtomCardinal', '2', 'WMSyncRequestCounter')
    request_counter_fields = ['Low', 'High']
    DO_REPLY_STRUCTURE('WMSyncRequestCounter', request_counter_fields)
    DO_SET_STRUCTURE('WMSyncRequestCounter', '_NET_WM_SYNC_REQUEST_COUNTER',
                     'WMSyncRequestCounter', request_counter_fields)

    #> _NET_WM_FULLSCREEN_MONITORS
    do_header('_NET_WM_FULLSCREEN_MONITORS')
    DO_GET_PROPERTY('WMFullscreenMonitors', '_NET_WM_FULLSCREEN_MONITORS',
                    'x.AtomCardinal', '4', 'WMFullscreenMonitors')
    fullscreen_monitors_fields = ['Top', 'Bottom', 'Left', 'Right']
    DO_REPLY_STRUCTURE('WMFullscreenMonitors', fullscreen_monitors_fields)
    DO_SET_STRUCTURE('WMFullscreenMonitors', '_NET_WM_FULLSCREEN_MONITORS',
                     'WMFullscreenMonitors', fullscreen_monitors_fields)

    end()


if __name__ == '__main__':
    main()
