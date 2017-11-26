package btcec

import "testing"

func TestSelectUint32(t *testing.T) {
	tests := []struct {
		v uint32
		x uint32
		y uint32
		a uint32
	}{
		{0, 1375, 17, 17},
		{1, 1375, 17, 1375},
		{0, 1 << 31, 131, 131},
		{1, 1 << 31, 131, 1 << 31},
		{0, 0, 131, 131},
		{1, 0, 131, 0},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		answer := SelectUint32(test.v, test.x, test.y)
		if test.a != answer {
			t.Errorf("SelectUint32 #%d wrong result\ngot: %v\n"+
				"want: %v", i, answer, test.a)
			continue
		}
	}
}

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
		{2 ^ 32, 0, 0},
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

func TestLessOrEqUint32(t *testing.T) {
	tests := []struct {
		x uint32
		y uint32
		a uint32
	}{
		{0, 1, 1},
		{2, 2, 1},
		{1 << 31, 1 << 31, 1},
		{17, 1 << 31, 1},
		{2 ^ 32, 0, 0},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		answer := LessOrEqUint32(test.x, test.y)
		if test.a != answer {
			t.Errorf("LessOrEqUint32 #%d wrong result\ngot: %v\n"+
				"want: %v", i, answer, test.a)
			continue
		}
	}
}

func TestEqUint32(t *testing.T) {
	tests := []struct {
		x uint32
		y uint32
		a uint32
	}{
		{0, 1, 0},
		{2, 2, 1},
		{1 << 31, 1 << 31, 1},
		{17, 1 << 31, 0},
		{2 ^ 32, 0, 0},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		answer := EqUint32(test.x, test.y)
		if test.a != answer {
			t.Errorf("LessOrEqUint32 #%d wrong result\ngot: %v\n"+
				"want: %v", i, answer, test.a)
			continue
		}
	}
}

func TestNotEqUint32(t *testing.T) {
	tests := []struct {
		x uint32
		y uint32
		a uint32
	}{
		{0, 1, 1},
		{2, 2, 0},
		{1 << 31, 1 << 31, 0},
		{17, 1 << 31, 1},
		{2 ^ 32, 0, 1},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		answer := NotEqUint32(test.x, test.y)
		if test.a != answer {
			t.Errorf("NotEqUint32 #%d wrong result\ngot: %v\n"+
				"want: %v", i, answer, test.a)
			continue
		}
	}
}

var MSBMask64 = int64(-2) * (int64(1)  >> 63)

func TestLessThanMSBInt64(t *testing.T) {
	tests := []struct {
		x int64
		y int64
		a int64
	}{
		{0, 1, MSBMask64},
		{2, 2, 0},
		{1 << 31, 1 << 31, 0},
		{17, 1 << 31, MSBMask64},
		{2 ^ 32, 0, 0},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		answer := LessThanMSBInt64(test.x, test.y) & MSBMask64
		if test.a != answer {
			t.Errorf("LessThanMSBInt64 #%d wrong result\ngot: %v\n"+
				"want: %v", i, answer, test.a)
			continue
		}
	}
}

func TestLessOrEqMSBInt64(t *testing.T) {
	tests := []struct {
		x int64
		y int64
		a int64
	}{
		{0, 1, MSBMask64},
		{2, 2, MSBMask64},
		{1 << 31, 1 << 31, MSBMask64},
		{17, 1 << 31, MSBMask64},
		{2 ^ 32, 0, 0},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		answer := LessOrEqMSBInt64(test.x, test.y) & MSBMask64
		if test.a != answer {
			t.Errorf("LessOrEqMSBInt64 #%d wrong result\ngot: %v\n"+
				"want: %v", i, answer, test.a)
			continue
		}
	}
}

func TestEqMSBInt64(t *testing.T) {
	tests := []struct {
		x int64
		y int64
		a int64
	}{
		{0, 1, 0},
		{2, 2, MSBMask64},
		{1 << 31, 1 << 31, MSBMask64},
		{17, 1 << 31, 0},
		{2 ^ 32, 0, 0},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		answer := EqMSBInt64(test.x, test.y) & MSBMask64
		if test.a != answer {
			t.Errorf("EqMSBInt64 #%d wrong result\ngot: %v\n"+
				"want: %v", i, answer, test.a)
			continue
		}
	}
}

func TestNotEqMSBInt64(t *testing.T) {
	tests := []struct {
		x int64
		y int64
		a int64
	}{
		{0, 1, MSBMask64},
		{2, 2, 0},
		{1 << 31, 1 << 31, 0},
		{17, 1 << 31, MSBMask64},
		{2 ^ 32, 0, MSBMask64},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		answer := NotEqMSBInt64(test.x, test.y) & MSBMask64
		if test.a != answer {
			t.Errorf("NotEqMSBInt64 #%d wrong result\ngot: %v\n"+
				"want: %v", i, answer, test.a)
			continue
		}
	}
}

func TestLessThanZeroMSBInt64(t *testing.T) {
	tests := []struct {
		x int64
		a int64
	}{
		{-1, MSBMask64},
		{0, 0},
		{1, 0},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		answer := LessThanZeroMSBInt64(test.x) & MSBMask64
		if test.a != answer {
			t.Errorf("LessThanZeroMSBInt64 #%d wrong result\ngot: %v\n"+
				"want: %v", i, answer, test.a)
			continue
		}
	}
}

func TestLessOrEqZeroMSBInt64(t *testing.T) {
	tests := []struct {
		x int64
		a int64
	}{
		{-1, MSBMask64},
		{0, MSBMask64},
		{1, 0},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		answer := LessOrEqZeroMSBInt64(test.x) & MSBMask64
		if test.a != answer {
			t.Errorf("LessOrEqZeroMSBInt64 #%d wrong result\ngot: %v\n"+
				"want: %v", i, answer, test.a)
			continue
		}
	}
}

func TestEqZeroMSBInt64(t *testing.T) {
	tests := []struct {
		x int64
		a int64
	}{
		{1, 0},
		{0, MSBMask64},
		{-1, 0},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		answer := EqZeroMSBInt64(test.x) & MSBMask64
		if test.a != answer {
			t.Errorf("EqZeroMSBInt64 #%d wrong result\ngot: %v\n"+
				"want: %v", i, answer, test.a)
			continue
		}
	}
}

func TestNotEqZeroMSBInt64(t *testing.T) {
	tests := []struct {
		x int64
		a int64
	}{
		{1, MSBMask64},
		{0, 0},
		{-1, MSBMask64},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		answer := NotEqZeroMSBInt64(test.x) & MSBMask64
		if test.a != answer {
			t.Errorf("NotEqZeroMSBInt64 #%d wrong result\ngot: %v\n"+
				"want: %v", i, answer, test.a)
			continue
		}
	}
}
