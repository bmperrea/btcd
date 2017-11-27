package btcec

import "testing"

func TestLessThanUint32(t *testing.T) {
	tests := []struct {
		x uint32
		y uint32
		a uint32
	}{
		{0, 1, 1},
		{2, 2, 0},
		{1 << 31, 1 << 31, 0},
		{17, 1 << 31, 1},
		{1 << 31, 0, 0},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		answer := LessThanUint32(test.x, test.y)
		if test.a != answer {
			t.Errorf("LessThanUint32 #%d wrong result\ngot: %v\n"+
				"want: %v", i, answer, test.a)
			continue
		}
	}
}

func TestIsZeroUint32(t *testing.T) {
	tests := []struct {
		x uint32
		a uint32
	}{
		{1, 0},
		{0, 1},
		{^uint32(0), 0},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		answer := IsZeroUint32(test.x)
		if test.a != answer {
			t.Errorf("IsZeroUint32 #%d wrong result\ngot: %v\n"+
				"want: %v", i, answer, test.a)
			continue
		}
	}
}

func TestNotZeroUint32(t *testing.T) {
	tests := []struct {
		x uint32
		a uint32
	}{
		{1, 1},
		{0, 0},
		{^uint32(0), 1},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		answer := NotZeroUint32(test.x)
		if test.a != answer {
			t.Errorf("NotZeroUint32 #%d wrong result\ngot: %v\n"+
				"want: %v", i, answer, test.a)
			continue
		}
	}
}
