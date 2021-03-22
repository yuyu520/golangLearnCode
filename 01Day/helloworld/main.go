package main

import "fmt"

func main() {
	sum := m1("sdsdwqefdewfedfewfe嗯嗯非非法人")
	fmt.Println(sum)
}

func m1(str string) int {
	index := 0
	for _, v := range str {
		if len(string(v)) >= 3 {
			index++
		}
	}
	return index
}
