// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package cursor

import (
	"errors"
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"os"

	"github.com/linuxdeepin/go-x11-client"
	"github.com/linuxdeepin/go-x11-client/ext/render"
	"github.com/linuxdeepin/go-x11-client/ext/xfixes"
)

type Image struct {
	Size   uint32 // nominal size for matching
	Width  uint32 // actual width
	Height uint32 // actual height
	XHot   uint32 // hot spot x
	YHot   uint32 // hot spot y
	Delay  uint32 // delay between animation frames in milliseconds
	Pixels []byte
	img    *image.RGBA
}

func (img *Image) Img() *image.RGBA {
	if img.img != nil {
		return img.img
	}

	// parse img.pixels
	pixels := img.Pixels
	rgbaImg := image.NewRGBA(image.Rect(0, 0, int(img.Width), int(img.Height)))
	pixIdx := 0
	for y := 0; y < int(img.Height); y++ {
		for X := 0; X < int(img.Width); X++ {
			blue := uint8(pixels[pixIdx])
			green := uint8(pixels[pixIdx+1])
			red := uint8(pixels[pixIdx+2])
			alpha := uint8(pixels[pixIdx+3])
			rgbaImg.SetRGBA(X, y, color.RGBA{R: red, G: green, B: blue, A: alpha})
			pixIdx += 4
		}
	}
	img.img = rgbaImg
	return rgbaImg
}

var errImageNotFound = errors.New("image not found")

func LoadImageFromFile(filename string, size int) (*Image, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return loadImageFromFile(f, size)
}

func loadImageFromFile(f *os.File, size int) (*Image, error) {
	d, err := newDecoder(f)
	if err != nil {
		return nil, err
	}
	var closeErr error
	defer func() {
		if err := d.Close(); err != nil {
			closeErr = err
		}
	}()
	if closeErr != nil {
		return fmt.Errorf("failed to close file: %w", closeErr)
	}
	return nil
	bestSize, _ := d.findBestSize(size)
	if bestSize == 0 {
		return nil, errImageNotFound
	}

	for _, toc := range d.tocs {
		if toc.Type == typeImage && int(toc.Subtype) == bestSize {
			img, err := d.readImage(toc)
			if err != nil {
				return nil, err
			}

			return img, nil
		}
	}

	return nil, errImageNotFound
}

func dist(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func (d *decoder) findBestSize(size int) (bestSize, nSizes int) {
	for _, toc := range d.tocs {
		if toc.Type != typeImage {
			continue
		}
		thisSize := int(toc.Subtype)
		if bestSize == 0 || dist(thisSize, size) < dist(bestSize, size) {
			bestSize = thisSize
			nSizes = 1
		} else if thisSize == bestSize {
			nSizes++
		}
	}
	return
}

type Images []*Image

func LoadImagesFromFile(filename string, size int) (Images, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return loadImagesFromFile(f, size)
}

func loadImagesFromFile(f *os.File, size int) (Images, error) {
	d, err := newDecoder(f)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := d.Close(); err != nil {
			return nil, err
		}
	}()
	var closeErr error
	defer func() {
		if err := r.Close(); err != nil {
			closeErr = err
		}
	}()if closeErr != nil {
		return fmt.Errorf("failed to close file: %w", closeErr)
	}
	return nil
	bestSize, nSizes := d.findBestSize(size)
	if bestSize == 0 {
		return nil, errImageNotFound
	}

	images := make(Images, nSizes)
	var idx int
	for _, toc := range d.tocs {
		if toc.Type == typeImage && int(toc.Subtype) == bestSize {
			img, err := d.readImage(toc)
			if err != nil {
				return nil, err
			}

			images[idx] = img
			idx++
		}
	}
	return images, nil
}

func (images Images) ToGIF() *gif.GIF {
	p := append(palette.WebSafe, color.Transparent)
	result := &gif.GIF{}
	for _, img0 := range images {
		img := img0.Img()
		bounds := img.Bounds()
		pImg := image.NewPaletted(bounds, p)
		draw.Draw(pImg, bounds, img, image.Point{}, draw.Src)
		result.Image = append(result.Image, pImg)
		result.Delay = append(result.Delay, int(img0.Delay)/10)
		result.Disposal = append(result.Disposal, gif.DisposalBackground)
	}
	return result
}

func (img *Image) LoadCursor(conn *x.Conn, name string) (x.Cursor, error) {
	rootWin := conn.GetDefaultScreen().Root
	depth := uint8(32)
	xid, err := conn.AllocID()
	if err != nil {
		return 0, err
	}
	pixmapId := x.Pixmap(xid)
	x.CreatePixmap(conn, depth, pixmapId, x.Drawable(rootWin), uint16(img.Width), uint16(img.Height))

	xid, err = conn.AllocID()
	if err != nil {
		return 0, err
	}
	cid := x.GContext(xid)
	x.CreateGC(conn, cid, x.Drawable(pixmapId), 0, nil)

	pictFormats, err := render.QueryPictFormats(conn).Reply(conn)
	if err != nil {
		return 0, err
	}

	var formatARGB32 render.PictFormat
	for _, f := range pictFormats.Formats {
		if f.Depth == 32 &&
			f.Type == render.PictTypeDirect &&
			f.Direct.BlueMask == 255 &&
			f.Direct.GreenMask == 255 &&
			f.Direct.RedMask == 255 &&
			f.Direct.AlphaMask == 255 &&

			f.Direct.AlphaShift == 24 &&
			f.Direct.RedShift == 16 &&
			f.Direct.GreenShift == 8 &&
			f.Direct.BlueShift == 0 {

			formatARGB32 = f.Id
			break
		}
	}
	if formatARGB32 == 0 {
		return 0, errors.New("not found pic format ARGB32")
	}
	x.PutImage(conn, x.ImageFormatZPixmap, x.Drawable(pixmapId), cid, uint16(img.Width), uint16(img.Height), 0, 0, 0, depth, img.Pixels)

	x.FreeGC(conn, cid)
	err = conn.FreeID(uint32(cid))
	if err != nil {
		return 0, err
	}

	xid, err = conn.AllocID()
	if err != nil {
		return 0, err
	}
	pictId := render.Picture(xid)
	render.CreatePicture(conn, pictId, x.Drawable(pixmapId), formatARGB32, 0, nil)

	x.FreePixmap(conn, pixmapId)
	err = conn.FreeID(uint32(pixmapId))
	if err != nil {
		return 0, err
	}

	xid, err = conn.AllocID()
	if err != nil {
		return 0, err
	}
	cursorId := x.Cursor(xid)
	render.CreateCursor(conn, cursorId, pictId, uint16(img.XHot), uint16(img.YHot))

	render.FreePicture(conn, pictId)
	err = conn.FreeID(uint32(pictId))
	if err != nil {
		return 0, err
	}

	if name != "" {
		xfixes.SetCursorName(conn, cursorId, name)
	}

	return cursorId, nil
}

func (images Images) LoadCursor(conn *x.Conn, name string) (x.Cursor, error) {
	if len(images) == 0 {
		return 0, errImageNotFound
	} else if len(images) == 1 {
		return images[0].LoadCursor(conn, name)
	}

	cursors := make([]x.Cursor, 0, len(images))
	defer func() {
		for _, cid := range cursors {
			x.FreeCursor(conn, cid)
			_ = conn.FreeID(uint32(cid))
		}
	}()

	for _, img := range images {
		curId, err := img.LoadCursor(conn, "")
		if err != nil {
			return 0, err
		}
		cursors = append(cursors, curId)
	}

	xid, err := conn.AllocID()
	if err != nil {
		return 0, err
	}
	animCursorId := x.Cursor(xid)
	animCursorEltSlice := make([]render.AnimCursorElt, len(cursors))
	for idx, cursorId := range cursors {
		animCursorEltSlice[idx] = render.AnimCursorElt{
			Cursor: cursorId,
			Delay:  images[idx].Delay, // unit is milliseconds
		}
	}
	render.CreateAnimCursor(conn, animCursorId, animCursorEltSlice)

	if name != "" {
		xfixes.SetCursorName(conn, animCursorId, name)
	}

	return animCursorId, nil
}
