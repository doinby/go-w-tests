package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	
	want := sum
	got := 4

	if want != got {
		t.Errorf("expected %d but got %d", want, got)
	}
}