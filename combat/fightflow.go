package combat

import (
	"RPG-Core/inputs"
	"RPG-Core/models"
	"fmt"
)

/*

Timer could be used for blocking:
https://tour.golang.org/concurrency/6
*/

//Fightflow come here to fight
func Fightflow(cp models.Player) models.Player {
	mlist, i := models.MonsterGetByLoc(cp.Loc)
	if i > 1 {
		fmt.Println("More monsters than one? Odd.")
	}
	//introduction
	cm := mlist[0]
	fmt.Println("You see a", cm.FullName)
	fmt.Println(cm.Details)
	//fight selector
	//Hearts, dice
	cp = devildice(cp, cm)
	return cp
}

//Devil dice - 7 sided dice, 5 of them. Three rolls. Maybe reduce to 4?
//Sun hurts them, Shadows hurt you, Moons reroll
//Day you roll first, night they roll first

var (
	ddice = []string{"Sun", "Sun", "Moon", "Moon", "Shadow", "Shadow", "Shadow"} //add one more shadow?
)

//works basically, need to add a bunch of flavor text
//and mosnter turn
func devildice(cp models.Player, cm models.Monster) models.Player {
	fmt.Println("Welcome to Devil Dice") //rules?
	fcont := true
	pturn := false
	_, solar := models.GetTime() //do you need to aquire dice first?
	if solar == "Day" {
		pturn = true
	}
	for fcont == true {
		switch pturn {
		case false:
			fmt.Println("The", cm.ShortName, "rolls.")
			roll := ddmonturn()
			mdmg, pdmg, _ := dout(roll) //monsters don't get rerolls
			cp.Health = cp.Health - pdmg
			cm.Health = cm.Health - mdmg
			fmt.Println("Test: Monster health and player health", cm.Health, cp.Health)
			pturn = true
		case true:
			fmt.Println(cp.Name, "rolls.")
			roll := ddturn()
			mdmg, pdmg, reroll := dout(roll)
			cp.Health = cp.Health - pdmg
			cm.Health = cm.Health - mdmg
			fmt.Println("Test: Monster health and player health", cm.Health, cp.Health)
			for reroll == true { //for rerolls
				roll = ddturn()
				mdmg, pdmg, reroll = dout(roll)
				cp.Health = cp.Health - pdmg
				cm.Health = cm.Health - mdmg
				fmt.Println("Test:", cm.Health, cp.Health)
			}
			pturn = false
		}
		if cp.Health <= 0 {
			fmt.Println("You rolled your last.")
			fcont = false
		}
		if cm.Health <= 0 {
			fmt.Println("The monster retreats")
			fcont = false
		}
		/*
			fmt.Println("Test, 1 to continue. 2 to quit.")
			test1 := inputs.Numberinput(2)
			if test1 == 2 {
				fcont = false
			}
		*/
	}
	return cp
}

type ddie struct {
	face string
	keep bool
}

func dout(roll []ddie) (int, int, bool) {
	mdmg := 0
	pdmg := 0
	moon := 0
	reroll := false
	for i := range roll {
		if roll[i].face == "Sun" {
			mdmg = mdmg + 10
		} else if roll[i].face == "Shadow" {
			pdmg = pdmg + 10
		} else if roll[i].face == "Moon" {
			moon++
		}
	}
	if moon >= 3 {
		reroll = true
		fmt.Println("Three moons, that's a reroll.")
	}
	return mdmg, pdmg, reroll
}

func ddmonturn() []ddie {
	roll := make([]ddie, 5)
	//one roll for monster
	for i := 0; i < 5; i++ {
		roll[i].face = diceroll()
		roll[i].keep = false
	}
	fmt.Println("Your foe rolls:")
	for _, v := range roll {
		fmt.Println(v.face)
	}
	return roll
}

func ddturn() []ddie {
	roll := make([]ddie, 5)
	//easier just walking through three rolls
	//first roll
	for i := 0; i < 5; i++ {
		roll[i].face = diceroll()
		roll[i].keep = false
	}
	fmt.Println("You have:")
	for _, v := range roll {
		fmt.Println(v.face)
	}
	sa := []string{"Roll all dice again", "Keep some dice", "Keep all dice"}
	r1 := inputs.StringarrayInput(sa)
	switch r1 {
	case 2:
		roll = dicekeep(roll)
	case 3:
		for i := range roll {
			roll[i].keep = true
		}
	}
	//second roll
	for i, v := range roll {
		if v.keep == false {
			roll[i].face = diceroll()
		}
	}
	fmt.Println("You have:")
	for i, v := range roll {
		fmt.Println(v.face)
		roll[i].keep = false
	}
	r1 = inputs.StringarrayInput(sa)
	switch r1 {
	case 2:
		roll = dicekeep(roll)
	case 3:
		for i := range roll {
			roll[i].keep = true
		}
	}
	//third roll
	for i, v := range roll {
		if v.keep == false {
			roll[i].face = diceroll()
		}
	}
	fmt.Println("Your final roll is:")
	for _, v := range roll {
		fmt.Println(v.face)
	}
	fmt.Println("test:", roll)
	return roll
}

func diceroll() string {
	n := Brng(7)
	s1 := ddice[n-1]
	return s1
}

func dicekeep(roll []ddie) []ddie {
	kont := true
	fmt.Println("Which die to keep?")
	for kont == true {
		fmt.Println("You have:")
		for i, v := range roll {
			fmt.Println(i+1, v.face)
		}
		r1 := inputs.Numberinput(5)
		roll[r1-1].keep = true
		sa := []string{"Keep more dice?", "That's enough"}
		r2 := inputs.StringarrayInput(sa)
		if r2 == 2 {
			kont = false
		}
	}
	return roll
}