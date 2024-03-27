package pattern

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

type Command interface {
	Execute()
}

type Light struct {
	IsOn bool
}

func (l *Light) TurnOn() {
	l.IsOn = true
}

func (l *Light) TurnOff() {
	l.IsOn = false
}

type ButtonOn struct {
	light *Light
}

func (o *ButtonOn) Execute() {
	o.light.TurnOn()
}

type ButtonOff struct {
	light *Light
}

func (o *ButtonOff) Execute() {
	o.light.TurnOff()
}
