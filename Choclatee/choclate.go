package main

import (
	"fmt"
)

func main() {
	s := []int32{2, 2, 1, 3, 2}
	var m, d = 2, 4
	cnt := 0

	for i := 0; i <= (len(s) - int(m)); i++ {

		slc := s[i : i+int(m)]

		tp := 0
		for _, v := range slc {

			tp = tp + int(v)

		}

		if tp == int(d) {

			cnt++

		}

	}
	fmt.Println(cnt)
}
