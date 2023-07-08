package convert

import (
	"bytes"
	"io"
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

func FromReaderAndSave(r io.Reader, destination string, target img.ImageFormat) error {
	image, err := img.ReadImage(r)
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

func FromReader(r io.Reader, target img.ImageFormat) (io.Reader, error) {
	image, err := img.ReadImage(r)
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer([]byte{})
	err = img.WriteImage(image, buffer, target)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(buffer.Bytes()), nil
}
