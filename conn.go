package x

import (
	"bufio"
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
	conn      net.Conn
	bufReader *bufio.Reader

	host          string
	display       string
	DisplayNumber int
	ScreenNumber  int
	setup         *Setup

	ioMu sync.Mutex
	in   *In
	out  *Out

	exts   *exts
	xidGen *xidGenerator
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

func (c *Conn) AddEventChan(eventChan chan<- GenericEvent) {
	c.in.addEventChan(eventChan)
}

func (c *Conn) RemoveEventChan(eventChan chan<- GenericEvent) {
	c.in.removeEventChan(eventChan)
}

func (c *Conn) AddErrorChan(errorChan chan<- Error) {
	c.in.addErrorChan(errorChan)
}

func (c *Conn) RemoveErrorChan(errorChan chan<- Error) {
	c.in.removeErrorChan(errorChan)
}
