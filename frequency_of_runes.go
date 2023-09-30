package main

import (
	"fmt"
	"sort"
	"strings"
)

var poem = `
those who do not feel this love
pulling them like a river
those who do not drink dawn
like a cup of spring water
or take in sunset like supper
those who do not want to change
let them sleep
`

func main() {
	// fmt.Println(strconv.Unquote(poem))
	poemChars := strings.Replace(poem, " ", "", -1)
	poemChars = strings.Replace(poemChars, "\n", "", -1)
	poemLen := float64(len(poemChars))

	freq := make(map[rune]int)
	for _, letter := range poemChars {
		freq[letter]++
	}

	chars := make([]rune, 0)
	for k := range freq {
		chars = append(chars, k)
	}

	var c1 rune
	var c2 rune
	sort.Slice(chars, func(i, j int) bool {
		c1, c2 = chars[i], chars[j]
		return freq[c1] > freq[c2]
	})

	var perc float64
	for _, char := range chars {
		perc = (float64(freq[rune(char)]) / poemLen) * 100.0
		fmt.Println(string(char) + ": " + fmt.Sprintf("%.1f", perc) + "%")
	}
	fmt.Println()
}
