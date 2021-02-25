package cf

import (
	"os"
	"os/user"
)

// GetHomeDir 获取homeDir
func GetHomeDir() string {
	u, err1 := user.Current()
	if err1 != nil {
		panic(err1)
	}
	return u.HomeDir
}

// IsDir 判断目录是否存在 return bool
func IsDir(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		panic(err)
	}
	return s.IsDir()
}

// CreateDir 创建目录
func CreateDir(path string) (string, error) {
	if !IsDir(path) {
		err := os.Mkdir(path, 0666)
		if err != nil {
			return "", err
		}
	}
	return path, nil
}
