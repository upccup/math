package main

import (
	"fmt"
)

/*
假设有四种面额的钱币，1 元、2 元、5 元和 10 元
有多少种排列方式使和为10元
*/

var rewards = []int{1, 2, 5, 10}

func main() {
	get(10, []int{})
}

func get(totalReward int, result []int) {
	if totalReward == 0 {
		fmt.Println(result)
		return
	} else if totalReward < 0 {
		return
	} else {
		for i := 0; i < len(rewards); i++ {
			newResult := append(result, rewards[i])
			get(totalReward-rewards[i], newResult)
		}
	}
}
