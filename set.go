package set

// Set implementation
type Set struct {
	set map[interface{}]struct{}
}

// OrderedPair is used as a tuple for the cartesian product implementation
type OrderedPair struct {
	A interface{}
	B interface{}
}

// Exists returns wether a value is present in the set
func (i *Set) Exists(v interface{}) bool {
	_, ok := i.set[v]
	return ok
}

// NewSet returns a new initialized instance of Set
func NewSet() *Set {
	return &Set{make(map[interface{}]struct{})}
}

// Add inserts a value to the set
func (i *Set) Add(v interface{}) {
	i.set[v] = struct{}{}
}

// Delete removes a value from the set
func (i *Set) Delete(v interface{}) {
	delete(i.set, v)
}

// Len returns the number of elements in the set
func (i *Set) Len() int {
	return len(i.set)
}

// Union returns a new set with all the elements that the caller has in common
// with s
func (i *Set) Union(s *Set) *Set {
	r := NewSet()

	for k := range i.set {
		r.set[k] = struct{}{}
	}

	for k := range s.set {
		r.set[k] = struct{}{}
	}

	return r
}

// Intersect returns a new set that contains the common elements of the caller
// with s
func (i *Set) Intersect(s *Set) *Set {
	r := NewSet()

	for k := range i.set {
		if _, ok := s.set[k]; ok {
			r.Add(k)
		}
	}

	return r
}

// Complement returns a new set with the elements that exists in the caller but
// not in s
func (i *Set) Complement(s *Set) *Set {
	r := NewSet()

	for k := range i.set {
		if _, ok := s.set[k]; !ok {
			r.Add(k)
		}
	}

	return r
}

// CartesianProduct returns the set of all OrderedPairs int the form {A B} so that A
// is an element of the caller and B is an element of s
func (i *Set) CartesianProduct(s *Set) *Set {
	r := NewSet()

	for p := range i.set {
		for q := range s.set {
			r.Add(OrderedPair{p, q})
		}
	}

	return r
}
