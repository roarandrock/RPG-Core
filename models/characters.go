package models

//Character base struct
type Character struct {
	Name     string
	Loc      int
	Friendly bool
	Status   string
	Depth    int
}

/* Character notes:
1 - campground

Need to save Characters, tied to having separate starting and in-game maps

Option - initial scripts and save scripts that define the characters/world.
Then when I create new or open, it goes to the script and runs it to "set the world"
Will need functions like CreateCharacter that call from the initial/saved model to create the character map

Dialog - Greeting, exit, panic
Separate conversation data

//Mike, Josh, Susie and Veronica
*/

//Charactermap maps characters by name
var charactermap = map[string]Character{}

//Characterset sets initial charactermap
func Characterset() {
	//defaults
	mike := Character{"Mike", 1, true, "Real", 3}
	josh := Character{"Josh", 1, true, "Real", 3}
	veronica := Character{"Veronica", 40, true, "Imaginary", 3}

	CharacterUpdate(mike)
	CharacterUpdate(josh)
	CharacterUpdate(veronica)
}

//CharacterGetByName grabs current item by number
func CharacterGetByName(c string) Character {
	cm := cmap()
	i := cm[c]
	return i
}

//CharacterGetByLoc grabs character by location
func CharacterGetByLoc(l int) ([]Character, int) {
	cm := cmap()
	cslice := []Character{}
	i := 0
	for _, v := range cm {
		if v.Loc == l {
			cslice = append(cslice, v)
			i++
		}
	}
	return cslice, i
}

func cmap() map[string]Character {
	return charactermap
}

//CharacterUpdate allows updates to characters
func CharacterUpdate(cc Character) Character {
	cm := cmap()
	cm[cc.Name] = cc
	return cc
}
