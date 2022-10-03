package generics

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		assertEqual(t, 1, 1)
		assertNotEqual(t, 1, 2)
	})

	// what if we want assert on strings
	t.Run("asserting on strings", func(t *testing.T) {
		assertEqual(t, "hello", "hello")
		assertNotEqual(t, "hello", "grace")
	})
}

func TestStack(t *testing.T) {

	t.Run("integer stack", func(t *testing.T) {
		myStackOfInts := new(Stack[int])

		// check stack is empty
		AssertTrue(t, myStackOfInts.IsEmpty())

		// add a thing, then check it's not empty
		myStackOfInts.Push(123)
		AssertFalse(t, myStackOfInts.IsEmpty())

		// add another thing, pop it back again
		myStackOfInts.Push(456)
		value, _ := myStackOfInts.Pop()
		assertEqual(t, value, 456)
		value, _ = myStackOfInts.Pop()
		assertEqual(t, value, 123)

		AssertTrue(t, myStackOfInts.IsEmpty())
	})

	t.Run("string stack", func(t *testing.T) {
		mystackOfStrings := new(Stack[string])

		// check stack is empty
		AssertTrue(t, mystackOfStrings.IsEmpty())

		// add a thing, then check it's not empty
		mystackOfStrings.Push("123")
		AssertFalse(t, mystackOfStrings.IsEmpty())

		// add another thing, pop it back again
		mystackOfStrings.Push("456")
		value, _ := mystackOfStrings.Pop()
		assertEqual(t, value, "456")
		value, _ = mystackOfStrings.Pop()
		assertEqual(t, value, "123")
		AssertTrue(t, mystackOfStrings.IsEmpty())
	})

	t.Run("interface stack dx is horrid", func(t *testing.T) {
		// can get the numbers we put in as numbers, not untyped interface{}
		myStackOfInts := new(Stack[int])
		myStackOfInts.Push(1)
		myStackOfInts.Push(2)
		firstNum, _ := myStackOfInts.Pop()
		secondNum, _ := myStackOfInts.Pop()
		assertEqual(t, firstNum+secondNum, 3)
	})
}

func assertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want %v", got)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}
