package main

import (
	"fmt"
)

func main() {

	var n, p int = 5, 1
	var i, f, b int = 0, 0, 0

	// if n%2 == 0 && n/2 <= p {
	// 	for i = n + 1; i >= p; i = i - 2 {
	// 		b++
	// 	}
	// 	fmt.Println(b)
	// }

	// if n%2 == 0 && n/2 >= p {
	// 	for i = 0; i <= p; i = i + 2 {
	// 		b++
	// 	}
	// 	fmt.Println(b)
	// }

	// if n%2 == 1 && n/2 <= p && (n-p) > 1 {
	// 	for i = n + 1; i >= p; i = i - 2 {
	// 		f++
	// 	}
	// 	fmt.Println(f)
	// }

	// if n%2 == 1 && n/2 >= p {
	// 	for i = 0; i < p+1; i = i + 2 {
	// 		f++
	// 	}
	// 	fmt.Println(f)
	// }
	if n%2 == 0 {
		for i = 1; i < p; i = i + 2 {
			f++
		}

		//fmt.Println(f)

		for i = n; i > p; i = i - 2 {
			b++
		}
		//fmt.Println(b)

		if f < b {
			fmt.Println(f)
		} else {
			fmt.Println(b)
		}
	}
	if n%2 == 1 {
		for i = 1; i < p; i = i + 2 {
			f++
		}

		//fmt.Println(f)

		for i = n - 1; i > p; i = i - 2 {
			b++
		}
		//fmt.Println(b)

		if f < b {
			fmt.Println(f)
		} else {
			fmt.Println(b)
		}
	}
}

//////// HACKERRANK DONE

// func pageCount(n int32, p int32) int32 {
//     // Write your code here
//     var f,b,i int32 = 0,0,0
//     if n%2 == 0 {
//         for i = 1; i < p; i = i + 2 {
//             f++
//         }

//         //fmt.Println(f)

//         for i = n; i > p; i = i - 2 {
//             b++
//         }
//         //fmt.Println(b)

//         if f < b {
//             return f
//         } else {
//             return b
//         }

//     }
//     if n%2 == 1 {
//         for i = 1; i < p; i = i + 2 {
//             f++
//         }

//         //fmt.Println(f)

//         for i = n - 1; i > p; i = i - 2 {
//             b++
//         }
//         //fmt.Println(b)

//         if f < b {
//             return f
//         } else {
//             return b
//         }

//     }
//         return 0

// }
