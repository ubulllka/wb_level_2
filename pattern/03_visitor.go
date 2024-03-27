package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

type Classs1 struct{}

func (c *Classs1) do(v Visitor) {
	v.Visit1(c)
}

type Classs2 struct{}

func (c *Classs2) do(v Visitor) {
	v.Visit2(c)
}

type Classs3 struct{}

func (c *Classs3) do(v Visitor) {
	v.Visit3(c)
}

type Visitor interface {
	Visit1(c *Classs1)
	Visit2(c *Classs2)
	Visit3(c *Classs3)
}

type VisitorOne struct{}

func (o VisitorOne) Visit1(c *Classs1) {
	fmt.Println("One classs1")
}

func (o VisitorOne) Visit2(c *Classs2) {
	fmt.Println("One classs2")
}

func (o VisitorOne) Visit3(c *Classs3) {
	fmt.Println("One classs3")
}

type VisitorTwo struct{}

func (t VisitorTwo) Visit1(c *Classs1) {
	fmt.Println("Two classs1")
}

func (t VisitorTwo) Visit2(c *Classs2) {
	fmt.Println("Two classs2")
}

func (t VisitorTwo) Visit3(c *Classs3) {
	fmt.Println("Two classs3")
}
