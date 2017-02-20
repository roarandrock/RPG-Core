package models

import "fmt"

/*
Need items:
Small backpack
Map
Flashlight
Headlamp?
Pocket Knife
Fishing Pole?
Walking stick
Bigger Backpack
?small backpack to start
Matches/Flare?

*/

//Item is basic item model
type Item struct {
	FullName  string
	ShortName string
	Size      string
	Iclass    string //battle, character, location
	Loc       int
	Details   string
	LightBuff int
	Ident     int //double digit ident
}

var itemmap = map[string]Item{}

//Itemset sets initial items
func Itemset() {
	//descriptions
	sb1 := "A backpack that can carry one small item and one medium. Or one large item."
	md1 := "A handy map showing all the campground. Including how to get there."
	fd1 := "A reliable flashlight."
	bb1 := "A big backpack that can carry 1 large and 1 medium item. Or 2 medium items. Or 3 small items."
	dd1 := "A set of devil dice. Each die is 8-sided. Four of the sides show a shadow figure, two sides a sun and two sides a moon."
	//defaults
	mapi := Item{"Camp Map", "map", "small", "location", 23, md1, 0, 10}
	flashlight := Item{"Flashlight", "flashlight", "medium", "combat", 22, fd1, 20, 11}
	smallbackpack := Item{"Small Backpack", "small pack", "large", "player", 19, sb1, 0, 12}
	largebackpack := Item{"Large Backpack", "large pack", "large", "player", 21, bb1, 0, 13}
	devildicebasic := Item{"Devil Dice", "devil dice", "small", "combat", 24, dd1, 0, 14}
	ItemUpdate(mapi)
	ItemUpdate(flashlight)
	ItemUpdate(smallbackpack)
	ItemUpdate(largebackpack)
	ItemUpdate(devildicebasic)
}

//ItemGetByName grabs current item by short name
func ItemGetByName(c string) Item {
	cm := imap()
	i := cm[c]
	return i
}

//ItemGetByLoc grabs Item by location
func ItemGetByLoc(l int) ([]Item, int) {
	cm := imap()
	cslice := []Item{}
	i := 0
	for _, v := range cm {
		if v.Loc == l {
			cslice = append(cslice, v)
			i++
		}
	}
	return cslice, i
}

//ItemGetByNumber grabs item by its ident number
func ItemGetByNumber(i int) Item {
	cm := imap()
	var ri Item
	for _, v := range cm {
		if v.Ident == i {
			ri = v
		}
	}
	return ri
}

func imap() map[string]Item {
	return itemmap
}

//ItemUpdate allows updates to Items
func ItemUpdate(cc Item) Item {
	cm := imap()
	cm[cc.ShortName] = cc
	return cc
}

//ItemSizeCheck to see if there's room in the backpack
func ItemSizeCheck(cc Item) {
	si := 0
	mi := 0
	li := 0
	full := false
	//default size is One small and one medium or one large
	inBackpack, _ := ItemGetByLoc(20)
	for _, v := range inBackpack {
		if v.Size == "small" {
			si++
		} else if v.Size == "medium" {
			mi++
		} else if v.Size == "large" {
			li++
		}
	}
	if si > 0 {
		fmt.Println("You have", si, "small items.")
	}
	if mi > 0 {
		fmt.Println("You have", mi, " medium items.")
	}
	if li > 0 {
		fmt.Println("You have", li, "large items.")
	}
	//got to do better with this
	if si == 2 {
		full = true
	}
	if mi == 1 {
		if si == 1 || cc.Size == "medium" || cc.Size == "large" {
			full = true
		}
	}
	if li == 1 {
		full = true
	}
	if full == true {
		fmt.Println("No room.")
	} else {
		fmt.Println("There's room. Item added to backpack")
		cc.Loc = 20
		ItemUpdate(cc)
	}
}
