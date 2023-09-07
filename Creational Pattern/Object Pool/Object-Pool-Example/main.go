package main

import (
	"fmt"
	Pool "test/labtopPool"
)

func main() {

	labtopPool := Pool.InitLabtopPool([]*Pool.Labtop{
		{Name: "MacBook Pro"},
		{Name: "MacBook Air"},
		{Name: "MacBook"},
	})

	labtopPool.AddToFashionList()
	labtopPool.AddToFashionList()

	fmt.Println("OnFashion: ", labtopPool.OnFashion[0].Name, " Length: ", len(labtopPool.OnFashion))
	fmt.Println("OldFashion: ", labtopPool.OldFashion[0].Name, " Length: ", len(labtopPool.OldFashion))
	fmt.Println("NotYetFashion: ", labtopPool.NotYetFashion[0].Name, " Length: ", len(labtopPool.NotYetFashion))

	labtopPool.InsertNotYetFashion(&Pool.Labtop{Name: "Khoa's Labtop"})
	labtopPool.AddToFashionList()

	fmt.Println("OnFashion: ", labtopPool.OnFashion[0].Name, " Length: ", len(labtopPool.OnFashion))
	fmt.Println("OldFashion: ", labtopPool.OldFashion[0].Name, " Length: ", len(labtopPool.OldFashion))
	fmt.Println("NotYetFashion: ", labtopPool.NotYetFashion[0].Name, " Length: ", len(labtopPool.NotYetFashion))

}
