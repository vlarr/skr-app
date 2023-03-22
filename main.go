package main

import (
	"log"
	"sort"
)

func main() {
	log.Println("hello there")

	contextInst := readCsvFiles("./local/effect.csv", "./local/ingrid.csv")

	pairIdToWorthMap := buildWorthCombinationMap(contextInst, true)
	worthInfoArr := convertWorthCombinationMapToResultArr(contextInst, pairIdToWorthMap)
	sort.Sort(byWorth(*worthInfoArr))

	//showResult(worthInfoArr)
	saveResultToFile(worthInfoArr, "./local/result.txt")
}
