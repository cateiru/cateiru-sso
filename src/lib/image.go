package lib

import (
	"bytes"
	"image"
	"image/png"
	"io"

	exifremove "github.com/scottleedavis/go-exif-remove"
	"golang.org/x/image/draw"

	_ "image/gif"
	_ "image/jpeg"
)

// 画像のリサイズをExif削除を行う
func ValidateImage(data io.Reader, width, height int) (*bytes.Buffer, error) {
	img, _, err := image.Decode(data)
	if err != nil {
		return nil, err
	}

	rctSrc := img.Bounds()

	imgDst := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.CatmullRom.Scale(imgDst, imgDst.Bounds(), img, rctSrc, draw.Over, nil)

	buffer := &bytes.Buffer{}
	writer := io.Writer(buffer)

	err = png.Encode(writer, imgDst)
	if err != nil {
		return nil, err
	}

	// exif 削除
	noExifBytes, err := exifremove.Remove(buffer.Bytes())
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(noExifBytes), nil
}
