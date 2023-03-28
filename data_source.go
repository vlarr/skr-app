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
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer func(f *os.File) {
		err := f.Close()
		checkErr(err)
	}(f)

	csvReader := csv.NewReader(f)
	csvReader.LazyQuotes = true
	csvReader.Comment = '#'
	records, err := csvReader.ReadAll()

	log.Printf("Read %d lines from %s.\n", len(records), filePath)
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	var result = map[int]*effectInfo{}

	for _, record := range records {
		id, _ := strconv.Atoi(strings.TrimSpace(record[0]))
		name := strings.TrimSpace(record[1])
		worth, _ := strconv.ParseFloat(strings.TrimSpace(record[2]), 64)
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
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer func(f *os.File) {
		err := f.Close()
		checkErr(err)
	}(f)

	csvReader := csv.NewReader(f)
	csvReader.LazyQuotes = true
	csvReader.Comment = '#'
	records, err := csvReader.ReadAll()

	log.Printf("Read %d lines from %s.\n", len(records), filePath)
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	var result = map[int]*ingridInfo{}

	for _, record := range records {
		id, _ := strconv.Atoi(strings.TrimSpace(record[0]))
		name := strings.TrimSpace(record[1])
		eff1, _ := strconv.Atoi(strings.TrimSpace(record[2]))
		eff2, _ := strconv.Atoi(strings.TrimSpace(record[3]))
		eff3, _ := strconv.Atoi(strings.TrimSpace(record[4]))
		eff4, _ := strconv.Atoi(strings.TrimSpace(record[5]))

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
