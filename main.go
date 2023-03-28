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
	showFlagPtr := flag.Bool("show", false, "show results")
	saveFlagPtr := flag.Bool("save", false, "save results")
	reduceCoefFlagPtr := flag.Bool("rc", false, "enable reduce coefficient by ingrid num")
	numIngridsStrPtr := flag.String("ni", "2", "num ingrids (ets. \"2\", \"2,3\", \"2,3,4\"). default: 2")
	flag.Parse()
	numIngrids := parseIntSet(*numIngridsStrPtr)

	contextPtr := readCsvFiles("./effect.csv", "./ingrid.csv")
	ingridIdsWithWorthTable := buildWorthOfCombinationTableForIngridNums(contextPtr, numIngrids, *reduceCoefFlagPtr)
	ingridNamesWithWorthTable := replaceIngridIdsToNames(contextPtr, ingridIdsWithWorthTable)
	sort.Sort(byWorth(*ingridNamesWithWorthTable))

	if *showFlagPtr {
		showResult(ingridNamesWithWorthTable)
	}
	if *saveFlagPtr {
		saveResultToFile(ingridNamesWithWorthTable, "./output.txt")
	}
}
