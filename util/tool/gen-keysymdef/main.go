package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

const (
	goPkgName     = "keysyms"
	keysymdefFile = "/usr/include/X11/keysymdef.h"
)

type Keysym struct {
	Name    string
	Value   string
	Unicode string
	Comment string
}

var reIfDef = regexp.MustCompile(`^#ifdef\s+(\w+)`)
var reDef0 = regexp.MustCompile(`^#define (XK_[a-zA-Z_0-9]+)\s+0x([0-9a-f]+)\s*/\* U\+([0-9A-F]{4,6}) (.*)\s*\*/`)
var reDef1 = regexp.MustCompile(`^#define (XK_[a-zA-Z_0-9]+)\s+0x([0-9a-f]+)\s*/\*\(U\+([0-9A-F]{4,6}) (.*)\)\*/`)
var reDef2 = regexp.MustCompile(`^#define (XK_[a-zA-Z_0-9]+)\s+0x([0-9a-f]+)\s*(/\*\s*(.*)\s*\*/)?`)

var sections = []string{
	"XK_MISCELLANY",
	"XK_XKB_KEYS",
	"XK_LATIN1",
	"XK_LATIN2",
	"XK_LATIN3",
	"XK_LATIN4",
	"XK_LATIN8",
	"XK_LATIN9",
	"XK_CAUCASUS",
	"XK_GREEK",
	"XK_KATAKANA",
	"XK_ARABIC",
	"XK_CYRILLIC",
	"XK_HEBREW",
	"XK_THAI",
	"XK_KOREAN",
	"XK_ARMENIAN",
	"XK_GEORGIAN",
	"XK_VIETNAMESE",
	"XK_CURRENCY",
	"XK_MATHEMATICAL",
	"XK_BRAILLE",
	"XK_SINHALA",
}

func definedSection(section string) bool {
	for _, value := range sections {
		if value == section {
			return true
		}
	}
	return false
}

func getComment(unicode_, comment string) string {
	if unicode_ != "" {
		return fmt.Sprintf("U+%s %s", unicode_, comment)
	}
	return comment
}

func main() {
	fmt.Println("package", goPkgName)
	fmt.Printf("import x %q\n", "github.com/linuxdeepin/go-x11-client")

	log.SetFlags(log.Lshortfile)

	fh, err := os.Open(keysymdefFile)
	if err != nil {
		log.Fatal(err)
	}

	var currentSection string
	var currentSectionOk bool = true
	scanner := bufio.NewScanner(fh)

	// stat vars
	var (
		def0Count     int
		def1Count     int
		def2Count     int
		defOtherCount int
	)

	var symbols []Keysym
	for scanner.Scan() {
		line := scanner.Text()
		result := reIfDef.FindStringSubmatch(line)
		if len(result) > 0 {
			currentSection = result[1]
			currentSectionOk = definedSection(currentSection)
			continue
		}

		if strings.HasPrefix(line, "#endif") {
			currentSection = ""
			currentSectionOk = false
			continue
		}

		// ignore some section
		if !currentSectionOk {
			log.Println("ignore ", line)
			continue
		}

		result = reDef0.FindStringSubmatch(line)
		if len(result) > 0 {
			def0Count++
			log.Printf("def0 result: %#v\n", result[1:])
			symbols = append(symbols, Keysym{
				Name:    result[1],
				Value:   result[2],
				Unicode: result[3],
				Comment: getComment(result[3], result[4]),
			})
			continue
		}

		result = reDef1.FindStringSubmatch(line)
		if len(result) > 0 {
			def1Count++
			log.Printf("def1 result: %#v\n", result[1:])
			symbols = append(symbols, Keysym{
				Name:    result[1],
				Value:   result[2],
				Unicode: result[3],
				Comment: getComment(result[3], result[4]),
			})
			continue
		}

		result = reDef2.FindStringSubmatch(line)
		if len(result) > 0 {
			def2Count++
			log.Printf("def2 result: %#v\n", result[1:])
			// no unicode
			symbols = append(symbols, Keysym{
				Name:    result[1],
				Value:   result[2],
				Comment: result[4],
			})
			continue
		}

		if strings.HasPrefix(line, "#define") {
			defOtherCount++
			log.Println("warn line: ", line)
			continue
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("stat:\ndef0: %d\ndef1: %d\ndef2: %d\ndef other: %d\n",
		def0Count, def1Count, def2Count, defOtherCount)

	// define constants
	fmt.Println("const (")
	for _, sym := range symbols {
		if sym.Comment != "" {
			fmt.Println("// ", sym.Comment)
		}
		fmt.Printf("%s = 0x%s\n", sym.Name, sym.Value)
	}
	fmt.Println(")")

	// map keysym => english str
	var tempMap = make(map[string]string, len(symbols))
	fmt.Println("var KeysymEngMap = map[x.Keysym]string{")
	for _, sym := range symbols {
		dupSym, ok := tempMap[sym.Value]
		if ok {
			fmt.Printf("// %s == %s # %s\n", sym.Name, dupSym, sym.Comment)
		} else {
			tempMap[sym.Value] = sym.Name
			fmt.Printf("%s:%q,\n", sym.Name, getEnglish(sym.Name))
		}
	}
	fmt.Println("}")

	// map english str => keysym
	fmt.Println("var EngKeysymMap = map[string]x.Keysym{")
	for _, sym := range symbols {
		fmt.Printf("%q:%s,\n", getEnglish(sym.Name), sym.Name)
	}
	fmt.Println("}")
}

func getEnglish(symName string) string {
	if strings.HasPrefix(symName, "XK_") {
		return symName[len("XK_"):]
	} else {
		panic("invalid symbol name:" + symName)
	}
}
