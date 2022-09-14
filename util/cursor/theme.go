// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package cursor

import (
	"bufio"
	"bytes"
	"errors"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func LoadImage(theme, name string, size int) (*Image, error) {
	f := scanTheme(theme, name, 0)
	if f == nil {
		f = scanTheme("default", name, 0)
	}
	if f == nil {
		return nil, errors.New("not found image file")
	}

	return loadImageFromFile(f, size)
}

func LoadImages(theme, name string, size int) (Images, error) {
	f := scanTheme(theme, name, 0)
	if f == nil {
		f = scanTheme("default", name, 0)
	}
	if f == nil {
		return nil, errors.New("not found image file")
	}

	return loadImagesFromFile(f, size)
}

func getLibraryPaths() []string {
	xcursorPath, ok := os.LookupEnv("XCURSOR_PATH")
	if ok {
		return strings.Split(xcursorPath, ":")
	}

	var paths []string
	home := getHome()
	// add ~/.icons
	if home != "" {
		paths = append(paths, filepath.Join(home, ".icons"))
	}
	paths = append(paths, "/usr/share/icons", "/usr/share/pixmaps")
	return paths
}

func getHome() string {
	home, ok := os.LookupEnv("HOME")
	if ok {
		return home
	}

	u, err := user.Current()
	if err != nil {
		return ""
	}
	return u.HomeDir
}

func scanTheme(theme, name string, depth int) *os.File {
	if depth > 50 {
		return nil
	}
	libPaths := getLibraryPaths()
	var inherits []string
loop:
	for _, libPath := range libPaths {
		imgFile := filepath.Join(libPath, theme, "cursors", name)
		f, err := os.Open(imgFile)
		if err == nil {
			return f
		}

		for _, base := range []string{"cursor.theme", "index.theme"} {
			themeFile := filepath.Join(libPath, theme, base)
			inherits, _ = getThemeInherits(themeFile)
			if len(inherits) > 0 {
				break loop
			}
		}
	}

	for _, i := range inherits {
		if i == theme {
			continue
		}
		f := scanTheme(i, name, depth+1)
		if f != nil {
			return f
		}
	}

	// not found
	return nil
}

func getThemeInherits(themeFile string) ([]string, error) {
	f, err := os.Open(themeFile)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fields := bytes.SplitN(scanner.Bytes(), []byte("="), 2)
		if len(fields) != 2 {
			continue
		}

		if "Inherits" == string(bytes.TrimSpace(fields[0])) {
			result := strings.Split(string(bytes.TrimSpace(fields[1])), ":")
			return result, nil
		}
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return nil, nil
}
