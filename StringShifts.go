package main

import "fmt"

func leftShift(a []rune, shift int32) []rune {
	for i := 0; int32(i) < shift; i++ {
		a = append(a[1:len(a)], a[0])
	}
	return a
}

func rightShift(a []rune, shift int32) []rune {
	for i := 0; int32(i) < shift; i++ {
		a = append([]rune{a[len(a)-1]}, a[0:len(a)-1]...)
	}
	return a
}

func shift(s string, leftShift int32, rightShift int32) string {
	shifts := leftShift - rightShift
	a := []rune(s)
	
	if(rotation > 0){
		for i := 0; int32(i) < shift; i++ {
			a = append(a[1:len(a)], a[0])
		}
	} else {
		for ; shift < 0; shift++ {
			a = append([]rune{a[len(a)-1]}, a[0:len(a)-1]...)
		}
	}
	return string(a)
}

func main() {
	var s string = "abcde"
	a := []rune(s)
	fmt.Printf("input: %s\n", s)
	fmt.Printf("Left Output: %s\n", string(leftShift(a, 1)))
	fmt.Printf("Right Output: %s\n", string(rightShift(a, 1)))
	fmt.Printf("Output: %s\n", shift(s, 1, 2))
}
