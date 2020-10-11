package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
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
	if err := parse(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	outFormat, err := ext2format(filepath.Ext(outFile))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	if err := convert(inFile, outFile, outFormat); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}

func parse() error {
	flag.Parse()

	if inFile == "" {
		return errors.New("'-i' is required")
	}
	if outFile == "" {
		return errors.New("'-o' is required")
	}
	return nil
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
		return "", fmt.Errorf("invalid file extension: %s", ext)
	}
}

func convert(src string, dst string, format string) error {
	sf, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sf.Close()

	img, _, err := image.Decode(sf)
	if err != nil {
		return err
	}

	df, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer df.Close()

	if err := encode(df, img, format); err != nil {
		os.Remove(dst)
		return err
	}

	return nil
}

func encode(w io.Writer, img image.Image, format string) error {
	switch format {
	case "png":
		return png.Encode(w, img)
	case "jpeg":
		opts := jpeg.Options{Quality: 100}
		return jpeg.Encode(w, img, &opts)
	case "gif":
		opts := gif.Options{}
		return gif.Encode(w, img, &opts)
	default:
		return fmt.Errorf("invalid format: %s", format)
	}
}
