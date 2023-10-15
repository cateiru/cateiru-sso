package lib_test

import (
	"image"
	"io"
	"os"
	"testing"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/stretchr/testify/require"
)

const IMAGE_PATH = "../test_sample_image.png"

func TestValidateImage(t *testing.T) {
	file, err := os.OpenFile(IMAGE_PATH, os.O_RDONLY, 0666)
	require.NoError(t, err)

	i, err := lib.ValidateImage(file, 100, 100)
	require.NoError(t, err)

	img, imgType, err := image.Decode(io.Reader(i))
	require.NoError(t, err)

	require.Equal(t, imgType, "png")

	bounds := img.Bounds()
	require.Equal(t, bounds.Max.X, 100)
	require.Equal(t, bounds.Max.Y, 100)
}
