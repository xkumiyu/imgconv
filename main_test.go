package main

import (
	"image"
	"image/color"
	"io/ioutil"
	"os"
	"testing"
)

func TestExt2format(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name   string
		ext    string
		expect string
	}{
		{name: "png", ext: ".png", expect: "png"},
		{name: "jpg", ext: ".jpg", expect: "jpeg"},
		{name: "jpeg", ext: ".jpeg", expect: "jpeg"},
		{name: "gif", ext: ".gif", expect: "gif"},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			actual, err := ext2format(c.ext)

			if err != nil {
				t.Errorf(`err="%s"`, err)
			}
			if actual != c.expect {
				t.Errorf(`expect="%s" actual="%s"`, c.expect, actual)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name   string
		format string
	}{
		{name: "png", format: "png"},
		{name: "jpeg", format: "jpeg"},
		{name: "gif", format: "gif"},
	}

	tmpFile, _ := ioutil.TempFile("", "tmp")
	defer os.Remove(tmpFile.Name())
	img := makeImage()

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			err := encode(tmpFile, img, c.format)
			if err != nil {
				t.Errorf(`err="%s"`, err)
			}
		})
	}
}

func makeImage() image.Image {
	width := 10
	height := 10
	color := color.RGBA{255, 255, 255, 255}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			img.Set(j, i, color)
		}
	}

	return img
}
