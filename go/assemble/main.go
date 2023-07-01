package main

import "math"

func main() {
	var a uint64 = math.MaxUint64
	println(a) //Output: 18446744073709551615

	var b = a + 1
	println(b) //Output: 0
}

//func main() {
//	a := math.MaxInt64
//	println(a) //Output: 9223372036854775807
//
//	b := a + 1
//	println(b) //Output: -9223372036854775808
//}
