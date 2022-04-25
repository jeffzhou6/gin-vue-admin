package utils

type Set struct {
	data map[interface{}]interface{}
}

func (s *Set) Add(v interface{}) {
	s.data[v] = struct{}{}
}

func (s *Set) Has(v interface{}) bool {
	_, ok := s.data[v]
	return ok
}

func (s *Set) Get() []interface{} {
	var keys = make([]interface{}, 0, s.Len())
	for k := range s.data {
		keys = append(keys, k)
	}
	return keys
}

func (s *Set) Del(v interface{}) {
	delete(s.data, v)
}

func (s *Set) Len() int {
	return len(s.data)
}

func NewSet() *Set {
	data := make(map[interface{}]interface{})
	return &Set{data}
}
