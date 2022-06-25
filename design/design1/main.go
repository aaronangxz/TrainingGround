package main

import (
	"fmt"
	"github.com/aaronangxz/TrainingGround/common"
	"math/rand"
)

type MyStruct struct {
	Items   []int
	HashMap map[int]int //mapping of element:index
}

func NewMyStruct() *MyStruct {
	return &MyStruct{
		Items:   []int{},
		HashMap: make(map[int]int),
	}
}

func (m *MyStruct) Insert(item int) {
	_, ok := m.HashMap[item]
	if !ok {
		//Insert if element does not exist in map
		//insert the index into the map
		m.Items = append(m.Items, item)
		m.HashMap[item] = len(m.Items) - 1
	}
}

func (m *MyStruct) Remove(item int) {
	index, ok := m.HashMap[item]
	if ok {
		//remove from the map
		delete(m.HashMap, m.HashMap[item])
		//swap element to be deleted with the last element
		common.Swap(&m.Items[index], &m.Items[len(m.Items)-1])
		//remove the last element
		m.Items = m.Items[:len(m.Items)-1]
		//update the swapped element's index in map
		m.HashMap[len(m.Items)-1] = index
	}
}

func (m *MyStruct) Search(item int) int {
	index, ok := m.HashMap[item]
	if ok {
		return index
	}
	return -1
}

func (m *MyStruct) GetRandom() int {
	r := rand.Intn(len(m.Items))
	return m.Items[r]
}

func main() {
	s := NewMyStruct()
	s.Insert(10)
	s.Insert(20)
	s.Insert(30)
	s.Insert(40)
	fmt.Println(s.Search(30))
	s.Remove(20)
	s.Insert(50)
	fmt.Println(s.Search(50))
	fmt.Println(s.GetRandom())
}
