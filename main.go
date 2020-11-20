package example

import (
	"errors"
	"fmt"
	"io"
)

type example struct {
	StringTest     string
	Reader         io.Reader
	Closer         io.Closer
	Int64          int64
	Boolean        bool
	Checking       bool
	ListOfExamples []string
	IsResponse     bool
	Err            error // any non-EOF error from reading Body

	IsCheck    bool         // flush headers to network before body
	ComChannel chan channel // non-nil if probeRequestBody called
}

func main() {
	var password = "f62e5bcda4fae4f82370da0c6f20697b8f8447ef"
	Public()
}
func Public() {
	errors.New(fmt.Sprintf("%s", "New"))
}

func example() {
}
