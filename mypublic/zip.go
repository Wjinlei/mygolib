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
	"sync"

	"github.com/Wjinlei/mygolib/myencode"
	"github.com/yeka/zip"
)

var wg sync.WaitGroup

// TGZ tar.gz压缩
func TGZ(srcpath, dstpath string) error {
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
	srcinfo, err := os.Lstat(srcpath)
	//srcinfo, err := srcfile.Stat()
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

				// Handle path
				header.Name = strings.TrimPrefix(path, fmt.Sprintf("%s/", srcpath))

				if pathinfo.IsDir() {
					header.Name = fmt.Sprintf("%s/", header.Name)
					err = tw.WriteHeader(header)
					if err != nil {
						return err
					}
				} else {
					// Open the file which will be written into the archive
					pathfile, err := os.Open(path)
					if err != nil {
						return err
					}
					defer pathfile.Close()

					symlink, err := os.Readlink(path)
					if err != nil {
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
					} else {
						// Set symlink name
						header.Linkname = symlink
						// Write file header to the tar archive
						err = tw.WriteHeader(header)
						if err != nil {
							return err
						}
						// Write ""
						_, err = tw.Write([]byte(""))
						if err != nil {
							return err
						}
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
		//header.Name = srcpath

		symlink, err := os.Readlink(srcpath)
		if err != nil {
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
		} else {
			// Set symlink name
			header.Linkname = symlink
			// Write file header to the tar archive
			err = tw.WriteHeader(header)
			if err != nil {
				return err
			}
			// Write ""
			_, err = tw.Write([]byte(""))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// ZIP zip压缩,支持设置文件名编码,Windows一般设置为GBK
func ZIP(srcpath, dstpath, encoding string) error {
	// encoder
	encoder := myencode.GetEncoder(encoding)
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
	srcinfo, err := os.Lstat(srcpath)
	//srcinfo, err := srcfile.Stat()
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

				// Handle path
				header.Name = strings.TrimPrefix(path, fmt.Sprintf("%s/", srcpath))
				header.Name = encoder.ConvertString(header.Name)
				if pathinfo.IsDir() {
					header.Name = fmt.Sprintf("%s/", header.Name)
					_, err := zw.CreateHeader(header)
					if err != nil {
						return err
					}
				} else {
					// Open the file which will be written into the archive
					pathfile, err := os.Open(path)
					if err != nil {
						return err
					}
					defer pathfile.Close()

					// Readlink
					symlink, err := os.Readlink(path)
					if err != nil {
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
					} else {
						header.SetMode(os.ModeSymlink)
						// Create file Header
						pathWriter, err := zw.CreateHeader(header)
						if err != nil {
							return err
						}
						_, err = pathWriter.Write([]byte(symlink))
						if err != nil {
							return err
						}
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

		// Use full path as name (FileInfoHeader only takes the basename)
		// If we don't do this the directory strucuture would
		// not be preserved
		// https://golang.org/src/archive/tar/common.go?#L626
		//header.Name = encoder.ConvertString(header.Name)

		// Readlink
		symlink, err := os.Readlink(srcpath)
		if err != nil {
			// Create file Header
			pathWriter, err := zw.CreateHeader(header)
			if err != nil {
				return err
			}
			// Copy file content to tar archive
			_, err = io.Copy(pathWriter, srcfile)
			if err != nil {
				return err
			}
		} else {
			header.SetMode(os.ModeSymlink)
			// Create file Header
			pathWriter, err := zw.CreateHeader(header)
			if err != nil {
				return err
			}
			_, err = pathWriter.Write([]byte(symlink))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// ZIPDecrypt 解压缩
func ZIPDecrypt(srcpath, destpath, password, charset string) error {
	// 编码器
	encoder := myencode.GetEncoder(charset)
	if encoder == nil {
		return fmt.Errorf("Charset error: [%s]", charset)
	}
	// 解码器
	decoder := myencode.GetDecoder(charset)
	if decoder == nil {
		return fmt.Errorf("Charset error: [%s]", charset)
	}
	password = encoder.ConvertString(password)

	// 读取源文件
	readCloser, err := zip.OpenReader(srcpath)
	if err != nil {
		fmt.Println("readCloser err:", err)
		return err
	}
	defer readCloser.Close()

	// 解压路径
	destpath = GetPath(destpath)

	for _, file := range readCloser.File {
		// 创建目录
		filepath := fmt.Sprintf("%s/%s",
			destpath, decoder.ConvertString(file.Name))
		if file.FileInfo().IsDir() {
			if err := MakeDir(filepath); err != nil {
				return err
			}
			continue
		}

		//  设置解压密码
		if file.IsEncrypted() {
			file.SetPassword(password)
		}

		// 打开原文件
		fmt.Printf("file.Open: %s\n", file.Name)
		src, err := file.Open()
		if err != nil {
			return err
		}

		// 创建目标文件
		fmt.Printf("os.Create: %s\n", filepath)
		dst, err := os.Create(filepath)
		if err != nil {
			src.Close()
			return err
		}

		// 写入数据
		fmt.Printf("io.Copy: %s, %s\n", dst.Name(), file.Name)
		_, err = io.Copy(dst, src)
		if err != nil {
			src.Close()
			dst.Close()
			return err
		}
		src.Close()
		dst.Close()
	}
	return nil
}

// TGZDecrypt tar.gz解压缩
func TGZDecrypt(srcpath, destpath, charset string) error {
	decoder := myencode.GetDecoder(charset)
	if decoder == nil {
		return fmt.Errorf("Charset error: [%s]", charset)
	}

	srcfile, err := os.Open(srcpath)
	if err != nil {
		return err
	}
	defer srcfile.Close()

	gr, err := gzip.NewReader(srcfile)
	if err != nil {
		return err
	}
	defer gr.Close()
	tr := tar.NewReader(gr)

	destpath = strings.TrimRight(destpath, "/")

	for {
		file, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}

		// 获取目标路径
		filepath := destpath + "/" + decoder.ConvertString(file.Name)

		// 创建目录
		if file.FileInfo().IsDir() {
			if err := MakeDir(filepath); err != nil {
				return err
			}
			continue
		}

		// 创建目标文件
		destfile, err := os.Create(filepath)
		if err != nil {
			return err
		}

		// 写入数据
		_, err = io.Copy(destfile, tr)
		if err != nil {
			destfile.Close()
			return err
		}
		destfile.Close()
	}
	return nil
}
