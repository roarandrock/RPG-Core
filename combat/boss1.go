package combat

import (
	"RPG-Core/check"
	"RPG-Core/inputs"
	"RPG-Core/models"
	"fmt"
)

/*
for fighting first boss -
three headed flying beast with
rabbit
turtle
fox
heads

Now can defeat them with just talking to turtle. Should not be that easy. Like leave Turtle and Rabbit and want to be only left Rabbit?
Or first convince Rabbit you will play dice with him before killing the others

Need exit text here

And flavor text for Devil Dice!

*/

//Boss1Flow for first boss fight
func Boss1Flow(cp models.Player) models.Player {
	//intro
	fmt.Println("As it flies closer you hear it talking to itself.")
	fmt.Println("\"You know I'm faster,\" says the rabbit head." +
		"\n\"Does not matter,\" replies the turtle head." +
		"\n\"What do you mean it doesn't matter? Of course it matters!\", rebukes the rabbit." +
		"\n\"Doesn't matter. I still win.\"" +
		"\n\"You win at what?\", asks the rabbit." +
		"\n\"Life is not a race. And I live longer, so I win that too.\"" +
		"\n\"You upside-down soup and bowl, my family and I will...\"" +
		"\n\"Shut up! You two idiots,\" interrupts the fox head, \"There is the human.\"" +
		"\nThe large wings flap. The monster hovers in front of you." +
		"\n\"Hello human. We know you have a set of devil dice. We want to know where you aquired it?\", asks the Fox." +
		"\n\"We were sent by the Professor,\" adds the Rabbit. The Fox snaps at him and the Rabbit head pulls back.") //need more intro, set up chance to betray veronica
	heads := []bool{true, true, true} //shows they are still alive. Rabbit, Turtle, Fox
	cm := models.Monster{
		FullName:  "Representative Zero Bird",
		ShortName: "zero bird",
		Health:    50, //for testing, should be like 70
		Loc:       40,
		Details:   "Flying three headed beast",
		Engaged:   true}
	//loop
	options := []string{"Talk", "Fight", "Run"}
	bfcont := true

	for bfcont == true {
		r1 := inputs.StringarrayInput(options)
		switch r1 {
		case 1:
			heads, cp.Health = b1Chat(heads, cp.Health)
			//bfcont = false //testing
		case 2:
			fchoice := FightingOptions()
			switch fchoice {
			case "Fists":
				fmt.Println("It's flying. Your fists cannot reach it.")
			case "Teeth":
				fmt.Println("It's flying. Your teeth cannot reach it.")
			case "Devil Dice":
				fmt.Println("You pull out your set of dice.")
				if heads[0] == true {
					fmt.Println("The ears on the rabbit twitch. It's eyes follow your hand full of dice.")
					if heads[2] == true {
						fmt.Println("The fox snaps at the rabbit, \"You idiot. We don't have time for that!\"")
					} else if heads[1] == true {
						fmt.Println("One of the bird's claws, pulls up into it's feathers. Then the turtle solemnly states, \"I don't think the Professor wants us to play with him.\"" +
							"\n\"What do you know?\", replies the rabbit. But it stops moving and the claw reappears empty.")
					} else {
						fmt.Println("The bird's claw reaches into the feathers and pulls out a set of appropriately large dice." +
							"\n\"Let's play,\" says the rabbit.")
						cp = DDFlow(cp, cm)
						if cp.Health > 0 {
							fmt.Println("\"No! How could I lose to a human? The Professor is going to kill me!\", yells the Rabbit head." +
								"\"Time for me to go into hiding. Good luck brat. You may have beat me, but I have kin. Many, many kin.\"")
							fmt.Println("The bird beast that once had three heads is defeated. It flies off howling in pain and cursing it's stupidity.")
						}
						bfcont = false
					}
				}
			case "Flashlight":
				fmt.Println("You shine the light on the beast. Some of its orange feathers start curling and blackening." +
					"\n\"Stop that\", says the rabbit and it not gently smacks you with it's claw.")
				cp.Health = cp.Health - 5
			default:
				fmt.Println("Oddly, no effect on the beast")
			}
		case 3:
			fmt.Println("You run")
			//need to implement death here.
			bfcont = false //for testing
		}
		if check.PContCheck(cp) == false {
			fmt.Println("You are defeated. The bird claw gently plucks the devil dice from your corpse and flies away.")
			bfcont = false
		}
	}
	//finales

	return cp //how to force exit when lose?
}

func b1Chat(heads []bool, cpH int) ([]bool, int) {
	options := []string{"Rabbit"}
	if heads[1] == true {
		options = append(options, "Turtle")
	}
	if heads[2] == true {
		options = append(options, "Fox")
	}
	fmt.Println("Whom to talk with?")
	r1 := inputs.StringarrayInput(options)
	switch r1 { //unnecessary, can just pass directly to flow? Or get startling line of dialog here and pass to flow?
	case 1:
		fmt.Println("You gesture at the Rabbit head and it swoops in to listen.")
		heads, cpH = b1ConvoFlow(1, heads, cpH)
	case 2:
		fmt.Println("You gesture at the Turtle head and it swoops in to listen.")
		heads, cpH = b1ConvoFlow(2, heads, cpH)
	case 3:
		fmt.Println("You gesture at the Fox head and it swoops in to listen.")
		heads, cpH = b1ConvoFlow(3, heads, cpH)
	}
	return heads, cpH
}

//need a simplified version of dialog from Josh
func b1ConvoFlow(head int, heads []bool, cpH int) ([]bool, int) {
	bconv := true
	var poptions []string
	var npD pD
	var nbD bD
	//define bD first
	//need a switch for different heads
	switch head {
	case 1:
		nbD = bDstartR
	case 2:
		nbD = bDstartT
	case 3:
		nbD = bDstartF
	}
	for bconv == true {
		poptions = choicemakerVB(nbD)
		r1 := inputs.StringarrayInput(poptions)
		npD = *nbD.branches[r1-1]
		nbD = *npD.branch
		v1 := "\"" + nbD.words + "\""
		fmt.Println(v1)
		//checks branches for possible events
		nbD, bconv, cpH, heads = dialogB1Events(nbD, cpH, heads)
	}
	return heads, cpH
}

/*
pKillT  = pD{}
pKillF  = pD{}
pKillR  = pD{}
pPlayDD = pD{}
pBetray = pD{}
*/

func dialogB1Events(nbD bD, cpH int, heads []bool) (bD, bool, int, []bool) {
	bconv := true
	switch nbD.branches[0] {
	case &pExit:
		bconv = false
		fmt.Println("The head you've been talking with writhes away. It can no longer hear you.")
	case &pBetray:
		//kill player
		bconv = false
		cpH = 0
		fmt.Println("Having betrayed Veronica, the beast tears you apart. Each head savoring a different piece of your body.")
	case &pKillT:
		bconv = false
		fmt.Println("The Rabbit head attacks the Turtle head. It's a manic attack, biting all along the Turtle's scaly head and long neck." +
			"\nThe Turtle eventually retaliates. It's beak bloodying the Rabbit.")
		if heads[2] == true {
			fmt.Println("The Fox barks trying to restore order. But it's too late. The fight is over. The Turtle head goes limp, the neck hanging straight down in the air." +
				"\nThe Rabbit is bleeding heavily. The Fox finishes the Rabbit in disgust. Snout covered in blood, it turns to you." +
				"\n\"You annoying little punk. I don't care anymore about who gave you those Dice. The Professor can figure it out himself." +
				"\nBecause I'm going to eat you.\"" +
				"\nThe Fox eats you in one swallow.")
			cpH = 0
			heads = []bool{false, false, true}
		} else if heads[2] == false {
			fmt.Println("The rabbit manages a narrow victory. The turtle is defeated and it joins the fox in death.")
			//add more and test, change rabbit on turtle dialog
			bD14R = bD{"I showed him. No one left to tell me what to do.", []*pD{&pExit}}
			heads = []bool{true, false, false}
		}
	case &pKillF:
		bconv = false
		fmt.Println("The Turtle writhes away to confer with Rabbit. They simultaneously attack Fox." +
			"He snaps and tears at them with his fangs. But he cannot hold off their combined attack." +
			"Fox goes limp, dead. It dangles limp at the end of a long feathery neck. The bird body dips in the air, the wings hold the weight but look to be struggling.") //lazy cheat
		heads = []bool{true, true, false}
		//need to redefine conversations then
		bD11R = bD{"He was a jerk. Now he's dead. Thanks for reminding me.", []*pD{&pD14R}} //too much of a one way tangent
		bD20T = bD{"I've got nothing to fear from Fox. Especially now that he's dead." +
			"Does that make you the clever one or the predator?", []*pD{&pExit}}
	case &pPlayDD:
		bconv = false
	}
	return nbD, bconv, cpH, heads
}

type pD struct {
	words  string
	branch *bD
}

type bD struct {
	words    string
	branches []*pD
}

var (
	bDstartR = bD{"Start!", []*pD{&pDs1R, &pDs2R, &pDs3R, &pDs4R}}
	bDstartT = bD{"Start!", []*pD{&pDs1T, &pDs2T, &pDs3T, &pDs4T}}
	bDstartF = bD{"Start!", []*pD{&pDs1F, &pDs2F, &pDs3F, &pDs4F}}

	pDstarts = []string{"Do you like Devil Dice?", "How are you?", "Fight me!", "Leave me alone."}
	pDs1R    = pD{pDstarts[0], &bD1R}
	pDs2R    = pD{pDstarts[1], &bD2R}
	pDs3R    = pD{pDstarts[2], &bD3R}
	pDs4R    = pD{pDstarts[3], &bD4R}

	pDs1T = pD{pDstarts[0], &bD1T}
	pDs2T = pD{pDstarts[1], &bD2T}
	pDs3T = pD{pDstarts[2], &bD3T}
	pDs4T = pD{pDstarts[3], &bD4T}

	pDs1F = pD{pDstarts[0], &bD1F}
	pDs2F = pD{pDstarts[1], &bD2F}
	pDs3F = pD{pDstarts[2], &bD3F}
	pDs4F = pD{pDstarts[3], &bD4F}

	//events
	pExit   = pD{}
	pKillT  = pD{}
	pKillF  = pD{}
	pKillR  = pD{}
	pPlayDD = pD{}
	pBetray = pD{}

	//Need to make these different per head
	bD1R = bD{"I used to play but I stopped. Fox said it was a stupid game." +
		"\nI used to play all the time. Fox said I was an idiot for playing it. He's the idiot.", []*pD{&pD10R, &pD11R}} //devil dice, need a check for other heads
	bD2R = bD{"I'm great. I'm here, which is not the Shadowlands. So that's new. I've got a mission. From the Professor." +
		"\nI'm attached to two total jerks. That's not great.", []*pD{&pD20R, &pD11R}} //How are you
	bD3R = bD{"Hahahahaha. Bwahahahaha. You wouldn't last long.", []*pD{&pExit}} //fight me
	bD4R = bD{"Can't do. I'd love to fly off, explore this bright land. But I'm on a mission." +
		"\nWho gave you those devil dice?", []*pD{&pD40R, &pD41R}} // leave me alone

	pD10R = pD{"I think it's a great game", &bD10R}
	pD11R = pD{"What is up with that Fox?", &bD11R}

	bD10R = bD{"Show's what you know. It's a stupid game. A stupid, fun game I want to play so badly." +
		"\nFox would eat me if I even thought of it. Quiet human! You're so distracting.", []*pD{&pExit}}

	bD11R = bD{"He's a jerk. Always tell me what to do.", []*pD{&pD13R}} //too much of a one way tangent
	pD13R = pD{"Why don't you do something about him?", &bD13R}
	bD13R = bD{"I would! But he has teeth, and I'm soft and furry. I'm not tough like the turtle.", []*pD{&pD14R}}
	pD14R = pD{"What about the turtle?", &bD14R}
	bD14R = bD{"He's also a jerk. So pretentious. And slow. But tough.", []*pD{&pD140R, &pD15R}}

	pD140R = pD{"I hear Turtle is great.", &bD15R}
	pD15R  = pD{"I bet you could beat Turtle.", &bD16R}

	bD15R  = bD{"Great at what?", []*pD{&pD150R, &pD151R, &pD152R}}
	bD16R  = bD{"Like in a fight?", []*pD{&pD16R, &pD161R}}
	pD16R  = pD{"Yes", &bD160R}
	pD161R = pD{"No", &bD161R}

	bD160R = bD{"Haha, you want me to fight Turtle? I the quick-witted rabbit see through your ruse." +
		"I do not do what anyone tells me! Go away human", []*pD{&pExit}}

	bD161R = bD{"Good. Because I would utterly win. Now leave me alone human.", []*pD{&pExit}}

	pD150R = pD{"Everything.", &bD150R}
	pD151R = pD{"Kicking your ass.", &bD151R}
	pD152R = pD{"Racing.", &bD152R}

	bD150R = bD{"No! He's great at nothing. He's stupid and arrogant.", []*pD{&pD153R, &pD15R}}
	bD151R = bD{"What? No. Turtle is arrogant but is he violent?", []*pD{&pD153R, &pD15R}}
	bD152R = bD{"No! He's slow and stupid.", []*pD{&pD153R, &pD15R}}

	pD153R = pD{"He says you couldn't beat him at anything. That you're weak.", &bD153R}
	bD153R = bD{"No!", []*pD{&pD154R, &pD15R}}

	pD154R = pD{"And stupid. And easily manipulated!", &bD154R}

	bD154R = bD{"I'll show him!", []*pD{&pKillT}}

	pD20R = pD{"Who is the Professor?", &bD20R}
	bD20R = bD{"I shouldn't tell you anything about him." +
		"\nHe's a real big shot. And he's coming here. But I cannot tell you anything.", []*pD{&pExit}}

	pD40R = pD{"Veronica. She's at the lake.", &bD40R}
	bD40R = bD{"Thanks a lot for the information. And betraying your friend. Now I'm going to kill you. Nothing personal.", []*pD{&pBetray}}

	pD41R = pD{"No one, I found them.", &bD41R}
	bD41R = bD{"I don't believe you. Those are not typical human dice. The Professor doesn't want you having them.", []*pD{&pExit}}

	//turtle
	bD1T = bD{"It's a nice game. Too fast and short for my tastes. Also not enough blood.", []*pD{&pD1T, &pD10T}}
	bD2T = bD{"Same as always. I am the rock. Not fast, not clever.", []*pD{&pD2T, &pD20T}}
	bD3T = bD{"Are you in a hurry to die?", []*pD{&pD3T}}
	bD4T = bD{"I will leave you in peace. Once you tell me who gave you those devil dice.", []*pD{&pD4T, &pD40T}}

	pD1T  = pD{"Blood?", &bD11T}
	pD10T = pD{"I hear rabbit likes to play.", &bD12T}
	bD11T = bD{"Yes. Blood.", []*pD{&pExit}}
	bD12T = bD{"Rabbit used to play often. Fox does not approve. How about you? Do you play?", []*pD{&pD12T, &pD13T}}
	pD12T = pD{"Yes", &bD13T}
	pD13T = pD{"No. Never.", &bD14T}
	bD14T = bD{"Usless you have a set then. Who gave you the dice? I can return it to them.", []*pD{&pD4T, &pD40T}}
	bD13T = bD{"You know there are better dice sets? The one you have is quite basic." +
		"I'll give you a better set if you tell me who gave you those dice?", []*pD{&pD4T, &pD40T}}

	pD2T  = pD{"What do you think about Fox?", &bD20T}
	bD20T = bD{"I've got nothing to fear from Fox.", []*pD{&pD23T, &pD24T, &pD25T}}
	pD23T = pD{"Are you sure? In my world Foxes eat Turtles.", &bD23T}
	pD24T = pD{"Of course. You'll be fine.", &bD24T}
	pD25T = pD{"Totally.", &bD24T}

	bD23T = bD{"I have a shell.", []*pD{&pD26T, &pD27T}}
	bD24T = bD{"Now I am concerned. Are you sure?", []*pD{&pD28T, &pD29T}}

	pD28T = pD{"Nope. I wouldn't trust a Fox.", &bD23T}
	pD29T = pD{"Yes. You'll be fine.", &bD25T}

	bD25T = bD{"I have lost my happy garden. I must withdraw for a moment", []*pD{&pExit}}

	pD26T = pD{"No you don't! You have a bird body. Attached to the Fox!", &bD26T}
	pD27T = pD{"Sure you do. You'll be fine.", &bD25T}

	bD26T = bD{"I forgot. Concerning. But I cannot conquer Fox alone.", []*pD{&pD201T, &pD202T}}

	pD201T = pD{"Rabbit will help!", &bD201T}
	pD202T = pD{"I will help!", &bD202T}

	bD201T = bD{"True. Together maybe we can free ourselves from clever tyranny.", []*pD{&pKillF}}
	bD202T = bD{"But you are so small and human? The human we were sent to kill. I am confused and it is an unpleasant state.", []*pD{&pExit}}

	pD20T = pD{"What do you think about Rabbit?", &bD21T}
	bD21T = bD{"Rabbit is quick. Quick and fearful of being manipulated." +
		"But his fear of manipulation is just a manipulation of fear", []*pD{&pD21T, &pD22T}}
	pD21T = pD{"That's what I thought!", &bD22T}
	pD22T = pD{"What?", &bD22T}
	bD22T = bD{"Indeed. Can you tell me where you got those devil dice?", []*pD{&pD4T, &pD40T}}

	//ready to die
	pD3T  = pD{"No.", &bD22T}
	pD31T = pD{"Yes.", &bD22T}

	//tell me who gave you those dice
	pD4T  = pD{"Veronica. She's at the lake.", &bD40T}
	pD40T = pD{"No one, I found them.", &bD41T}

	bD40T = bD{"Thank you little human for your honesty. Unfortunately we will kill you now.", []*pD{&pBetray}}
	bD41T = bD{"That is unlikely. Possibly impossible. I will withdraw now and let you think about honesty.", []*pD{&pExit}}

	bD1F = bD{"It's an idiotic game. All luck. Tell me, who gave you the dice?", []*pD{&pD4F, &pD41F}}
	bD2F = bD{"I'm clever. And I'm not going to answer any of your questions until you answer mine. Who gave you those dice?", []*pD{&pD4F, &pD41F}}
	bD3F = bD{"That is not a smart request. You are the same size as my smallest fang. Instead of death, you should tell me who gave you those dice.", []*pD{&pD4F, &pD41F}}
	bD4F = bD{"We will. I just need to know who gave you those devil dice.", []*pD{&pD4F, &pD41F}}

	//pD1F  = pD{}
	//pD2F  = pD{}
	//pD3F  = pD{}
	pD4F  = pD{"Veronica. She's at the lake.", &bD40F}
	pD41F = pD{"No one, I found them.", &bD41F}

	bD40F = bD{"Thank you. For making this easy.", []*pD{&pBetray}}
	bD41F = bD{"Impossible. I will get the truth. Tell me!", []*pD{&pD4F, &pD42F}}

	pD42F = pD{"You're crazy. I found them!", &bDByeF}

	bDByeR = bD{"I don't have time for this.", []*pD{&pExit}}
	bDByeT = bD{"I will reflect on your words.", []*pD{&pExit}}
	bDByeF = bD{"Quiet human.", []*pD{&pExit}}
)

func choicemakerVB(tang bD) []string {
	var options []string
	n := len(tang.branches)
	for i := 0; i < n; i++ {
		br1 := *tang.branches[i]
		options = append(options, br1.words)
	}
	return options
}
