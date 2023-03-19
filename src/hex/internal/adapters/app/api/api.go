package api

import "hexArch/internal/ports"

type Adapter struct {
	arth ports.ArithmeticPort
	db ports.DBPort
}

func NewAdapter(db ports.DBPort, arth ports.ArithmeticPort) *Adapter {
	return &Adapter{
		arth: arth,
		db: db,
	}
}

func (apia Adapter) GetAdd(a, b int32) (int32, error) {
	ans, err := apia.arth.Add(a, b)
	if err != nil {
		return 0, err
	}

	err = apia.db.AddToHistory(ans, "add")
	if err != nil {
		return 0, nil
	}

	return ans, nil
}
func (apia Adapter) GetSub(a, b int32) (int32, error) {
	ans, err := apia.arth.Sub(a, b)
	if err != nil {
		return 0, err
	}
	err = apia.db.AddToHistory(ans, "sub")
	if err != nil {
		return 0, nil
	}
	return ans, nil
}
func (apia Adapter) GetMul(a, b int32) (int32, error) {
	ans, err := apia.arth.Mul(a, b)
	if err != nil {
		return 0, err
	}
	err = apia.db.AddToHistory(ans, "mul")
	if err != nil {
		return 0, nil
	}
	return ans, nil
}
func (apia Adapter) GetDiv(a, b int32) (int32, error) {
	ans, err := apia.arth.Div(a, b)
	if err != nil {
		return 0, err
	}
	err = apia.db.AddToHistory(ans, "div")
	if err != nil {
		return 0, nil
	}
	return ans, nil
}
