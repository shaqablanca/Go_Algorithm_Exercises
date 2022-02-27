




func migratoryBirds(arr []int32) int32 {
	// Write your code here
	dict := make(map[int32]int)

	for _, v := range arr {
		dict[v] += 1
	}

	lowerCount := arr[0]
	counter := dict[arr[0]]
	for k, v := range dict {
		if lowerCount < int32(k) && counter < v {
			counter = v
			lowerCount = int32(k)
		}
	}

	return lowerCount
}