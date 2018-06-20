package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"strings"

	"pkg.deepin.io/lib/strv"
)

var optIsExt bool

func init() {
	flag.BoolVar(&optIsExt, "e", false, "is ext")
}

var xPrefix string

func astNodeToStr(fset *token.FileSet, node interface{}) (string, error) {
	var buf bytes.Buffer
	err := printer.Fprint(&buf, fset, node)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

type Generator struct {
	buf bytes.Buffer
}

func (g *Generator) p(format string, args ...interface{}) (int, error) {
	return fmt.Fprintf(&g.buf, format, args...)
}

func (g *Generator) format() []byte {
	src, err := format.Source(g.buf.Bytes())
	if err != nil {
		log.Println("warning: internal error: invalid Go generated:", err)
		return g.buf.Bytes()
	}
	return src
}

type requestFunc struct {
	name          string
	writeFuncName string
	writeFuncArgs string
	args          string
	noReply       bool
}

func main() {
	log.SetFlags(log.Lshortfile)

	goPackage := os.Getenv("GOPACKAGE")
	log.Println(goPackage)
	// parse file

	if goPackage != "x" {
		xPrefix = "x."
	}

	fs := token.NewFileSet()
	flag.Parse()
	name := flag.Arg(0)
	f, err := parser.ParseFile(fs, name, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	var requestFuncs []*requestFunc
	var hasReplyRequests strv.Strv
	ast.Inspect(f, func(node ast.Node) bool {
		funcDel, ok := node.(*ast.FuncDecl)
		if !ok {
			return true
		}
		if funcDel.Recv != nil {
			// is method
			return false
		}

		funcName := funcDel.Name.Name
		if strings.HasPrefix(funcName, "write") {

			if !strings.Contains(funcDel.Doc.Text(), "#WREQ") {
				return false
			}

			params := funcDel.Type.Params.List
			var args []string
			var writeFuncArgs = []string{"w"}
			for idx, param := range params {
				if idx == 0 {
					// skip first param
					continue
				}

				log.Println(param.Names)

				var paramNames []string
				for _, name := range param.Names {
					paramNames = append(paramNames, name.Name)
				}
				typeStr, err := astNodeToStr(fs, param.Type)
				if err != nil {
					log.Fatal(err)
				}

				argNamesType := strings.Join(paramNames, ",") + " " + typeStr

				args = append(args, argNamesType)
				writeFuncArgs = append(writeFuncArgs, paramNames...)
			}

			requestFuncs = append(requestFuncs, &requestFunc{
				writeFuncName: funcName,
				writeFuncArgs: strings.Join(writeFuncArgs, ","),
				name:          strings.TrimPrefix(funcName, "write"),
				args:          strings.Join(args, ","),
			})
		} else if strings.HasPrefix(funcName, "read") &&
			strings.HasSuffix(funcName, "Reply") {
			reqName := strings.TrimPrefix(funcName, "read")
			reqName = strings.TrimSuffix(reqName, "Reply")
			hasReplyRequests = append(hasReplyRequests, reqName)
		}

		return false
	})

	log.Println(hasReplyRequests)
	for _, reqFunc := range requestFuncs {
		if !hasReplyRequests.Contains(reqFunc.name) {
			reqFunc.noReply = true
		}
	}
	//spew.Dump(requestFuncs)

	g := &Generator{}

	g.p("package %s\n", goPackage)

	if goPackage != "x" {
		g.p("import x \"github.com/linuxdeepin/go-x11-client\"")
	}

	for _, reqFunc := range requestFuncs {
		if reqFunc.noReply {
			g.pRequestFunc(reqFunc, false)
			g.pRequestFunc(reqFunc, true)
		} else {
			g.pRequestFunc(reqFunc, true)
			// cookie Reply method
			g.p("\nfunc (cookie %sCookie) Reply(conn *%sConn) (*%sReply, error) {\n",
				reqFunc.name, xPrefix, reqFunc.name)
			g.p(`replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, conn.NewError(replyBuf)
	}
	r := %sNewReaderFromData(replyBuf)
	var reply %sReply
	err := read%sReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}
`, xPrefix, reqFunc.name, reqFunc.name)
		}

	}

	os.Stdout.Write(g.format())
}

func (g *Generator) pRequestFunc(reqFunc *requestFunc, check bool) {
	returnType := xPrefix + "VoidCookie"
	if !reqFunc.noReply {
		// has reply
		returnType = reqFunc.name + "Cookie"
	} else if !check {
		// no reply + no check
		returnType = ""
	}
	funcName := reqFunc.name
	sendReqFlag := "0"
	if check {
		sendReqFlag = xPrefix + "RequestChecked"
	}

	if reqFunc.noReply && check {
		funcName += "Checked"
	}

	g.p("\nfunc %s(conn *%sConn, %s) %s {\n", funcName, xPrefix, reqFunc.args,
		returnType)

	g.p("w := %sNewWriter()\n", xPrefix)
	g.p("%s(%s)\n", reqFunc.writeFuncName, reqFunc.writeFuncArgs)
	g.p("d := w.Bytes()\n")
	g.p("req := &%sProtocolRequest{\n", xPrefix)

	if optIsExt {
		g.p("Ext: _ext,\n")
	}

	if reqFunc.noReply {
		g.p("NoReply: true,\n")
	}

	g.p("Opcode: %sOpcode,\n", reqFunc.name)
	g.p("}\n")

	if returnType == "" {
		g.p("conn.SendRequest(%s, d, req)\n", sendReqFlag)
	} else {
		g.p("seq := conn.SendRequest(%s, d, req)\n", sendReqFlag)
		g.p("return %s(seq)", returnType)
	}

	g.p("}\n")
}
