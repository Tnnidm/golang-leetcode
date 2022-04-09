package sumSubRegion

import "testing"

func Test_sumSubRegion(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		matrix := [][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}}
		obj := Constructor(matrix)
		ret1 := obj.SumRegion(2, 1, 4, 3)
		ret2 := obj.SumRegion(1, 1, 2, 2)
		ret3 := obj.SumRegion(1, 2, 2, 4)
		if ret1 != 8 {
			t.Fatalf("expect 8 but got %v", ret1)
		}
		if ret2 != 11 {
			t.Fatalf("expect 11 but got %v", ret2)
		}
		if ret3 != 12 {
			t.Fatalf("expect 12 but got %v", ret3)
		}
	})
}
