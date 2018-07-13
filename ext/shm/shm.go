package shm

import "github.com/linuxdeepin/go-x11-client"

type BadSegError struct {
	x.ResourceErrorBase
}

func (e BadSegError) Error() string {
	return "shm.BadSegError" + e.Msg()
}

func readBadSegError(r *x.Reader) x.Error {
	return BadSegError{x.ReadResourceErrorBase(r)}
}

type CompletionEvent struct {
}

func readCompletionEvent(r *x.Reader, v *CompletionEvent) error {
	return nil
}

// #WREQ
func writeQueryVersion(w *x.Writer) {
	w.WritePad(4)
}

type QueryVersionReply struct {
	SharedPixmaps bool
	MajorVersion  uint16
	MinorVersion  uint16
	Uid           uint16
	Gid           uint16
	PixmapFormat  uint8
}

func readQueryVersionReply(r *x.Reader, v *QueryVersionReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.SharedPixmaps = x.Uint8ToBool(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MajorVersion = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MinorVersion = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Uid = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Gid = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.PixmapFormat = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(15)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func writeAttach(w *x.Writer, shmSeg Seg, shmId uint32, readOnly bool) {
	w.WritePad(4)
	w.Write4b(uint32(shmSeg))
	w.Write4b(shmId)
	w.Write1b(x.BoolToUint8(readOnly))
	w.WritePad(3)
}

// #WREQ
func writeDetach(w *x.Writer, shmSeg Seg) {
	w.WritePad(4)
	w.Write4b(uint32(shmSeg))
}

// #WREQ
func writePutImage(w *x.Writer, drawable x.Drawable, gc x.GContext, totalWidth,
	totalHeight, srcX, srcY, srcWidth, srcHeight uint16, dstX, dstY int16,
	depth, format uint8, sendEvent bool, shmSeg Seg, offset uint32) {

	w.WritePad(4)
	w.Write4b(uint32(drawable))
	w.Write4b(uint32(gc))
	w.Write2b(totalWidth)
	w.Write2b(totalHeight)
	w.Write2b(srcX)
	w.Write2b(srcY)
	w.Write2b(srcWidth)
	w.Write2b(srcHeight)
	w.Write2b(uint16(dstX))
	w.Write2b(uint16(dstY))

	w.Write1b(depth)
	w.Write1b(format)
	w.Write1b(x.BoolToUint8(sendEvent))
	w.WritePad(1)

	w.Write4b(uint32(shmSeg))
	w.Write4b(offset)
}

// #WREQ
func writeGetImage(w *x.Writer, drawable x.Drawable, X, y int16, width, height uint16,
	planeMask uint32, format uint8, shmSeg Seg, offset uint32) {
	w.WritePad(4)
	w.Write4b(uint32(drawable))
	w.Write2b(uint16(X))
	w.Write2b(uint16(y))
	w.Write2b(width)
	w.Write2b(height)
	w.Write4b(planeMask)
	w.Write1b(format)
	w.WritePad(3)
	w.Write4b(uint32(shmSeg))
	w.Write4b(offset)
}

type GetImageReply struct {
	Depth  uint8
	Visual x.VisualID
	Size   uint32
}

func readGetImageReply(r *x.Reader, v *GetImageReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Depth = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Visual = x.VisualID(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Size = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func writeCreatePixmap(w *x.Writer, pid x.Pixmap, drawable x.Drawable,
	width, height uint16, depth uint8, shmSeg Seg, offset uint32) {
	w.WritePad(4)
	w.Write4b(uint32(pid))
	w.Write4b(uint32(drawable))
	w.Write2b(width)
	w.Write2b(height)

	w.Write1b(depth)
	w.WritePad(3)

	w.Write4b(uint32(shmSeg))
	w.Write4b(offset)
}
