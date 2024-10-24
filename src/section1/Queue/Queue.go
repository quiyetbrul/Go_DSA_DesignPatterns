package main

import "fmt"

type Queue struct {
	Items []int
}

func (q *Queue) Enqueue(i int) {
	q.Items = append(q.Items, i)
}

func (q *Queue) Dequeue() int {
	size := len(q.Items)
	if size == 0 {
		return -1
	}
	toRemove := q.Items[0]
	q.Items = q.Items[1:]
	return toRemove
}

func main() {
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())

}
