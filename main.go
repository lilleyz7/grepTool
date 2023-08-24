package main

import (
	"fmt"

	"github.com/alexflint/go-arg"
	"github.com/lillez7/grepTool/controllers"
)

type arguments struct {
	SearchTerm     string   `arg:"positional" help:"term you would like to search"`
	FilePaths      []string `arg:"positional" help:"all of the file paths you would like to check for"`
	StandardSearch bool     `arg:"positional" help:"enter false to search where word is not found" default:"true"`
}

func (arguments) Description() string {
	return "This program replicates the functionality of grep using golang"
}

func Run() {
	var result arguments
	parser := arg.MustParse(&result)

	if result.SearchTerm == "" || len(result.FilePaths) == 0 {
		parser.Fail("First entry must be search term and the following will be file paths")
	}

	if result.StandardSearch == true {
		final, err := controllers.SearchFiles(result.SearchTerm, result.FilePaths)
		if err != nil {
			fmt.Println(err)
		}
		for key, entry := range final.FilesContaining {
			fmt.Printf("-----%s----- \n", key)
			for _, val := range entry {
				fmt.Printf("Line %d: %s \n", val.LineNumber, val.Line)
			}

		}
	}

	if result.StandardSearch == false {
		//
	}

}

func main() {
	Run()
}
