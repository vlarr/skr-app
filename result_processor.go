package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func showResult(result *[]worthInfoWithNames) {
	for _, worthInfo := range *result {
		fmt.Printf("worth=%-6v ingrid=%v\n", worthInfo.worth, strings.Join(*worthInfo.ingridNames, ", "))
	}
}

func saveResultToFile(result *[]worthInfoWithNames, fileName string) {
	f, _ := os.Create(fileName)
	defer f.Close()

	for _, worthInfo := range *result {
		str := fmt.Sprintf("worth=%-6v ingrid=%v\n", worthInfo.worth, strings.Join(*worthInfo.ingridNames, ", "))
		f.WriteString(str)
	}

	log.Printf("Write %v lines to %v\n", len(*result), fileName)

	f.Sync()
}
