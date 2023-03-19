package ports

type ArithmeticPort interface {
	Add(int32, int32) (int32, error)
	Sub(int32, int32) (int32, error)
	Mul(int32, int32) (int32, error)
	Div(int32, int32) (int32, error)
}