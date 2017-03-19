package flow

import (
	"RPG-Core/check"
	"RPG-Core/inputs"
	"RPG-Core/models"
	"fmt"
	"strconv"
)

//Mainflow is the main loop of the game
func Mainflow(cp models.Player) (models.Player, error) {
	//setting loop continue constant
	cont := cp.Cont
	//cheating to make an error with error
	_, err := strconv.Atoi("-42")
	for cont == true {
		cp, err = gameloop(cp)
		cont = check.PContCheck(cp) //checks health and player continue
	}
	return cp, err
}

//Intro starts the game
func Intro() (models.Player, error) {
	fmt.Println("1. Make a new game\n2. Open Saved Game")
	r1 := inputs.Numberinput(2)
	cp := models.Player{}

	//cheating to make an error with error
	_, err := strconv.Atoi("-42")

	switch r1 {
	case 1:
		models.Characterset() // cheating, maybe should make one call to a list for this
		models.Itemset()
		models.Monsterset()
		models.Storyblobset()
		cp, err = newgame()
		//need an intro
		i1 := models.StoryblobGetByName(1)
		fmt.Println(i1.Story)
		i1.Shown = true
		models.StoryblobUpdate(i1)
	case 2:
		models.Characterset() // cheating
		models.Itemset()
		models.Monsterset()
		models.Storyblobset()
		cp, err = opensave()
	}
	return cp, err
}

func newgame() (models.Player, error) {
	np := newplayer()
	//create savefile
	err := saveStart(np)
	return np, err
}

func newplayer() models.Player {
	np := models.Player{}
	fmt.Println("What is your first name?")
	np.Name = inputs.Stringinput(1)
	fmt.Println("Are you tall(1) or short(2)?")
	np.Height = inputs.Numberinput(2)
	np.Loc = 1
	np.Cont = true
	np.Health = 100
	models.UpdateCurrentPlayer(np)
	return np
}

func gameloop(cp models.Player) (models.Player, error) {
	cd := models.CalendarCheck()
	ct, cs := models.GetTime()
	//cl := models.LocationGet(cp.Loc)
	//fmt.Println("You are in the", cl.Name)
	fmt.Println("It is", ct, ",", cs, "time. And day number", cd) //cheat
	//can implement a battle check or something here for events
	s := actielist(cp)
	fmt.Println("You", s)
	cp, err := actieSelector(s, cp)
	models.UpdateCurrentPlayer(cp)
	return cp, err
}

//Finale for end of game, after mainflow is completed. Clean up and stuff
func Finale(cp models.Player) (models.Player, error) {
	//cheating to make an error with error
	_, err := strconv.Atoi("-42")
	if cp.Health <= 0 {
		fmt.Println("No more fight is left. You have been defeated.")
	}
	if cp.Cont == false {
		fmt.Println("You are no longer able to continue. Oddly you have come to an end.")
	}
	return cp, err
}
