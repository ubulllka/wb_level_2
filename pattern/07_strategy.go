package pattern

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

type Strategy interface {
	Do(a, b int) int
}

type Plus struct{}

func (p *Plus) Do(a, b int) int {
	return a + b
}

type Minus struct{}

func (m Minus) Do(a, b int) int {
	return a - b
}

type SomeOperation struct {
	Strategy
}

func (o *SomeOperation) DoOperation(a, b int) int {
	return o.Strategy.Do(a, b)
}
