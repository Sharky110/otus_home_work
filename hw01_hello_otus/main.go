package main

import (
	"fmt"

	strUtil "golang.org/x/example/stringutil"
)

func main() {
	rtr := strUtil.Reverse("Hello, OTUS!")
	fmt.Println(rtr)
}
