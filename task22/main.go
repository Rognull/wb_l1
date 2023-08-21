package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)

	_, ok := a.SetString("100000000000000000", 10)
	if !ok {
		fmt.Println("Wrong data")
	}

	_, ok = b.SetString("99999999999999999", 10)
	if !ok {
		fmt.Println("Wrong data")
	}

	result := new(big.Int)
	fmt.Println(result)


	result.Add(a,b)

	fmt.Println(result)

	result.Sub(result,b)

	fmt.Println(result)

	result.Mul(a,b)

	fmt.Println(result)


	if b.Cmp(big.NewInt(0)) == 0 {
		fmt.Println("Divison by zero, select another operation")
	} else {
		result.Div(a,b)
		fmt.Println(result)
	}

}
