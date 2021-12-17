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

	"github.com/linuxdeepin/go-lib/strv"
)

var optIsExt bool
var optExtraExts string

func init() {
	flag.BoolVar(&optIsExt, "e", false, "is ext")
	flag.StringVar(&optExtraExts, "extra-exts", "", "")
}

var xPrefix string

func astNodeToStr(fileSet *token.FileSet, node interface{}) (string, error) {
	var buf bytes.Buffer
	err := printer.Fprint(&buf, fileSet, node)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

type Generator struct {
	buf bytes.Buffer
}

func (g *Generator) p(format string, args ...interface{}) {
	_, err := fmt.Fprintf(&g.buf, format, args...)
	if err != nil {
		log.Fatal(err)
	}
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
	name           string
	encodeFuncName string
	encodeFuncArgs string
	args           string
	noReply        bool
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
		if strings.HasPrefix(funcName, "encode") {

			docText := funcDel.Doc.Text()
			if !strings.Contains(docText, "#WREQ") {
				return false
			}
			wreqOption := strings.TrimSpace(strings.TrimPrefix(docText, "#WREQ"))
			wreqOptions := strings.Split(wreqOption, " ")
			var varargs []string
			for _, value := range wreqOptions {
				if strings.HasPrefix(value, "vararg:") {
					varargs = append(varargs, value[len("vararg:"):])
				}
			}

			params := funcDel.Type.Params.List
			var args []string
			var encodeFuncArgs []string
			for _, param := range params {
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
				for _, value := range paramNames {
					if strv.Strv(varargs).Contains(value) {
						value = "options..."
					}
					encodeFuncArgs = append(encodeFuncArgs, value)
				}
			}

			requestFuncs = append(requestFuncs, &requestFunc{
				encodeFuncName: funcName,
				encodeFuncArgs: strings.Join(encodeFuncArgs, ","),
				name:           strings.TrimPrefix(funcName, "encode"),
				args:           strings.Join(args, ","),
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

	const pkgPathBase = "github.com/linuxdeepin/go-x11-client"
	if goPackage != "x" {
		g.p("import x \"%s\"\n", pkgPathBase)
	}

	extraExts := strings.Split(optExtraExts, ",")
	for _, e := range extraExts {
		if e == "" {
			continue
		}
		g.p("import \"%s/ext/%s\"\n", pkgPathBase, e)
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
			g.p(`replyBuf, err := conn.WaitForReply(%sSeqNum(cookie))
	if err != nil {
		return nil, err
	}
	r := %sNewReaderFromData(replyBuf)
	var reply %sReply
	err = read%sReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}
`, xPrefix, xPrefix, reqFunc.name, reqFunc.name)
		}

	}

	_, err = os.Stdout.Write(g.format())
	if err != nil {
		log.Fatal(err)
	}
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

	if optIsExt {
		g.p("body := %s(%s)\n", reqFunc.encodeFuncName, reqFunc.encodeFuncArgs)
	} else {
		g.p("headerData, body := %s(%s)\n", reqFunc.encodeFuncName, reqFunc.encodeFuncArgs)
	}
	g.p("req := &%sProtocolRequest{\n", xPrefix)

	if optIsExt {
		g.p("Ext: _ext,\n")
	}

	if reqFunc.noReply {
		g.p("NoReply: true,\n")
	}

	g.p("Header: %sRequestHeader{\n", xPrefix)

	if optIsExt {
		g.p("    Data: %sOpcode,\n", reqFunc.name)
	} else {
		g.p("    Opcode: %sOpcode,\n", reqFunc.name)
		g.p("    Data: headerData,\n")
	}

	g.p("},\n")
	g.p("Body: body,\n")
	g.p("}\n")

	if returnType == "" {
		g.p("conn.SendRequest(%s, req)\n", sendReqFlag)
	} else {
		g.p("seq := conn.SendRequest(%s, req)\n", sendReqFlag)
		g.p("return %s(seq)", returnType)
	}

	g.p("}\n")
}
