package ewmh

import (
	"errors"
	"fmt"

	x "github.com/linuxdeepin/go-x11-client"
)

const LENGTH_MAX = 0xFFFF

func getUTF8StrFromReply(reply *x.GetPropertyReply) (string, error) {
	if reply.Format != 8 {
		return "", errors.New("bad reply")
	}

	return string(reply.Value), nil
}

func getUTF8StrsFromReply(reply *x.GetPropertyReply) ([]string, error) {
	if reply.Format != 8 {
		return nil, errors.New("bad reply")
	}

	data := reply.Value
	var strs []string
	sstart := 0
	for i, c := range data {
		if c == 0 {
			strs = append(strs, string(data[sstart:i]))
			sstart = i + 1
		}
	}
	if sstart < len(data) {
		strs = append(strs, string(data[sstart:]))
	}
	return strs, nil
}

func getWmIconFromReply(reply *x.GetPropertyReply) ([]WmIcon, error) {
	if reply.Format != 32 {
		return nil, errors.New("bad reply")
	}
	return getIcons(reply.Value)
}

type WmIcon struct {
	Width, Height uint32
	Data          []byte
}

func (icon *WmIcon) String() string {
	if icon == nil {
		return "nil"
	}

	return fmt.Sprintf("ewmh.WmIcon{ size: %dx%d, dataLength: %d }", icon.Width, icon.Height, len(icon.Data))
}

func getIcons(p []byte) ([]WmIcon, error) {
	var icons []WmIcon
	for len(p) >= 2*4 {
		width := x.Get32(p)
		height := x.Get32(p[4:])
		area := width * height

		if len(p) >= int(2+area)*4 {
			if area > 0 {
				icon := WmIcon{
					Width:  width,
					Height: height,
					Data:   p[2*4 : (2+area)*4],
				}
				icons = append(icons, icon)
			}

			if len(p) > int(2+area)*4 {
				// next icon
				p = p[(2+area)*4:]
				continue
			} else {
				// end
				break
			}
		} else {
			return nil, errors.New("bad icon data")
		}
	}
	return icons, nil
}
