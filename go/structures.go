package main

type Cords struct {
	x int16
	y int16
}

type Stack []Cords

func (s *Stack) Push(v Cords) {
	*s = append(*s, v)
}

func (s *Stack) Pop() Cords {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}
