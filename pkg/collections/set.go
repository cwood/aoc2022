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

// All checks if all values are this value
func (c Counter) All(v int) bool {

	var meta = Counter{}

	for _, v := range c {
		meta.Add(v)
	}

	if !meta.Contains(v) {
		return false
	}

	return meta[v] != len(c)-1
}

func (c Counter) Add(v any) {
	if c.Contains(v) {
		c[v] += 1
		return
	}

	c[v] = 1
}
