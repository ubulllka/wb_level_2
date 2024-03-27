package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

type State interface {
	Handle()
}

type StateA struct{}

func (s *StateA) Handle() {
	fmt.Println("Состояние A")
}

type StateB struct{}

func (s *StateB) Handle() {
	fmt.Println("Состояние B")
}

type Context struct {
	state State
}

func (c *Context) SetState(state State) {
	c.state = state
}

func (c *Context) Request() {
	c.state.Handle()
}
