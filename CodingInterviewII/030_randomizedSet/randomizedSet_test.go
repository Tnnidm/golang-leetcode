package randomizedSet

import "testing"

func Test_randomizeSet(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		obj := Constructor()
		ret1 := obj.Remove(0)
		ret2 := obj.Insert(0)
		ret3 := obj.GetRandom()
		ret4 := obj.Remove(0)
		if ret1 || !ret2 || ret3 != 0 || !ret4 {
			t.Fatal("Error!")
		}
	})
}
