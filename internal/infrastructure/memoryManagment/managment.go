package memoryManagment

import (
	"fmt"
	"reflect"
)

type StaticManagment struct {
	storage map[string]interface{}
}

type IStaticManagment interface {
	Create(key string, value string) error
	Get(key string) (*interface{}, error)
}

var managment StaticManagment

func InstanceStaticManagment() *StaticManagment {
	if managment.storage == nil {
		managment = StaticManagment{
			storage: map[string]interface{}{},
		}
	}
	return &managment
}

func (s *StaticManagment) Create(key string, value interface{}) error {
	valueType := reflect.TypeOf(value).Kind()
	if valueType != reflect.Interface {
		return fmt.Errorf("invalid type %s\n ", valueType)
	}
	s.storage[key] = value
	return nil
}

func (s *StaticManagment) Get(key string) (interface{}, error) {
	value, exists := s.storage[key]
	if !exists {
		return nil, fmt.Errorf("the value not exists")
	}
	return value, nil
}
