package generics


/*

type stackOfInts struct {
    values []int
}

func (s *stackOfInts) Push(value int) {
    s.values = append(s.values, value)
}

func (s *stackOfInts) IsEmpty() bool {
    return len(s.values) == 0
}

func (s *stackOfInts) Pop() (int, bool) {
    if s.IsEmpty() {
        return 0, false
    }

    index := len(s.values) - 1
    el := s.values[index]
    s.values = s.values[:index]

    return el, true
} 

type stackOfStrings struct {
    values []string
}

func (s *stackOfStrings) Push(value string) {
    s.values = append(s.values, value)
}

func (s *stackOfStrings) IsEmpty() bool {
    return len(s.values) == 0
}

func (s *stackOfStrings) Pop() (string, bool) {
    if s.IsEmpty() {
        return "", false
    }

    index := len(s.values) - 1
    el := s.values[index]
    s.values = s.values[:index]

    return el, true
} 
*/

/*
Stack implementation with interfaces but this is not good since  the compiler has no idea what type of interface is each function returning this can cause compiler error

type stackOfInts = Stack
type stackOfStrings = Stack

type Stack struct {
    values []interface{}
}


func (s *Stack) Push(value interface{}) {
    s.values = append(s.values, value)
}

func (s *Stack) IsEmpty() bool {
    return len(s.values) == 0
}

func (s *Stack) Pop() (interface{}, bool) {
    if s.IsEmpty() {
        var zero interface{}
        return  zero, false
    }

    index := len(s.values) - 1
    el := s.values[index]
    s.values = s.values[:index]
    return el, true
}
*/

type Stack[T any] struct {
    values []T
}

func (s *Stack[T]) Push(value T) {
    s.values = append(s.values, value)
}

func (s *Stack[T]) IsEmpty() bool {
    return len(s.values) == 0
}

func (s *Stack[T]) Pop() (T, bool) {
    if s.IsEmpty() {
        var zero T
        return zero, false
    }

    index := len(s.values) - 1
    el := s.values[index]
    s.values = s.values[:index]
    return el, true
}
