package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Foundlocation struct {
	Line       string
	LineNumber uint16
}

type Result struct {
	FilesContaining map[string][]Foundlocation
}

func SearchFiles(searchTerm string, searchedFiles []string) (Result, error) {
	result := Result{
		FilesContaining: make(map[string][]Foundlocation),
	}

	for _, name := range searchedFiles {
		lineNum := 0
		found := false
		allLocation := []Foundlocation{}
		readFile, err := os.Open(name)

		if err != nil {
			fmt.Println(err)
		}
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		var fileLines []string

		for fileScanner.Scan() {
			fileLines = append(fileLines, fileScanner.Text())
		}

		readFile.Close()

		for _, line := range fileLines {
			if strings.Contains(line, searchTerm) {
				allLocation = append(allLocation, Foundlocation{Line: line,
					LineNumber: uint16(lineNum)})
				found = true
			}
			lineNum++
		}

		if found {
			result.FilesContaining[name] = allLocation
		}
	}
	return result, nil
}
