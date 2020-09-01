package service

// DataProvider interface for data providers
type DataProvider interface {
	GetValue(key string) string
}
