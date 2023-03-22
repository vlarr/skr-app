package main

const UnknownEffectId = 0
const OtherEffectId = 1

type worthInfoWithIds struct {
	ingridIds []int
	worth     float64
}

type worthInfoWithNames struct {
	ingridNames *[]string
	worth       float64
}

type byWorth []worthInfoWithNames

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

func calculateWorth(contextInst *context, ingridId1 int, ingridId2 int) float64 {
	if ingridId1 == ingridId2 {
		panic(1)
	}

	ingridInfo1 := contextInst.ingridIdToInfoMap[ingridId1]
	ingridInfo2 := contextInst.ingridIdToInfoMap[ingridId2]
	effectIds := findEffectIdPair(ingridInfo1.effectIdArr, ingridInfo2.effectIdArr)
	var result = 0.0
	for id := range effectIds {
		result = result + contextInst.effectIdToInfoMap[id].worth
	}
	return result
}

func buildWorthCombinationArray(contextInst *context, isFilterZeroWorth bool) *[]worthInfoWithIds {
	result := make([]worthInfoWithIds, 0)

	for id1, ingrid1 := range contextInst.ingridIdToInfoMap {
		for id2, ingrid2 := range contextInst.ingridIdToInfoMap {
			if id2 > id1 {
				combinationWorth := calculateWorth(contextInst, id1, id2)
				idPair := []int{ingrid1.id, ingrid2.id}
				if isFilterZeroWorth {
					if combinationWorth > 0 {
						result = append(result, worthInfoWithIds{ingridIds: idPair, worth: combinationWorth})
					}
				} else {
					result = append(result, worthInfoWithIds{ingridIds: idPair, worth: combinationWorth})
				}
			}
		}
	}
	return &result
}

func replaceIngridIdsToNames(contextInst *context, worthInfoWithIdsArr *[]worthInfoWithIds) *[]worthInfoWithNames {
	result := make([]worthInfoWithNames, len(*worthInfoWithIdsArr))
	i := 0
	for _, elem := range *worthInfoWithIdsArr {
		ingridNames := make([]string, len(elem.ingridIds))

		for j := 0; j < len(ingridNames); j++ {
			ingridNames[j] = contextInst.ingridIdToInfoMap[elem.ingridIds[j]].name
		}

		result[i] = worthInfoWithNames{
			ingridNames: &ingridNames,
			worth:       elem.worth,
		}
		i++
	}
	return &result
}
