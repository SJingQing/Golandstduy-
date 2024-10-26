package main

import (
	"fmt"
)

// isPrime 函数判断一个整数是否为质数
func isPrime(number int) bool {
	if number <= 1 {
		return false // 0和1不是质数
	}
	if number <= 3 {
		return true // 2和3是质数
	}
	if number%2 == 0 || number%3 == 0 {
		return false // 排除能被2和3整除的数
	}
	for i := 5; i*i <= number; i += 6 {
		if number%i == 0 || number%(i+2) == 0 {
			return false // 排除能被其他数整除的数
		}
	}
	return true
}

func main() {
	// 测试几个数字
	testNumbers := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 13, 17, 19, 23, 29, 97}
	for _, number := range testNumbers {
		if isPrime(number) == true {
			fmt.Printf("%d 是质数\n", number)
		} else {
			fmt.Printf("%d 不是质数\n", number)
		}
	}
}
