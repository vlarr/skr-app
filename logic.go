package main

const unknownEffectId = 0
const otherEffectId = 1

type IngridIdsWithWorth struct {
	ingridIds []int
	worth     float64
}

type IngridNamesWithWorth struct {
	ingridNamesPtr *[]string
	worth          float64
}

type byWorth []IngridNamesWithWorth

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
				if value1 == unknownEffectId || value1 == otherEffectId {
					continue
				}
				for _, value2 := range a[j] {
					if value2 == unknownEffectId || value2 == otherEffectId {
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

func validateIds(ids []int) bool {
	for i := 0; i < len(ids)-1; i++ {
		if ids[i+1] <= ids[i] {
			return false
		}
	}
	return true
}

func buildWorthOfCombinationTable(contextPtr *context, isFilterZeroWorth bool) *[]IngridIdsWithWorth {
	result := make([]IngridIdsWithWorth, 0)

	ingridIds := make([]int, len(contextPtr.ingridIdToInfoMap))
	{
		i := 0
		for id := range contextPtr.ingridIdToInfoMap {
			ingridIds[i] = id
			i++
		}
	}

	iter := createIter(2, &ingridIds)
	isNext := true

	for isNext {
		ids := iter.getValues()
		if validateIds(ids) {
			combinationExists, combinationWorth := calculateWorth(contextPtr, ids[0], ids[1])
			if combinationExists {
				if isFilterZeroWorth {
					if combinationWorth > 0 {
						result = append(result, IngridIdsWithWorth{ingridIds: ids, worth: combinationWorth})
					}
				} else {
					result = append(result, IngridIdsWithWorth{ingridIds: ids, worth: combinationWorth})
				}
			}
		}

		isNext, iter = iter.next()
	}

	return &result
}

func replaceIngridIdsToNames(contextPtr *context, ingridIdsWithWorthPtr *[]IngridIdsWithWorth) *[]IngridNamesWithWorth {
	result := make([]IngridNamesWithWorth, len(*ingridIdsWithWorthPtr))
	i := 0
	for _, elem := range *ingridIdsWithWorthPtr {
		ingridNames := make([]string, len(elem.ingridIds))

		for j := 0; j < len(ingridNames); j++ {
			ingridNames[j] = contextPtr.ingridIdToInfoMap[elem.ingridIds[j]].name
		}

		result[i] = IngridNamesWithWorth{
			ingridNamesPtr: &ingridNames,
			worth:          elem.worth,
		}
		i++
	}
	return &result
}
