package core

type StaticManagment struct {
	Storage map[string]int
}

type IStaticManagment interface {
}

var managment StaticManagment

func InstanceStaticManagment() *StaticManagment {
	if managment.Storage == nil {
		managment = StaticManagment{
			Storage: map[string]int{},
		}
	}
	return &managment
}

func (s *StaticManagment) Create(key string, value string) (int64, error) {
	return 0, nil
}

func (s *StaticManagment) Get(key string) (interface{}, error) {
	return nil, nil
}
