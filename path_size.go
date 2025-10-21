package code

import (
	"fmt"
	"log"
	"os"
)

func GetSize(path string) string {
	fi, err := os.Lstat(path)
	if err != nil {
		log.Fatal(err)
	}

	var size int64
	switch mode := fi.Mode(); {
	case mode.IsRegular():
		size = fi.Size()
		/*case mode.IsDir():
		files, err := os.ReadDir(path)
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			if !file.IsDir() {
				fmt.Println(file.Info())
			}
		}
		*/
	}
	return fmt.Sprintf("size: %d b	path: %s", size, path)
}
