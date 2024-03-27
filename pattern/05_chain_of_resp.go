package pattern

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

import "fmt"

// Handler - интерфейс обработчика
type Handler interface {
	SetNext(handler Handler)
	Handle(request string)
}

// BaseHandler - базовая реализация обработчика
type BaseHandler struct {
	nextHandler Handler
}

func (h *BaseHandler) SetNext(handler Handler) {
	h.nextHandler = handler
}

// HandlerA - конкретная реализация обработчика A
type HandlerA struct {
	BaseHandler
}

func (h *HandlerA) Handle(request string) {
	if request == "A" {
		fmt.Println("ConcreteHandlerA: обработан запрос", request)
	} else if h.nextHandler != nil {
		h.nextHandler.Handle(request)
	} else {
		fmt.Println("ConcreteHandlerA: невозможно обработать запрос", request)
	}
}

// HandlerB - конкретная реализация обработчика B
type HandlerB struct {
	BaseHandler
}

func (h *HandlerB) Handle(request string) {
	if request == "B" {
		fmt.Println("ConcreteHandlerB: обработан запрос", request)
	} else if h.nextHandler != nil {
		h.nextHandler.Handle(request)
	} else {
		fmt.Println("ConcreteHandlerB: невозможно обработать запрос", request)
	}
}
