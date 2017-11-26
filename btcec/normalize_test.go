package btcec

import (
	"testing"
	//"time"
	//"fmt"
	"math/rand"
)

//func TestNormalize2(t *testing.T){
//	nums := [][10]uint32{
//		[10]uint32{0x00000005, 0, 0, 0, 0, 0, 0, 0, 0, 0},
//		[10]uint32{0xffffffff, 0xffffffc0, 0xffffffc0, 0x3ffc0, 0, 0, 0, 0, 0, 0},
//		[10]uint32{0xfffffc30, 0xffffff86, 0xffffffc0, 0xffffffc0, 0xffffffc0, 0xffffffc0, 0xffffffc0, 0xffffffc0, 0xffffffc0, 0x3fffc0},
//		[10]uint32{0x03ffffff, 0x03fffeff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x0013fffff},
//		[10]uint32{0x03ffffff, 0x03fffaff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x0013fffff},
//		[10]uint32{0x03fffc30, 0x03ffffbf, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x07ffffff, 0x003fffff},
//		[10]uint32{0x148f6, 0x3ffffc0, 0x3ffffff, 0x3ffffff, 0x3ffffff, 0x3ffffff, 0x3ffffff, 0x3ffffff, 0x3ffffff, 0x000007},
//		[10]uint32{0x148f6, 0x0000013, 0x3ffffff, 0x3ffffff, 0x3ffffff, 0x3ffffff, 0x3ffffff, 0x3ffffff, 0x3ffffff, 0x000000},
//		[10]uint32{0x03fffc30, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03fffff0, 0x003fffff},
//	}
//
//	runs := 40000000
//
//	for i, num := range nums {
//		f := new(fieldVal)
//		f.n = num
//
//		t9 := f.n[9]
//		m := t9 >> fieldMSBBits
//		t9 = t9 & fieldMSBMask
//		t0 := f.n[0] + m*977
//		t1 := (t0 >> fieldBase) + f.n[1] + (m << 6)
//		t0 = t0 & fieldBaseMask
//		t2 := (t1 >> fieldBase) + f.n[2]
//		t1 = t1 & fieldBaseMask
//		t3 := (t2 >> fieldBase) + f.n[3]
//		t2 = t2 & fieldBaseMask
//		t4 := (t3 >> fieldBase) + f.n[4]
//		t3 = t3 & fieldBaseMask
//		t5 := (t4 >> fieldBase) + f.n[5]
//		t4 = t4 & fieldBaseMask
//		t6 := (t5 >> fieldBase) + f.n[6]
//		t5 = t5 & fieldBaseMask
//		t7 := (t6 >> fieldBase) + f.n[7]
//		t6 = t6 & fieldBaseMask
//		t8 := (t7 >> fieldBase) + f.n[8]
//		t7 = t7 & fieldBaseMask
//		t9 = (t8 >> fieldBase) + t9
//		t8 = t8 & fieldBaseMask
//
//		aa := t9 == fieldMSBMask
//		bb := t2&t3&t4&t5&t6&t7&t8 == fieldBaseMask
//		cc := ((t0+977)>>fieldBase + t1 + 64) > fieldBaseMask
//		dd := t9>>fieldMSBBits != 0
//
//		start := time.Now()
//		for i := 0; i < runs; i++ {
//			f.Normalize()
//		}
//		elapsed := float32(time.Since(start).Nanoseconds()) / float32(runs)
//		fmt.Printf("%v took %.1f ns/op for branch conditions %v %v %v %v \n", i, elapsed, bool2int(aa), bool2int(bb), bool2int(cc), bool2int(dd))
//	}
//}

func bool2int(b bool) uint32 {
	if b {
		return uint32(1)
	} else {
		return uint32(0)
	}
}



// generates an slice of somewhat random numbers (actually only num_rand to keep it quick)
func randUint32s(n int) []uint32 {
	arr := make([]uint32, n)
	for i := range arr {
		arr[i] = rand.Uint32()
	}
	return arr
}

func BenchmarkNormaliz1(b *testing.B) {
	x := randUint32s(b.N)
	f := new(fieldVal)
	f.n = [10]uint32{0xffffffff, 0xffffffc0, 0xfc0, 0, 0, 0, 0, 0, 0, 0}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.Normalize()
		f.n[1] = f.n[1] + x[i]
		f.n[9] = f.n[9] + x[i]
	}
	b.StopTimer()
	if (f.n[1] == 73) {
		print("whatever")
	}
}

func BenchmarkNormaliz2(b *testing.B) {
	x := randUint32s(b.N)
	f := new(fieldVal)
	f.n = [10]uint32{0xfffffc30, 0xffffff86, 0xffffffc0, 0xffffffc0, 0xffffffc0, 0xffffffc0, 0xffffffc0, 0xffffffc0, 0xffffffc0, 0x3fffc0}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.Normalize()
		f.n[0] = f.n[0] + x[i]
	}
	b.StopTimer()
	if (f.n[1] == 73) {
		print("whatever")
	}
}

func BenchmarkNormaliz3(b *testing.B) {
	f := new(fieldVal)
	f.n = [10]uint32{0xffffffff, 0xffffffc0, 0xfc0, 0, 0, 0, 0, 0, 0, 0}
	for i := 0; i < b.N; i++ {
		f.Normalize()
	}
}

func BenchmarkNormaliz4(b *testing.B) {
	f := new(fieldVal)
	f.n = [10]uint32{0xfffffc30, 0xffffff86, 0xffffffc0, 0xffffffc0, 0xffffffc0, 0xffffffc0, 0xffffffc0, 0xffffffc0, 0xffffffc0, 0x3fffc0}
	for i := 0; i < b.N; i++ {
		f.Normalize()
	}
}


/*

		nums := [][10]uint32{
			[10]uint32{0x00000005, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			[10]uint32{0x04000000, 0x0, 0, 0, 0, 0, 0, 0, 0, 0},
			[10]uint32{0x04000001, 0x0, 0, 0, 0, 0, 0, 0, 0, 0},
			[10]uint32{0xffffffff, 0x00, 0, 0, 0, 0, 0, 0, 0, 0},
			[10]uint32{0x04000000, 0x3f, 0, 0, 0, 0, 0, 0, 0, 0},
			[10]uint32{0x04000001, 0x3f, 0, 0, 0, 0, 0, 0, 0, 0},
			[10]uint32{0xffffffff, 0xffffffc0, 0xfc0, 0, 0, 0, 0, 0, 0, 0},
			[10]uint32{0x04000000, 0x03ffffff, 0x0fff, 0, 0, 0, 0, 0, 0, 0},
			[10]uint32{0x04000001, 0x03ffffff, 0x0fff, 0, 0, 0, 0, 0, 0, 0},
			[10]uint32{0xffffffff, 0xffffffc0, 0xffffffc0, 0x3ffc0, 0, 0, 0, 0, 0, 0},
			[10]uint32{0x04000000, 0x03ffffff, 0x03ffffff, 0x3ffff, 0, 0, 0, 0, 0, 0},
			[10]uint32{0xffffffff, 0xffffffc0, 0xffffffc0, 0xffffffc0, 0xffffc0, 0, 0, 0, 0, 0},
			[10]uint32{0x04000000, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x0ffffff, 0, 0, 0, 0, 0},
			[10]uint32{0xfffffc2f, 0xffffff80, 0xffffffc0, 0xffffffc0, 0xffffffc0, 0xffffffc0, 0xffffffc0, 0xffffffc0, 0xffffffc0, 0x3fffc0},
			[10]uint32{0xfffffc30, 0xffffff86, 0xffffffc0, 0xffffffc0, 0xffffffc0, 0xffffffc0, 0xffffffc0, 0xffffffc0, 0xffffffc0, 0x3fffc0},
			[10]uint32{0xfffffc2a, 0xffffff87, 0xffffffc0, 0xffffffc0, 0xffffffc0, 0xffffffc0, 0xffffffc0, 0xffffffc0, 0xffffffc0, 0x3fffc0},
			[10]uint32{0xffffffff, 0xffffffc0, 0xffffffc0, 0xffffffc0, 0xffffffc0, 0xffffffc0, 0xffffffc0, 0xffffffc0, 0xffffffc0, 0x3fffc0},
			[10]uint32{0x03fffc2f, 0x03ffffbf, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x003fffff},
			[10]uint32{0x03fffc30, 0x03ffffbf, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x003fffff},
			[10]uint32{0x03fffc2f, 0x03ffffc0, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x003fffff},
			[10]uint32{0x03ffffff, 0x03fffeff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x0013fffff},
			[10]uint32{0x03fffc30, 0x03ffffbf, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x07ffffff, 0x003fffff},
			[10]uint32{0x03fffc2f, 0x03ffffc0, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x3ffffff, 0x07ffffff, 0x003fffff},
			[10]uint32{0x03fffc30, 0x03ffffc0, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x03ffffff, 0x07ffffff, 0x003fffff},
			[10]uint32{0x148f6, 0x3ffffc0, 0x3ffffff, 0x3ffffff, 0x3ffffff, 0x3ffffff, 0x3ffffff, 0x3ffffff, 0x3ffffff, 0x3fffff},
			[10]uint32{0x148f6, 0x3ffffc0, 0x3ffffff, 0x3ffffff, 0x3ffffff, 0x3ffffff, 0x3ffffff, 0x3ffffff, 0x3ffffff, 0x000007},
		}

 */
