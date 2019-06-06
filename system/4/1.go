package main

type A struct {
	next  *A
	value interface{}
}

type Box struct {
	first *A
	last  *A
}

func Head(va interface{}) {
	a := new(Box)
	value := a.first
	if value.next.value == va {

	}

}
