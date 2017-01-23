package check

import (
	"fmt"
)

//for errors
func Check(e error) {
	if e != nil {
		fmt.Println("Error!")
		panic(e)
	}
}
