// Package set provides a set implementation written purely in Go and with
// support for basic set operations
package set

// Set implementation
type Set map[interface{}]struct{}

// Exists returns whether a value is present in the set
func (i *Set) Exists(v interface{}) bool {
	_, ok := (*i)[v]
	return ok
}

// NewSet returns a new initialized instance of Set
func NewSet(elem ...interface{}) *Set {
	n := make(Set)

	for _, v := range elem {
		n[v] = struct{}{}
	}

	return &n
}

// Add inserts a value to the set
func (i *Set) Add(v interface{}) {
	(*i)[v] = struct{}{}
}

// Delete removes a value from the set
func (i *Set) Delete(v interface{}) {
	delete(*i, v)
}

// Len returns the number of elements in the set
func (i *Set) Len() int {
	return len(*i)
}

// Union returns a new set with all the elements the caller has in common
// with s
func (i *Set) Union(s *Set) *Set {
	r := NewSet()

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
func (i *Set) Intersect(s *Set) *Set {
	r := NewSet()

	for k := range *i {
		if _, ok := (*s)[k]; ok {
			r.Add(k)
		}
	}

	return r
}

// Complement returns a new set with the elements that exists in the caller but
// not in s
func (i *Set) Complement(s *Set) *Set {
	r := NewSet()

	for k := range *i {
		if _, ok := (*s)[k]; !ok {
			r.Add(k)
		}
	}

	return r
}

// CartesianProduct returns the set of all ordered pairs int the form {A B} so
// that A is an element of the caller and B is an element of s
func (i *Set) CartesianProduct(s *Set) *Set {
	r := NewSet()

	for p := range *i {
		for q := range *s {
			r.Add([2]interface{}{p, q})
		}
	}

	return r
}
