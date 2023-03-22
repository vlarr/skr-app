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

	contextInst := readCsvFiles("./effect.csv", "./ingrid.csv")

	pairIdToWorthMap := buildWorthCombinationArray(contextInst, true)
	worthInfoArr := replaceIngridIdsToNames(contextInst, pairIdToWorthMap)
	sort.Sort(byWorth(*worthInfoArr))

	showFlagPtr := flag.Bool("show", false, "show results")
	saveFlagPtr := flag.Bool("save", false, "save results")
	flag.Parse()

	if *showFlagPtr {
		showResult(worthInfoArr)
	}
	if *saveFlagPtr {
		saveResultToFile(worthInfoArr, "./output.txt")
	}
}
