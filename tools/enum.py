from common import *

_enum_to_str_list = [
    # xproto.xml
    ('xcb', 'VisualClass'),
    ('xcb', 'EventMask'),
    ('xcb', 'BackingStore'),
    ('xcb', 'ImageOrder'),
    ('xcb', 'ModMask'),
    ('xcb', 'ButtonMask'),
    ('xcb', 'KeyButMask'),
    ('xcb', 'Motion'),
    ('xcb', 'NotifyDetail'),
    ('xcb', 'NotifyMode'),
    ('xcb', 'Visibility'),
    ('xcb', 'Place'),
    ('xcb', 'ColormapState'),
    ('xcb', 'Mapping'),
    ('xcb', 'WindowClass'),
    ('xcb', 'Gravity'),
    ('xcb', 'MapState'),
    ('xcb', 'ConfigWindow'),
    ('xcb', 'GrabStatus'),
    ('xcb', 'FontDraw'),
    ('xcb', 'StackMode'),
    ('xcb', 'Blanking'),
    ('xcb', 'Exposures'),
    ('xcb', 'MappingStatus'),

    # TODO

    # screensaver.xml
    ('xcb', 'ScreenSaver', 'Kind'),
    ('xcb', 'ScreenSaver', 'State'),
]

def get_enum_prefix(name):
    pkg_name, name = split_pkg_and_type(name)
    # check result
    enum_prefix = ''.join(name)
    assert enum_prefix[0].isupper()

    _ns = get_namespace()
    if _ns.is_ext and pkg_name != '' and pkg_name != get_go_pkg_name(_ns.ext_name):
        # add pkg ident
        enum_prefix = pkg_name + '.' + enum_prefix

    return enum_prefix


_cname_special_cases = {
    'VISUALID': 'VisualID',        
}

_item_name_special_list = [
        # xproto.xml
        'YX',
        'XY',
        'DE',
        'RGB',
        'LSB',
        'MSB',
        'WM',

        # render.xml
        'BGR',
        'HSL',
]

def get_enum_item_name(name):
    if name in _cname_special_cases:
        return _cname_special_cases[name]

    name_parts = split_name(name)
    result = []
    for x in name_parts:
        if x in _item_name_special_list:
            result.append(x)
        else:
            result.append(x.title())
    return ''.join(result)


def define_go_enum_to_string_func(prefix, self):
    if len(self.values) < 2:
        return
    if not self.name in _enum_to_str_list:
        return

    l('\nfunc %sEnumToStr(v int) string {', prefix)

    if len(self.bits) > 0:
        l('strv := []string{}')
        l('switch {')
        for (item_name, val) in self.values:
            val = int(val)
            if val == 0:
                continue
            l('case v & %s%s > 0:', prefix, item_name)
            l('    strv = append(strv, "%s")', item_name)
        l('}') # end switch
        l('return "[" + strings.Join(strv, "|") + "]"')

    else:
        items = dict()
        # key is val, value is list of item_name
        for (item_name, val) in self.values:
            val = int(val)
            if val in items:
                items[val].append(item_name)
                # items[val] = items[val].append(item_name)
            else:
                items[val] = [ item_name ]

        l('switch v {')
        for k in sorted(items.keys()):
            l('case %d:', k)
            l('    return "%s"', ' or '.join(items[k]))

        # case default
        l('default:')
        l('return "<Unknown>"')
        l('}') # end switch

    l('}') # end func
