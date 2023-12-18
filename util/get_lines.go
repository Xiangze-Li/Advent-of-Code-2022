package util

import (
	"bufio"
	"os"
)

// GetLines reads lines from file and returns them as a slice of strings.
func GetLines(filename string) []string {
	f := Must(os.Open(filename))
	defer f.Close()

	scanner := bufio.NewScanner(f)
	ret := make([]string, 0)
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}
	return ret
}

// GetBlocks reads lines from file, splits them with empty lines, and returns them as 2D slice of strings.
func GetBlocks(filename string) [][]string {
	lines := GetLines(filename)
	ret := make([][]string, 0)
	blockStart := 0
	for i := 0; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			ret = append(ret, lines[blockStart:i])
			blockStart = i + 1
		}
	}
	if blockStart < len(lines) {
		ret = append(ret, lines[blockStart:])
	}
	return ret
}
