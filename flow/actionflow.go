package flow

import (
	"RPG-Core/inputs"
	"RPG-Core/models"
	"fmt"
	"strconv"
)

func actielist(cp models.Player) string {
	cl := models.LocationGet(cp.Loc)
	alist := cl.Actions
	alist = append(alist, "Menu")
	for i := 0; i < len(alist); i++ {
		fmt.Printf("%v. %v\n", i+1, alist[i])
	}
	r1 := inputs.Numberinput(len(alist))
	s := alist[r1-1]
	return s
}

func actieSelector(act string, cp models.Player) (models.Player, error) {
	_, err := strconv.Atoi("-42")
	switch act {
	case "Look":
		//describe place, people, and items here
		loc := models.LocationGet(cp.Loc)
		fmt.Println("You are in the", loc.Name)
		scene := models.SceneryGet(cp.Loc)
		fmt.Println(scene)
		//people list
		cl, n := models.CharacterGetByLoc(cp.Loc)
		if n != 0 {
			fmt.Printf("You see ")
			for i := 0; i < n; i++ {
				il := cl[i]
				if i == 0 {
					fmt.Printf("%s", il.Name)
				} else {
					fmt.Printf(" and %s", il.Name)
				}
			}
			fmt.Printf("\n")
		}
		//items available
	case "Walk":
		//Go somewhere else casually
		loclist := models.TravelGet(cp.Loc)
		fmt.Println("You can walk to the:")
		for i := 0; i < len(loclist); i++ {
			iloc := models.LocationGet(loclist[i])
			if i == 0 {
				fmt.Printf("1. %s\n", iloc.Name)
			} else {
				fmt.Printf("%v. %s\n", i+1, iloc.Name)
			}
		}
		w1 := inputs.Numberinput(len(loclist))
		if w1 == cp.Loc {
			fmt.Println("You stay still")
		} else {
			dt := models.TravelTime(cp.Loc, w1)
			cp.Loc = w1
			models.UpdateTime(dt)
			fmt.Println("You travel. It takes roughly one hour.")
		}
	case "Talk":
		//people list
		cl, n := models.CharacterGetByLoc(cp.Loc)
		if n != 0 {
			fmt.Println("Whom to converse with?")
			for i := 0; i < n; i++ {
				il := cl[i]
				if i == 0 {
					fmt.Printf("1. %s\n", il.Name)
				} else {
					fmt.Printf("%v. %s\n", i+1, il.Name)
				}
			}
			r1 := inputs.Numberinput(n)
			cc := cl[r1-1]
			converser(cc)
		} else {
			fmt.Println("No one to speak with.")
		}
	case "Run":
		//Go somewhere else fast!
		fmt.Println("Nowhere to run")
	case "Menu":
		//save, inventory, stats, quit
		fmt.Println("1. Save\n2. Quit")
		m1 := inputs.Numberinput(2)
		switch m1 {
		case 1:
			fmt.Println("This is where you would save")
		case 2:
			fmt.Println("Bye bye")
			cp.Cont = false
		}
	default:
		fmt.Println("Odd. This action does nothing.")
	}
	return cp, err
}

//Convo is base struct
type convo struct {
	Character    models.Character
	depth        int
	qa           string
	stilltalking bool
}

func converser(cChar models.Character) {
	//need to separate this out. Initial responses and character specific, player responses
	//only talk to a character once? instead of just exhausting all the options? more real
	playerH := []string{"Hi. Fucker.", "Hello.", "Good to see you.", "Damn good to see you."}
	playerQ := []string{"How are you?", "Want do drugs?"}
	//playerA := []string{"I'm good", "Up and down", "Feeling shit", "Don't want to talk about it. Bye."}
	playerGB := []string{"Bye"}
	//defaults, starting point
	cc := convo{}
	cc.Character = cChar
	cc.stilltalking = true
	cc.depth = 3 //can pull from character later
	cc.qa = "H"
	var cr1 string

	for cc.stilltalking == true {
		if cc.qa == "H" {
			d1 := inputs.StringarrayInput(playerH)
			cr1, cc = characterA(d1, cc)
		} else if cc.qa == "CA" {
			options := append(playerQ, playerGB...)
			d1 := inputs.StringarrayInput(options) //need to add other options here, like GB. Seems to work with append
			cc.qa = "PQ"
			cr1, cc = characterA(d1, cc)
		} else if cc.qa == "GB" { //don't need, redundant
			inputs.StringarrayInput(playerGB)
			cc.stilltalking = false
		} else {
			cc.stilltalking = false
		}
		fmt.Println(cr1, cc.qa)
	}
}

//Need a table for each character, responses and questions
//Make it's own file, package maybe
//qa = Hello, Player Question, Character Question, Player Answer, Character Answer, Character Tangent, Awkward Silence, Goodbye
// qa = H, PQ, CQ, PA, CA, CT, AS, GB
//depth 1-5 = Hostile, Dislike, Neutral, Like, Love

func characterA(d1 int, cc convo) (string, convo) {
	var r1 string
	i := 1
	//need a table per character and a function to grab their relevant dialog
	//going to cheat
	//need to break this up into some useful functions. and codify it to make it easier for multiple characters.
	//like each response needs at least five options. They can all be the same, but it means the same response regardless of depth.
	//can add counters too, for each character and within each convo
	characterH := []string{"Fuck you too.", "Hello", "Good to see you too."}
	characterA1 := []string{"I'm good.", "Actually, I'm not so good."} //need one stack for each possible PQ
	characterA2 := []string{"No dude, of course not.", "Actually, that sounds great."}
	characterGB := [][]string{{"Later"}, {"Bye"}}
	switch cc.qa {
	case "H":
		if d1 == 1 {
			i = 0
			cc.depth = cc.depth - 1 //need to set only to 1-5, maybe a seperate function
		} else if d1 > 2 {
			i = 2
		}
		r1 = characterH[i]
		cc.qa = "CA"
	case "PQ":
		if cc.depth < 4 {
			i = 0
		} else {
			i = 1
		}
		switch d1 { //can remove this if combine characterA arrrays?
		case 1:
			r1 = characterA1[i]
			cc.qa = "CA"
		case 2:
			r1 = characterA2[i]
			cc.qa = "CA"
		case 3:
			r1 = characterGB[0][0]
			cc.qa = "GB"
			cc.stilltalking = false
		}

	}
	r1 = "\"" + r1 + "\""
	return r1, cc
}

/* First attempt at convo. No longer needed

func characterQuestions(d1 int, depth int) (string, string, int, bool) {
	resp := "I've got nothing to say. Oddly."
	rqa := "a"
	newdepth := depth
	stilltalking := true
	switch d1 {
	case 1:
		switch depth {
		case 1:
			resp = "Good to hear man."
			rqa = "a"
		case 2:
		}
	case 2:
		switch depth {
		case 1:
			resp = "I like your honesty."
			newdepth = 2
		case 2:
		}
	case 3:
		switch depth {
		case 1:
			resp = "Bummer dude."
		case 2:
			resp = "Want to do some drugs?"
		}
	case 4:
		switch depth {
		case 1:
			resp = "Your life dude."
			stilltalking = false
		case 2:
			resp = "Sorry to hear that man."
			stilltalking = false
			newdepth = 1
		}
	}
	return resp, rqa, newdepth, stilltalking
}

func cResponse(d1 int, depth int) (string, string, int, bool) {
	resp := "I've got nothing to say. Oddly."
	rqa := "a"
	newdepth := depth
	stilltalking := true
	switch d1 {
	case 1:
		switch depth {
		case 1:
			resp = "How are you doing?"
			rqa = "q"
		case 2:
		}
	case 2:
		switch depth {
		case 1:
			resp = "I'm good bro."
			newdepth = 2
		case 2:
		}
	case 3:
		switch depth {
		case 1:
			resp = "No little man. Not cool."
		case 2:
			resp = "Why the fuck not!"
		}
	case 4:
		switch depth {
		case 1:
			resp = "Bye dude."
			stilltalking = false
		case 2:
			resp = "Later bro."
			stilltalking = false
		}
	}
	return resp, rqa, newdepth, stilltalking
}
*/
