package collections

type Set map[any]struct{}

func (s Set) Contains(v any) bool {
	_, ok := s[v]
	return ok
}

func (s Set) Add(v any) {
	s[v] = struct{}{}
}

type Counter map[any]int

func (c Counter) Contains(v any) bool {
	_, ok := c[v]
	return ok
}

func (c Counter) Add(v any) {
	if c.Contains(v) {
		c[v] += 1
		return
	}

	c[v] = 1
}
