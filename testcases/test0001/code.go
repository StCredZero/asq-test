package test0001

func rando_code(y, z thingy) {
	var foo, bar, baz thingy
	y.Foo().String() // we should find this line
	z.Foo().String() // we should find this line
}

func (x thing) Foo() {}