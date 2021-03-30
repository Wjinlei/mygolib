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

// TarGZ tar.gz压缩
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

// ZIP zip压缩,支持设置文件名编码,Windows一般设置为GBK
func ZIP(srcpath, dstpath, encoding string) error {
	// encoder
	encoder := mahonia.NewEncoder(encoding)
	if encoder == nil {
		return errors.New("encoding error")
	}

	// Create output file
	dstfile, err := os.Create(dstpath)
	if err != nil {
		return err
	}
	defer dstfile.Close()

	// Create new Writers for zip
	// These writers are chained. Writing to the tar writer will
	// write to the zip writer which in turn will write to
	// the "buf" writer
	zw := zip.NewWriter(dstfile)
	defer zw.Close()

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
				header, err := zip.FileInfoHeader(pathinfo)
				if err != nil {
					return err
				}
				// Set Header
				header.Name = strings.TrimPrefix(path, fmt.Sprintf("%s/", srcpath))
				header.Name = encoder.ConvertString(header.Name)
				if pathinfo.IsDir() {
					header.Name = fmt.Sprintf("%s/", header.Name)
				} else {
					// Open the file which will be written into the archive
					pathfile, err := os.Open(path)
					if err != nil {
						return err
					}
					defer pathfile.Close()
					// Create file Header
					pathWriter, err := zw.CreateHeader(header)
					if err != nil {
						return err
					}
					// Copy file content to tar archive
					_, err = io.Copy(pathWriter, pathfile)
					if err != nil {
						return err
					}
				}
			}
			return nil
		})
	} else {
		// Create a tar Header from the FileInfo data
		header, err := zip.FileInfoHeader(srcinfo)
		if err != nil {
			return err
		}
		// Set Header
		header.Name = encoder.ConvertString(srcpath)
		// Create file Header
		destWriter, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		// Copy file content to tar archive
		_, err = io.Copy(destWriter, srcfile)
		if err != nil {
			return err
		}
	}
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
