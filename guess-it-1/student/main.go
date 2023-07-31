package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func av(nums []float64) float64 {
	sum := 0.0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	av := sum / float64(len(nums))
	return av
}

func vars(nums []float64, avg float64) float64 {
	sum := 0.0
	for _, v := range nums {
		sum += math.Pow((float64(v) - avg), 2.0)
	}
	return (sum / float64(len(nums)))
}

func std(nums float64) float64 {
	return math.Sqrt(nums)
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	var inputNums []float64
	var min, max int
	for reader.Scan() {
		temp1, err := strconv.ParseFloat(reader.Text(), 64)
		if err != nil {
			fmt.Println("Error: can not convert to number")
			return
		}
		if len(inputNums) > 2 && (temp1 > 2*av(inputNums) || temp1 < 0.1*av(inputNums)) {
			temp1 = av(inputNums)
		}
		inputNums = append(inputNums, temp1)
		min = int(math.Round(av(inputNums) - std(vars(inputNums, av(inputNums)))))
		max = int(math.Round(av(inputNums) + std(vars(inputNums, av(inputNums)))))
		fmt.Println(min, max)
	}
}
