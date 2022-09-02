package core

var managment []int

type StaticManagment struct{}

func InstanceStaticManagment() *StaticManagment {
	if managment == nil {
		management = {}
	}
}
