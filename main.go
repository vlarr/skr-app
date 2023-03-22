package main

import (
	"log"
	"os"
	"sort"
)

func containArg(args *[]string, testArg string) bool {
	for _, arg := range *args {
		if testArg == arg {
			return true
		}
	}
	return false
}

func main() {
	log.Println("hello there")

	contextInst := readCsvFiles("./local/effect.csv", "./local/ingrid.csv")

	pairIdToWorthMap := buildWorthCombinationArray(contextInst, true)
	worthInfoArr := replaceIngridIdsToNames(contextInst, pairIdToWorthMap)
	sort.Sort(byWorth(*worthInfoArr))

	args := os.Args[1:]

	if containArg(&args, "--show") {
		showResult(worthInfoArr)
	}

	if containArg(&args, "--save") {
		saveResultToFile(worthInfoArr, "./local/output.txt")
	}
}
