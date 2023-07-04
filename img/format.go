package img

import (
	"strings"
)

type ImageFormat string

const (
	Jpeg ImageFormat = "jpeg"
	Png  ImageFormat = "png"
	Gif  ImageFormat = "gif"
	Webp ImageFormat = "webp"
)

func IsImage(filename string) bool {
	format := GetImageFormat(filename[strings.LastIndex(filename, ".")+1:])

	return format != ""
}

func GetImageFormat(format string) ImageFormat {
	switch format {
	case "jpg":
		return Jpeg
	case "jpeg":
		return Jpeg
	case "png":
		return Png
	case "gif":
		return Gif
	case "webp":
		return Webp
	default:
		return ""
	}
}
