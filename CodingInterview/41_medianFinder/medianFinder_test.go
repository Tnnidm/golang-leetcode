package MedianFinder

import "testing"

func TestMedianFinder(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		obj := Constructor()
		obj.AddNum(1)
		obj.AddNum(2)
		ret1 := obj.FindMedian()
		if ret1 != 1.5 {
			t.Fatalf("expect: %v, but got: %v", 1.5, ret1)
		}
		obj.AddNum(3)
		ret2 := obj.FindMedian()
		if ret2 != 1.5 {
			t.Fatalf("expect: %v, but got: %v", 2, ret2)
		}
	})
}
