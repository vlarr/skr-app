package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func showResult(resultPtr *[]IngridNamesWithWorth) {
	fmt.Printf("results: %v\n", len(*resultPtr))
	for _, worthInfo := range *resultPtr {
		fmt.Printf("worth=%-6.1f ingrid=%v\n", worthInfo.worth, strings.Join(*worthInfo.ingridNamesPtr, ", "))
	}
}

func saveResultToFile(resultPtr *[]IngridNamesWithWorth, fileName string) {
	f, _ := os.Create(fileName)
	defer func(f *os.File) {
		err := f.Close()
		checkErr(err)
	}(f)

	_, err := f.WriteString(fmt.Sprintf("results: %v\n", len(*resultPtr)))
	checkErr(err)

	for _, worthInfo := range *resultPtr {
		str := fmt.Sprintf("worth=%-6.1f ingrid=%v\n", worthInfo.worth, strings.Join(*worthInfo.ingridNamesPtr, ", "))
		_, err = f.WriteString(str)
		checkErr(err)
	}

	log.Printf("Write %v lines to %v\n", len(*resultPtr), fileName)

	err = f.Sync()
	checkErr(err)
}
