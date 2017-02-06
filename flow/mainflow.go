package flow

import (
	"RPG-Core/inputs"
	"RPG-Core/models"
	"fmt"
	"os"
	"strconv"
)

//Mainflow is the main loop of the game
func Mainflow(cp models.Player) (models.Player, error) {
	fmt.Println("Welcome to the game", cp.Name)
	//setting loop continue constant
	cont := cp.Cont
	//cheating to make an error with error
	_, err := strconv.Atoi("-42")
	for cont == true {
		cp, err = gameloop(cp)
		cont = cp.Cont
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
		cp, err = newgame()
		models.Characterset() // cheating
		models.Itemset()
		models.Monsterset()
	case 2:
		cp, err = opensave()
		models.Characterset() // cheating
		models.Itemset()
		models.Monsterset()
	}
	return cp, err
}

func newgame() (models.Player, error) {

	fmt.Println("Make a new game")

	np := newplayer()
	//create savefile
	save := "save1"
	savef, err := os.Create(save)
	defer savef.Close()
	fmt.Println(savef.Name(), " created")
	//save name first then height. But Name can be variable length
	//Can clean this up. Only one write string per item?
	//Make seperate save function to be called now and later
	savef.WriteString("Name")
	n1 := len(np.Name)
	s2 := strconv.Itoa(n1)
	savef.WriteString(s2)
	savef.WriteString(np.Name)
	h1 := strconv.Itoa(np.Height)
	savef.WriteString("Height")
	savef.WriteString(h1)
	l1 := strconv.Itoa(np.Loc)
	savef.WriteString("Loc")
	savef.WriteString(l1)
	savef.Sync()
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
	return np
}

func opensave() (models.Player, error) {

	fmt.Println("Open saved game")
	//create player from data
	np := models.Player{}
	savef, err := os.Open("save1")
	defer savef.Close()
	if err != nil {
		fmt.Println("No saved game found.")
	}
	fmt.Println(savef.Name(), " opened")
	//only reads bytes now. Can work by making the right size byte for the data and then reading from the right space
	//could change later to a full read and then using string functions or buffered reader
	b1 := make([]byte, 1)
	savef.ReadAt(b1, 4)
	n, _ := strconv.Atoi(string(b1))
	b2 := make([]byte, n)
	savef.ReadAt(b2, 5)
	fmt.Println("You are", string(b2))
	np.Name = string(b2)
	n64 := int64(n + 5 + 6)
	savef.ReadAt(b1, n64)
	np.Height, _ = strconv.Atoi(string(b1))
	n64 = int64(n + 5 + 6 + 1 + 3)
	savef.ReadAt(b1, n64)
	np.Loc, err = strconv.Atoi(string(b1))
	np.Cont = true
	np.Health = 100 //cheat
	savef.Sync()

	return np, err
}

func gameloop(cp models.Player) (models.Player, error) {
	//for testing list a few things
	cd := models.CalendarCheck()
	ct, cs := models.GetTime()
	cl := models.LocationGet(cp.Loc)
	fmt.Println("You are in the", cl.Name)
	fmt.Println("It is", ct, ",", cs, "on day", cd)
	//can implement a battle check or something here for events
	s := actielist(cp)
	fmt.Println("You", s)
	cp, err := actieSelector(s, cp)
	return cp, err
}

//Finale for end of game, after mainflow is completed. Clean up and stuff
func Finale(cp models.Player) (models.Player, error) {
	//cheating to make an error with error
	_, err := strconv.Atoi("-42")
	return cp, err
}
