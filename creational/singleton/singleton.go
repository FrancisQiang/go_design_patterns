package singleton

import (
	"sync"
)

var (
	h *home
	once sync.Once
)

type home struct{

}

// @Description 懒汉
func GetInstance() *home {
	once.Do(func() {
		h = &home{}
	})
	return h
}

var (
	h1 = &home{}
)

func GetInstance1() *home {
	return h1
}
