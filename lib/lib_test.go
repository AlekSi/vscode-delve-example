package lib

import "testing"

func TestLess(t *testing.T) {
	versions := []string{
		"0.0.0", "0.0.1", "0.0.2",
		"0.1.0", "0.1.1", "0.1.2",
		"1.0.0", "1.0.1", "1.0.2",
		"1.1.0", "1.1.1", "1.1.2",
		"1.1.10", "1.10.1", "10.1.1", "10.10.10",
	}
	for i, vi := range versions {
		left, err := NewVersion(vi)
		if err != nil {
			t.Fatal(err)
		}

		for _, vj := range versions[:i] {
			right, err := NewVersion(vj)
			if err != nil {
				t.Fatal(err)
			}
			if left.Less(right) {
				t.Errorf("Expected %q >= %q", left, right)
			}
		}
		for _, vj := range versions[i+1:] {
			right, err := NewVersion(vj)
			if err != nil {
				t.Fatal(err)
			}
			if !left.Less(right) {
				t.Errorf("Expected %q < %q", left, right)
			}
		}
	}
}
