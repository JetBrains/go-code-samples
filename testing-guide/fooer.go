package main

import "strconv"

// Fooer If the number is divisible by 3, write "Foo" otherwise, the number
func Fooer(input int) string {

	isfoo := (input % 3) == 0

	if isfoo {
		return "Foo"
	}

	return strconv.Itoa(input)
}
