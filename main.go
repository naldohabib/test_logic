package main

import (
	"fmt"
	"math/rand"
)

func main() {

	player := 4

	var hashMap = make(map[int][]int)
	var point = map[int]int{1: 0, 2: 0, 3: 0, 4: 0}
	var daddu = map[int]int{1: 4, 2: 4, 3: 4, 4: 4}

	for i := 1; i < player; i++ {
		arr := random(4)
		hashMap[i] = append(hashMap[i], arr...)
		fmt.Println(hashMap[i])
	}

	for i := 0; i < player; i++ {
		arrays := hashMap[i]
		temp := 1

		for j := 0; j < len(arrays); j++ {
			if arrays[j] == 6 {
				daddu[temp] = daddu[temp] - 1
				point[temp] = point[temp] + 1
			} else if arrays[j] == 1 {
				daddu[temp] = daddu[temp] - 1
				daddu[temp+1] = daddu[temp+1] + 1
			}
		}
	}
}

func random(n int) []int {
	temp := make([]int, 0)
	min := 1
	max := 6

	for i := 0; i < n; i++ {
		result := min + rand.Intn(max-min)
		temp = append(temp, result)
	}

	return temp
}