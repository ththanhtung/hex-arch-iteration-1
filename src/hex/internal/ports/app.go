package ports

type APIPort interface {
	GetAdd(int32, int32) (int32, error)
	GetSub(int32, int32) (int32, error)
	GetMul(int32, int32) (int32, error)
	GetDiv(int32, int32) (int32, error)
}
