package flow

import (
	"RPG-Core/check"
	"RPG-Core/models"
	"fmt"
	"os"
	"strconv"
)

//for save stuff

type saveData struct {
	//Player
	name   string
	height int
	loc    int
	//World
	time int
	day  int
	//events
	blobbool []bool
	//Characters
	//Items
}

//separate variable for number of events

//for starting a new save file
func saveStart(np models.Player) error {
	//Here it will pull in all defaults
	//player
	var newF saveData
	newF.name = np.Name
	newF.height = np.Height
	newF.loc = np.Loc
	//time
	newF.time, _ = models.GetTime()
	newF.day = models.CalendarCheck()
	//events, needs to be updated to reflect the number of events. Can automate?
	newF.blobbool = []bool{false, false, false, false, false}
	for i := 0; i < 5; i++ {
		newF.blobbool[i] = check.Eventcheck(i + 1)
	}
	//call save function
	err := saveSave(newF)
	return err
}

//need a save function to properly test the open
func saveFlow() {
	var newF saveData
	//player
	np := models.GetCurrentPlayer()
	newF.name = np.Name
	newF.height = np.Height
	newF.loc = np.Loc
	//time
	newF.time, _ = models.GetTime()
	newF.day = models.CalendarCheck()
	//events, needs to be updated to reflect the number of events. Can automate?
	newF.blobbool = []bool{false, false, false, false, false}
	for i := 0; i < 5; i++ {
		newF.blobbool[i] = check.Eventcheck(i + 1)
	}
	err := saveSave(newF)
	check.Check(err)
}

//separate save that takes in structure
func saveSave(newF saveData) error {
	//create file
	save := "save1"
	savef, err := os.Create(save)
	defer savef.Close()
	fmt.Println(savef.Name(), " created")
	//start writing
	//player
	n1 := len(newF.name)
	s2 := strconv.Itoa(n1)
	savef.WriteString(s2)
	savef.WriteString(newF.name)
	h1 := strconv.Itoa(newF.height)
	savef.WriteString(h1)
	l1 := strconv.Itoa(newF.loc)
	savef.WriteString(l1)
	//time
	t1 := strconv.Itoa(newF.time)
	d1 := strconv.Itoa(newF.day)
	//time needs to be four digits
	if len(t1) == 3 {
		savef.WriteString(string(0))
	}
	savef.WriteString(t1)
	savef.WriteString(d1)
	//events
	for i := 0; i < 5; i++ { //0 for false, 1 for true
		b1 := 0
		if newF.blobbool[i] == true {
			b1 = 1
		}
		b2 := strconv.Itoa(b1)
		//b1 := strconv.FormatBool(newF.blobbool[i])
		savef.WriteString(b2)
	}
	//sync
	savef.Sync()
	return err
}

//for opening a previously saved game. Should modify to work with Save structure?
func opensave() (models.Player, error) {
	var oldF saveData

	savef, err := os.Open("save1")
	defer savef.Close()

	if err != nil {
		fmt.Println("No saved game found.")
	}
	fmt.Println(savef.Name(), " opened")
	//only reads bytes now. Can work by making the right size byte for the data and then reading from the right space
	//player
	b1 := make([]byte, 1)
	savef.ReadAt(b1, 0)
	n, _ := strconv.Atoi(string(b1))
	b2 := make([]byte, n)
	savef.ReadAt(b2, 1)
	fmt.Println("You are", string(b2))
	oldF.name = string(b2)
	n64 := int64(n + 1)
	savef.ReadAt(b1, n64)
	oldF.height, _ = strconv.Atoi(string(b1))
	n64 = int64(n + 1 + 1)
	savef.ReadAt(b1, n64)
	oldF.loc, _ = strconv.Atoi(string(b1))
	//time
	b4 := make([]byte, 4)
	n64 = int64(n + 3)
	savef.ReadAt(b4, n64)
	fmt.Println("Time test:", string(b4))
	oldF.time, _ = strconv.Atoi(string(b4))
	n64 = int64(n + 7)
	savef.ReadAt(b1, n64) //only works for single digit days
	oldF.day, _ = strconv.Atoi(string(b1))
	//events
	oldF.blobbool = []bool{false, false, false, false, false}
	for i := 0; i < 5; i++ {
		n64 = int64(n + 8 + i)
		savef.ReadAt(b1, n64)
		bb, _ := strconv.Atoi(string(b1))
		if bb == 1 {
			oldF.blobbool[i] = true
		}
	}
	savef.Sync()
	fmt.Println("test:", oldF)
	//update player model
	np := models.Player{}
	np.Name = oldF.name
	np.Loc = oldF.loc
	np.Height = oldF.height
	np.Cont = true
	np.Health = 100 //cheat
	//update time
	models.SetTime(oldF.time)
	models.UpdateDay(oldF.day)
	//update events
	for i, v := range oldF.blobbool {
		if v == true {
			cb := models.StoryblobGetByName(i + 1)
			cb.Shown = true
			models.StoryblobUpdate(cb)
			check.EventLoad(i + 1) //triggers the impact of events
		}
	}
	return np, err
}

/*
byte 0 = length of name
1-n = name
n+1 = height
n+2 = Location
n+3 = time
n+7 = day
n+8 = first event bool
*/
