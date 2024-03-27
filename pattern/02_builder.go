package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

type Phone struct {
	Camera  string
	Memory  int
	Battery int
}

type PhoneBuilder interface {
	SetCamera()
	SetMemory()
	SetBattery()
	GetPhone() Phone
}

type CheapPhone struct {
	Phone
}

func NewCheapPhone() *CheapPhone {
	return &CheapPhone{}
}

func (c *CheapPhone) SetCamera() {
	c.Camera = "bad"
}

func (c *CheapPhone) SetMemory() {
	c.Memory = 100
}

func (c *CheapPhone) SetBattery() {
	c.Battery = 100
}

func (c *CheapPhone) GetPhone() Phone {
	return c.Phone
}

type ExpensivePhone struct {
	Phone
}

func NewExpensivePhone() *ExpensivePhone {
	return &ExpensivePhone{}
}

func (e *ExpensivePhone) SetCamera() {
	e.Camera = "good"
}

func (e *ExpensivePhone) SetMemory() {
	e.Memory = 2000
}

func (e *ExpensivePhone) SetBattery() {
	e.Battery = 2000
}

func (e *ExpensivePhone) GetPhone() Phone {
	return e.Phone
}

func GetPhone(pType string) PhoneBuilder {
	switch pType {
	case "cheap":
		return NewCheapPhone()
	case "expensive":
		return NewExpensivePhone()
	default:
		return nil
	}
}

type Director struct {
	PhoneBuilder
}

func NewDirector(phoneBuilder PhoneBuilder) *Director {
	return &Director{PhoneBuilder: phoneBuilder}
}

func (d *Director) BuildPhone() {
	d.SetCamera()
	d.SetMemory()
	d.SetBattery()
}

func (d *Director) GetPhone() Phone {
	return d.PhoneBuilder.GetPhone()

}
