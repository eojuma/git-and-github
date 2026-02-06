package main

import "fmt"

func main() {
	fmt.Println(firstword("         Juma is likes playing video games"))
	fmt.Println(firstword("         Juma is likes playing video games"))
	fmt.Println(firstword("  hello!       Juma is likes playing video games"))
}

func firstword(str string) string {
	i := 0
	for i < len(str) && str[i] == ' ' {
		i++
	}

	start := i

	for i < len(str) && str[i] != ' ' {
		i++
	}
	return str[start:i]
}
