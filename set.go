// Package set provides a set implementation written purely in Go and with
// support for basic set operations
package set

// Set implementation
type Set[T comparable] map[T]struct{}

// Exists returns whether a value is present in the set
func (i *Set[T]) Exists(v T) bool {
	_, ok := (*i)[v]
	return ok
}

// NewSet returns a new initialized instance of Set
func NewSet[T comparable](elem ...T) *Set[T] {
	n := make(Set[T])

	for _, v := range elem {
		n[v] = struct{}{}
	}

	return &n
}

// Add inserts a value to the set
func (i *Set[T]) Add(v T) {
	(*i)[v] = struct{}{}
}

// Delete removes a value from the set
func (i *Set[T]) Delete(v T) {
	delete(*i, v)
}

// Len returns the number of elements in the set
func (i *Set[T]) Len() int {
	return len(*i)
}

// Union returns a new set with all the elements from both the caller and s
func (i *Set[T]) Union(s *Set[T]) *Set[T] {
	r := NewSet[T]()

	for k := range *i {
		(*r)[k] = struct{}{}
	}

	for k := range *s {
		(*r)[k] = struct{}{}
	}

	return r
}

// Intersect returns a new set that contains the common elements of the caller
// with s
func (i *Set[T]) Intersect(s *Set[T]) *Set[T] {
	r := NewSet[T]()

	for k := range *i {
		if _, ok := (*s)[k]; ok {
			r.Add(k)
		}
	}

	return r
}

// Complement returns a new set with the elements that exists in the caller but
// not in s
func (i *Set[T]) Complement(s *Set[T]) *Set[T] {
	r := NewSet[T]()

	for k := range *i {
		if _, ok := (*s)[k]; !ok {
			r.Add(k)
		}
	}

	return r
}

// CartesianProduct returns the set of all ordered pairs int the form {A B} so
// that A is an element of the caller and B is an element of s
func (i *Set[T]) CartesianProduct(s *Set[T]) *Set[interface{}] {
	// can't instantiate generic array type [2]T
	r := NewSet[interface{}]()

	for p := range *i {
		for q := range *s {
			r.Add([2]T{p, q})
		}
	}

	return r
}
