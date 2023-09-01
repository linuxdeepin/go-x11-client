// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package x

/*
auth.go contains functions to facilitate the parsing of .Xauthority files.

It is largely unmodified from the original XGB package that I forked.
*/

import (
	"encoding/binary"
	"errors"
	"io"
	"log"
	"os"
)

// readAuthority reads the X authority file for the DISPLAY.
// If hostname == "" or hostname == "localhost",
// then use the system's hostname (as returned by os.Hostname) instead.
func readAuthority(hostname, display string) (
	name string, data []byte, err error) {

	// b is a scratch buffer to use and should be at least 256 bytes long
	// (i.e. it should be able to hold a hostname).
	b := make([]byte, 256)

	// As per /usr/include/X11/Xauth.h.
	const familyLocal = 256
	const familyWild = 65535

	if len(hostname) == 0 || hostname == "localhost" {
		hostname, err = os.Hostname()
		if err != nil {
			return
		}
	}

	fname := os.Getenv("XAUTHORITY")
	if len(fname) == 0 {
		home := os.Getenv("HOME")
		if len(home) == 0 {
			err = errors.New("Xauthority not found: $XAUTHORITY, $HOME not set")
			return
		}
		fname = home + "/.Xauthority"
	}

	r, err := os.Open(fname)
	if err != nil {
		return
	}
	defer func() {
		closeErr := r.Close()
		if err == nil {
			err = closeErr
		} else {
			log.Printf("error on close file %v: %v", r.Name(), closeErr)
		}
	}()

	for {
		var family uint16
		if err = binary.Read(r, binary.BigEndian, &family); err != nil {
			return
		}

		var addr string
		addr, err = getString(r, b)
		if err != nil {
			return
		}

		var disp string
		disp, err = getString(r, b)
		if err != nil {
			return
		}

		name, err = getString(r, b)
		if err != nil {
			return
		}

		data, err = getBytes(r, b)
		if err != nil {
			return
		}

		addrmatch := (family == familyWild) ||
			(family == familyLocal && addr == hostname)
		dispmatch := (disp == "") || (disp == display)

		if addrmatch && dispmatch {
			return
		}
	}
}

func getBytes(r io.Reader, b []byte) ([]byte, error) {
	var n uint16
	if err := binary.Read(r, binary.BigEndian, &n); err != nil {
		return nil, err
	} else if n > uint16(len(b)) {
		return nil, errors.New("bytes too long for buffer")
	}

	if _, err := io.ReadFull(r, b[0:n]); err != nil {
		return nil, err
	}
	return b[0:n], nil
}

func getString(r io.Reader, b []byte) (string, error) {
	b, err := getBytes(r, b)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
