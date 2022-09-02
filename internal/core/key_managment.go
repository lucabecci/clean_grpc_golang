package core

type KeyManagment struct {
	Managment string
}

type IKeyManagment interface {
	Get(key string) (map[string]interface{}, error)
}

func (k *KeyManagment) Get(key string) (map[string]interface{}, error) {
	return nil, nil
}
