package main

import (
	"fmt"
	pools "object-pool/Pools"
)

func main() {
	num := func() interface{} {
		return 10.0
	}

	pool := pools.NewPool(num)
	object := pool.Acquire()

	fmt.Println(pool.Inuse, "... Pool Inuse")
	fmt.Println(pool.Available, "... Pool Available")

	pool.Release(object)
	fmt.Println(pool.Inuse, "... Pool Inuse")
	fmt.Println(pool.Available, "... Pool Available")
}
