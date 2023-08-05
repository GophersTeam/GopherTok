package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func FileSha256(file multipart.File) string {
	_sha256 := sha256.New()
	io.Copy(_sha256, file)
	return hex.EncodeToString(_sha256.Sum(nil))
}
func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()

	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}

	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}

	return nil
}
