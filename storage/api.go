package storage

import (
	"encoding/json"
	"math/rand"
)

// GetOrder returns random order token and
// increments request count
func (s *Storage) GetOrder() string {
	index := rand.Intn(s.Limit)
	token := s.OrdersIndex[index]
	s.inc(token)
	return token
}

// GetList returns all order tokens with
// request count (negative value means disabled)
func (s *Storage) GetStats() (data []byte) {
	s.RLock()
	defer s.RUnlock()
	data, _ = json.Marshal(s.Orders)
	return
}

// Shake func disables random active item
// and adds new unique item
func (s *Storage) Shake() {
	for {
		index := rand.Intn(s.Limit)
		token := s.OrdersIndex[index]

		newToken := getUniqueToken()
		_, exist := s.read(newToken)
		if !exist {
			s.write(newToken, 1)
			s.disableRequest(token)

			for i, req := range s.OrdersIndex {
				if req == token {
					s.OrdersIndex[i] = newToken
					break
				}
			}
			break
		}
	}
}
