package public

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/axgle/mahonia"
	"github.com/yeka/zip"
)

// 不加密压缩
func ZIP(srcpath, destpath string) error {
	// 判断传入的源路径是否是目录
	ok, err := IsDir(srcpath)
	if err != nil {
		return err
	}
	if !ok {
		// 文件
		if err := ZIPFile(srcpath, destpath); err != nil {
			return err
		}
	} else {
		// 目录
		if err := ZIPDir(srcpath, destpath); err != nil {
			return err
		}
	}
	return nil
}

// 压缩文件
func ZIPFile(srcpath, destpath string) error {
	dest, err := os.Create(destpath)
	if err != nil {
		return err
	}
	defer dest.Close()
	srcWriter := zip.NewWriter(dest)
	defer srcWriter.Close()
	destIoWriter, err := srcWriter.Create(srcpath)
	if err != nil {
		return err
	}
	src, err := os.Open(srcpath)
	if err != nil {
		return err
	}
	defer src.Close()
	_, err = io.Copy(destIoWriter, src)
	if err != nil {
		return err
	}
	srcWriter.Flush()
	return nil
}

// 压缩目录
func ZIPDir(srcpath, destpath string) error {
	dest, err := os.Create(destpath)
	if err != nil {
		return err
	}
	defer dest.Close()
	srcWriter := zip.NewWriter(dest)
	defer srcWriter.Close()
	srcpath = strings.TrimRight(srcpath, "/")
	filepath.Walk(srcpath, func(path string, info os.FileInfo, err error) error {
		if path != srcpath {
			header, _ := zip.FileInfoHeader(info)
			header.Name = strings.TrimPrefix(path, srcpath+"/")
			if info.IsDir() {
				header.Name += "/"
			}
			destIoWriter, _ := srcWriter.CreateHeader(header)
			if !info.IsDir() {
				file, _ := os.Open(path)
				defer file.Close()
				io.Copy(destIoWriter, file)
			}
		}
		return nil
	})
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
