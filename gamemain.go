/*
Author: roarandrock

Actions:
Provide a dynamic list. not always the same.
Look, Move, Use, Save

Content:
Location,  need a rough map

Models:
Locales, Items, Characters


Notes:
Can declare fmt.Println as a variable for ease!
var p = fmt.Println
p("Contains:  ", s.Contains("test", "es"))

And packages:
import s "strings"
s.Contains("test", "es")

Need to get better at placing vars in fmt
*/

package main

import (
	"RPG-Core/check"
	"RPG-Core/flow"
	"fmt"
	"log"
)

func main() {

	fmt.Println("Game start")
	log.Println(log.Ldate)
	//call starting screen, then pass player to mainflow. Go from there.
	cp, err := flow.Intro()
	check.Check(err)
	cp, err = flow.Mainflow(cp)
	check.Check(err)
	_, err = flow.Finale(cp)
	check.Check(err)
}
