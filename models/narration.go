package models

//Storyblob is base story struct
type Storyblob struct {
	Story string
	Ident int
	Shown bool
}

//Storyblobmap maps Storyblobs by name
var Storyblobmap = map[int]Storyblob{}

var (
	i1 = "Welcome to the Stone Hills forest. You are a 13 year old member of Adventure Crew." +
		"\nYou have backpacked deep into the New Mexico wilderness. Camp has been made. You are free to explore." +
		"\nYour best friend has not joined on this trip. The other Crew members are friendly strangers." +
		"\nMike is the group leader. He seems nice." +
		"\nThe adults keep to themselves. They have been distracted the last couple days." +
		"\nThe ghost stories told at night are excellent. This place is bursting with scenery and life." +
		"\nYet, you have trouble enjoying yourself. You are lonely. The forests looks more threatening than welcoming." +
		"\nAnd at night, you are scared to be anywhere without a light." +
		"\nWhich is a problem. Because you cannot find your flashlight. Last night you had it, this morning it was gone."
	sm = "A small critter scurries around the forest. It has two little arms and two little legs." +
		"\nIt has too many little teeth in a perpetual smile. Everything is murky shadow except for that white smile."
	mm = "Mike is friendly but you are not friends. Your only friend in the Crew was Steve and he's not here." +
		"\nMike is a little older, more athletic and confident than you see yourself."
	jm = "You know nothing about Josh." +
		"\nHe appears awkward. A little loud when he does talk. His laughter strained from trying too hard." //add more?n
	vm = "\"Hi there,\" you hear from the shore. You turn, surprised and embarrassed." +
		"\nThere is a woman on the shore. You swear, she wasn't there before." +
		"\nShe's bare foot and wearing a green dress. She has long black hair." +
		"\nShe looks older, probably in high school." +
		"\nShe's smiling."
	susm = "Susie is sitting on a rock. Her feet dangling in the air, far above the forest." +
		"\nShe is sun tanned. Another Crew member you have seen around but do not know well." +
		"\nFearless and fit, you've never seen her turn down a challenge or activity."
	cabinp = "\"It's here. You can get there through the forest.\""
	boss1  = "Half way down the mountain you hear a squabble. A large creature approaches from around the mountainside." +
		"\nCovered in bright orange and yellow feathers, it's like a bird. Except for the heads." +
		"\nIt has three heads. Each one is at the end of a long, serpentine neck. And each one is different." +
		"\nOne is beaked and scaly like a turtle. One has long ears and buck teeth like a rabbit." +
		"\nAnd one has a snout and short ears like a fox."
)

//Storyblobset sets initial Storyblobmap
func Storyblobset() {
	//defaults
	intro := Storyblob{i1, 1, false}
	shadowmeet := Storyblob{sm, 2, false}
	//need one for the other campers
	mikemeet := Storyblob{mm, 3, false}
	joshmeet := Storyblob{jm, 4, false}
	veronicameet := Storyblob{vm, 5, false}
	susiemeet := Storyblob{susm, 6, false}
	//events
	cabinpath := Storyblob{cabinp, 7, false}
	//boss 1
	boss1blob := Storyblob{boss1, 8, false}

	StoryblobUpdate(intro)
	StoryblobUpdate(shadowmeet)
	StoryblobUpdate(mikemeet)
	StoryblobUpdate(joshmeet)
	StoryblobUpdate(veronicameet)
	StoryblobUpdate(susiemeet)
	StoryblobUpdate(cabinpath)
	StoryblobUpdate(boss1blob)

}

//StoryblobGetByName grabs current item by number
func StoryblobGetByName(c int) Storyblob {
	cm := sbmap()
	i := cm[c]
	return i
}

//StorySizeGet returns how many events are in the game
func StorySizeGet() int {
	sb := sbmap()
	t1 := len(sb)
	return t1
}

func sbmap() map[int]Storyblob {
	return Storyblobmap
}

//StoryblobUpdate allows updates to Storyblobs
func StoryblobUpdate(cc Storyblob) Storyblob {
	cm := sbmap()
	cm[cc.Ident] = cc
	return cc
}
