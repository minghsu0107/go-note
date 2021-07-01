package main

import (
	"container/list"
	"fmt"
)

/*
type Element struct {
    Value interface{}
}

func (e *Element) Next() *Element
func (e *Element) Prev() *Element

func (l *List) InsertAfter(v interface{}, mark *Element) *Element
func (l *List) InsertBefore(v interface{}, mark *Element) *Element
func (l *List) Len() int
func (l *List) MoveToBack(e *Element)
func (l *List) MoveToFront(e *Element)
func (l *List) PushBack(v interface{}) *Element
func (l *List) PushBackList(other *List)
func (l *List) PushFront(v interface{}) *Element
func (l *List) PushFrontList(other *List)
func (l *List) Remove(e *Element) interface{}
*/

type Person struct {
	Name string
	Age  int
}

func printAll(lt *list.List) {
	for e := lt.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func printAllPerson(persons *list.List) {
	for e := persons.Front(); e != nil; e = e.Next() {
		p := e.Value.(*Person)
		fmt.Printf("姓名：%s\t年齡：%d\n", p.Name, p.Age)
	}
}

func main() {
	lt := list.New()
	for i := 1; i <= 10; i++ {
		lt.PushBack(i)
	}
	printAll(lt)

	persons := list.New()

	persons.PushBack(&Person{"Irene", 12})
	persons.PushBack(&Person{"Justin", 45})
	persons.PushFront(&Person{"Monica", 42})
	persons.InsertAfter(&Person{"Ming", 20}, persons.Back())
	printAllPerson(persons)
	fmt.Println(persons.Back().Value.(*Person).Name) // Ming
}
