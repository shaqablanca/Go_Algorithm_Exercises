package main

import "fmt"

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {

	for i, v := range apples {
		apples[i] = v + a
	}

	for i, v := range oranges {
		oranges[i] = v + b
	}
	count1 := 0
	for _, v := range apples {
		if v >= s && v <= t {
			count1 = count1 + 1
		}

		count2 := 0
		for _, v := range oranges {
			if v >= s && v <= t {
				count2 = count2 + 1
			}

			
		}
		
		fmt.Println(count1)
		fmt.Println(count2)

	}

	// elma := []int32{-2, 6, 5, 10}
	// portakal := []int32{-5, 3, -7, -4}
	// var a, b, s, t int32

	// s = 7
	// t = 11

	// rline := make([]int32, t-s+1)
	// for i := range rline {
	// 	rline[i] = s + int32(i)
	// }

	// fmt.Println(s, t)
	// fmt.Println(len(elma), len(portakal))

	// a = 4
	// b = 12

	// fmt.Println(a, b)
	// fmt.Println(elma)
	// fmt.Println(portakal)

	// for i := 0; i < len(elma); i++ {
	// 	elma[i] = elma[i] + a
	// 	portakal[i] = portakal[i] + b

	// }

}

/*for _, j := range portakal {
		if j == 10 {

			fmt.Println("birtane geldi")
		}
	}

	//for _, h := range elma {
		if h == 10 {

			fmt.Println("birtane geldi", h)
		}
	}

}

*/
