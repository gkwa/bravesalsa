package core

import (
	"bytes"
	"strings"
	"testing"
)

func TestSortFiles(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		reverse bool
		want    string
	}{
		{
			name:    "empty input",
			input:   "",
			reverse: false,
			want:    "",
		},
		{
			name:    "single file",
			input:   "testdata/file1.txt\n",
			reverse: false,
			want:    "testdata/file1.txt\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := strings.NewReader(tt.input)
			output := &bytes.Buffer{}

			err := SortFiles(input, output, tt.reverse)
			if err != nil {
				t.Errorf("SortFiles() error = %v", err)
				return
			}

			if got := output.String(); got != tt.want {
				t.Errorf("SortFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}
