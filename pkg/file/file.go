package file

import (
	"io/ioutil"
	"mime/multipart" //主要实现了 MIME 的 multipart 解析，主要适用于 HTTP 和常见浏览器生成的 multipart 主体
	"os"
	"path"
)

func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)
	return len(content), err
}

func GetExt(fileName string) string {
	return path.Ext(fileName)
}

// 能够接受ErrNotExist、syscall的一些错误，它会返回一个布尔值，能够得知文件不存在或目录不存在

func CheckNotExist(src string) bool {
	_, err := os.Stat(src) //返回文件信息结构描述文件。如果出现错误，会返回*PathError
	return os.IsNotExist(err)
}

// 能够接受ErrPermission、syscall的一些错误，它会返回一个布尔值，能够得知权限是否满足

func CheckPermission(src string) bool {
	_, err := os.Stat(src)
	return os.IsPermission(err)
}

func IsNotExistMkDir(src string) error {
	if notExist := CheckNotExist(src); notExist == true {
		if err := mkDir(src); err != nil {
			return err
		}
	}
	return nil
}

func mkDir(src string) error { //返回与当前目录对应的根路径名
	err := os.MkdirAll(src, os.ModePerm) //创建对应的目录以及所需的子目录，若成功则返回nil，否则返回error。/const定义ModePerm FileMode = 0777
	if err != nil {
		return err
	}
	return nil
}

func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	//调用文件，支持传入文件名称、指定的模式调用文件、文件权限，返回的文件的方法可以用于 I/O。如果出现错误，则为*PathError。
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}
	return f, nil
}
