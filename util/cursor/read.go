package cursor

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
	"os"
)

const (
	typeImage   = 0xfffd0002
)

type decoder struct {
	f       *os.File
	br      *bufio.Reader
	version uint32
	tocs    []tocEntry
	tmp     [36]byte
}

func newDecoder(f *os.File) (*decoder, error) {
	var d decoder
	d.f = f
	d.br = bufio.NewReader(f)
	err := d.readHeader()
	if err != nil {
		return nil, err
	}

	return &d, nil
}

func (d *decoder) close() error {
	return d.f.Close()
}

// table of contents entry
type tocEntry struct {
	Type     uint32 // entry type
	Subtype  uint32 //
	Position uint32 // absolute byte position of table in file
}

func readFull(r io.Reader, b []byte) error {
	_, err := io.ReadFull(r, b)
	if err == io.EOF {
		err = io.ErrUnexpectedEOF
	}
	return err
}

func (d *decoder) readHeader() error {
	err := readFull(d.br, d.tmp[:16])
	if err != nil {
		return err
	}

	magic := string(d.tmp[:4])
	if magic != "Xcur" {
		return errors.New("magic not match")
	}

	d.version = binary.LittleEndian.Uint32(d.tmp[8:12])
	nToc := binary.LittleEndian.Uint32(d.tmp[12:16])
	d.tocs = make([]tocEntry, nToc)
	for i := range d.tocs {
		d.tocs[i], err = d.readTocEntry()
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *decoder) readTocEntry() (tocEntry, error) {
	err := readFull(d.br, d.tmp[:12])
	if err != nil {
		return tocEntry{}, err
	}

	var e tocEntry
	e.Type = binary.LittleEndian.Uint32(d.tmp[:4])
	e.Subtype = binary.LittleEndian.Uint32(d.tmp[4:8])
	e.Position = binary.LittleEndian.Uint32(d.tmp[8:12])
	return e, nil
}

func (d *decoder) setPos(pos uint32) error {
	_, err := d.f.Seek(int64(pos), io.SeekStart)
	if err != nil {
		return err
	}
	d.br.Reset(d.f)
	return nil
}

func (d *decoder) readImage(toc tocEntry) (*Image, error) {
	err := d.setPos(toc.Position)
	if err != nil {
		return nil, err
	}

	err = readFull(d.br, d.tmp[:36])
	if err != nil {
		return nil, err
	}

	// chuck header
	headerLen := binary.LittleEndian.Uint32(d.tmp[:4])
	if headerLen != 36 {
		return nil, errors.New("image type chunk head len is not 36")
	}

	type0 := binary.LittleEndian.Uint32(d.tmp[4:8])
	if type0 != typeImage {
		return nil, errors.New("chunk type not match")
	}

	subtype := binary.LittleEndian.Uint32(d.tmp[8:12])

	version := binary.LittleEndian.Uint32(d.tmp[12:16])
	if version != 1 {
		return nil, errors.New("version not supported")
	}

	var img Image
	img.Size = subtype
	img.Width = binary.LittleEndian.Uint32(d.tmp[16:20])
	img.Height = binary.LittleEndian.Uint32(d.tmp[20:24])
	img.XHot = binary.LittleEndian.Uint32(d.tmp[24:28])
	if img.XHot > img.Width {
		img.XHot = img.Width
	}

	img.YHot = binary.LittleEndian.Uint32(d.tmp[28:32])
	if img.YHot > img.Height {
		img.YHot = img.Height
	}

	img.Delay = binary.LittleEndian.Uint32(d.tmp[32:36])

	// read pixels
	img.Pixels = make([]byte, 4*int(img.Width)*int(img.Height))
	err = readFull(d.br, img.Pixels)
	if err != nil {
		return nil, err
	}

	return &img, nil
}
