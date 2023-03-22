package main

const UnknownEffectId = 0
const OtherEffectId = 1

type ingridIdsWithWorth struct {
	ingridIds []int
	worth     float64
}

type ingridNamesWithWorth struct {
	ingridNamesPtr *[]string
	worth          float64
}

type byWorth []ingridNamesWithWorth

func (s byWorth) Len() int {
	return len(s)
}

func (s byWorth) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byWorth) Less(i, j int) bool {
	return s[i].worth > s[j].worth
}

func findEffectIdPair(a ...[4]int) map[int]bool {
	pairIdMap := map[int]bool{}
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			for _, value1 := range a[i] {
				if value1 == UnknownEffectId || value1 == OtherEffectId {
					continue
				}
				for _, value2 := range a[j] {
					if value2 == UnknownEffectId || value2 == OtherEffectId {
						continue
					}
					if value1 == value2 {
						pairIdMap[value1] = true
					}
				}
			}
		}
	}

	return pairIdMap
}

func calculateWorth(contextPtr *context, ingridIds ...int) (exists bool, worth float64) {
	ingridIdMap := map[int]bool{}
	for _, ingridId := range ingridIds {
		ingridIdMap[ingridId] = true
	}

	if len(ingridIdMap) < 2 || len(ingridIds) != len(ingridIdMap) {
		panic(1)
	}

	var ingridEffectsTable [][4]int
	for _, ingridId := range ingridIds {
		effectIdArr := contextPtr.ingridIdToInfoMap[ingridId].effectIdArr
		ingridEffectsTable = append(ingridEffectsTable, effectIdArr)
	}

	effectIds := findEffectIdPair(ingridEffectsTable...)

	if len(effectIds) == 0 {
		return false, 0.0
	} else {
		var result = 0.0
		for id := range effectIds {
			result = result + contextPtr.effectIdToInfoMap[id].worth
		}
		return true, result
	}
}

func buildWorthOfCombinationTable(contextPtr *context, isFilterZeroWorth bool) *[]ingridIdsWithWorth {
	result := make([]ingridIdsWithWorth, 0)

	for id1, ingrid1 := range contextPtr.ingridIdToInfoMap {
		for id2, ingrid2 := range contextPtr.ingridIdToInfoMap {
			if id2 > id1 {
				_, combinationWorth := calculateWorth(contextPtr, id1, id2)
				idPair := []int{ingrid1.id, ingrid2.id}
				if isFilterZeroWorth {
					if combinationWorth > 0 {
						result = append(result, ingridIdsWithWorth{ingridIds: idPair, worth: combinationWorth})
					}
				} else {
					result = append(result, ingridIdsWithWorth{ingridIds: idPair, worth: combinationWorth})
				}
			}
		}
	}
	return &result
}

func replaceIngridIdsToNames(contextPtr *context, ingridIdsWithWorthPtr *[]ingridIdsWithWorth) *[]ingridNamesWithWorth {
	result := make([]ingridNamesWithWorth, len(*ingridIdsWithWorthPtr))
	i := 0
	for _, elem := range *ingridIdsWithWorthPtr {
		ingridNames := make([]string, len(elem.ingridIds))

		for j := 0; j < len(ingridNames); j++ {
			ingridNames[j] = contextPtr.ingridIdToInfoMap[elem.ingridIds[j]].name
		}

		result[i] = ingridNamesWithWorth{
			ingridNamesPtr: &ingridNames,
			worth:          elem.worth,
		}
		i++
	}
	return &result
}
