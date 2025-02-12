package usecases

import (
	"errors"
	"github.com/rhuandantas/re-partners-home-test/internal/adapters/repository/cache"
	"strconv"
	"strings"
	"time"
)

type StorePackSize interface {
	Execute([]int) error
}

type storePackSize struct {
	cacheClient cache.MemcacheClient
}

func NewStorePackSize(cacheClient cache.MemcacheClient) StorePackSize {
	return &storePackSize{
		cacheClient: cacheClient,
	}
}

func (s *storePackSize) Execute(packs []int) error {
	var strNumbers []string
	for _, num := range packs {
		strNumbers = append(strNumbers, strconv.Itoa(num))
	}
	// Join them with commas
	result := strings.Join(strNumbers, ",")
	ttl := time.Hour * 24
	err := s.cacheClient.Set("pack-size", []byte(result), ttl)
	if err != nil {
		return errors.New("error storing pack sizes")
	}

	return nil
}
