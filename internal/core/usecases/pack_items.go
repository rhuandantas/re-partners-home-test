package usecases

import (
	"errors"
	"github.com/rhuandantas/re-partners-home-test/internal/adapters/repository/cache"
	"sort"
	"strconv"
	"strings"
)

type PackItem interface {
	Execute(items int) (map[int]int, error)
}

type packItem struct {
	cacheClient cache.MemcacheClient
}

func NewPackItem(cacheClient cache.MemcacheClient) PackItem {
	return &packItem{
		cacheClient: cacheClient,
	}
}

func (p *packItem) Execute(items int) (map[int]int, error) {
	packItems := items
	packedItems := make(map[int]int) // map to store packed items

	value, err := p.cacheClient.Get("pack-size")
	if err != nil {
		return nil, errors.New("error getting pack sizes")
	}

	if len(value) == 0 {
		return nil, errors.New("pack sizes not found. please add pack sizes")
	}

	packs := make([]int, 0)
	for _, v := range strings.Split(string(value), ",") {
		pack, err := strconv.Atoi(v)
		if err != nil {
			return nil, errors.New("error parsing pack sizes")
		}
		packs = append(packs, pack)
	}

	sortedPacks := sort.IntSlice(packs)
	// sort in descending order to start with the largest pack size
	sort.Sort(sort.Reverse(sortedPacks))
	for _, pack := range sortedPacks {
		isLastPackSize := pack == packs[len(packs)-1]
		rest := packItems % pack
		if rest > 0 {
			if packItems/pack >= 1 { // check if the pack size can be used
				packedItems[pack] = packItems / pack
			}
			packItems = rest
			if isLastPackSize {
				pi := packedItems[pack]
				if pi > 0 {
					packedItems[pack] += 1
					packItems = rest // update the remaining items
				} else {
					packedItems[pack] = 1
				}
			}
		} else {
			packedItems[pack] = packItems / pack
			break
		}
	}

	return packedItems, nil
}
