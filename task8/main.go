package main

import (
	"fmt"
	"strconv"
	// "math"

)

func setBit(num int64, pos int64, bitValue int) int64 {
	mask := int64(1) << pos
	if bitValue == 1 {
		// if pos == 63{
		// 	 return int64(math.Abs(float64(num)) * -1)  //  в Go, отрицательные числа представляются в формате дополнительного кода. В языке C это происходит позже, как я думал
		// }
		num |= mask
	} else if bitValue == 0 {
		// if pos ==63{
		// 	return int64(math.Abs(float64(num)))
		// }
		mask = ^mask
		num &= mask
	}
	return num
}

func main() {
	var num int64

	fmt.Print("Number: ")
	fmt.Scan(&num)

	binaryString := strconv.FormatInt(int64(num), 2)

	fmt.Printf("Your number in binary is : %s\n",binaryString)

	var pos int64
	fmt.Print("Position of bit: ")
	fmt.Scan(&pos)

	var bitValue int
	fmt.Print("Value of bit: ")
	fmt.Scan(&bitValue)

	if (pos < 0 || pos > 63) || (bitValue != 0 && bitValue != 1) {
		fmt.Println("wrong position or value of bit")
		return
	}

 
	result := setBit(num, pos, bitValue)
	binaryString = strconv.FormatInt(int64(result), 2)
	fmt.Printf("result: %d\n , in binary is %s", result,binaryString)
}