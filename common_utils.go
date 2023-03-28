package main

import "log"

func checkUniqueIds(ids []int) bool {
	idMap := map[int]bool{}
	for _, id := range ids {
		idMap[id] = true
	}
	return len(ids) == len(idMap)
}

func equalIdMaps(map1 map[int]bool, map2 map[int]bool) bool {
	for i, elem1 := range map1 {
		if map2[i] != elem1 {
			return false
		}
	}
	return true
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
