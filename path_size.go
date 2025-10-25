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

func FormatSize(bytes int64, flag string) string {
	if flag == "human" {
		return humanReadableSize(bytes)
	}

	return fmt.Sprintf("%dB", bytes)
}

func humanReadableSize(bytes int64) string {
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
		TB = GB * 1024
		PB = TB * 1024
		EB = PB * 1024
	)

	switch {
	case bytes >= EB:
		return fmt.Sprintf("%.1fEB", float64(bytes)/float64(EB))
	case bytes >= PB:
		return fmt.Sprintf("%.1fPB", float64(bytes)/float64(PB))
	case bytes >= TB:
		return fmt.Sprintf("%.1fTB", float64(bytes)/float64(TB))
	case bytes >= GB:
		return fmt.Sprintf("%.1fGB", float64(bytes)/float64(GB))
	case bytes >= MB:
		return fmt.Sprintf("%.1fMB", float64(bytes)/float64(MB))
	case bytes >= KB:
		return fmt.Sprintf("%.1fKB", float64(bytes)/float64(KB))
	default:
		return fmt.Sprintf("%dB", bytes)
	}
}
