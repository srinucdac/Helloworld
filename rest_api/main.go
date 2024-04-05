package main
import (
	"fmt"

	"github.com/srinucdac/rest_api/utils"

)

func main() {
	s := "Hello, World!"
	reversed := utils.ReverseString(s)
	fmt.Println(reversed) // Outputs: !dlroW ,olleH
}
