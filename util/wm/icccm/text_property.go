package icccm

import (
	"errors"

	x "github.com/linuxdeepin/go-x11-client"
	"golang.org/x/text/encoding/charmap"
)

func convertLatin1ToUTF8(p []byte) (string, error) {
	dec := charmap.ISO8859_1.NewDecoder()
	result, err := dec.Bytes(p)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

type GetTextCookie x.GetPropertyCookie

type TextProperty struct {
	Value    []byte
	Encoding x.Atom
	Format   uint8
}

func (tp *TextProperty) GetStr() (string, error) {
	if tp.Format != 8 || tp.Encoding != x.AtomString {
		return "", errors.New("unsupported encoding")
	}
	return convertLatin1ToUTF8(tp.Value)
}

func getTextProperty(c *x.Conn, window x.Window, property x.Atom) GetTextCookie {
	cookie := x.GetProperty(c, false, window, property, x.AtomAny, 0, getPropertyMaxLength)
	return GetTextCookie(cookie)
}

func (cookie GetTextCookie) Reply(c *x.Conn) (TextProperty, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c)
	if err != nil {
		return TextProperty{}, err
	}
	return getTextPropertyFromReply(reply)
}

func getTextPropertyFromReply(reply *x.GetPropertyReply) (TextProperty, error) {
	if reply.Type == x.None {
		return TextProperty{}, errors.New("not found property")
	}

	return TextProperty{
		Value:    reply.Value,
		Encoding: reply.Type,
		Format:   reply.Format,
	}, nil
}
