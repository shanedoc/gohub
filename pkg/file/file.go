package file

import "os"

//package file辅助函数

//将数据存入文件
func Put(data []byte, to string) error {
	if err := os.WriteFile(to, data, 0644); err != nil {
		return err
	}
	return nil
}

//判断文件是否存在
func Exists(fileToCheck string) bool {
	if _, err := os.Stat(fileToCheck); os.IsNotExist(err) {
		return false
	}
	return true
}
