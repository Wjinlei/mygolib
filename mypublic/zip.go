package mypublic

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/axgle/mahonia"
	"github.com/yeka/zip"
)

// ZIP 不加密压缩
func ZIP(srcpath, destpath, encoding string) error {
	// 判断传入的源路径是否是目录
	ok := IsDir(srcpath)
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

// ZIPFile 压缩文件
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

// ZIPDir 压缩目录
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

// ZIPDecrypt 解压缩
func ZIPDecrypt(srcpath, destpath, password, charset string) error {
	encoder := mahonia.NewEncoder(charset)
	if encoder == nil {
		return fmt.Errorf("Charset error: [%s]", charset)
	}
	password = encoder.ConvertString(password)
	decoder := mahonia.NewDecoder(charset)
	if decoder == nil {
		return fmt.Errorf("Charset error: [%s]", charset)
	}
	readCloser, err := zip.OpenReader(srcpath)
	if err != nil {
		return err
	}
	defer readCloser.Close()
	destpath = strings.TrimRight(destpath, "/")
	for _, file := range readCloser.File {
		filepath := destpath + "/" + decoder.ConvertString(file.Name)
		if file.FileInfo().IsDir() {
			if err := MakeDir(filepath); err != nil {
				return err
			}
			continue
		}
		if file.IsEncrypted() {
			file.SetPassword(password)
		}
		src, err := file.Open()
		if err != nil {
			src.Close()
			return err
		}
		dest, err := os.Create(filepath)
		if err != nil {
			dest.Close()
			return err
		}
		_, err = io.Copy(dest, src)
		if err != nil {
			src.Close()
			dest.Close()
			return err
		}
		src.Close()
		dest.Close()
	}
	return nil
}

func TarGZ(srcpath, dstpath string) error {
	// Create output file
	dstfile, err := os.Create(dstpath)
	if err != nil {
		return err
	}
	defer dstfile.Close()

	// Create new Writers for gzip and tar
	// These writers are chained. Writing to the tar writer will
	// write to the gzip writer which in turn will write to
	// the "buf" writer
	gw := gzip.NewWriter(dstfile)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()

	// Open the file which will be written into the archive
	srcfile, err := os.Open(srcpath)
	if err != nil {
		return err
	}
	defer srcfile.Close()

	// Get FileInfo about our file providing file size, mode, etc.
	srcinfo, err := srcfile.Stat()
	if err != nil {
		return err
	}

	if srcinfo.IsDir() {
		filepath.Walk(srcpath, func(path string, pathinfo os.FileInfo, err error) error {
			if path != srcpath {
				// Create a tar Header from the FileInfo data
				header, err := tar.FileInfoHeader(pathinfo, pathinfo.Name())
				if err != nil {
					return err
				}

				header.Name = strings.TrimPrefix(path, fmt.Sprintf("%s/", srcpath))

				if pathinfo.IsDir() {
					header.Name = fmt.Sprintf("%s/", header.Name)
				} else {
					// Open the file which will be written into the archive
					pathfile, err := os.Open(path)
					if err != nil {
						return err
					}
					defer pathfile.Close()

					// Write file header to the tar archive
					err = tw.WriteHeader(header)
					if err != nil {
						return err
					}

					// Copy file content to tar archive
					_, err = io.Copy(tw, pathfile)
					if err != nil {
						return err
					}
				}
			}
			return nil
		})
	} else {
		// Create a tar Header from the FileInfo data
		header, err := tar.FileInfoHeader(srcinfo, srcinfo.Name())
		if err != nil {
			return err
		}

		// Use full path as name (FileInfoHeader only takes the basename)
		// If we don't do this the directory strucuture would
		// not be preserved
		// https://golang.org/src/archive/tar/common.go?#L626
		header.Name = srcpath

		// Write file header to the tar archive
		err = tw.WriteHeader(header)
		if err != nil {
			return err
		}

		// Copy file content to tar archive
		_, err = io.Copy(tw, srcfile)
		if err != nil {
			return err
		}
	}
	return nil
}
