package pattern

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

import (
	"fmt"
)

type Thing struct {
	state State
}

func (c *Thing) Request() {
	c.state.Handle()
}

func (c *Thing) SetState(state State) {
	c.state = state
}

type State interface {
	Handle()
}

type ConcreteStateA struct{}

func (s *ConcreteStateA) Handle() {
	fmt.Println("ConcreteStateA.Handle()")
}

type ConcreteStateB struct{}

func (s *ConcreteStateB) Handle() {
	fmt.Println("ConcreteStateB.Handle()")
}

// func main() {
// 	context := Context{new(ConcreteStateA)}
// 	context.Request()
// 	context.SetState(new(ConcreteStateB))
// 	context.Request()
// }
