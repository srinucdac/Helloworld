package main

import (
	"fmt"

	"github.com/srinucdac/utils"

)

func main() {
	s := "Hello, World!"
	reversed := utils.ReverseString(s)
	fmt.Println(reversed) // Outputs: !dlroW ,olleH
}
