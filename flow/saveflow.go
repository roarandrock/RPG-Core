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
	health int
	//World
	time1 int
	time2 int
	time3 int
	time4 int
	day   int
	//events
	blobbool []bool
	//Characters
	//Items
	iOnP      []int
	iBackpack int
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
	newF.health = np.Health
	//time
	ct, _ := models.GetTime()
	newF.time4 = ct % 10
	ct = ct / 10
	newF.time3 = ct % 10
	ct = ct / 10
	newF.time2 = ct % 10
	ct = ct / 10
	newF.time1 = ct % 10
	newF.day = models.CalendarCheck()
	//events, needs to be updated to reflect the number of events. Can automate?
	newF.blobbool = []bool{false, false, false, false, false}
	for i := 0; i < 5; i++ {
		newF.blobbool[i] = check.Eventcheck(i + 1)
	}
	//items
	itemsOnP, n := models.ItemGetByLoc(20)
	newF.iOnP = make([]int, n)
	for i := 0; i < n; i++ {
		newF.iOnP[i] = itemsOnP[i].Ident
	}
	backpack, _ := models.ItemGetByLoc(19)
	newF.iBackpack = backpack[0].Ident
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
	newF.health = np.Health
	//time
	ct, _ := models.GetTime()
	newF.time4 = ct % 10
	ct = ct / 10
	newF.time3 = ct % 10
	ct = ct / 10
	newF.time2 = ct % 10
	ct = ct / 10
	newF.time1 = ct % 10
	newF.day = models.CalendarCheck()
	//events, needs to be updated to reflect the number of events. Can automate?
	newF.blobbool = []bool{false, false, false, false, false}
	for i := 0; i < 5; i++ {
		newF.blobbool[i] = check.Eventcheck(i + 1)
	}
	//items`
	itemsOnP, n := models.ItemGetByLoc(20)
	newF.iOnP = make([]int, n)
	for i := 0; i < n; i++ {
		newF.iOnP[i] = itemsOnP[i].Ident
	}
	backpack, _ := models.ItemGetByLoc(19)
	newF.iBackpack = backpack[0].Ident
	//call save function
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
	t1 := strconv.Itoa(newF.time1)
	t2 := strconv.Itoa(newF.time2)
	t3 := strconv.Itoa(newF.time3)
	t4 := strconv.Itoa(newF.time4)
	d1 := strconv.Itoa(newF.day)
	//writing
	savef.WriteString(t1)
	savef.WriteString(t2)
	savef.WriteString(t3)
	savef.WriteString(t4)
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
	//items
	ni := len(newF.iOnP) //how many items
	s3 := strconv.Itoa(ni)
	savef.WriteString(s3)
	for i := 0; i < ni; i++ {
		item1 := newF.iOnP[i]
		item2 := strconv.Itoa(item1)
		savef.WriteString(item2)
	}
	bp := strconv.Itoa(newF.iBackpack)
	savef.WriteString(bp)
	//sync
	savef.Sync()
	fmt.Println("test:", newF)
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
	bname := make([]byte, n)
	savef.ReadAt(bname, 1)
	fmt.Println("You are", string(bname))
	oldF.name = string(bname)
	n64 := int64(n + 1)
	savef.ReadAt(b1, n64)
	oldF.height, _ = strconv.Atoi(string(b1))
	n64 = int64(n + 1 + 1)
	savef.ReadAt(b1, n64)
	oldF.loc, _ = strconv.Atoi(string(b1))
	//time
	n64 = int64(n + 3)
	savef.ReadAt(b1, n64)
	oldF.time1, _ = strconv.Atoi(string(b1))
	n64 = int64(n + 4)
	savef.ReadAt(b1, n64)
	oldF.time2, _ = strconv.Atoi(string(b1))
	n64 = int64(n + 5)
	savef.ReadAt(b1, n64)
	oldF.time3, _ = strconv.Atoi(string(b1))
	n64 = int64(n + 6)
	savef.ReadAt(b1, n64)
	oldF.time4, _ = strconv.Atoi(string(b1))
	n64 = int64(n + 7)
	savef.ReadAt(b1, n64) //only works for single digit days
	oldF.day, _ = strconv.Atoi(string(b1))
	//events
	oldF.blobbool = []bool{false, false, false, false, false} //test for character intros
	for i := 0; i < 5; i++ {
		n64 = int64(n + 8 + i)
		savef.ReadAt(b1, n64)
		bb, _ := strconv.Atoi(string(b1))
		if bb == 1 {
			oldF.blobbool[i] = true
		}
	}
	//items
	b2 := make([]byte, 2)
	n64 = int64(n + 13)
	savef.ReadAt(b1, n64)
	ni, _ := strconv.Atoi(string(b1))
	if ni != 0 {
		oldF.iOnP = make([]int, ni)
		for i := 0; i < ni; i++ {
			n64 = int64(n + 14 + ni)
			savef.ReadAt(b2, n64)
			oldF.iOnP[i], _ = strconv.Atoi(string(b2))
		}
	}
	//backpack
	n64 = int64(n + 14 + ni*2)
	savef.ReadAt(b2, n64)
	oldF.iBackpack, _ = strconv.Atoi(string(b2))
	//save
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
	ct := oldF.time4
	ct = oldF.time3*10 + ct
	ct = oldF.time2*100 + ct
	ct = oldF.time1*1000 + ct
	models.SetTime(ct)
	models.SetDay(oldF.day)
	//update events
	for i, v := range oldF.blobbool {
		if v == true {
			cb := models.StoryblobGetByName(i + 1)
			cb.Shown = true
			models.StoryblobUpdate(cb)
			check.EventLoad(i + 1) //triggers the impact of events
		}
	}
	//update items
	for _, v := range oldF.iOnP {
		citem := models.ItemGetByNumber(v)
		citem.Loc = 20
		models.ItemUpdate(citem)
	}
	cbackpack := models.ItemGetByNumber(oldF.iBackpack)
	cbackpack.Loc = 19
	models.ItemUpdate(cbackpack)
	return np, err
}

/*
byte 0 = length of name
1-n = name
n+1 = height
n+2 = Location
n+3 = time1
n+4
n+5
n+6 = time4
n+7 = day
n+8 = first event bool
n+9
n+10
n+11
n+12 = fifth event bool
n+13 = item length, two digits?
n+14 = first item
*/
