package test0001

type thingy struct{}

func (x thingy) msg1() thingy { return x }
func (x thingy) msg2() thingy { return x }
func (x thingy) Foo() thingy  { return x }

func asq_query(x thingy) {
	//asq_start
	/***/ x. /***/ msg1().String()
	//asq_end
}
