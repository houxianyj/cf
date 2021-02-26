package cf

import (
	"io"
	"log"
	"os"
)

// ReadAll 读取文件所有内容
func ReadAll(filePath string) []byte {
	f, err := os.Open(filePath)
	if err != nil {
		//panic(err)
		log.Panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)
	//func ReadAll(r io.Reader) ([]byte, error)  r io.Reader 为什么 传入f *File 也可以呢？ 因为 f *File  继承了 type Reader interface
	//也就是实现了 type Reader interface 的 Read 方法
	res, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return res
}

type File struct {
}

func (f *File) ReadAll(file string) []byte {
	return ReadAll(file)
}

// NewFile 出始化File struct
func NewFile() *File {
	return new(File)
}
