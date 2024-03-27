package pattern

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

type Person interface {
	GetName() string
}

const (
	ONE = "One"
	TWO = "Two"
)

type PersonOne struct{}

func NewPersonOne() *PersonOne {
	return &PersonOne{}
}

func (o *PersonOne) GetName() string {
	return "Hi, I'm One"
}

type PersonTwo struct{}

func NewPersonTwo() *PersonTwo {
	return &PersonTwo{}
}

func (t *PersonTwo) GetName() string {
	return "Hi, I'm Two"
}

func GetPerson(pType string) Person {
	switch pType {
	case ONE:
		return NewPersonOne()
	case TWO:
		return NewPersonTwo()
	default:
		return nil
	}
}
