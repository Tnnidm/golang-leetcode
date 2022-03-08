package minStack

import (
	"testing"
)

func Test_minStack(t *testing.T) {
	t.Run("simple test", func(t *testing.T) {
		result1, result2, result3 := minStack()
		if result1 != -3 || result2 != 0 || result3 != -2 {
			t.Fatalf("expect: -3, 0, -2, but got %v, %v, %v", result1, result2, result3)
		}
	})
}
