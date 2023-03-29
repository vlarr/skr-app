package main

import (
	"flag"
	"log"
	"sort"
	"strconv"
	"strings"
)

func parseIntArray(str string) []int {
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
	effectCsvFileNamePtr := flag.String("effect-csv", "effects.csv", "effect csv file name.")
	ingridCsvFileNamePtr := flag.String("ingrid-csv", "ingredients.csv", "ingrid csv file name.")
	outputFileNamePtr := flag.String("output-file", "output.txt", "output file name.")
	limitPtr := flag.Int("limit", 20, "limit first values.")

	flag.Parse()
	numIngrids := parseIntArray(*numIngridsStrPtr)

	contextPtr := readCsvFiles(*effectCsvFileNamePtr, *ingridCsvFileNamePtr)
	ingridIdsWithWorthTable := buildWorthOfCombinationTableForIngridNums(contextPtr, numIngrids, *reduceCoefFlagPtr)
	ingridNamesWithWorthTable := replaceIngridIdsToNames(contextPtr, ingridIdsWithWorthTable)
	sort.Sort(byWorth(*ingridNamesWithWorthTable))

	if *showFlagPtr {
		showResult(ingridNamesWithWorthTable, *limitPtr)
	}
	if *saveFlagPtr {
		saveResultToFile(ingridNamesWithWorthTable, *outputFileNamePtr, *limitPtr)
	}
}
