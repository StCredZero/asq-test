package test0001

type thingy struct{}

func (x thingy) msg1() thingy   { return x }
func (x thingy) msg2() thingy   { return x }
func (x thingy) Foo() thingy    { return x }
func (x thingy) String() string { return "" }

func rando_code(y, z thingy) {
	y.Foo().String() // we should find this line
	z.Foo().String() // we should find this line
}
