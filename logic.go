package main

const UnknownEffectId = 0
const OtherEffectId = 1

type worthInfo struct {
	ingridNames *[]string
	worth       float64
}

type byWorth []worthInfo

func (s byWorth) Len() int {
	return len(s)
}

func (s byWorth) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byWorth) Less(i, j int) bool {
	return s[i].worth > s[j].worth
}

func findEffectIdPair(a [4]int, b [4]int) *[]int {
	var result []int
	for _, value1 := range a {
		if value1 == UnknownEffectId || value1 == OtherEffectId {
			continue
		}
		for _, value2 := range b {
			if value2 == UnknownEffectId || value2 == OtherEffectId {
				continue
			}
			if value1 == value2 {
				result = append(result, value1)
			}
		}
	}

	return &result
}

func calculateWorthCombination(contextInst context, info1 ingridInfo, info2 ingridInfo) float64 {
	ids := findEffectIdPair(info1.effectIdArr, info2.effectIdArr)
	var result = 0.0
	for _, id := range *ids {
		result = result + contextInst.effectIdToInfoMap[id].worth
	}
	return result
}

func buildWorthCombinationMap(contextInst context, isFilterZeroWorth bool) map[*[]int]float64 {
	result := make(map[*[]int]float64)

	for id1, ingrid1 := range contextInst.ingridIdToInfoMap {
		for id2, ingrid2 := range contextInst.ingridIdToInfoMap {
			if id2 > id1 {
				combinationWorth := calculateWorthCombination(contextInst, ingrid1, ingrid2)
				idPair := []int{ingrid1.id, ingrid2.id}
				if isFilterZeroWorth {
					if combinationWorth > 0 {
						result[&idPair] = combinationWorth
					}
				} else {
					result[&idPair] = combinationWorth
				}
			}
		}
	}
	return result
}

func convertWorthCombinationMapToResultArr(contextInst context, worthMap map[*[]int]float64) *[]worthInfo {
	result := make([]worthInfo, len(worthMap))
	i := 0
	for ingridIds, worth := range worthMap {
		ingridNames := make([]string, len(*ingridIds))

		for j := 0; j < len(*ingridIds); j++ {
			ingridNames[j] = contextInst.ingridIdToInfoMap[(*ingridIds)[j]].name
		}

		result[i] = worthInfo{
			ingridNames: &ingridNames,
			worth:       worth,
		}
		i++
	}
	return &result
}
