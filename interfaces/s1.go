package main

import (
	"fmt"
)


type RuntimeObj interface {
	Test()
}


type Creater interface {
	New() int
	Create(age int) int
}

type NamedCreater interface {
	New() int
	Create(age int, name string) int
}


type Etcd struct {
	name string
}

type Service struct {
	age int
	name string
}


func (e Etcd) Test() {
	fmt.Println("Test Etcd")
}

func (e Etcd) New() int {
	return 10
}

func (e Etcd) Create(age int) int {
	return 20
}



func (e Service) Test() {
	fmt.Println("Test Service")
}

func (e Service) New() int {
	return 1000
}

func (e Service) Create(age int, name string) int {
	return 2000
}



func test(obj RuntimeObj) {
	a, b := obj.(Creater)
	fmt.Println(a)
	fmt.Println(b)

	c, d := obj.(NamedCreater)
	fmt.Println(c)
	fmt.Println(d)
}


func main() {
	fmt.Println("start")

	fmt.Println("Test Etcd")
	e := Etcd {"my etcd"}
	test(e)

	fmt.Println("Test Service")
	s := Service{100, "my service"}
	test(s)

	fmt.Println("end")
}
