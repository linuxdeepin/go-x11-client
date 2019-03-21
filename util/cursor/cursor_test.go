package cursor

import (
	"image/gif"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadImage(t *testing.T) {
	_, err := os.Stat("/usr/share/icons/deepin/cursors/left_ptr")
	if err == nil {
		_, err = LoadImage("deepin", "left_ptr", 24)
		assert.Nil(t, err)
	} else {
		t.Skip(err)
	}
}

func TestLoadImages(t *testing.T) {
	_, err := os.Stat("/usr/share/icons/deepin/cursors/watch")
	if err == nil {
		_, err = LoadImages("deepin", "watch", 24)
		assert.Nil(t, err)
	} else {
		t.Skip(err)
	}
}

func TestLoadImageFromFile(t *testing.T) {
	img, err := LoadImageFromFile("testdata/left_ptr", 24)
	assert.Nil(t, err)
	_ = img.Img()
	assert.EqualValues(t, 24, img.Size)
	assert.EqualValues(t, 24, img.Width)
	assert.EqualValues(t, 24, img.Height)
	assert.EqualValues(t, 7, img.XHot)
	assert.EqualValues(t, 4, img.YHot)
	assert.EqualValues(t, 50, img.Delay)
}

func TestLoadImagesFromFile(t *testing.T) {
	images, err := LoadImagesFromFile("testdata/watch", 24)
	assert.Nil(t, err)
	assert.Len(t, images, 16)

	for _, img := range images {
		_ = img.Img()
		assert.EqualValues(t, 24, img.Size)
		assert.EqualValues(t, 24, img.Width)
		assert.EqualValues(t, 24, img.Height)
		assert.EqualValues(t, 60, img.Delay)
	}
}

func BenchmarkLoadImageFromFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		img, err := LoadImageFromFile("testdata/left_ptr", 24)
		if err != nil {
			b.Fatal(err)
		}
		_ = img.Img()
	}
}

func BenchmarkLoadImagesFromFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		imgs, err := LoadImagesFromFile("testdata/watch", 24)
		if err != nil {
			b.Fatal(err)
		}
		for _, img := range imgs {
			_ = img.Img()
		}
	}
}

func TestImages_ToGIF(t *testing.T) {
	images, err := LoadImagesFromFile("testdata/watch", 24)
	assert.Nil(t, err)
	outGif := images.ToGIF()
	err = gif.EncodeAll(ioutil.Discard, outGif)
	assert.Nil(t, err)
}
