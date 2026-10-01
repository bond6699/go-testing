package even

import "testing"

func TestEven(t *testing.T) {
	data := map[int]bool{
		0:  true,
		1:  false,
		2:  true,
		3:  false,
		4:  true,
		5:  false,
		6:  true,
		7:  false,
		8:  true,
		9:  false,
		10: true,
		11: false,
		12: true,
		13: false,
		14: true,
		15: false,
	}

	for k, want := range data {
		got := IsEven(k)

		if got != want {
			t.Errorf("%d isEven == %t but receved %t", k, want, got)
		}
	}
}
