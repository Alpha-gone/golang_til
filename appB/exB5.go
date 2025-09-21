package main

import "fmt"

func main() {
	fac := NewFlyweightFactory(1000)

	for i := 0; i < 1000; i++ {
		obj := fac.Create()
		obj.Somedata = "Somedata"
		fac.Dispose(obj)
	}

	fmt.Println("AllocCnt:", fac.AllocCnt)
}

type FlyWeightFactory struct {
	pool     []*FlyWeight
	AllocCnt int
}

func NewFlyweightFactory(initSize int) *FlyWeightFactory {
	return &FlyWeightFactory{pool: make([]*FlyWeight, 0, initSize)}
}

func (fac *FlyWeightFactory) Create() *FlyWeight {
	var obj *FlyWeight

	if len(fac.pool) > 0 {
		obj, fac.pool = fac.pool[len(fac.pool)-1], fac.pool[:len(fac.pool)-1]
		obj.Reuse()
	} else {
		obj = &FlyWeight{}
		fac.AllocCnt++
	}

	return obj
}

func (fac *FlyWeightFactory) Dispose(obj *FlyWeight) {
	obj.Dispose()
	fac.pool = append(fac.pool, obj)
}

type FlyWeight struct {
	Somedata   string
	isDisposed bool
}

func (f *FlyWeight) Reuse() {
	f.isDisposed = false
}

func (f *FlyWeight) Dispose() {
	f.isDisposed = true
}

func (f *FlyWeight) IsDisposed() bool {
	return f.isDisposed
}
