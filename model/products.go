package model

type Products struct {
	id    int
	name  string
	value float32
}

func (p *Products) GetProduct() (*int, *string, *float32) {
	return &p.id, &p.name, &p.value
}

func (p *Products) GetProductId() *int {
	return &p.id
}

func (p *Products) GetProductName() *string {
	return &p.name
}

func (p *Products) GetProductValue() *float32 {
	return &p.value
}

func (p *Products) SetName(name string) {
	p.name = name
}

func (p *Products) SetValue(value float32) {
	p.value = value
}


