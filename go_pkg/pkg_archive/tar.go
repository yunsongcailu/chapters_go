package pkg_archive

import (
	"archive/tar"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
)

func DemoTar() {
	var keyword string
	for {
		goback := false
		fmt.Printf("please input keywords:\n"+
			"查看标题:%s,查看常量:%s,查看header结构体:%s\n"+
			"打包文件:%s,解包文件:%s,演示字符串打包解包:%s\n", "title", "const", "header", "tarfiles", "untar", "tarstring")
		fmt.Scanln(&keyword)
		switch keyword {
		case "title":
			showTitle()
		case "const":
			showConst()
		case "header":
			showHeaderStruct()
		case "tarfiles":
			err := tarDiskToFiles()
			if err != nil {
				fmt.Println(err)
			}
		case "untar":
			err := unTarFileToDisk()
			if err != nil {
				fmt.Println(err)
			}
		case "tarstring":
			err := tarStringToFiles()
			if err != nil {
				fmt.Println(err)
			}
		default:
			goback = true
		}
		if !goback {
			break
		}
	}

}

// 打包 指定路径下文件
func tarDiskToFiles() error {
	// 目标文件
	src := []string{"./pkg_archive/source/document.txt", "./pkg_archive/source/document2.txt", "./pkg_archive/source/document3.txt"}
	// 打包后文件
	dst := "./pkg_archive/source/source.tar"
	// 创建打包文件
	fw, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer fw.Close()

	// 创建 writer
	tw := tar.NewWriter(fw)
	defer func() {
		if err := tw.Close(); err != nil {
			fmt.Println("tw close error")
		}
	}()

	// 打包文件
	for _, fileName := range src {
		// 文件状态
		f, err := os.Stat(fileName)
		if err != nil {
			fmt.Printf("fileName:%s,stat error:%v\n", fileName, err)
			continue
		}
		// 写文件头信息
		hdr, err := tar.FileInfoHeader(f, "")
		// 写入tw
		if err := tw.WriteHeader(hdr); err != nil {
			return err
		}
		// 打开路径下文件 并把文件内容 写入 tar包中
		fs, err := os.Open(fileName)
		if err != nil {
			return err
		}
		if _, err := io.Copy(tw, fs); err != nil {
			return err
		}
		fs.Close()
	}
	return nil
}

// 解包
func unTarFileToDisk() error {
	srcFile := "./pkg_archive/source/source.tar"
	// 打开tar包
	fr, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer fr.Close()
	tr := tar.NewReader(fr)
	var resErr error
	for hdr, err := tr.Next(); err != io.EOF; hdr, err = tr.Next() {
		if err != nil {
			message := fmt.Sprintf("error:%s,tar reader file header:%s, error:%s\n", resErr.Error(), hdr.Name, err.Error())
			resErr = errors.New(message)
			continue
		}
		// 读取文件信息
		fi := hdr.FileInfo()
		// 创建一个空文件 写入解包后的数据
		fw, err := os.Create(fi.Name())
		if err != nil {
			message := fmt.Sprintf("error:%s,create new file:%s, error:%s\n", resErr.Error(), fi.Name, err.Error())
			resErr = errors.New(message)
			continue
		}

		// 拷贝数据
		if _, err := io.Copy(fw, tr); err != nil {
			message := fmt.Sprintf("error:%s,copy data:%s to file:%s, error:%s\n", resErr.Error(), hdr.Name, fi.Name, err.Error())
			resErr = errors.New(message)
		}
		// 赋予文件系统权限
		os.Chmod(fi.Name(), fi.Mode().Perm())
		// 关闭
		fw.Close()
	}
	return resErr
}

// 字符串内容 创建压缩文件
func tarStringToFiles() error {
	// 创建一个缓冲区来写入我们的存档。
	buf := new(bytes.Buffer)

	// 创建一个新的tar存档。
	tw := tar.NewWriter(buf)

	// 将一些文件添加到存档中。
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling license."},
	}
	for _, file := range files {
		hdr := &tar.Header{
			Name: file.Name,
			Mode: 0600,
			Size: int64(len(file.Body)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			return err
		}
		if _, err := tw.Write([]byte(file.Body)); err != nil {
			return err
		}
	}
	// 确保在Close时检查错误。
	if err := tw.Close(); err != nil {
		return err
	}

	// 打开tar档案以供阅读。
	r := bytes.NewReader(buf.Bytes())
	tr := tar.NewReader(r)
	var resErr error
	// 迭代档案中的文件。
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			// tar归档结束
			break
		}
		if err != nil {
			resErr = err
			continue
		}
		fmt.Printf("Contents of %s:\n", hdr.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			resErr = err
			continue
		}
		fmt.Println()
	}
	return resErr
}

func showTitle() {
	title := "tar包实现了tar格式压缩文件的存取。本包目标是覆盖大多数tar的变种，包括GNU和BSD生成的tar文件。"
	fmt.Printf("Title:%s\n", title)
}

func showConst() {
	fmt.Printf("Constants:\n%s\n", ConstType)
	fmt.Printf("Variables:\n%s\n", Variables)
}

func showHeaderStruct() {
	fmt.Printf("Header struct:\n%s\n", HeaderStruct)
}
