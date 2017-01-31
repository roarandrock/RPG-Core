package models

//Location base struct
type Location struct {
	Name    string
	Loc     int
	Actions []string
}

/* Location notes:
1 - campground
Need forest
multiple camps
bathrooms

Actions:
Look, walk, run, hide, sleep, swim, use, talk,menu?

What are the classic adventure ones? Use, talk, Look, walk, combine?
Or can keep specific. Few locations, each has one or two unique options.
These give the player special things, like sleeping in the camp. Or swimming in the river.

Currently locations are static, no changes to be saved

Change actions, make one Leave?

Allow people to walk anywhere? Just time dependent? that could be good
*/

var (
	campground   = Location{"Campground", 1, []string{"Look", "Walk", "Talk", "Sleep"}}
	riverside    = Location{"Riverside", 2, []string{"Look", "Walk", "Talk", "Swim"}}
	mountainbase = Location{"Mountain Base", 3, []string{"Look", "Walk", "Climb"}}
)

//Locationmap maps locations
var Locationmap = map[int]Location{
	campground.Loc:   campground,
	riverside.Loc:    riverside,
	mountainbase.Loc: mountainbase,
}

var (
	s0      = "The Endless Void"
	s1      = "A comfortable campground, shaded by trees and full of people."
	s2      = "A babbling brook. Good place to take a bath or catch a fish."
	s3      = "A grand mountain rises before you. It's top rocky and free of trees. You see the trailhead."
	scenery = []string{s0, s1, s2, s3}
)

//WorldMap makes the connections for the world
//func WorldMap() {
//Or this, but it's stupid messy

var (
	campgroundadj = []int{1, 2, 3}
	riversideadj  = []int{1, 2}
	mountainbasej = []int{1}
)

//TravelGet gets the possible travel options
func TravelGet(l int) []int {
	werld := make([][]int, 5)
	werld[1] = campgroundadj
	werld[2] = riversideadj
	werld[3] = mountainbasej
	i := werld[l]
	return i
}

//SceneryGet returns scenery
func SceneryGet(l int) string {
	s := scenery[l]
	return s
}

//LocationGet grabs current item by number
func LocationGet(l int) Location {
	i := Locationmap[l]
	return i
}

//TravelTime tells you how long it takes to get somewhere
func TravelTime(cp int, w int) int {
	dt := 100 //currently just 1 hour to get somewhere, can change later
	return dt
}
