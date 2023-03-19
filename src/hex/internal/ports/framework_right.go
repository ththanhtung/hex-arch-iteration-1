package ports

type DBPort interface {
	CloseDBConnection()
	AddToHistory(ans int32, operation string) error
}