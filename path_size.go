package code

import (
	"fmt"
	"os"
)

func GetSize(path string) (string, error) {
	if len(path) == 0 {
		return "The path to the file or directory has not been entered. Run the program with the -h flag to read the help.", nil
	}
	fi, err := os.Lstat(path)
	if err != nil {
		return fmt.Sprintf("Error: %s", err), err
	}

	var size int64
	switch mode := fi.Mode(); {
	case mode.IsRegular():
		size = fi.Size()
	case mode.IsDir():
		files, err := os.ReadDir(path)
		if err != nil {
			return fmt.Sprintf("Error: %s", err), err
		}

		for _, file := range files {
			if !file.IsDir() {
				filePath := path + "/" + file.Name()
				fi, err := os.Lstat(filePath)
				if err != nil {
					return fmt.Sprintf("Error: %s", err), err
				}
				size += fi.Size()
			}
		}
	}
	return fmt.Sprintf("%dB	%s", size, path), err
}
