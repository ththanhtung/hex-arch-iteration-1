package arithmetic

type Arithmetic struct{}

func NewArithmetic() *Arithmetic {
	return &Arithmetic{}
}

func (arth Arithmetic) Add(a, b int32) (int32, error) {
	return a + b, nil
}

func (arth Arithmetic) Sub(a, b int32) (int32, error) {
	return a - b, nil
}

func (arth Arithmetic) Mul(a, b int32) (int32, error) {
	return a * b, nil
}

func (arth Arithmetic) Div(a, b int32) (int32, error) {
	return a / b, nil
}
