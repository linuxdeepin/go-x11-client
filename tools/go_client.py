import sys
from common import *
from enum import *
import argparse
import subprocess

error_names = []
request_names = []
max_error_code = 0
x_prefix = ''

parser = argparse.ArgumentParser()
parser.add_argument('xml_file')
parser.add_argument('-o', '--output', dest='output', default='/dev/stdout')
parser.add_argument('-F', '--no-fmt', dest='no_fmt', action='store_true')
parser.add_argument('-p', '--package', dest='package')
parser.add_argument(
    '-I', '--no-import-x', dest='no_import_x', action='store_true')

args = parser.parse_args()

def go_open(self):
    set_namespace(self.namespace)
    _ns = get_namespace()

    global args
    if args.package is not None:
        pkg = args.package
    else:
        # pkg = _ns.header
        if _ns.is_ext:
            pkg = get_go_pkg_name(_ns.ext_name)
        else:
            pkg = 'x'

    set_level(0)
    l('package %s', pkg)

    if pkg != 'x':
        l('import x "github.com/linuxdeepin/go-x11-client"')
        global x_prefix
        x_prefix = 'x.'

    if _ns.is_ext:
        l('// _ns.ext_name: %s', _ns.ext_name)
        l('const MajorVersion = %d', int(_ns.major_version))
        l('const MinorVersion = %d', int(_ns.minor_version))

        l('var _ext *%sExtension', x_prefix)
        l('func Ext() *x.Extension {')
        l('    return _ext')
        l('}')

def go_close(self):
    # handle errors
    _ns = get_namespace()
    global error_names
    global request_names
    global x_prefix

    if len(error_names) > 0:
        l("var errorCodeNameMap = map[uint8]string{")
        for name in error_names:
            error_code = get_type_name(name) + "ErrorCode"
            error_name = get_type_name(name)
            if not error_name.startswith("Bad"):
                error_name = "Bad" + error_name
            if error_name == "BadDeviceBusy":
                error_name = "DeviceBusy"
            l("%s : \"%s\",", error_code, error_name)
        l("}") # end error code name map

    if len(request_names) > 0:
        l("var requestOpcodeNameMap = map[uint]string{")
        for name in request_names:
            l("%sOpcode : \"%s\",", name, name)
        l("}") # end request opcode name map

    if _ns.is_ext:
        l("func init() {")
        errMap = "nil"
        if len(error_names) > 0:
            errMap = "errorCodeNameMap"
        reqMap = "nil"
        if len(request_names) > 0:
            reqMap = "requestOpcodeNameMap"
    
        l('_ext = %sNewExtension("%s", %d, %s, %s)', x_prefix, _ns.ext_xname, max_error_code, errMap, reqMap)
        l("}") # end func init

    # write source file
    global args
    gofile = open(args.output, 'w')
    if args.no_fmt:
        output_lines(gofile)
    else:
        pipe = subprocess.Popen(
            'gofmt', bufsize=4096, stdin=subprocess.PIPE, stdout=gofile)
        output_lines(pipe.stdin)
        pipe.wait()
        if pipe.returncode != 0:
            # write no formated code and raise exception.
            output_lines(gofile)
            raise Exception("gofmt exit with error code %d" % pipe.returncode)

def go_simple(self, name):
    if self.name != name:
        type_name = get_type_name(name)
        source_type_name = get_type_name(self.name)
        set_level(0)
        l("// simple %s", name)
        l("type %s %s", type_name, source_type_name)

def go_enum(self, name):
    prefix = get_enum_prefix(name)
    set_level(0)
    l("// enum %s", prefix)
    l('const (')
    for (item_name, val) in self.values:
        item_name = get_enum_item_name(item_name)
        l("%s%s = %s", prefix, item_name, val)
    # newline
    l(')\n')


def go_struct(self, name):
    pass

def go_union(self, name):
    pass

def go_request(self, name):
    global request_names
    name_base = get_request_name(name)
    opcode_name = name_base + "Opcode"
    l("const %s = %s", opcode_name, self.opcode)
    request_names.append(name_base)

    if self.reply:
        l("type %s %sSeqNum", name_base + "Cookie", x_prefix)

def go_event(self, name):
    type_name = get_type_name(name) + "Event"
    code_name = type_name + "Code"
    code = self.opcodes[name]
    l('const %s = %d', code_name, int(code))
    define_go_event_new_func(type_name)


def define_go_event_new_func(go_type):
    global x_prefix
    l('\nfunc New%s(data []byte) (*%s, error) {', go_type, go_type)
    l('var ev %s', go_type)
    l('r := %sNewReaderFromData(data)', x_prefix)
    l('err := read%s(r, &ev)', go_type)
    l('if err != nil {')
    l('    return nil, err')
    l('}')
    l('return &ev, nil')
    l('}')

def go_error(self, name):
    global max_error_code
    global error_names
    type_name = get_type_name(name) + 'Error'
    code_name = type_name + "Code"
    code = int(self.opcodes[name])
    l('const %s = %d', code_name, code)
    error_names.append(name)
    if code > max_error_code:
        max_error_code = code

def go_eventstruct(self, name):
    pass

output = {
    'open': go_open,
    'close': go_close,
    'simple': go_simple,
    'enum': go_enum,
    'struct': go_struct,
    'union': go_union,
    'request': go_request,
    'event': go_event,
    'error': go_error,
    'eventstruct': go_eventstruct,
}
# Import the module class
from xcbgen.state import Module
from xcbgen.xtypes import *

# Parse the xml header
module = Module(args.xml_file, output)

# Build type-registry and resolve type dependencies
module.register()
module.resolve()

# Output the code
module.generate()
