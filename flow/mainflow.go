package flow

import (
	"RPG-Core/inputs"
	"RPG-Core/models"
	"fmt"
	"os"
	"strconv"
)

func Mainflow(cp models.Player) (models.Player, error) {
	fmt.Println("Welcome to the game", cp.Name)

	//cheating to make an error with error
	i, err := strconv.Atoi("-42")
	fmt.Println(i)
	switch cp.Height {
	case 1:
		fmt.Println(cp.Name, " is tall")
	case 2:
		fmt.Println(cp.Name, " is short")
	}
	if cp.Loc == 1 {
		fmt.Println("You are in the campground.")
	} else {
		fmt.Println("You are in limbo.")
	}

	return cp, err
}

func Intro() (models.Player, error) {
	fmt.Println("1. Make a new game\n2. Open Saved Game")
	r1 := inputs.Numberinput(2)
	cp := models.Player{}

	//cheating to make an error with error
	_, err := strconv.Atoi("-42")

	switch r1 {
	case 1:
		cp, err = newgame()
	case 2:
		cp, err = opensave()
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
	_, err = savef.WriteString("Name")
	n1 := len(np.Name)
	s2 := strconv.Itoa(n1)
	_, err = savef.WriteString(s2)
	_, err = savef.WriteString(np.Name)
	h1 := strconv.Itoa(np.Height)
	_, err = savef.WriteString("Height")
	n1, err = savef.WriteString(h1)
	l1 := strconv.Itoa(np.Loc)
	_, err = savef.WriteString("Loc")
	_, err = savef.WriteString(l1)
	savef.Sync()
	return np, err
}

func newplayer() models.Player {
	np := models.Player{}
	fmt.Println("Are you tall(1) or short(2)?")
	np.Height = inputs.Numberinput(2)
	fmt.Println("What is your name?")
	np.Name = inputs.Stringinput(1)
	np.Loc = 1
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
	_, err = savef.ReadAt(b1, 4)
	n, err := strconv.Atoi(string(b1))
	b2 := make([]byte, n)
	_, err = savef.ReadAt(b2, 5)
	fmt.Println("You are", string(b2))
	np.Name = string(b2)
	n64 := int64(n + 5 + 6)
	_, err = savef.ReadAt(b1, n64)
	np.Height, err = strconv.Atoi(string(b1))
	n64 = int64(n + 5 + 6 + 1 + 3)
	_, err = savef.ReadAt(b1, n64)
	np.Loc, err = strconv.Atoi(string(b1))
	savef.Sync()

	return np, err
}
