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
	lightused := false //cheat for flashlight
	for cm.Engaged == true {
		fmt.Println("What would you like to do?")
		options := []string{"Talk", "Fight", "Run"}
		r1 := inputs.StringarrayInput(options)
		switch r1 {
		case 1:
			cm, cp = monsterTalk(cm, cp)
			cm.Engaged = check.PContCheck(cp) //in case player dies
		case 2: //sepearte into own function
			fchoice := combat.FightingOptions()
			switch fchoice {
			case "Devil Dice":
				fmt.Println("The", cm.ShortName, "pulls a set of dice out...does it have pockets?")
				cp = combat.DDFlow(cp, cm)
				cm.Engaged = false
			case "Fists":
				fmt.Println("You try to pummel the beast.")
				if cm.ShortName == smiler {
					fmt.Println("It deftly dodges your blows, smiling the whole time.")
				}
			case "Teeth":
				fmt.Println("You try to bite it.")
				if cm.ShortName == smiler {
					fmt.Println("It tastes awful. And returns the favor by taking a bite out of you.")
					cp.Health = cp.Health - 10
					cm.Engaged = check.PContCheck(cp) //in case player dies
				}
			case "Flashlight":
				fmt.Println("You shine the light on the creature.")
				if cm.ShortName == smiler {
					switch lightused {
					case false:
						fmt.Println("The grin disappears and the creature to sizzles under the light. You've weakened the thing.")
						lightused = true
						cm.Health = cm.Health / 2
					case true:
						fmt.Println("The smiler deflty dodges the beam of light. Not again.")
					}
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
	cp.Cont = check.PContCheck(cp)
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
			b1.Shown = true
			models.StoryblobUpdate(b1)
		}
	} else if cp.Loc == 8 && dest == 7 {
		fmt.Println("Test! Triggered")
		if check.Eventcheck(8) == true { //not showing as true!
			fmt.Println("You see the abandoned cabin. As you approach out steps a rabbit. Or a person with a rabbit head." +
				"It's very smooth looking. It has no eyes and no mouth. Just a rabbit nose and long ears. It turns in your direction." +
				"And that's it! I'm done. For now. Thanks for playing.")
			cp.Cont = false
		}
	}
	return cp
}
