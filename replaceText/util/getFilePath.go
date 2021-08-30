package util

import (
	"io/fs"
	"io/ioutil"
	"os"
)

func GetFilePath(pathname string) (files []string, err error) {

	files = make([]string, 0)

	var rd []fs.FileInfo
	if isDir(pathname) {
		// if path, read it
		rd, err = ioutil.ReadDir(pathname)
		if err != nil {
			return nil, err
		}

		for _, fi := range rd {
			if fi.IsDir() {
				continue
			}
			files = append(files, pathname+"/"+fi.Name())
		}

	} else {
		files = append(files, pathname)
	}

	return
}

func isDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func isFile(path string) (isfielbool, err error) {
	return
}
