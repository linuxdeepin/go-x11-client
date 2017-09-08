package client

import (
	"io/ioutil"
	"log"
	"net"
	"os"
	"sync"
)

var Logger *log.Logger

func init() {
	// setup logger
	logOut := ioutil.Discard
	if os.Getenv("DEBUG_XGB") == "1" {
		logOut = os.Stderr
	}
	Logger = log.New(logOut, "[xgb]", log.Lshortfile)
}

type Conn struct {
	conn net.Conn

	host          string
	display       string
	DisplayNumber int
	ScreenNumber  int
	setup         *Setup

	writeChan chan []byte

	iolock sync.Mutex
	in     *In
	out    *Out

	exts *exts
	xg   *xidGenerator

	EventHandlers   map[uint8]EventHandler
	EventHandlersMu sync.RWMutex
}

func (c *Conn) GetSetup() *Setup {
	return c.setup
}

func (c *Conn) GetDefaultScreen() *Screen {
	return &c.setup.Roots[c.ScreenNumber]
}

func (c *Conn) Close() {
	c.conn.Close()
	c.in.close()
}
