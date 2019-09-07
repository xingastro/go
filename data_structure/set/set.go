package set

type Set struct {
	items map[string]bool
}

func (s *Set) Add (item string) {
	if s.items == nil {
		s.items = make(map[string]bool)
	}
	_, ok := s.items[item]
	if !ok {
		s.items[item] = true
	}
}

func (s *Set) Clear() {
	s.items = make(map[string]bool)
}

func (s *Set) Delete(item string){
	delete(s.items, item)
}

func (s *Set) Has (item string) bool {
	_, ok := s.items[item]
	return ok
}

func (s *Set) Items() []string {
	items := []string{}
	for i := range s.items {
		items = append(items, i)
	}
	return items
}

func (s *Set) Size() int {
	return len(s.items)
}