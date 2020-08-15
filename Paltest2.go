package main

import (
	"fmt"
	"math/big"
	"os"
	"time"
)

//Settings
const (
	StartNumber      string = "1"
	TopNumber        string = "1000"
	TopIterations    int    = 300
	RecordIterations int    = 290
)

func main() {
	t0 := time.Now()
	prNum, revNum, number := new(big.Int), new(big.Int), new(big.Int)
	var numOfIteration int
	topNum, _ := new(big.Int).SetString(TopNumber, 10)
	for number, _ = new(big.Int).SetString(StartNumber, 10); -1 == number.Cmp(topNum); number.Add(number, big.NewInt(1)) {
		prNum.Set(number)
		numOfIteration = 1
	loop:
		for numOfIteration <= TopIterations {
			revNum.Set(reverse(prNum))
			if 0 == prNum.Cmp(revNum) { //if prNum == revNum {
				if numOfIteration >= RecordIterations {
					fmt.Println("OMG!!!LUCKY")
					t1 := time.Now()
					fmt.Printf("Elapsed time: %v", t1.Sub(t0))
					os.Exit(0)
				}
				fmt.Printf("number:%v, prNum:%v, revNum:%v, numOfIteration:%v\n\n", number, prNum, revNum, numOfIteration)
				break loop
			} else {
				fmt.Printf("number:%v, prNum:%v, revNum:%v, numOfIteration:%v\n", number, prNum, revNum, numOfIteration)
				prNum.Add(prNum, revNum)
				numOfIteration++
			}
		}
		if numOfIteration >= TopIterations {
			fmt.Printf("OverTop[%v]-----------------\n\n", TopIterations)
		}
	}
	t1 := time.Now()
	fmt.Printf("NOnfound,Elapsed time: %v", t1.Sub(t0))
}

func reverse(number *big.Int) *big.Int {
	revNum := big.NewInt(0)
	ans := new(big.Int)
	ans.Set(number)
	for 0 != ans.Cmp(big.NewInt(0)) {
		revNum.Add(new(big.Int).Mul(revNum, big.NewInt(10)), new(big.Int).Mod(ans, big.NewInt(10)))
		ans.Div(ans, big.NewInt(10))
	}
	return revNum
}
