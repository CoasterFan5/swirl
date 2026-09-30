package main

type A struct {
	name string
}

func Load() A {
	a := A{
		name: "hello",
	}
	return a
}
