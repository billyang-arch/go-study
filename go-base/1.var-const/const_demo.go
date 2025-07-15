package main

// const同时声明多个常量时，如果省略了值则表示和上面一行的值相同
const (
	n1 = 100
	n2
	n3
)

const (
	m1 = iota //0
	m2        //1
	m3        //2
	m4        //3
)

const (
	k1 = iota //0
	k2        //1
	_
	k4 //3
)

const (
	l1 = iota //0
	l2 = 100  //100
	l3 = iota //2
	l4        //3
)
const n5 = iota //0

const (
	a, b = iota + 1, iota + 2 //1,2
	c, d                      //2,3
	e, f                      //3,4
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {

}
