package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func showResult(resultPtr *[]IngridNamesWorth, limit int) {
	fmt.Printf("results: %v\n", len(*resultPtr))
	for i, worthInfo := range *resultPtr {
		if i >= limit {
			break
		}
		fmt.Printf("worth=%-6.1f ingrid=%v\n", worthInfo.worth, strings.Join(*worthInfo.ingridNamesPtr, ", "))
	}
}

func saveResultToFile(resultPtr *[]IngridNamesWorth, fileName string, limit int) {
	f, _ := os.Create(fileName)
	defer func(f *os.File) {
		err := f.Close()
		checkErr(err, "")
	}(f)

	_, err := f.WriteString(fmt.Sprintf("results: %v\n", len(*resultPtr)))
	checkErr(err, "")

	for i, worthInfo := range *resultPtr {
		if i >= limit {
			break
		}
		str := fmt.Sprintf("worth=%-6.1f ingrid=%v\n", worthInfo.worth, strings.Join(*worthInfo.ingridNamesPtr, ", "))
		_, err = f.WriteString(str)
		checkErr(err, "")
	}

	log.Printf("Write %v lines to %v\n", len(*resultPtr), fileName)

	err = f.Sync()
	checkErr(err, "")
}
