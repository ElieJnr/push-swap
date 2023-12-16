package tools

func (s *Stack) Shift(item int) {
	temp := []int{item}
	temp = append(temp, s.items...)
	s.items = temp
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		panic("Empty stack")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack) push(item int) {
	s.items = append(s.items, item)
}
