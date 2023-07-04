package convert

import (
	"os"

	"github.com/kalogs-c/turbosizenator/img"
)

func FromFilePath(fpath string, destination string, target img.ImageFormat) error {
	f, err := os.Open(fpath)
	if err != nil {
		return err
	}
	defer f.Close()

	image, err := img.ReadImage(f)
	if err != nil {
		return err
	}

	writer, err := os.Create(destination)
	if err != nil {
		return err
	}

	err = img.WriteImage(image, writer, target)
	if err != nil {
		return err
	}

	return nil
}
