package check

import (
	"RPG-Core/models"
)

//Eventcheck is an unclear function that should check where the player is in the story
func Eventcheck(eventid int) bool {
	//check storyblobs for now
	event := models.StoryblobGetByName(eventid)
	return event.Shown
}

//EventLoad should trigger the actions of events for opening from a save
func EventLoad(eventid int) {
	switch eventid {
	case 5:
		vc := models.CharacterGetByName("Veronica")
		vc.Loc = 2
		models.CharacterUpdate(vc)
	}
}
