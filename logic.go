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

func findActiveEffectsByIngridEffects(ingridEffectsTable ...[4]int) map[int]bool {
	pairIdMap := map[int]bool{}
	for i := 0; i < len(ingridEffectsTable); i++ {
		for j := i + 1; j < len(ingridEffectsTable); j++ {
			for _, value1 := range ingridEffectsTable[i] {
				if value1 == unknownEffectId || value1 == otherEffectId {
					continue
				}
				for _, value2 := range ingridEffectsTable[j] {
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

func findActiveEffectsByIngridIds(contextPtr *context, ingridIds ...int) map[int]bool {
	var ingridEffectsTable [][4]int
	for _, ingridId := range ingridIds {
		effectIdArr := contextPtr.ingridIdToInfoMap[ingridId].effectIdArr
		ingridEffectsTable = append(ingridEffectsTable, effectIdArr)
	}

	return findActiveEffectsByIngridEffects(ingridEffectsTable...)
}

func calculateWorth(contextPtr *context, enableReduceCoef bool, ingridIds ...int) (exists bool, worth float64) {
	if !checkUniqueIds(ingridIds) || len(ingridIds) < 2 {
		panic(1)
	}

	effectIds := findActiveEffectsByIngridIds(contextPtr, ingridIds...)
	if len(effectIds) == 0 {
		return false, 0.0
	}

	var result = 0.0
	for id := range effectIds {
		result = result + contextPtr.effectIdToInfoMap[id].worth
	}

	if enableReduceCoef {
		reduceWorthCoef := 2.0 / float64(len(ingridIds))
		result = result * reduceWorthCoef
	}

	return true, result
}

func buildWorthOfCombinationTable(contextPtr *context, ingridNum int, enableReduceCoef bool) *[]IngridIdsWithWorth {
	result := make([]IngridIdsWithWorth, 0)

	ingridIds := make([]int, len(contextPtr.ingridIdToInfoMap))
	{
		i := 0
		for id := range contextPtr.ingridIdToInfoMap {
			ingridIds[i] = id
			i++
		}
	}

	iter := createIter(ingridNum, &ingridIds)
	isNext := true

	for isNext {
		ids := iter.getValues()
		if simpleValidateIngridIdsByOrder(ids) && validateIngridByActiveEffects(contextPtr, ids) {
			combinationExists, combinationWorth := calculateWorth(contextPtr, enableReduceCoef, ids...)
			if combinationExists {
				if combinationWorth > 0 {
					result = append(result, IngridIdsWithWorth{ingridIds: ids, worth: combinationWorth})
				}
			}
		}

		isNext, iter = iter.next()
	}

	return &result
}

func buildWorthOfCombinationTableForIngridNums(contextPtr *context, ingridNums []int, enableReduceCoef bool) *[]IngridIdsWithWorth {
	result := make([]IngridIdsWithWorth, 0)

	for _, num := range ingridNums {
		resultElem := buildWorthOfCombinationTable(contextPtr, num, enableReduceCoef)
		for _, worth := range *resultElem {
			result = append(result, worth)
		}
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
