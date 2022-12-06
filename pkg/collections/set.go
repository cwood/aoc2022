package collections

type Set map[any]struct{}

func (s Set) Contains(v any) bool {
	_, ok := s[v]
	return ok
}

func (s Set) Add(v any) {
	s[v] = struct{}{}
}
