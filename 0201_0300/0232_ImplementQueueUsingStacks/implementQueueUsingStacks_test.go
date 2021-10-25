package implementQueueUsingStacks

import (
	"reflect"
	"testing"
)

func implementQueueUsingStacks() (int, int, bool) {
	myStack := Constructor()
	myStack.Push(1)
	myStack.Push(2)
	result1 := myStack.Peek()  // return 1
	result2 := myStack.Pop()   // return 1
	result3 := myStack.Empty() // return False
	return result1, result2, result3
}

func Test(t *testing.T) {
	t.Run(
		"test case 1",
		func(t *testing.T) {
			result1, result2, result3 := implementQueueUsingStacks()
			if !(reflect.DeepEqual(result1, 1) && reflect.DeepEqual(result2, 1) && reflect.DeepEqual(result3, false)) {
				t.Fatalf("expected: 1, 1, false. but got: %v, %v, %v", result1, result2, result3)
			}
		},
	)
}
