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
		case 2:
			fmt.Println("How would you like to fight?")
			options = []string{"Fists", "Teeth"}
			ilist, _ := models.ItemGetByLoc(20)
			for _, v := range ilist {
				if v.Iclass == "combat" {
					options = append(options, v.FullName)
				}
			}
			r2 := inputs.StringarrayInput(options)
			fchoice := options[r2-1]
			switch fchoice {
			case "Devil Dice":
				cp = combat.Fightflow(cp, cm)
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
				fmt.Println("In the sun, you run back to camp succesfully.")
				cp.Loc = 1
				dt := models.TravelTime(cp.Loc, 1)
				models.UpdateTime(dt)
				cm.Engaged = false
			} else {
				fmt.Println("In the fading light the beast moves fast. It blocks your escape.")
			}
		}
	}
	return cp
}

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
