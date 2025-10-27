package code

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func GetDirSize(files []os.DirEntry, path string, flags []string) (int64, error) {
	var bytes int64
	for _, file := range files {
		if !file.IsDir() {
			fp := filepath.Join(path, file.Name())
			fi, err := os.Lstat(fp)
			if err != nil {
				return 0, fmt.Errorf("error from getDirSize 18: %s", err)
			}
			fileSize, err := GetFileSize(fi, flags)
			if err != nil {
				return 0, fmt.Errorf("error getDirSize 22: %s", err)
			}
			bytes += fileSize
		} else if file.IsDir() && slices.Contains(flags, "recursive") {
			fp := filepath.Join(path, file.Name())
			files, err := os.ReadDir(fp)
			if err != nil {
				return 0, fmt.Errorf("error from getDirSize 30: %s", err)
			}
			dirSize, err := GetDirSize(files, fp, flags)
			if err != nil {
				return 0, fmt.Errorf("error from getDirSize 36: %s", err)
			}
			bytes += dirSize
		}
	}
	return bytes, nil
}

func GetFileSize(fi os.FileInfo, flags []string) (int64, error) {
	if strings.HasPrefix(fi.Name(), ".") && !slices.Contains(flags, "all") {
		return 0, nil
	} else {
		return fi.Size(), nil
	}
}

func GetPathSize(path string, flags []string) (int64, error) {
	if len(path) == 0 {
		return 0, fmt.Errorf("the path to the file or directory has not been entered")
	}
	fi, err := os.Lstat(path)
	if err != nil {
		return 0, fmt.Errorf("error from getSize 59: %s", err)
	}
	var bytes int64
	switch mode := fi.Mode(); {
	case mode.IsRegular():
		fileSize, err := GetFileSize(fi, flags)
		if err != nil {
			return 0, fmt.Errorf("error from getSize 66: %s", err)
		}
		bytes += fileSize
	case mode.IsDir():
		files, err := os.ReadDir(path)
		if err != nil {
			return 0, fmt.Errorf("error from getSize 72: %s", err)
		}
		dirSize, err := GetDirSize(files, path, flags)
		if err != nil {
			return 0, fmt.Errorf("error from getSize 76: %s", err)
		}
		bytes += dirSize
	}
	return bytes, err
}

func FormatSize(bytes int64, flags []string) string {
	if slices.Contains(flags, "human") {
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
