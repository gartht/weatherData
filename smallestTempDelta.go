package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	file, error := os.Open("github.com/gartht/weatherData/weather.dat")
	checkError(error)
	defer file.Close()

	const valueSet = 8

	file.Seek(97, 0)

	b := make([]byte, valueSet)

	var smallestDelta int64

	evaluator := deltaEvaluator()

	for i := 0; i < 30; i++ {
		_, error := io.ReadAtLeast(file, b, 2)
		checkError(error)
		smallestDelta = evaluator(stringsToInts(strings.Split(strings.Trim(string(b), " "), "   ")))
		file.Seek(90-valueSet, 1)
	}
	fmt.Println(smallestDelta)
}

func deltaEvaluator() func(lowHi []int64) (smallestDelta int64) {
	var smallestDelta int64
	smallestDelta = 1000
	return func(lowHi []int64) int64 {
		low := lowHi[1]
		hi := lowHi[0]
		delta := hi - low
		if delta < smallestDelta {
			smallestDelta = delta
		}
		return smallestDelta
	}
}

func stringsToInts(hiAndLow []string) (tempSet []int64) {
	var err error
	tempSet = make([]int64, 2)
	tempSet[0], err = strconv.ParseInt(strings.Trim(hiAndLow[0], "*"), 10, 8)
	checkError(err)
	tempSet[1], err = strconv.ParseInt(strings.Trim(hiAndLow[1], " "), 10, 8)
	checkError(err)
	return
}

type highLow struct {
	high, low int
}
