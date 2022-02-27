package main

import "fmt"

func main() {
	path := "UDDDUDDUU"

	//pt := make(map[int]string)

	// for i, v := range path {
	// 	pt[i] = string(v)
	// }
	var steps, loc, count int32
	steps = int32(len(path))
	//pt := make(map[int32]string)

	for i := 0; int32(i) < steps; i++ {
		// pt[int32(i)] = string(path[int32(i)])
		// fmt.Println(pt)

		if string(path[i]) == "D" {
			loc--
		} else {
			loc++
		}

		if loc <= 0 && loc == 0 {

			count++

		}

	}

	fmt.Println(count)
}

// 		if pt[i] == "D" && pt[i]!="U" {

// 			for pt[i==0]; i<steps;i++{
// 			a++
// 			y=y
// 		}

// 		}
// 		if pt[i] == "U" && pt[i]!="D" {
// 			a=a
// 			y++
// 	}

// }
