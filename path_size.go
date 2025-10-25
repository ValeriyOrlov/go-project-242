package code

import (
	"fmt"
	"os"
)

func GetSize(path string) (int64, error) {
	if len(path) == 0 {
		return 0, fmt.Errorf("the path to the file or directory has not been entered")
	}
	fi, err := os.Lstat(path)
	if err != nil {
		return 0, fmt.Errorf("error: %s", err)
	}

	var bytes int64
	switch mode := fi.Mode(); {
	case mode.IsRegular():
		bytes = fi.Size()
	case mode.IsDir():
		files, err := os.ReadDir(path)
		if err != nil {
			return 0, fmt.Errorf("error: %s", err)
		}

		for _, file := range files {
			if !file.IsDir() {
				filePath := path + "/" + file.Name()
				fi, err := os.Lstat(filePath)
				if err != nil {
					return 0, fmt.Errorf("error: %s", err)
				}
				bytes += fi.Size()
			}
		}
	}
	return bytes, err
}
