package main

import (
	"fmt"
	"sort"
)

func main() {
	packItems(12500)
}

func packItems(i int) {
	items := i
	packs := []int{5000, 250, 2000, 500, 1000}
	packedItems := make(map[int]int) // map to store packed items
	sortedPacks := sort.IntSlice(packs)
	// sort in descending order to start with the largest pack size
	sort.Sort(sort.Reverse(sortedPacks))
	for _, p := range sortedPacks {
		isLastPackSize := p == packs[len(packs)-1]
		rest := items % p
		if rest > 0 {
			if items/p >= 1 { // check if the pack size can be used
				packedItems[p] = items / p
			}
			items = rest
			if isLastPackSize {
				pi := packedItems[p]
				if pi > 0 {
					packedItems[p] += 1
					items = rest // update the remaining items
				} else {
					packedItems[p] = 1
				}
			}
		} else {
			packedItems[p] = items / p
			break
		}
	}
	fmt.Println(packedItems)
}
