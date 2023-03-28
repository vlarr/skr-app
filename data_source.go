package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
)

type effectInfo struct {
	id    int
	name  string
	worth float64
}

type ingridInfo struct {
	id          int
	name        string
	effectIdArr [4]int
}

type context struct {
	effectIdToInfoMap map[int]*effectInfo
	ingridIdToInfoMap map[int]*ingridInfo
}

func readEffectCsvFile(filePath string) map[int]*effectInfo {
	f, err := os.Open(filePath)
	checkErr(err)

	csvReader := csv.NewReader(f)
	csvReader.LazyQuotes = true
	csvReader.Comment = '#'
	records, err := csvReader.ReadAll()
	checkErr(err)
	log.Printf("Read %d lines from %s.\n", len(records), filePath)

	var result = map[int]*effectInfo{}

	for _, record := range records {
		id, err := strconv.Atoi(strings.TrimSpace(record[0]))
		checkErr(err)
		name := strings.TrimSpace(record[1])
		worth, err := strconv.ParseFloat(strings.TrimSpace(record[2]), 64)
		checkErr(err)
		result[id] = &effectInfo{
			id:    id,
			name:  name,
			worth: worth,
		}
	}

	return result
}

func readIngridCsvFile(filePath string) map[int]*ingridInfo {
	f, err := os.Open(filePath)
	checkErr(err)

	csvReader := csv.NewReader(f)
	csvReader.LazyQuotes = true
	csvReader.Comment = '#'
	records, err := csvReader.ReadAll()
	checkErr(err)
	log.Printf("Read %d lines from %s.\n", len(records), filePath)

	var result = map[int]*ingridInfo{}

	for _, record := range records {
		id, err := strconv.Atoi(strings.TrimSpace(record[0]))
		checkErr(err)
		name := strings.TrimSpace(record[1])
		eff1, err := strconv.Atoi(strings.TrimSpace(record[2]))
		checkErr(err)
		eff2, err := strconv.Atoi(strings.TrimSpace(record[3]))
		checkErr(err)
		eff3, err := strconv.Atoi(strings.TrimSpace(record[4]))
		checkErr(err)
		eff4, err := strconv.Atoi(strings.TrimSpace(record[5]))
		checkErr(err)

		result[id] = &ingridInfo{
			id:          id,
			name:        name,
			effectIdArr: [4]int{eff1, eff2, eff3, eff4},
		}
	}

	return result
}

func readCsvFiles(effectCsvPath string, ingridCsvPath string) *context {
	effectIdToInfoMap := readEffectCsvFile(effectCsvPath)
	ingridIdToInfoMap := readIngridCsvFile(ingridCsvPath)

	return &context{
		effectIdToInfoMap: effectIdToInfoMap,
		ingridIdToInfoMap: ingridIdToInfoMap,
	}
}
