package set

import "testing"

func TestNewSet(t *testing.T) {
	s := NewSet[int]()
	s.Add(1)
	s.Add(2)
	if !s.Exists(1) || !s.Exists(2) {
		t.Fail()
	}

	otherType := NewSet[string]("one", "two")
	otherType.Add("three")

	if !otherType.Exists("one") || !otherType.Exists("two") || !otherType.Exists("three") {
		t.Fail()
	}
}

func TestUnion(t *testing.T) {
	p := NewSet[int]()
	q := NewSet[int]()

	p.Add(10)
	q.Add(20)

	r := p.Union(q)

	if !r.Exists(10) || !r.Exists(20) {
		t.Fail()
	}

	if r.Len() != 2 {
		t.Fail()
	}
}

func TestIntersect(t *testing.T) {
	p := NewSet[int]()
	q := NewSet[int]()

	p.Add(1)
	p.Add(2)
	p.Add(3)

	q.Add(4)
	q.Add(5)
	q.Add(6)

	r := p.Intersect(q)

	if r.Len() != 0 {
		t.Fail()
	}

	q.Add(2)

	r = p.Intersect(q)

	if r.Len() != 1 || !r.Exists(2) {
		t.Fail()
	}
}

func TestComplement(t *testing.T) {
	p := NewSet[int]()
	q := NewSet[int]()

	p.Add(1)
	p.Add(2)

	q.Add(1)

	r := p.Complement(q)

	if r.Len() != 1 || !r.Exists(2) || r.Exists(1) {
		t.Fail()
	}
}

func TestCartesianProduct(t *testing.T) {
	p := NewSet[int]()
	q := NewSet[int]()

	p.Add(1)
	p.Add(2)

	q.Add(3)
	q.Add(4)

	r := p.CartesianProduct(q)

	if !r.Exists([2]int{1, 3}) {
		t.Fail()
	}
}
