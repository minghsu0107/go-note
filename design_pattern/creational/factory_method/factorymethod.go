package factorymethod

type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

type OperatorBase struct {
	a, b int
}

func (o *OperatorBase) SetA(a int) {
	o.a = a
}

func (o *OperatorBase) SetB(b int) {
	o.b = b
}

type PlusOperator struct {
	*OperatorBase
}

type MinusOperator struct {
	*OperatorBase
}

func (o PlusOperator) Result() int {
	return o.a + o.b
}

func (o MinusOperator) Result() int {
	return o.a - o.b
}

type OperatorFactory interface {
	Create() Operator
}

type PlusOperatorFactory struct{}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

type MinusOperatorFactory struct{}

func (MinusOperatorFactory) Create() Operator {
	return &MinusOperator{
		OperatorBase: &OperatorBase{},
	}
}
