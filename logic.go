package main

const unknownEffectId = 0
const otherEffectId = 1

type Potion struct {
	ingridIds   []int
	ingridNames []string
	worth       float64
	cost        float64
	profit      float64
}

func findPotionsWithWorthForIngridNums(contextPtr *context, ingridNums []int) []*Potion {
	result := make([]*Potion, 0)
	for _, num := range ingridNums {
		resultElem := findPotionsWithWorthForIngridNum(contextPtr, num)
		for _, worth := range resultElem {
			result = append(result, worth)
		}
	}
	return result
}

func findPotionsWithWorthForIngridNum(contextPtr *context, ingridNum int) []*Potion {
	result := make([]*Potion, 0)

	stockIngridIds := make([]int, len(contextPtr.ingridIdToStockInfoMap))
	{
		i := 0
		for id := range contextPtr.ingridIdToStockInfoMap {
			stockIngridIds[i] = id
			i++
		}
	}

	idSliceIterator := createIterator(ingridNum, &stockIngridIds)
	isNext := true

	for isNext {
		currentIngridIds := idSliceIterator.getValues()
		potionPtr := tryMakePotion(contextPtr, currentIngridIds)
		if potionPtr != nil {
			potionPtr.cost = calculateCost(contextPtr, currentIngridIds...)
			potionPtr.profit = potionPtr.worth - potionPtr.cost
			potionPtr.ingridNames = findIngridNames(contextPtr, potionPtr.ingridIds)
			result = append(result, potionPtr)
		}

		isNext, idSliceIterator = idSliceIterator.next()
	}

	return result
}

func tryMakePotion(contextPtr *context, ids []int) *Potion {
	if simpleValidateIngridIdsByOrder(ids) && validateIngridByActiveEffects(contextPtr, ids) {
		combinationExists, combinationWorth := calculateWorth(contextPtr, ids...)
		if combinationExists {
			if combinationWorth > 0 {
				return &Potion{ingridIds: ids, worth: combinationWorth}
			}
		}
	}
	return nil
}

func calculateWorth(contextPtr *context, ingridIds ...int) (exists bool, worth float64) {
	if !checkUniqueIds(ingridIds) || len(ingridIds) < 2 {
		panic(1)
	}

	effectIds := findActiveEffectsByIngridIds(contextPtr, ingridIds...)
	if len(effectIds) == 0 {
		return false, 0.0
	}

	var result = 0.0
	for id := range effectIds {
		result += contextPtr.effectIdToInfoMap[id].worth
	}

	return true, result
}

func findActiveEffectsByIngridIds(contextPtr *context, ingridIds ...int) map[int]bool {
	var ingridEffectsTable [][4]int
	for _, ingridId := range ingridIds {
		effectIdArr := contextPtr.ingridIdToInfoMap[ingridId].effectIds
		ingridEffectsTable = append(ingridEffectsTable, effectIdArr)
	}

	return findActiveEffectsByIngridEffects(ingridEffectsTable...)
}

func findActiveEffectsByIngridEffects(ingridEffectsTable ...[4]int) map[int]bool {
	effectCounterMap := map[int]int{}

	for _, ingridEffectsRow := range ingridEffectsTable {
		for _, effectId := range ingridEffectsRow {
			effectCounterMap[effectId]++
		}
	}

	result := map[int]bool{}
	for effectId, i := range effectCounterMap {
		if effectId != unknownEffectId && effectId != otherEffectId && i >= 2 {
			result[effectId] = true
		}
	}
	return result
}

func calculateCost(contextPtr *context, ingridIds ...int) float64 {
	result := 0.0
	for _, id := range ingridIds {
		cost := contextPtr.ingridIdToInfoMap[id].cost / contextPtr.ingridIdToStockInfoMap[id].amount
		result += cost
	}
	return result
}

func findIngridNames(contextPtr *context, ingridIds []int) []string {
	ingridNames := make([]string, len(ingridIds))
	for j := 0; j < len(ingridNames); j++ {
		defName := contextPtr.ingridIdToInfoMap[ingridIds[j]].name
		alternativeName := contextPtr.ingridIdToInfoMap[ingridIds[j]].alternativeName

		if len(alternativeName) > 0 {
			ingridNames[j] = alternativeName
		} else {
			ingridNames[j] = defName
		}
	}
	return ingridNames
}
