package main

import "fmt"

func main() {
	s1 := make([]int, 5)
	fmt.Printf("the length of s1: %d\n", len(s1))
	fmt.Printf("the capacity of s1: %d\n", cap(s1))
	fmt.Printf("the value pf s1: %d\n", s1)
	s2 := make([]int, 5, 8)
	fmt.Printf("the length of s2: %d\n", len(s2))
	fmt.Printf("the capacity of s2: %d\n", cap(s2))
	fmt.Printf("the value pf s2: %d\n", s2)
	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6]
	fmt.Printf("the length of s4: %d\n", len(s4))
	fmt.Printf("the capacity of s4: %d\n", cap(s4))
	fmt.Printf("the value pf s4: %d\n", s4)
	s4 = s4[0:cap(s4)]
	fmt.Printf("the length of s4: %d\n", len(s4))
	fmt.Printf("the capacity of s4: %d\n", cap(s4))
	fmt.Printf("the value pf s4: %d\n", s4)
}
