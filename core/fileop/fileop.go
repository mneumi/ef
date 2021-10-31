package fileop

import (
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

const (
	O_CREATE int = syscall.O_CREAT
	O_TRUNC  int = syscall.O_TRUNC
)

func WriteFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0666)
}

// 判断路径是否存在
func ExistPath(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

// 创建目录
func CreateDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

// 根据相对路径创建目录
func CreateDirRelative(relativePath string) error {
	dir := filepath.Dir(GetPathAbs(relativePath))
	return CreateDir(dir)
}

// 根据绝对路径创建目录
func CreateDirAbsolute(path string) error {
	dir := filepath.Dir(path)
	return CreateDir(dir)
}

// 根据相对路径获取绝对路径
func GetPathAbs(relativePath string) string {
	pwd, _ := os.Getwd()

	return filepath.Join(pwd, relativePath)
}

// 获取文件名的后缀
func GetFileExt(filename string) string {
	s := strings.Split(filename, ".")
	if len(s) < 2 {
		return ""
	}
	return s[1]
}

// 读取文件的全部内容，作为字符串返回
func ReadFile(path string) (string, error) {
	byte, err := os.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(byte), nil
}

// 获取路径分隔符
func GetSeparator() string {
	return string(filepath.Separator)
}

// 遍历目录
func Walk(root string, fn filepath.WalkFunc) error {
	return filepath.Walk(root, fn)
}

// 打开文件
func OpenFile(name string, flag int) (*os.File, error) {
	return os.OpenFile(name, flag, 0666)
}
