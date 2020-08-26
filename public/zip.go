package public

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/axgle/mahonia"
	"github.com/yeka/zip"
)

// 不加密压缩
func ZIP(srcpath, destpath string) error {
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
	srcWriter := zip.NewWriter(dest)
	defer srcWriter.Close()
	destIoWriter, err := srcWriter.Create(srcpath)
	if err != nil {
		return err
	}
	_, err = io.Copy(destIoWriter, src)
	if err != nil {
		return err
	}
	srcWriter.Flush()
	return nil
}

// 加密压缩
func ZIPEncrypt(srcpath, destpath, password string) error {
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
	srcWriter := zip.NewWriter(dest)
	defer srcWriter.Close()
	destIoWriter, err := srcWriter.Encrypt(srcpath, password, zip.StandardEncryption)
	if err != nil {
		return err
	}
	_, err = io.Copy(destIoWriter, src)
	if err != nil {
		return err
	}
	srcWriter.Flush()
	return nil
}

// 解压缩
func ZIPDecrypt(srcpath, destpath, password, charset string) error {
	encoder := mahonia.NewEncoder(charset)
	if encoder == nil {
		return errors.New(fmt.Sprintf("Charset error: [%s]", charset))
	}
	password = encoder.ConvertString(password)
	decoder := mahonia.NewDecoder(charset)
	if decoder == nil {
		return errors.New(fmt.Sprintf("Charset error: [%s]", charset))
	}
	readCloser, err := zip.OpenReader(srcpath)
	if err != nil {
		return err
	}
	defer readCloser.Close()
	destpath = strings.TrimRight(destpath, "/")
	for _, file := range readCloser.File {
		if file.FileInfo().IsDir() {
			continue
		}
		if file.IsEncrypted() {
			file.SetPassword(password)
		}
		filepath := destpath + "/" + decoder.ConvertString(file.Name)
		if err := MakeDirAll(filepath); err != nil {
			return err
		}
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()
		dest, err := os.Create(filepath)
		if err != nil {
			return err
		}
		defer dest.Close()
		_, err = io.Copy(dest, src)
		if err != nil {
			return err
		}
	}
	return nil
}
