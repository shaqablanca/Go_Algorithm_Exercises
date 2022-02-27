package main

func main() {

	//////////////////*** First Shot ***///////////////
	/*
		grades := []int32{23, 38, 67, 54, 75, 91}

		fmt.Println(len(grades))

		 for i, v := range grades {

		 	if v >= 38 && v%5 != 0 {
		 		if (5 - v%5) < 3 {
		 			grades[i] = grades[i] + (5 - v%5)
		 			fmt.Println(grades[i])
		 		}
		 		if (5 - v%5) >= 3 {
		 			fmt.Println(grades[i])
		 		}
		 	}
		 	if v%5 == 0 {
		 		fmt.Println(grades[i])
		 	}
		 	if v < 38 {
		 		fmt.Println(grades[i])
		 	}
	 }
	*/

	//////////////////*** Second Shot ***///////////////

	for i, v := range grades {

		if v >= 38 && v%5 != 0 {
			if (5 - v%5) < 3 {
				grades[i] = grades[i] + (5 - v%5)
			}
			if (5 - v%5) >= 3 {
			}
		}
		if v%5 == 0 {
		}
		if v < 38 {
		}
	}

}

///////////****** DONE ******//////////

/*

  for i, v := range grades {

        if v >= 38 && v%5 != 0 {
            if (5 - v%5) < 3 {
                grades[i] = grades[i] + (5 - v%5)
            }
            if (5 - v%5) >= 3 && v%5 == 0 && v < 38 {
            }
        }

    }

    return grades
*/
