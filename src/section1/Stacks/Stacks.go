package main

type Stack struct{
	items []int
}

func (s *Stack) Push(i int){
	s.items = append(s.items, i)
}

func (s *Stack) Pop() int{
	size := len(s.items)
	if(size == 0){
		return -1
	}
	l := size - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

func main(){
	s := Stack{}

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	println(s.Pop())
	println(s.Pop())
	println(s.Pop())
	println(s.Pop())
	println(s.Pop())

}
