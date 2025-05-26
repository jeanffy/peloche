package domain

import (
	"image"
	"os"
	"peloche/domain/ports"
	"reflect"
	"runtime"
	"strings"

	"github.com/adrium/goheif"
	"github.com/nfnt/resize"
)

type Photo struct {
	Name   string
	Ext    string
	Path   string
	Width  int
	Height int
	Buffer image.Image
}

func NewPhoto(name string, ext string, filePath string) *Photo {
	return &Photo{
		Name:   name,
		Ext:    ext,
		Path:   filePath,
		Width:  0,
		Height: 0,
		Buffer: nil,
	}
}

func (x *Photo) LoadBuffer(log ports.LogPort) {
	if x.Buffer == nil {
		img := getDecodedPhoto(log, x.Path, x.Ext)
		x.Buffer = img
	}
}

func getDecodedPhoto(log ports.LogPort, filePath string, ext string) image.Image {
	reader, err := os.Open(filePath)
	if err != nil {
		log.Error(ports.LogPortErrorParams{
			Module: reflect.TypeOf(PhotoList{}).Name(),
			Error:  err,
			Msg:    filePath,
		})
		return nil
	}

	defer reader.Close()

	var imgDecoded image.Image = nil

	if strings.ToLower(ext) == ".heic" {
		img, err := goheif.Decode(reader)
		if err != nil {
			log.Error(ports.LogPortErrorParams{
				Module: reflect.TypeOf(PhotoList{}).Name(),
				Error:  err,
				Msg:    filePath,
			})
			return nil
		}
		imgDecoded = img
	} else {
		img, _, err := image.Decode(reader)
		if err != nil {
			log.Error(ports.LogPortErrorParams{
				Module: reflect.TypeOf(PhotoList{}).Name(),
				Error:  err,
				Msg:    filePath,
			})
			return nil
		}
		imgDecoded = img
	}

	resized := resize.Resize(500, 0, imgDecoded, resize.Lanczos3)

	imgDecoded = nil
	runtime.GC()

	return resized
}
