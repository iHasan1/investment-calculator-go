package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount, expectedReturnRate, years float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue :=calcFutureValues(investmentAmount, expectedReturnRate, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}

func calcFutureValues (investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
