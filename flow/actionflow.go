package flow

import (
	"RPG-Core/combat"
	"RPG-Core/conversation"
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
		loclist := models.TravelGet(cp.Loc) //returns array of ints, each a location number
		fmt.Println("You can walk to the:")
		for i := 0; i < len(loclist); i++ {
			iloc := models.LocationGet(loclist[i])
			if i == 0 {
				fmt.Printf("1. %s\n", iloc.Name)
			} else {
				fmt.Printf("%v. %s\n", i+1, iloc.Name)
			}
		}
		dest := loclist[inputs.Numberinput(len(loclist))-1]
		if dest == cp.Loc {
			fmt.Println("You stay still")
		} else {
			dt := models.TravelTime(cp.Loc, dest)
			cp.Loc = dest
			models.UpdateTime(dt)
			fmt.Println("You travel. It takes roughly one hour.")
		}
		mc := monstercheck(cp)
		if mc == true {
			fmt.Println("Monsters are here!")
			cp = combat.Fightflow(cp)
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
			conversation.Converser(cc)
		} else {
			fmt.Println("No one to speak with.")
		}
	case "Run":
		//Go somewhere else fast!
		fmt.Println("Nowhere to run")
	case "Menu":
		//save, inventory, quit
		fmt.Println("1. Save\n2. Inventory\n3. Quit")
		m1 := inputs.Numberinput(3)
		switch m1 {
		case 1:
			fmt.Println("This is where you would save")
		case 2:
			blist, _ := models.ItemGetByLoc(19)
			fmt.Println("You are wearing a ", blist[0].ShortName)
			fmt.Println("You are carrying: ")
			ilist, n := models.ItemGetByLoc(20)
			if n == 0 {
				fmt.Println("Nothing.")
			}
			for i, v := range ilist {
				fmt.Println(i+1, v.FullName)
			}
		case 3:
			fmt.Println("Bye bye")
			cp.Cont = false
		}
	default:
		fmt.Println("Odd. This action does nothing.")
	}
	return cp, err
}

//simple monster check for now
func monstercheck(cp models.Player) bool {
	monsters := false
	_, i := models.MonsterGetByLoc(cp.Loc)
	if i > 0 {
		monsters = true
	}
	return monsters
}
