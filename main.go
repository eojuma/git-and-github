package main

import "fmt"

func main() {
	fmt.Println(firstword("         Juma is likes playing video games"))
	fmt.Println(firstword("         Juma is likes playing video games"))
	fmt.Println(firstword("  hello!       Juma is likes playing video games"))
}

func firstword(str string) string {
	// this will ensure that leading spaces are well taken care of
	// the iteration will occur until the character is not a space.
	i := 0
	for i < len(str) && str[i] == ' ' {
		i++
	}
//  at this point we now tell the iterator that we are beginnig to encounter character which are not spaces
	start := i

	for i < len(str) && str[i] != ' ' {
		i++
	}
	return str[start:i]
}
