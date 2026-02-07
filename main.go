package main

import "fmt"

func main() {
	fmt.Println(firstword("         Juma is likes playing video games"))
	fmt.Println(firstword("         Juma is likes playing video games"))
	fmt.Println(firstword("  hello!       Juma is likes playing video games"))
}
// this function finds the fast word of any sring passed to it
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

	// this will print characters immediately after  the last space character to the last character before the next closest space character


	return str[start:i]
}
