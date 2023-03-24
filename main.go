package main

import (
	"flag"
	"log"
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

	contextPtr := readCsvFiles("./effect.csv", "./ingrid.csv")

	pairIdToWorthMap := buildWorthOfCombinationTable(contextPtr, true)
	worthInfoPtr := replaceIngridIdsToNames(contextPtr, pairIdToWorthMap)
	sort.Sort(byWorth(*worthInfoPtr))

	showFlagPtr := flag.Bool("show", false, "show results")
	saveFlagPtr := flag.Bool("save", false, "save results")
	flag.Parse()

	if *showFlagPtr {
		showResult(worthInfoPtr)
	}
	if *saveFlagPtr {
		saveResultToFile(worthInfoPtr, "./output.txt")
	}
}
