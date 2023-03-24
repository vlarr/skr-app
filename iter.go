package main

type Iter struct {
	allValues *[]int
	indexes   []int
}

func createIter(size int, allValues *[]int) Iter {
	return Iter{
		indexes:   make([]int, size),
		allValues: allValues,
	}
}

func (p Iter) getValues() []int {
	result := make([]int, len(p.indexes))

	for i, index := range p.indexes {
		result[i] = (*p.allValues)[index]
	}

	return result
}

func (p Iter) incIndex(indexNum int) bool {
	if indexNum < 0 || indexNum >= len(p.indexes) {
		return false
	}

	p.indexes[indexNum] += 1
	if p.indexes[indexNum] >= len(*p.allValues) {
		p.indexes[indexNum] = 0
		incResult := p.incIndex(indexNum + 1)
		if incResult == false {
			return false
		}
	}
	return true
}

func (p Iter) next() (exists bool, iterInst Iter) {
	nextIter := p
	nextIterResult := nextIter.incIndex(0)
	return nextIterResult, nextIter
}
