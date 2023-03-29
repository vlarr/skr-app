package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func convPotionToString(potion *Potion) string {
	return fmt.Sprintf("worth=%-6.1f cost=%-6.1f profit=%-6.1f ingrid=%v\n", potion.worth, potion.cost, potion.profit, strings.Join(potion.ingridNames, ", "))
}

func showResult(potions []*Potion, limit int) {
	fmt.Printf("all records: %v, limit: %v\n", len(potions), limit)
	for i, potion := range potions {
		if i >= limit {
			break
		}
		fmt.Print(convPotionToString(potion))
	}
}

func saveResultToFile(potions []*Potion, fileName string, limit int) {
	f, _ := os.Create(fileName)
	defer func(f *os.File) {
		err := f.Close()
		checkErr(err, "")
	}(f)

	stringCounter := 0

	_, err := f.WriteString(fmt.Sprintf("all records: %v, limit: %v\n", len(potions), limit))
	checkErr(err, "")

	for i, potion := range potions {
		if i >= limit {
			break
		}
		_, err = f.WriteString(convPotionToString(potion))
		checkErr(err, "")
		stringCounter++
	}

	log.Printf("Write %v records to %v\n", stringCounter, fileName)

	err = f.Sync()
	checkErr(err, "")
}
