package main

import (
	"fmt"
)

func main() {
	cnt := 0
	ar := []int32{1, 3, 2, 6, 1, 2}
	var n, k int32 = 6, 3

	//for i,_ := range ar {

	for i := 0; i < int(n); i++ {

		for j := (i + 1); j < int(n); j++ {

			if (ar[i]+ar[j])%k == 0 {
				cnt++
			}
		}
	}
	fmt.Println(cnt)

}
