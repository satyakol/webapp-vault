package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	fmt.Printf("%q\n", re.FindStringSubmatch("/doc/articles/wiki/"))
	fmt.Printf("%q\n", re.FindStringSubmatch("/view/bdg/"))
}
