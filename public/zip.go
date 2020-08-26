package public

import (
	"github.com/yeka/zip"
	"io"
	"os"
)

// 不加密压缩
func ZIP(srcpath string, destpath string) error {
	src, err := os.Open(srcpath)
	if err != nil {
		return err
	}
	defer src.Close()
	dest, err := os.Create(destpath)
	if err != nil {
		return err
	}
	defer dest.Close()
	// 包装dest
	zipWriter := zip.NewWriter(dest)
	defer zipWriter.Close()
	ioWriter, err := zipWriter.Create(srcpath)
	if err != nil {
		return err
	}
	_, err = io.Copy(ioWriter, src)
	if err != nil {
		return err
	}
	zipWriter.Flush()
	return nil
}

// 加密压缩
func ZIPEncrypt(srcpath string, destpath string, key string) error {
	src, err := os.Open(srcpath)
	if err != nil {
		return err
	}
	defer src.Close()
	dest, err := os.Create(destpath)
	if err != nil {
		return err
	}
	defer dest.Close()
	// 包装dest
	zipWriter := zip.NewWriter(dest)
	defer zipWriter.Close()
	ioWriter, err := zipWriter.Encrypt(srcpath, key, zip.StandardEncryption)
	if err != nil {
		return err
	}
	_, err = io.Copy(ioWriter, src)
	if err != nil {
		return err
	}
	zipWriter.Flush()
	return nil
}
