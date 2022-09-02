package core

type HashValue struct{}

type IHashValue interface {
	Generate(key string) int64
}

func (h *HashValue) Decode(key string) int64 {
	return 1000
}
func (h *HashValue) Generate(key string) int64 {
	return 1000
}
