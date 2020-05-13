package storage

import (
	"math/rand"
	"sync"
	"time"
)

const (
	DefaultLimit    = 50
	TokenLength     = 2
	TokenCountLimit = 26 * 26 // 2 latin letters
)

// Storage contains taxi orders data
// and has high level functions around
type Storage struct {
	sync.RWMutex
	Orders      map[string]int
	OrdersIndex []string
	Limit       int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// InitStorage create storage instance and fill by unique orders
func InitStorage(limit int) (storage Storage) {
	if limit <= 0 {
		limit = DefaultLimit
	}

	storage.Limit = limit
	storage.OrdersIndex = make([]string, limit)
	storage.Orders = map[string]int{}

	for i := 0; i < limit; i++ {
		token := getUniqueToken()
		_, exist := storage.read(token)
		if exist {
			i--
		} else {
			storage.write(token, 1)
			storage.OrdersIndex[i] = token
		}
	}
	return
}
