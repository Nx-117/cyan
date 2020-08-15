package file

import (
	"bufio"
	"io"
	"os"
)

/**
读取文件内容
filePath = 文件路径
*/
func Read(filePath string) ([]byte, error) {
	fi, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	r := bufio.NewReader(fi)

	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if 0 == n {
			break
		}
	}
	return buf, nil
}

/**
读取文件到字符串
filePath = 文件路径
*/
func ReadToString(filePath string) (string, error) {
	v, err := Read(filePath)
	if err != nil {
		return "", err
	}
	return string(v), nil
}

/**
判断文件是否存在  存在返回 true 不存在返回false
fileName = 文件路径+文件名字+文件后缀
*/
func CheckFileIsExist(fileName string) bool {
	var exist = true
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
