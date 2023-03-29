package main

//type byWorth []*Potion
//
//func (s byWorth) Len() int {
//	return len(s)
//}
//
//func (s byWorth) Swap(i, j int) {
//	s[i], s[j] = s[j], s[i]
//}
//
//func (s byWorth) Less(i, j int) bool {
//	return s[i].worth > s[j].worth
//}

type byProfit []*Potion

func (s byProfit) Len() int {
	return len(s)
}

func (s byProfit) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byProfit) Less(i, j int) bool {
	return s[i].profit > s[j].profit
}
