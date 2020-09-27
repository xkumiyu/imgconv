package main

import (
	"errors"
	"flag"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
	"path/filepath"
)

var (
	inFile  string
	outFile string
)

func init() {
	flag.StringVar(&inFile, "i", "", "path of input image file")
	flag.StringVar(&outFile, "o", "", "path of output image file")
}

func main() {
	parse()
	outFormat, err := ext2format(filepath.Ext(outFile))
	if err != nil {
		// Consider error output
		log.Fatal(err)
	}
	convert(inFile, outFile, outFormat)
}

func parse() {
	flag.Parse()
	// TODO: check required options
}

func ext2format(ext string) (string, error) {
	switch ext {
	case ".png":
		return "png", nil
	case ".jpg", ".jpeg":
		return "jpeg", nil
	case ".gif":
		return "gif", nil
	default:
		return "", errors.New("error: invalid file extension")
	}
}

func convert(src string, dst string, format string) {
	sf, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer sf.Close()

	img, _, err := image.Decode(sf)

	df, err := os.Create(dst)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := df.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := encode(df, img, format); err != nil {
		// TODO: delete file
		log.Fatal(err)
	}
}

func encode(w io.Writer, img image.Image, format string) error {
	var err error

	switch format {
	case "png":
		err = png.Encode(w, img)
	case "jpeg":
		opts := jpeg.Options{Quality: 100}
		err = jpeg.Encode(w, img, &opts)
	case "gif":
		opts := gif.Options{}
		err = gif.Encode(w, img, &opts)
	default:
		err = errors.New("error: invalid format")
	}

	return err
}
