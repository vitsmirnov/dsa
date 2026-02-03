package main

type MinMaxVal struct{ val, min, max int }
type MinMaxStack struct{ items []MinMaxVal }

func (s *MinMaxStack) Len() int { return len(s.items) }
func (s *MinMaxStack) Push(val int) {
	minVal, maxVal := val, val
	if s.Len() != 0 {
		top := s.Top()
		minVal = min(minVal, top.min)
		maxVal = max(maxVal, top.max)
	}
	s.items = append(s.items, MinMaxVal{val: val, min: minVal, max: maxVal})
}
func (s *MinMaxStack) Pop() int {
	res := s.Top().val
	s.items = s.items[:s.Len()-1]
	return res
}
func (s *MinMaxStack) Top() MinMaxVal { return s.items[s.Len()-1] }
func (s *MinMaxStack) Min() int       { return s.Top().min }
func (s *MinMaxStack) Max() int       { return s.Top().max }

type MinMaxQueue struct {
	front, back MinMaxStack
}

func (q *MinMaxQueue) Len() int    { return q.front.Len() + q.back.Len() }
func (q *MinMaxQueue) Put(val int) { q.back.Push(val) }
func (q *MinMaxQueue) Get() int {
	if q.front.Len() == 0 {
		q.migrate()
	}
	return q.front.Pop()
}
func (q *MinMaxQueue) Min() int {
	if q.front.Len() == 0 {
		return q.back.Min()
	} else if q.back.Len() == 0 {
		return q.front.Min()
	} else {
		return min(q.front.Min(), q.back.Min())
	}
}
func (q *MinMaxQueue) Max() int {
	if q.front.Len() == 0 {
		return q.back.Max()
	} else if q.back.Len() == 0 {
		return q.front.Max()
	} else {
		return max(q.front.Max(), q.back.Max())
	}
}
func (q *MinMaxQueue) migrate() {
	for q.back.Len() != 0 {
		q.front.Push(q.back.Pop())
	}
}
