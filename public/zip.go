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
func ZIP(srcpath, destpath, encoding string) error {
	// 判断传入的源路径是否是目录
	ok, err := IsDir(srcpath)
	if err != nil {
		return err
	}
	if !ok {
		// 文件
		if err := ZIPFile(srcpath, destpath, encoding); err != nil {
			return err
		}
	} else {
		// 目录
		if err := ZIPDir(srcpath, destpath, encoding); err != nil {
			return err
		}
	}
	return nil
}

// 压缩文件
func ZIPFile(srcpath, destpath, encoding string) error {
	// 编码器
	encoder := mahonia.NewEncoder(encoding)
	if encoder == nil {
		return errors.New("encoding error")
	}
	// 设置文件名编码
	newPath := encoder.ConvertString(srcpath)
	// 创建目标
	dest, err := os.Create(destpath)
	if err != nil {
		return err
	}
	defer dest.Close()
	srcWriter := zip.NewWriter(dest)
	defer srcWriter.Close()
	// 使用编码后的文件名创建目标Writer
	destIoWriter, err := srcWriter.Create(newPath)
	if err != nil {
		return err
	}
	// 但打开文件还是要用源路径,否则无法找到文件
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
func ZIPDir(srcpath, destpath, encoding string) error {
	// 编码器
	encoder := mahonia.NewEncoder(encoding)
	if encoder == nil {
		return errors.New("encoding error")
	}
	// 创建目标
	dest, err := os.Create(destpath)
	if err != nil {
		return err
	}
	defer dest.Close()
	srcWriter := zip.NewWriter(dest)
	defer srcWriter.Close()
	// 遍历目录并压缩
	srcpath = strings.TrimRight(srcpath, "/")
	filepath.Walk(srcpath, func(path string, info os.FileInfo, err error) error {
		if path != srcpath {
			header, _ := zip.FileInfoHeader(info)
			header.Name = strings.TrimPrefix(path, srcpath+"/")
			if info.IsDir() {
				header.Name += "/"
			}
			// 设置文件名编码
			header.Name = encoder.ConvertString(header.Name)
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
