package flow

import (
	"RPG-Core/check"
	"RPG-Core/combat"
	"RPG-Core/inputs"
	"RPG-Core/models"
	"fmt"
)

var smiler = "smiler" //cheating, lame

//simple monster check for now
func monstercheck(cp models.Player) bool {
	monsters := false
	_, i := models.MonsterGetByLoc(cp.Loc)
	if i > 0 {
		monsters = true
	}
	return monsters
}

//move more of this to combat?
func monsterFlow(cp models.Player) models.Player {
	mlist, i := models.MonsterGetByLoc(cp.Loc)
	if i > 1 {
		fmt.Println("More monsters than one? Odd.")
	}
	//introduction
	cm := mlist[0]
	fmt.Println("You see a", cm.FullName)
	if cm.ShortName == smiler {
		met := check.Eventcheck(2)
		if met == false {
			sb := models.StoryblobGetByName(2)
			fmt.Println(sb.Story)
			sb.Shown = true
			models.StoryblobUpdate(sb)
		}
	}
	cm.Engaged = true
	for cm.Engaged == true {
		fmt.Println("What would you like to do?")
		options := []string{"Talk", "Fight", "Run"}
		r1 := inputs.StringarrayInput(options)
		switch r1 {
		case 1:
			cm, cp = monsterTalk(cm, cp)
		case 2: //sepearte into own function
			fchoice := combat.FightingOptions()
			switch fchoice {
			case "Devil Dice":
				cp = combat.DDFlow(cp, cm)
				cm.Engaged = false
			case "Fists":
				fmt.Println("You try to pummel the beast.")
				if cm.ShortName == smiler {
					fmt.Println("It deftly dodges your blows, smiling the whole time.")
					fmt.Println("Eventually it gets bored and wanders off.") //but then it's not an obstacle, need the player to leave.
					cm.Engaged = false
				}
			case "Teeth":
				fmt.Println("You try to bite it.")
				if cm.ShortName == smiler {
					fmt.Println("It tastes awful. And returns the favor by taking a bite out of you.")
					cp.Health = cp.Health - 10
					fmt.Println("It wanders off into the forest. It doesn't like how you taste apparently.")
					cm.Engaged = false
				}
			case "Flashlight":
				fmt.Println("You shine the light on the creature.")
				if cm.ShortName == smiler {
					fmt.Println("The grin disappears and the creature takes off, sizzling under the beam of light.")
					cm.Engaged = false
				}
			}

		case 3:
			_, day := models.GetTime()
			if day == "Day" {
				fmt.Println("In the sun, you run and succesfully escape the beast.")
				cp.Loc = 1
				dt, _ := models.TravelTime(cp.Loc, 1)
				models.UpdateTime(dt)
				cm.Engaged = false
			} else {
				fmt.Println("In the fading light the beast moves fast. It blocks your escape.")
			}
		}
	}
	return cp
}

//move to combat?
func monsterTalk(cm models.Monster, cp models.Player) (models.Monster, models.Player) {
	switch cm.ShortName {
	case smiler:
		fmt.Println("The creature ignores your attempts to speak with it.")
		options := []string{"Pet it", "Laugh at it", "Ignore it"}
		r1 := inputs.StringarrayInput(options)
		switch r1 {
		case 1:
			inputs.StringarrayInput([]string{"Palm up", "Palm down"})
			fmt.Println("The smiler approaches cautiously. Then with confidence bites your hand.")
			cp.Health = cp.Health - 10
		case 2:
			fmt.Println("You laugh at thing. It opens it's mouth and imitates your laughter." +
				"\nIt's almost adorable.")
		case 3:
			fmt.Println("You ignore it. It seems to stare at you. It's not going away.")
		}
	default:
		fmt.Println("The shadow has an odd smile.")
	}
	return cm, cp
}

func locEventCheck(cp models.Player, dest int) models.Player {
	//first boss - need to be coming down the mountain after finding out where the cabin is
	if cp.Loc == 4 && dest == 3 {
		if check.Eventcheck(8) == false && check.Eventcheck(7) == true {
			b1 := models.StoryblobGetByName(8)
			fmt.Println(b1.Story)
			cp = combat.Boss1Flow(cp)
		}
	}
	return cp
}
