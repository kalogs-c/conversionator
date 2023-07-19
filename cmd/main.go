package main

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"

import (
	"bytes"
	"log"
	"unsafe"

	"github.com/kalogs-c/turbosizenator/convert"
	"github.com/kalogs-c/turbosizenator/img"
)

//export ConvertImageFromFilePath
func ConvertImageFromFilePath(fpath *C.char, destination *C.char, target *C.char) {
	targetFormat := img.GetImageFormat(C.GoString(target))
	defer C.free(unsafe.Pointer(target))

	imagePath := C.GoString(fpath)
	defer C.free(unsafe.Pointer(fpath))

	destinationPath := C.GoString(destination)
	defer C.free(unsafe.Pointer(destination))

	convert.FromFilePath(imagePath, destinationPath, targetFormat)
}

//export ConvertImageFromBytes
func ConvertImageFromBytes(
	imageb unsafe.Pointer,
	filesize C.int,
	target *C.char,
) *C.char {
	imageBytes := C.GoBytes(imageb, filesize)
	// defer C.free(imageb)

	targetFormat := img.GetImageFormat(C.GoString(target))
	// defer C.free(unsafe.Pointer(target))

	reader := bytes.NewReader(imageBytes)

	buf, err := convert.FromReader(reader, targetFormat)
	if err != nil {
		log.Fatalf("error while converting the image %s\n", err.Error())
		panic(err)
	}

	cbytes := C.CBytes(buf.Bytes())

	return (*C.char)(cbytes)
}

func main() {}
