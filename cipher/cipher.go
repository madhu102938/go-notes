package main

import (
	"errors"
)

var stackUnderflow error = errors.New("stack underflow")
var oneElement error =  errors.New("only one element")


type stack struct {
	slice []byte
	top   int
}

func New(n int) *stack {
	slice := make([]byte, n)
	st := &stack{slice, -1}
	return st
}

func (s *stack) Push(b byte) {
	s.top++
	s.slice[s.top] = b
}

func (s *stack) Top() (byte, error) {
	if s.top == -1 {
		return 0, stackUnderflow
	}
	return s.slice[s.top], nil
}

func (s *stack) SecondTop() (byte, error) {
	if s.top == 0 {
		return 0, oneElement
	}

	return s.slice[s.top-1], nil
}

func (s *stack) Pop() (byte, error) {
	if s.top == -1 {
		return 0, stackUnderflow
	}

	b := s.slice[s.top]
	s.top--
	return b, nil
}

func (s *stack) Empty() (bool) {
	if s.top == -1 {
		return true
	}
	return false
}

func reverseStringASCII(s string) string {
    bytes := []byte(s)
    n := len(bytes)
    for i := 0; i < n/2; i++ {
        bytes[i], bytes[n-1-i] = bytes[n-1-i], bytes[i]
    }
    return string(bytes)
}

func makeFancyString(s string) string {
	st := New(len(s))

	for i := 0; i < len(s); i++ {
		char := s[i]
		if st.Empty() {
			st.Push(char)
		}

		if topchar, err := st.Top(); err == nil {
			if topchar == char {
				if secondTopchar, err := st.SecondTop(); err == nil {
					if secondTopchar != char {
						st.Push(char)
					}
				}
			} else {
                st.Push(char)
            }
		} else if errors.Is(err, stackUnderflow) || errors.Is(err, oneElement) {
			st.Push(char)
		}
	}

	var ans string
	for !st.Empty() {
		char, err := st.Pop()
		if err != nil {
			ans += string(char)
		}
	}

	return reverseStringASCII(ans)
}

func main() {
	s := "aaabaaaa"
	ans := makeFancyString(s)
	println(ans)
}