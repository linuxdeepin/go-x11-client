import re

_no_import_x = True
_ns = None
_lines = []
_level = 0

_cname_re = re.compile('([A-Z0-9][a-z]+|[A-Z0-9]+(?![a-z])|[a-z]+)')
_cname_special_cases = {
    # names in xproto.xml
    'GCONTEXT': 'GContext',
    'VISUALID': 'VisualID',
    'VISUALTYPE': 'VisualType',
    'RGB': 'RGB',
    'FONTPROP': 'FontProp',
    'CHARINFO': 'CharInfo',
    'TIMECOORD': 'TimeCoord',
    'COLORITEM': 'ColorItem',
    'DECnet': 'DECnet',

    # in render.xml
    'GLYPHSET': 'GlyphSet',
    'PICTFORMAT': 'PictFormat',
    'DIRECTFORMAT': 'DirectFormat',
    'PICTFORMINFO': 'PictFormInfo',
    'PICTVISUAL': 'PictVisual',
    'PICTDEPTH': 'PictDepth',
    'PICTSCREEN': 'PictScreen',
    'INDEXVALUE': 'IndexValue',
    'POINTFIX': 'PointFix',
    'LINEFIX': 'LineFix',
    'GLYPHINFO': 'GlyphInfo',
    'ANIMCURSORELT': 'AnimCursorElt',
    'SPANFIX': 'SpanFix',
}

_simple_type_map = {
    'uint8_t': 'uint8',
    'uint16_t': 'uint16',
    'uint32_t': 'uint32',
    'uint64_t': 'uint64',
    'int8_t': 'int8',
    'int16_t': 'int16',
    'int32_t': 'int32',
    'int64_t': 'int64',
    'char': 'byte',  # []byte => string
    'float': 'float32',
    'double': 'float64',
}

# readonly
# ingore bigreq
_ext_names = ['Shm', 'XF86Dri', 'RandR', 'Xv', 'DPMS', 'Xinerama',
        'Test', 'XCMisc', 'Res', 'XvMC', 'Composite', 'ScreenSaver',
        'SELinux', 'XF86VidMode', 'Present', 'DRI3', 'Input',
        'GenericEvent', 'Render', 'Xevie', 'XFixes', 'Glx', 'Sync',
        'XPrint', 'Shape', 'xkb', 'Damage', 'DRI2', 'Record' ]

# xml filename => go pkg name
# if the value is '-', it means that the value is equal to the key
_ext_header_pkgname_map = {
        'composite': '-',
        'damage': '-',
        'dpms': '-',
        'dri2': '-',
        'dri3': '-',
        'ge': 'genericevent',
        'glx': '-',
        'present': '-',
        'randr': '-',
        'record': '-',
        'render': '-',
        'res': '-',
        'screensaver': '-',
        'shape': '-',
        'shm': '-',
        'sync': '-',
        'xc_misc': 'xcmisc',
        'xevie': '-',
        'xf86dir': '-',
        'xf86vidmode': '-',
        'xfixes': '-',
        'xinerama': '-',
        'xinput': 'input',
        'xkb': '-',
        'xprint': '-',
        'xselinux': 'selinux',
        'xtest': 'test',
        'xv': '-',
        'xvmc': '-',
}

def is_int(x):
    try:
        int(x)
        return True
    except:
        return False


def l(fmt, *args):
    _lines[_level].append(fmt % args)

def todo(msg=''):
    l('// TODO %s', msg)

def set_level(idx):
    global _level
    while len(_lines) <= idx:
        _lines.append([])
    _level = idx


def output_lines(fh):
    for list in _lines:
        for line in list:
            fh.write(line)
            fh.write('\n')
    fh.close()


def set_namespace(val):
    global _ns
    _ns = val


def get_namespace():
    global _ns
    return _ns


def set_no_import_x(val):
    global _no_import_x
    _no_import_x = val


def lib_id(_id):
    global _no_import_x
    if _no_import_x:
        return _id
    else:
        return "x." + _id


def get_type_name(name):
    # type of name is tuple of str
    if len(name) == 1:
        # uintX intX
        return _simple_type_map.get(name[0], name[0])

    pkg_name, name = split_pkg_and_type(name)
    type_name = ''.join([n_item(x) for x in name])

    # check result
    assert type_name[0].isupper()

    global _ns
    if _ns.is_ext and pkg_name != '' and pkg_name != get_go_pkg_name(_ns.ext_name):
        # add pkg ident
        type_name = pkg_name + '.' + type_name

    return type_name

def split_pkg_and_type(name):
    assert len(name) >= 2

    if len(name) == 2 :
        # ex. xcb COLORMAP
        assert name[0] == 'xcb'
        return ('x', name[1:])
    elif len(name) >= 3:
        assert name[0] == 'xcb'
        global _ns
        global _ext_names
        if name[1] == _ns.ext_name or name[1] in _ext_names:
            # name[1] is ext name
            # ex. xcb RandR ScreenSize
            return (get_go_pkg_name(name[1]), name[2:])
        else:
            # ex. xcb CreateWindow ValueList
            return ('x', name[1:])

# go pkg name is _ns.ext_name.lower()
def get_go_pkg_name(ext_name):
    # TODO adjust package name
    return ext_name.lower()

def header_to_go_pkg_name(header):
    global _ext_header_pkgname_map
    pkg_name = _ext_header_pkgname_map[header]
    if pkg_name == '-':
        pkg_name = header

    return pkg_name

def remove_prefix(name):
    # remove xcb
    if name[0] == 'xcb':
        name = name[1:]

    # remove ext name
    if _ns.is_ext and name[0] == _ns.ext_name:
        name = name[1:]

    return name


def split_name(str):
    split = _cname_re.finditer(str)
    name_parts = [match.group(0) for match in split]
    return name_parts

# value_list => value, list => ValueList
# CreateWindow => Create, Window => CreateWindow
def n_item(str):
    '''
    Does C-name conversion on a single string fragment.
    Uses a regexp with some hard-coded special cases.
    '''
    if str in _cname_special_cases:
        return _cname_special_cases[str]
    else:
        split = _cname_re.finditer(str)
        name_parts = [match.group(0) for match in split]
        return ''.join([x.title() for x in name_parts])


def get_request_name(name):
    # print "get_request_name, name:", name
    name = remove_prefix(name)
    req_name = ''.join(name)
    if not req_name[0].isupper():
        raise Exception('first char of request name is not upper case')
    return req_name


# field_name_test => FieldNameTest
def get_field_name(name):
    if name is None:
        return ''
    return ''.join([x.title() for x in name.split('_')])

class FieldClassifyModel():
    def do(self, field):
        if field.type.is_pad:
            # pad
            if field.type.fixed_size():
                return self.pad_FS(field)
            else:
                return self.pad_VS(field)

        elif field.type.is_simple:
            # simple
            return self.simple(field)

        elif field.type.is_container:
            # containter
            if field.type.fixed_size():
                return self.container_FS(field)
            else:
                return self.container_VS(field)

        elif field.type.is_list:
            # list
            if field.type.nmemb is None:
                # VL
                if field.type.member.fixed_size():
                    # MFS
                    return self.list_VL_MFS(field)
                else:
                    # MVS
                    return self.list_VL_MVS(field)
            else:
                # CL
                if field.type.member.fixed_size():
                    #MFS
                    return self.list_CL_MFS(field)
                else:
                    #MVS
                    return self.list_CL_MVS(field)

    def pad_FS(self, field):
        raise Exception('abstract method not overridden!')

    def pad_VS(self, field):
        raise Exception('abstract method not overridden!')

    def simple(self, field):
        raise Exception('abstract method not overridden!')

    def container_FS(self, field):
        raise Exception('abstract method not overridden!')

    def container_VS(self, field):
        raise Exception('abstract method not overridden!')

    def list_CL_MFS(self, field):
        raise Exception('abstract method not overridden!')

    def list_CL_MVS(self, field):
        raise Exception('abstract method not overridden!')

    def list_VL_MFS(self, field):
        raise Exception('abstract method not overridden!')

    def list_VL_MVS(self, field):
        raise Exception('abstract method not overridden!')

