package main

//#include <stdio.h>
import "C"

import (
	"github.com/kalogs-c/turbosizenator/convert"
	"github.com/kalogs-c/turbosizenator/img"
)

//export ResizeImageFromFilePath
func ResizeImageFromFilePath(fpath *C.char, destination *C.char, target *C.char) {
	targetFormat := img.GetImageFormat(C.GoString(target))
	imagePath := C.GoString(fpath)
	destinationPath := C.GoString(destination)

	convert.FromFilePath(imagePath, destinationPath, targetFormat)
}

func main() {}
