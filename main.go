package example

import (
	"errors"
	"fmt"
)

func main() {
	Public()
}
func Public() {
	errors.New(fmt.Sprintf("%s", "New"))
}

func example() {
}
