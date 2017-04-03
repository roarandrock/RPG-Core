package models

//Location base struct
type Location struct {
	Name    string
	Loc     int
	Actions []string
}

/* Location notes:
1 - campground
2 - Lake
3 - Mountain base
4 - Mountain top
5- Mesa base
6 - mesa top
7 - Abandond Cabin
8 - forest

19 - Player's backpack
20 - Player
21 - Mike
22 - Josh
23 - Susie
24 - Veronica

40 - Does not exist
50 - Destroyed


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
	lake         = Location{"Lake", 2, []string{"Look", "Walk", "Talk", "Swim"}}
	mountainbase = Location{"Mountain Base", 3, []string{"Look", "Walk", "Hike"}}
	mountaintop  = Location{"Mountain Top", 4, []string{"Look", "Walk", "Talk", "Admire Scenery"}}
	mesabase     = Location{"Mesa Base", 5, []string{"Look", "Walk", "Climb"}}
	mesatop      = Location{"Mesa Top", 6, []string{"Look", "Walk", "Finish"}}
	cabin        = Location{"Abandoned Cabin", 7, []string{"Look", "Walk", "Investigate"}}
	forest       = Location{"Forest", 8, []string{"Look", "Walk", "Forage"}}
)

//change to be like other structures with an initial and an updater? and do descriptions like items

//Locationmap maps locations
var Locationmap = map[int]Location{
	campground.Loc:   campground,
	lake.Loc:         lake,
	mountainbase.Loc: mountainbase,
	mountaintop.Loc:  mountaintop,
	mesabase.Loc:     mesabase,
	mesatop.Loc:      mesatop,
	cabin.Loc:        cabin,
	forest.Loc:       forest,
}

var (
	s0      = "The Endless Void"
	s1      = "A comfortable campground, shaded by trees. Normally full of people but now there are only two."
	s2      = "A deep calm lake. Good place to take a bath or catch a fish."
	s3      = "A grand mountain rises before you. It's top rocky and free of trees. You see the trailhead."
	s4      = "You can see the whole camp from up here."
	s5      = "You are at the base of the mesa. It's rugged trail to the top."
	s6      = "Top of the mesa. The haunted mesa. Why are you here?"
	s7      = "An abandoned cabin, something happened here a long time ago."
	s8      = "The majestic and awesome forest. Seemingly endless. Full of life and shadows."
	scenery = []string{s0, s1, s2, s3, s4, s5, s6, s7, s8}
)

//WorldMap makes the connections for the world
//func WorldMap() {
//Or this, but it's stupid messy

var (
	campgroundadj = []int{1, 2, 8}
	lakeadj       = []int{1, 2, 8}
	mountainbasej = []int{3, 4, 8}
	mountaintopj  = []int{3, 4}
	//mountain top and mesa out
	cabinadj  = []int{7, 8}
	forestadj = []int{1, 2, 3, 8}
)

//TravelGet gets the possible travel options
//can open up things here based on events
func TravelGet(l int) []int {
	werld := make([][]int, 9)
	werld[1] = campgroundadj
	werld[2] = lakeadj
	werld[3] = mountainbasej
	werld[4] = mountaintopj
	werld[7] = cabinadj
	werld[8] = forestadj
	if l == 8 {
		if StoryblobGetByName(7).Shown == true {
			werld[8] = append(werld[8], 7)
		}
	}
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
func TravelTime(cp int, w int) (int, string) {
	dt := 100 //currently just 1 hour to get somewhere, can change later
	var dtext string
	if cp == 3 {
		if w == 4 {
			dt = 300
			dtext = "It's a long hike to the top." //can change these to output instead of printing them here
		}
	} else if cp == 4 {
		if w == 3 {
			dt = 200
			dtext = "You make good time going down the mountain."
		}
	} else {
		dl := LocationGet(w)
		dtext = "You arrive at the " + dl.Name
	}
	return dt, dtext
}
