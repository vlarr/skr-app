package main

import (
	"flag"
	"log"
	"sort"
	"strconv"
	"strings"
)

func parseIntSet(str string) []int {
	var result []int
	for _, s := range strings.Split(str, ",") {
		num, err := strconv.Atoi(strings.TrimSpace(s))
		checkErr(err, "")
		result = append(result, num)
	}
	return result
}

func main() {
	log.Println("hello there")
	showFlagPtr := flag.Bool("show", false, "show results in console")
	saveFlagPtr := flag.Bool("save", false, "save results to output file")
	reduceCoefFlagPtr := flag.Bool("rc", false, "enable reduce coefficient by ingrid num")
	numIngridsStrPtr := flag.String("ni", "2,3", "num ingrids (ets. \"2\", \"2,3\", \"2,3,4\").")
	effectCsvFileNamePtr := flag.String("effect-csv", "effect.csv", "effect csv file name.")
	ingridCsvFileNamePtr := flag.String("ingrid-csv", "ingrid.csv", "ingrid csv file name.")
	outputFileNamePtr := flag.String("output-file", "output.txt", "output file name.")

	flag.Parse()
	numIngrids := parseIntSet(*numIngridsStrPtr)

	contextPtr := readCsvFiles(*effectCsvFileNamePtr, *ingridCsvFileNamePtr)
	ingridIdsWithWorthTable := buildWorthOfCombinationTableForIngridNums(contextPtr, numIngrids, *reduceCoefFlagPtr)
	ingridNamesWithWorthTable := replaceIngridIdsToNames(contextPtr, ingridIdsWithWorthTable)
	sort.Sort(byWorth(*ingridNamesWithWorthTable))

	if *showFlagPtr {
		showResult(ingridNamesWithWorthTable)
	}
	if *saveFlagPtr {
		saveResultToFile(ingridNamesWithWorthTable, *outputFileNamePtr)
	}
}
