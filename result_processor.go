package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func showResult(resultPtr *[]worthInfoWithNames) {
	for _, worthInfo := range *resultPtr {
		fmt.Printf("worth=%-6v ingrid=%v\n", worthInfo.worth, strings.Join(*worthInfo.ingridNamesPtr, ", "))
	}
}

func saveResultToFile(resultPtr *[]worthInfoWithNames, fileName string) {
	f, _ := os.Create(fileName)
	defer f.Close()

	for _, worthInfo := range *resultPtr {
		str := fmt.Sprintf("worth=%-6v ingrid=%v\n", worthInfo.worth, strings.Join(*worthInfo.ingridNamesPtr, ", "))
		f.WriteString(str)
	}

	log.Printf("Write %v lines to %v\n", len(*resultPtr), fileName)

	f.Sync()
}
