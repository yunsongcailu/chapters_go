package pkg_archive

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func DemoZip() {
	//zipFiles()
	unZipFiles()
}

// zip 压缩文件
func zipFiles() error {
	// 源文件目录
	src := "./pkg_archive/source"
	// 创建准备写入的文件
	dst := "./pkg_archive/source/source.zip"
	fw, err := os.Create(dst)
	defer fw.Close()
	if err != nil {
		return err
	}
	// zip writer
	zw := zip.NewWriter(fw)
	defer func() {
		if err := zw.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	err = filepath.Walk(src, func(path string, fi os.FileInfo, errBack error) (err error) {
		if errBack != nil {
			return errBack
		}
		// 通过源文件信息,创建目标zip的 文件信息
		fh, err := zip.FileInfoHeader(fi)
		if err != nil {
			return err
		}
		// 替换文件信息中的文件名
		fh.Name = strings.TrimPrefix(path, string(filepath.Separator)) // 去掉斜杠 '/'
		// 如果是个目录 则要加上尾斜杠'/'
		if fi.IsDir() {
			fh.Name += "/"
		}
		// 写入文件信息,并返回一个 write
		w, err := zw.CreateHeader(fh)
		if err != nil {
			return err
		}
		// 检测，如果不是标准文件就只写入头信息，不写入文件数据到 w
		// 如目录，也没有数据需要写
		if !fh.Mode().IsRegular() {
			return nil
		}
		// 打开要压缩的文件
		fr, err := os.Open(path)
		defer fr.Close()
		if err != nil {
			return err
		}
		// 将打开的文件 Copy 到 w
		n, err := io.Copy(w, fr)
		if err != nil {
			return err
		}
		// 输出压缩的内容
		fmt.Printf("成功压缩文件： %s, 共写入了 %d 个字符的数据\n", path, n)
		return nil
	})
	return err
}

// zip 解压缩
func unZipFiles() error {
	// 打开压缩文件，这个 zip 包有个方便的 ReadCloser 类型
	// 这个里面有个方便的 OpenReader 函数，可以比 tar 的时候省去一个打开文件的步骤
	// 目标路径
	src := "./pkg_archive/source/source.zip"
	zr, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer zr.Close()
	// 解压后保存目录
	dst := "./pkg_archive/source/unzip/"
	// 默认当前目录,如果解压后不是放在当前目录就按照保存目录去创建目录
	if dst != "" {
		if err = os.MkdirAll(dst, 0755); err != nil {
			return err
		}
	}
	// 遍历 zr, 将文件写入磁盘
	for _, file := range zr.File {
		path := filepath.Join(dst, file.Name) // 文件路径
		// 如果是目录则创建
		if file.FileInfo().IsDir() {
			if err = os.MkdirAll(path, file.Mode()); err != nil {
				return err
			}
			// 目录创建完成则 跳出当次循环
			continue
		}
		// 处理非目录的文件
		// 打开文件 获取reader
		fr, err := file.Open()
		if err != nil {
			return err
		}
		// 创建要写入的文件对应的writer
		fw, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, file.Mode())
		if err != nil {
			_ = fr.Close()
			return err
		}
		n, err := io.Copy(fw, fr)
		if err != nil {
			_ = fw.Close()
			_ = fr.Close()
			return err
		}
		// 输出解压结果
		// 将解压的结果输出
		fmt.Printf("成功解压 %s ，共写入了 %d 个字符的数据\n", path, n)
		_ = fw.Close()
		_ = fr.Close()

	} // 因为在循环中 不能defer close
	return nil
}
