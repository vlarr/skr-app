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

	ingridIdsWithWorthTable := buildWorthOfCombinationTableForIngridNums(contextPtr, []int{2, 3}, true)
	ingridNamesWithWorthTable := replaceIngridIdsToNames(contextPtr, ingridIdsWithWorthTable)
	sort.Sort(byWorth(*ingridNamesWithWorthTable))

	showFlagPtr := flag.Bool("show", false, "show results")
	saveFlagPtr := flag.Bool("save", false, "save results")
	flag.Parse()

	if *showFlagPtr {
		showResult(ingridNamesWithWorthTable)
	}
	if *saveFlagPtr {
		saveResultToFile(ingridNamesWithWorthTable, "./output.txt")
	}
}
