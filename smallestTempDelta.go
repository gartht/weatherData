package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gartht/minDeltaUtil"
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

	reader := bufio.NewReader(file)

	deltaFinder := minDeltaUtil.Finder(1, 2, 0)

	outputString := ""

	//read past the first two lines (header and an empty row)
	reader.ReadString('\n')
	reader.ReadString('\n')

	for {
		line, error := reader.ReadString('\n')

		if error == nil {
			line = strings.Replace(line, "*", " ", -1)
			outputString = deltaFinder(strings.Fields(line))
			continue
		}

		if error == io.EOF {
			break
		}

		panic(error)
	}
	fmt.Println(outputString)
}
