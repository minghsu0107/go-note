package main

import (
	"container/list"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type PersonQueue struct {
	list *list.List
}

func NewPersonQueue() *PersonQueue {
	return &PersonQueue{list.New()}
}

func (q *PersonQueue) Len() int {
	return q.list.Len()
}

func (q *PersonQueue) Offer(p *Person) {
	q.list.PushBack(p)
}

func (q *PersonQueue) Peek() *Person {
	if q.list.Len() == 0 {
		return nil
	}

	e := q.list.Remove(q.list.Front())
	return e.(*Person)
}

func main() {
	q := NewPersonQueue()

	q.Offer(&Person{"Irene", 12})
	q.Offer(&Person{"Justin", 45})
	q.Offer(&Person{"Monica", 42})

	for p := q.Peek(); p != nil; p = q.Peek() {
		fmt.Printf("姓名：%s\t年齡：%d\n", p.Name, p.Age)
	}
}
