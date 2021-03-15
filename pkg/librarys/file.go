package librarys

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

// GetSize get the file size
func FileGetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)

	return len(content), err
}

// GetExt get the file ext
func FileGetExt(fileName string) string {
	return path.Ext(fileName)
}

// CheckNotExist check if the file exists
func FileCheckNotExist(src string) bool {
	_, err := os.Stat(src)

	return os.IsNotExist(err)
}

// CheckPermission check if the file has permission
func FileCheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

// IsNotExistMkDir create a directory if it does not exist
func FileIsNotExistMkDir(src string) error {
	if notExist := FileCheckNotExist(src); notExist == true {
		if err := FileMkDir(src); err != nil {
			return err
		}
	}

	return nil
}

// MkDir create a directory
func FileMkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// Open a file according to a specific mode
func FileOpen(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// MustOpen maximize trying to open the file
func FileMustOpen(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := dir + "/" + filePath
	perm := FileCheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	err = FileIsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	f, err := FileOpen(src+fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
	}

	return f, nil
}
