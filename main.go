package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func grabDigits(i string, d string, c int) string {
	if c != 0 {
		var result string
		switch d {
		case "1":
			result = "0 "
		case "0":
			result = "00 "
		}
		var interim string
		for j := 0; j < c; j++ {
			interim = fmt.Sprintf("%s0", interim)
		}

		if i != "" {
			i = fmt.Sprintf("%s ", i)
		}
		result = fmt.Sprintf("%s%s%s", i, result, interim)
		return result
	}
	return i
}

func getBin(i string) string {
	var result string
	for _, v := range i {
		bar := fmt.Sprintf("%b", v)
		bar = strings.Trim(bar, "[")
		bar = strings.Trim(bar, "]")
		if len(bar) < 7 {
			bar = fmt.Sprintf("0%s ", bar)
		}

		result = fmt.Sprintf("%s%s", result, bar)
	}
	return result
}

func chuckNorrisEncode(i string) string {
	bar := getBin(i)

	var r string
	oneCount := 0
	zeroCount := 0

	for k := range bar {
		switch string(bar[k]) {
		case "0":
			zeroCount++
			if oneCount > 0 {
				r = grabDigits(r, "1", oneCount)
				oneCount = 0
			}
		case "1":
			oneCount++
			if zeroCount > 0 {
				r = grabDigits(r, "0", zeroCount)
				zeroCount = 0
			}
		default:
			r = fmt.Sprintf("%s", r)
		}
	}
	r = grabDigits(r, "0", zeroCount)
	r = grabDigits(r, "1", oneCount)
	return r
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	scanner.Scan()
	M := scanner.Text()

	fmt.Println(chuckNorrisEncode(M))

}
