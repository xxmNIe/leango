package main

import "fmt"

func findPoisonedDuration(timeSeries []int, duration int) int {
	cnt := 0
	for i := 0; i < len(timeSeries); i++ {

		if i == len(timeSeries)-1 {
			cnt += duration
			break
		}

		if timeSeries[i]+duration-1 < timeSeries[i+1] {
			cnt += duration
		} else {
			cnt += (timeSeries[i+1] - timeSeries[i])
		}
	}

	return cnt
}

func main() {
	fmt.Println(findPoisonedDuration([]int{1}, 2))
}
