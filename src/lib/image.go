package lib

import (
	"image"
	"io"

	"golang.org/x/image/draw"
)

// TODO
func ValidateImage(data io.Reader) error {
	imgSrc, _, err := image.Decode(data)
	if err != nil {
		return err
	}

	rctSrc := imgSrc.Bounds()

	imgDst := image.NewRGBA(image.Rect(0, 0, rctSrc.Dx()/4, rctSrc.Dy()/4))
	draw.CatmullRom.Scale(imgDst, imgDst.Bounds(), imgSrc, rctSrc, draw.Over, nil)

	return nil
}
