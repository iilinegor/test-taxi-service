package storage

func (s *Storage) read(key string) (value int, ok bool) {
	s.RLock()
	defer s.RUnlock()
	value, ok = s.Orders[key]
	return
}

func (s *Storage) inc(key string) {
	s.Lock()
	s.Orders[key]++
	s.Unlock()
}

func (s *Storage) write(key string, value int) {
	s.Lock()
	s.Orders[key] = value
	s.Unlock()
}

func (s *Storage) disableRequest(key string) {
	s.Lock()
	s.Orders[key] = s.Orders[key] * -1
	s.Unlock()
}
