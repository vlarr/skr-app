package main

func simpleValidateIngridIdsByOrder(ingridIds []int) bool {
	for i := 0; i < len(ingridIds)-1; i++ {
		if ingridIds[i+1] <= ingridIds[i] {
			return false
		}
	}
	return true
}

func validateIngridByActiveEffects(contextPtr *context, ingridIds []int) bool {
	if !checkUniqueIds(ingridIds) || len(ingridIds) < 2 {
		return false
	}

	effectIds := findActiveEffectsByIngridIds(contextPtr, ingridIds...)
	if len(effectIds) == 0 {
		return false
	}

	if len(ingridIds) > 2 {
		for _, id1 := range ingridIds {
			var tempIngridIds []int
			for _, id2 := range ingridIds {
				if id1 != id2 {
					tempIngridIds = append(tempIngridIds, id2)
				}
			}
			tempEffectIds := findActiveEffectsByIngridIds(contextPtr, tempIngridIds...)
			if equalIdMaps(effectIds, tempEffectIds) {
				return false
			}
		}
	}

	return true
}
