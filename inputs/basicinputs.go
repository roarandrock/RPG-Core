package inputs

import "fmt"

//Numberinput takes in an expected range and returns the number
func Numberinput(l int) int {
	var i int
	_, err := fmt.Scan(&i)
	for err != nil || i > l || i <= 0 {
		if err != nil {
			fmt.Println("Input failed ", err, "\nTry again.")
		} else if i > l || i <= 0 {
			fmt.Println("Not a valid choice.\nTry again.")
		}
		_, err = fmt.Scan(&i)
	}
	return i
}

//Stringinput takes in a string with an expected number of words, returns the string
func Stringinput(l int) string {
	var s string
	i, err := fmt.Scan(&s)
	//doesn't really work, would have to try a different scan. But it works for now
	for err != nil || i > l {
		fmt.Println("no idea what I'm doing")
		if err != nil {
			fmt.Println("Input failed", err, "\nTry again.")
		} else if i > l {
			fmt.Println("Too many words.\nTry again.")
		}
		i, err = fmt.Scan(&s)
	}
	return s
}

//StringarrayInput takes in string array and outputs the player's selection
func StringarrayInput(sa []string) int {
	//can make a standalone function that takes a string array, makes a loop and a number input
	for i := 0; i < len(sa); i++ {
		fmt.Printf("%v. %s\n", i+1, sa[i])
	}
	d1 := Numberinput(len(sa))
	return d1
}
