package controllers

import "fmt"

func Run() {
	result, err := SearchFiles("pal", []string{"cheese.txt"})
	if err != nil {
		fmt.Println("Trash")
	}
	fmt.Println(result)
}
