package check

import (
	"fmt"
)

//Check for errors
func Check(e error) {
	if e != nil {
		fmt.Println("Error!")
		panic(e)
	}
}
