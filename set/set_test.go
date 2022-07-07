package set

import "testing"

func TestNewSet(t *testing.T) {
	eles := []int{1, 2, 3, 4, 5}
	s := NewSafe(eles...)
	for _, e := range eles {
		if !s.Has(e) {
			t.Error("Should contain element")
		}
	}
}

func TestRemoveEle(t *testing.T) {
	eles := []int{1, 2, 3, 4, 5}
	s := NewSafe(eles...)
	s.Remove(1)
	if s.Has(1) {
		t.Error("Should not contain element")
	}
}

func TestClearSet(t *testing.T) {
	eles := []int{1, 2, 3, 4, 5}
	s := NewSafe(eles...)
	s.Clear()
	l := s.Len()
	if l != 0 {
		t.Errorf("Length not equal, want: %d, got: %d", 0, l)
	}
}

func TestSubset(t *testing.T) {
	s1 := NewSafe(1, 2, 3)
	s2 := NewSafe(1, 2)
	if s1.IsSubset(s2) {
		t.Error("Should not be subset")
	}
	if !s2.IsSubset(s1) {
		t.Error("Should be subset")
	}
	s1.Remove(3)
	if !s1.IsSubset(s2) || !s2.IsSubset(s1) {
		t.Error("Should be subset")
	}
}

func TestDisjoint(t *testing.T) {
	s1 := NewSafe(1, 2, 3)
	s2 := NewSafe(4, 5, 6)
	if !s1.IsDisjoint(s2) {
		t.Error("Should be disjoint")
	}
	s1.Add(4)
	if s1.IsDisjoint(s2) {
		t.Error("Should not be disjoint")
	}
}

func TestDiff(t *testing.T) {
	s1 := NewSafe(1, 2, 3)
	s2 := NewSafe(3, 4)
	want := NewSafe(1, 2)
	diff := s1.Diff(s2)
	if !diff.IsIdentical(want) {
		t.Errorf("Diff not equal, want: %s, got: %s", want, diff)
	}
}

func TestCopy(t *testing.T) {
	s := NewSafe(1, 2, 3)
	cpy := Copy[int](s)
	if !cpy.IsIdentical(s) {
		t.Errorf("Copy not equal, want: %s, got: %s", s, cpy)
	}
}

func TestUnion(t *testing.T) {
	s1 := NewSafe(1, 2, 3)
	s2 := NewSafe(3, 4, 5)
	want := NewSafe(1, 2, 3, 4, 5)
	union := Union[int](s1, s2)
	if !union.IsIdentical(want) {
		t.Errorf("Union not equal, want: %s, got: %s", want, union)
	}
}

func TestIntersection(t *testing.T) {
	s1 := NewSafe(1, 2, 3)
	s2 := NewSafe(3, 4)
	want := NewSafe(3)
	union := Intersection[int](s1, s2)
	if !union.IsIdentical(want) {
		t.Errorf("Intersection not equal, want: %s, got: %s", want, union)
	}
}
