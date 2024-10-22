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

type FileSorter struct {
	input   io.Reader
	output  io.Writer
	reverse bool
}

func NewFileSorter(input io.Reader, output io.Writer, reverse bool) *FileSorter {
	return &FileSorter{
		input:   input,
		output:  output,
		reverse: reverse,
	}
}

func (fs *FileSorter) ReadFiles() ([]FileInfo, error) {
	scanner := bufio.NewScanner(fs.input)
	var files []FileInfo

	for scanner.Scan() {
		path := scanner.Text()
		file, err := fs.getFileInfo(path)
		if err != nil {
			fmt.Fprintf(fs.output, "Error reading file %s: %v\n", path, err)
			continue
		}
		files = append(files, file)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading input: %w", err)
	}

	return files, nil
}

func (fs *FileSorter) getFileInfo(path string) (FileInfo, error) {
	info, err := os.Stat(path)
	if err != nil {
		return FileInfo{}, err
	}

	return FileInfo{
		Path:    path,
		ModTime: info.ModTime().UnixNano(),
	}, nil
}

func (fs *FileSorter) sortFiles(files []FileInfo) {
	sort.Slice(files, func(i, j int) bool {
		if fs.reverse {
			return files[i].ModTime > files[j].ModTime
		}
		return files[i].ModTime < files[j].ModTime
	})
}

func (fs *FileSorter) writeResults(files []FileInfo) error {
	for _, file := range files {
		if _, err := fmt.Fprintln(fs.output, file.Path); err != nil {
			return fmt.Errorf("error writing output: %w", err)
		}
	}
	return nil
}

func SortFiles(input io.Reader, output io.Writer, reverse bool) error {
	sorter := NewFileSorter(input, output, reverse)

	files, err := sorter.ReadFiles()
	if err != nil {
		return err
	}

	sorter.sortFiles(files)
	return sorter.writeResults(files)
}
