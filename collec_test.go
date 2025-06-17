package collection

import (
	"fmt"
	"testing"
)

type User struct {
	Name    string
	Age     int
	Address string
}

func TestName(t *testing.T) {
	fmt.Println(nc(1).LtE(2))
}

func TestFirst(t *testing.T) {
	users := []User{
		{Name: "John", Age: 25, Address: "123 Main St"},
		{Name: "Jane", Age: 30, Address: "456 Elm St"},
		{Name: "Bob", Age: 40, Address: "789 Oak St"},
	}

	c := NewCollect(users...)
	fmt.Println(c.First())
	str := "hello world"
	fmt.Println(NewCollect(str).First())

	m := []map[any]any{
		{"name": "张三", "age": 18, 1: "aaa"},
		{"name": "张三2", "age": 19, 2: "bbb"},
		{"name": "i love you", "age": 20, 2: "ccc"},
		{"name": "i 爱 you", "age": 21, 2: "ddd"},
	}
	fmt.Println(NewCollect(m...).Where("name", "like", "%you%").All())
}

func TestFilter(t *testing.T) {
	users := []User{
		{Name: "John", Age: 25, Address: "123 Main St"},
		{Name: "Jane", Age: 30, Address: "456 Elm St"},
		{Name: "Bob", Age: 40, Address: "789 Oak St"},
	}

	first := NewCollect(users...).Filter(func(item, value interface{}) bool {
		user, ok := value.(User)
		return ok && user.Name == "Jane"
	}).First()
	fmt.Println(first, first.Name)
}

func TestPop(t *testing.T) {
	users := []User{
		{Name: "John", Age: 25, Address: "123 Main St"},
		{Name: "Jane", Age: 30, Address: "456 Elm St"},
		{Name: "Bob", Age: 40, Address: "789 Oak St"},
	}
	c := NewCollect(users...)
	pop := c.Pop()
	fmt.Println(pop, c.All())
}

func TestShift(t *testing.T) {
	users := []User{
		{Name: "John", Age: 25, Address: "123 Main St"},
		{Name: "Jane", Age: 30, Address: "456 Elm St"},
		{Name: "Bob", Age: 40, Address: "789 Oak St"},
	}
	c := NewCollect(users...)
	pop := c.Shift()
	fmt.Println(pop, c.All())
}

func TestShuffle(t *testing.T) {
	users := []User{
		{Name: "John", Age: 25, Address: "123 Main St"},
		{Name: "Jane", Age: 30, Address: "456 Elm St"},
		{Name: "Bob", Age: 40, Address: "789 Oak St"},
	}
	c := NewCollect(users...)

	fmt.Println(c.Shuffle().All())

	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	all := NewCollect([]byte(str)...).Shuffle().All()
	fmt.Println(string(all[0:7]))
}

func TestSortBy(t *testing.T) {
	m := []map[any]any{
		{"name": "i 爱 you", "age": 21, 2: "ddd"},
		{"name": "i love you", "age": 20, 2: "ccc"},
		{"name": "张三", "age": 18, 1: "aaa"},
		{"name": "张三2", "age": 19, 2: "bbb"},
	}
	fmt.Println(NewCollect(m...).SortBy("age").All())
}

func TestSort(t *testing.T) {
	s1 := []string{"b", "a", "c", "d"}
	fmt.Println(NewCollect(s1...).Sort().All())

	s2 := []int{8, 6, 10, 1}
	fmt.Println(NewCollect(s2...).Sort().All())

	s3 := []uint{8, 6, 10, 1}
	fmt.Println(NewCollect(s3...).Sort().All())

	s4 := []int8{8, 6, 10, 1}
	fmt.Println(NewCollect(s4...).Sort().All())

	s5 := []uint8{8, 6, 10, 1}
	fmt.Println(NewCollect(s5...).Sort().All())

	s6 := []float32{8.1, 6.1, 10.2, 1.1, 1.11}
	fmt.Println(NewCollect(s6...).Sort().All())

	s7 := []float64{8.1, 6.1, 10.2, 1.1, 1.11}
	fmt.Println(NewCollect(s7...).Sort().All())
}

func TestSortDesc(t *testing.T) {
	s1 := []string{"b", "a", "c", "d"}
	fmt.Println(NewCollect(s1...).SortDesc().All())

	s2 := []int{8, 6, 10, 1}
	fmt.Println(NewCollect(s2...).SortDesc().All())

	s3 := []uint{8, 6, 10, 1}
	fmt.Println(NewCollect(s3...).SortDesc().All())

	s4 := []int8{8, 6, 10, 1}
	fmt.Println(NewCollect(s4...).SortDesc().All())

	s5 := []uint8{8, 6, 10, 1}
	fmt.Println(NewCollect(s5...).SortDesc().All())

	s6 := []float32{8.1, 6.1, 10.2, 1.1, 1.11}
	fmt.Println(NewCollect(s6...).SortDesc().All())

	s7 := []float64{8.1, 6.1, 10.2, 1.1, 1.11}
	fmt.Println(NewCollect(s7...).SortDesc().All())
}

func TestValues(t *testing.T) {
	m := []map[any]any{
		{"name": "i 爱 you", "age": 21, 2: "ddd"},
		{"name": "i love you", "age": 20, 2: "ccc"},
		{"name": "张三", "age": 18, 2: "aaa"},
		{"name": "张三2", "age": 19, 2: "bbb"},
	}
	fmt.Println(NewCollect(m...).Values(2))
}
