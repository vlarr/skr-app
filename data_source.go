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
	code  string
	name  string
	worth float64
}

type ingridInfo struct {
	name        string
	cost        float64
	effectCodes [4]string
	id          int
	effectIds   [4]int
}

type context struct {
	effectCodeToInfoMap map[string]*effectInfo
	effectIdToInfoMap   map[int]*effectInfo
	ingridIdToInfoMap   map[int]*ingridInfo
}

func (c *context) readEffectCsvFile(filePath string) {
	f, err := os.Open(filePath)
	checkErr(err, "")

	csvReader := csv.NewReader(f)
	csvReader.LazyQuotes = true
	csvReader.Comment = '#'
	records, err := csvReader.ReadAll()
	checkErr(err, "")
	log.Printf("Read %d lines from %s.\n", len(records), filePath)

	c.effectIdToInfoMap = map[int]*effectInfo{}
	c.effectCodeToInfoMap = map[string]*effectInfo{}
	id := 0
	for _, record := range records {
		code := strings.TrimSpace(record[0])
		name := strings.TrimSpace(record[1])
		worth, err := strconv.ParseFloat(strings.TrimSpace(record[2]), 64)
		checkErr(err, "")
		effectInfoPtr := &effectInfo{
			id:    id,
			code:  code,
			name:  name,
			worth: worth,
		}
		c.effectIdToInfoMap[effectInfoPtr.id] = effectInfoPtr
		c.effectCodeToInfoMap[effectInfoPtr.code] = effectInfoPtr
		id++
	}
}

func (c *context) readIngridCsvFile(filePath string) {
	f, err := os.Open(filePath)
	checkErr(err, "")

	csvReader := csv.NewReader(f)
	csvReader.LazyQuotes = true
	csvReader.Comment = '#'
	records, err := csvReader.ReadAll()
	checkErr(err, "")
	log.Printf("Read %d lines from %s.\n", len(records), filePath)

	c.ingridIdToInfoMap = map[int]*ingridInfo{}
	id := 0
	for _, record := range records {
		name := strings.TrimSpace(record[0])

		cost, err := strconv.ParseFloat(strings.TrimSpace(record[1]), 64)
		checkErr(err, "")

		eff1Code := strings.TrimSpace(record[2])
		eff1Id := c.effectCodeToInfoMap[eff1Code].id
		eff2Code := strings.TrimSpace(record[3])
		eff2Id := c.effectCodeToInfoMap[eff2Code].id
		eff3Code := strings.TrimSpace(record[4])
		eff3Id := c.effectCodeToInfoMap[eff3Code].id
		eff4Code := strings.TrimSpace(record[5])
		eff4Id := c.effectCodeToInfoMap[eff4Code].id

		ingridInfoPtr := &ingridInfo{
			id:          id,
			name:        name,
			cost:        cost,
			effectCodes: [4]string{eff1Code, eff2Code, eff3Code, eff4Code},
			effectIds:   [4]int{eff1Id, eff2Id, eff3Id, eff4Id},
		}

		c.ingridIdToInfoMap[id] = ingridInfoPtr

		id++
	}
}

func readCsvFiles(effectCsvPath string, ingridCsvPath string) *context {
	result := new(context)
	result.readEffectCsvFile(effectCsvPath)
	result.readIngridCsvFile(ingridCsvPath)
	return result
}
