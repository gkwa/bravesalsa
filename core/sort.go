package core

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

type FileInfo struct {
	Path    string
	ModTime int64
}

func SortFiles(input io.Reader, output io.Writer, reverse bool) error {
	scanner := bufio.NewScanner(input)
	var files []FileInfo

	for scanner.Scan() {
		path := scanner.Text()
		info, err := os.Stat(path)
		if err != nil {
			fmt.Fprintf(output, "Error reading file %s: %v\n", path, err)
			continue
		}
		fileInfo := FileInfo{
			Path:    path,
			ModTime: info.ModTime().UnixNano(),
		}
		files = append(files, fileInfo)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading input: %w", err)
	}

	sort.Slice(files, func(i, j int) bool {
		if reverse {
			return files[i].ModTime > files[j].ModTime
		}
		return files[i].ModTime < files[j].ModTime
	})

	for _, file := range files {
		fmt.Fprintln(output, file.Path)
	}

	return nil
}
