package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fls := make(map[int]float32)
	for i := 0; i < 100000; i++ {
		fls[i] = rand.Float32()
	}

	sum := func(fls map[int]float32) float32 {
		var sum float32
		for _, f := range fls {
			// goのmapのループは順番がランダム
			sum += f
		}
		return sum
	}

	// 同じfloat32のmapに対してsumを取ってるので、同じになるはず・・・？
	fmt.Println(sum(fls))
	fmt.Println(sum(fls))
	fmt.Println(sum(fls))
	fmt.Println(sum(fls))
	fmt.Println(sum(fls))

	// Output:
	// 49923.273
	// 49922.965
	// 49923.14
	// 49923.47
	// 49922.89

	// ならないですね〜😇
}
