package test0001

type thing struct{}

func (x thing) msg1() thing { return x }
func (x thing) msg2() thing { return x }

func asq_query(x thing) {
	//asq_start
	/***/x./***/msg1().String()
	//asq_end
}
