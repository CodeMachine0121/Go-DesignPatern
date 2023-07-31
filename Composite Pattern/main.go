package main

import (
	composites "composit-pattern/Composites"
	leaves "composit-pattern/Leaves"
)

func main() {

	composit := composites.NewComposite()
	leaf1 := leaves.NewLeaf(1)
	composit.Add(leaf1)

	leaf2 := leaves.NewLeaf(2000)
	composit.Add(leaf2)

	leaf3 := composites.NewComposite()
	composit.Add(leaf3)

	composit.Excute()
}
