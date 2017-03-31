package flow

import (
	"RPG-Core/check"
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
		fmt.Println("You are at the", loc.Name)
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
			cp = locEventCheck(cp, dest)
			dt, dtext := models.TravelTime(cp.Loc, dest)
			cp.Loc = dest
			models.UpdateTime(dt)
			if cp.Health > 0 { //cheating, for if player dies on the way
				fmt.Println(dtext)
			}
		}
		mc := monstercheck(cp)
		if mc == true {
			fmt.Println("Monsters are here!")
			cp = monsterFlow(cp)
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
	case "Sleep":
		fmt.Println("You crawl into your tent and get ready for sleep." +
			"\nSet an alarm?")
		r1 := inputs.StringarrayInput([]string{"Yes", "No"})
		switch r1 {
		case 1:
			fmt.Println("How many hours would you like to sleep for?")
			dt := inputs.Numberinput(23)
			models.UpdateTime(dt * 100)
			cp.Health = 100
		case 2:
			fmt.Println("You pass out and sleep a solid six.")
			models.UpdateTime(6 * 100)
			cp.Health = 100
		}
	case "Hike":
		dt := 400
		models.UpdateTime(dt)
		fmt.Println("You spend three hours ascending and descending the switchbacks. You rise above the trees." +
			"Over the trees and onto rock under open sky. You can see different parts of camp." +
			"There is a cabin, abandoned. And past that a large mesa. The ghost stories told around the camp fire were about these places." +
			"You arrive back on the mountain base. Feeling healthier.")
		cp.Health = cp.Health + 15
	case "Swim":
		vm := models.StoryblobGetByName(5)
		if vm.Shown == false {
			fmt.Println("You look around, no one in sight. Take off your clothes are go swimming?")
			r1 := inputs.StringarrayInput([]string{"Yes", "No"})
			if r1 == 1 {
				fmt.Println("You remove your dirty camp clothes and jump into the cold water. The shock awakens your body and you laugh.")
				fmt.Println(vm.Story)
				//actualy update
				check.EventLoad(5)
				vc := models.CharacterGetByName("Veronica")
				models.CharacterUpdate(vc)
				conversation.Converser(vc)
				vm.Shown = true
				models.StoryblobUpdate(vm)
			} else {
				fmt.Println("You lack a swimswuit. So you stand on the shoreline and look at the calm, cool water.")
			}
		} else {
			fmt.Println("Veronica watches you swim. It's awkward.") //cheating
		}
	case "Forage":
		fmt.Println("You find some random berries. Do you want to eat them?")
		options := []string{"Yes", "No"}
		r1 := inputs.StringarrayInput(options)
		switch r1 {
		case 1:
			fmt.Println("Ouch. That burns going in. Probably when it goes out too.")
			cp.Health = cp.Health - 20
			cp.Cont = check.PContCheck(cp)
		case 2:
			fmt.Println("You return them to the forest.")
		}
	case "Menu":
		//save, inventory, quit
		fmt.Println("1. Save\n2. Inventory\n3. Quit")
		m1 := inputs.Numberinput(3)
		switch m1 {
		case 1:
			saveFlow() //testing
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
