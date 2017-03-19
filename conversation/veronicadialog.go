package conversation

import (
	"RPG-Core/check"
	"RPG-Core/inputs"
	"RPG-Core/models"
	"fmt"
)

/*
Veronica

Need to setup initial conversations.

Let's say covering the first day in Camp. Introductions, explanations, and maybe something about shadow monsters.

Maybe too many options? I mean most of this stuff only happens on multiple playthroughs. Why not restrict it?
But then there's less engagement, less choice, and thinner story

Too flirty now. Very anime, need to dial it back
*/

//player introductions
var (
	//place intro lines here. Need to separate for convenience. Going to have a lot of dialog
	vht01 = dialog{branch1: &pht1, branch2: &pht12}
	vht02 = dialog{branch1: &pht2, branch2: &pht21}

	ph1    = []string{"What the hell? Who are you?", "Hello, can you excuse me for a moment?"}
	pht1   = dialog{words: ph1[0], branch1: &vht1}
	pht12  = dialog{words: ph1[1], branch1: &vht2}
	ph12   = []string{"Not funny.", "Can you turn around so I can get out?", "Can you think faster?"}
	pht121 = dialog{words: ph12[0], branch1: &vht121}
	pht122 = dialog{words: ph12[1], branch1: &vht123}
	pht123 = dialog{words: ph12[2], branch1: &vht122}

	ph13   = []string{"Still not funny.", "Please?"}
	pht131 = dialog{words: ph13[0], branch1: &vht123}
	pht132 = dialog{words: ph13[1], branch1: &vht123}

	//name1 = "Default"//no way to grab player's name yet
	ph14  = []string{"Whatever. My clothes are soaked.", "That's a nice name."}
	ph141 = dialog{words: ph14[0], branch1: &vh151} //needs a response
	ph142 = dialog{words: ph14[1], branch1: &vh152}
	//ph143 = dialog{words: ph140[2], branch1: &vtest1}

	ph15  = []string{"I think I know, but there's no way I'm saying it.", "Do you work at the camp?", "How old are you?"}
	ph151 = dialog{words: ph15[0], branch1: &vh161}
	ph152 = dialog{words: ph15[1], branch1: &vh162}
	ph153 = dialog{words: ph15[2], branch1: &vh163}

	ph16  = []string{"Ha, I'm going to go now.", "Do you need help?", "Who are you?"}
	ph161 = dialog{words: ph16[0], branch1: &vh171}
	ph162 = dialog{words: ph16[1], branch1: &vh172}
	ph163 = dialog{words: ph16[2], branch1: &vh173}

	ph17  = []string{"Are you sure? I'm pretty awesome.", "Well, let me know if I can help.", "I forget nothing!"}
	ph171 = dialog{words: ph17[0], branch1: &vh181}
	ph172 = dialog{words: ph17[1], branch1: &vh182}
	ph173 = dialog{words: ph17[2], branch1: &vh183}

	vh18  = []string{"Really? I haven't heard of you.", "Such a gentleman.", "So basically, you're an elephant?"}
	vh181 = dialog{words: vh18[0], branch1: &ph161} //make an event here to change depth? cannot go back to who are you? it creates a loop
	vh182 = dialog{words: vh18[1], branch1: &ph161}
	vh183 = dialog{words: vh18[2], branch1: &ph161, branch2: &ph162}

	vh17  = []string{"Good luck little man", "Possibly. Can you provide it? Probably not.", "I said before, I am Veronica. Did you forget already?"} //when does she reveal more? Need a depth check
	vh171 = dialog{words: vh17[0], branch1: &pGBt0}
	vh172 = dialog{words: vh17[1], branch1: &ph161, branch2: &ph171, branch3: &ph172}
	vh173 = dialog{words: vh17[2], branch1: &ph161, branch2: &ph162, branch3: &ph173}

	vh16  = []string{"Free you mind little one.", "I've heard of camp. Does that count?", "Old enough to know you're not supposed to ask a lady that."}
	vh161 = dialog{words: vh16[0], branch1: &ph152, branch2: &ph153, branch3: &ph162}
	vh162 = dialog{words: vh16[1], branch1: &ph161, branch2: &ph162, branch3: &ph163}
	vh163 = dialog{words: vh16[2], branch1: &ph161, branch2: &ph162, branch3: &ph163}

	ph1Test  = dialog{words: "test", branch1: &vGBt0}
	ph1Event = dialog{words: "Event!", branch1: &vtest1} //for triggering events
	ph1Exit  = dialog{words: "odd test", branch1: &vtest1}

	ph2   = []string{"You again?", "Hello Veronica."}
	pht2  = dialog{words: ph2[0], branch1: &vht21}
	pht21 = dialog{words: ph2[1], branch1: &vht22}

	pdef = []string{"How are you?", "What's happening?", "Can you help me?"}
	pdt1 = dialog{words: pdef[0], branch1: &vct1}
	pdt2 = dialog{words: pdef[1], branch1: &vct2}
	pdt3 = dialog{words: pdef[2], branch1: &vct3}

	vca1 = []string{"I'm feeling like p words. Perky, playful, poignant. What letter are you?",
		"That's cruelly vague. With what sweetheat?", "Probably not. But I can humor you."}
	vct1 = dialog{words: vca1[0], branch1: &pqt11, branch2: &pqt12, branch3: &pqt13}
	vct2 = dialog{words: vca1[1], branch1: &pqt21, branch2: &pqt22}
	//need to check for smiler event, kind of annoying this happens in two places
	vct3 = dialog{words: vca1[2], branch1: &ph4smiler}
	//for when player has not seen the smiler
	vct36 = dialog{words: vca1[2], branch1: &pqt31}
	//when the player has seen the smiler
	vct37     = dialog{words: vca1[2], branch1: &pqt31, branch2: &pqt32}
	ph4smiler = dialog{words: "Smiler Event x2!", branch1: &vtest1} //arg! events need stuff or they don't work

	pqv1  = []string{"F for frustrated.", "H for happy.", "L for lost."} //event with entering letters?
	pqt11 = dialog{words: pqv1[0], branch1: &vct21}
	pqt12 = dialog{words: pqv1[1], branch1: &vct22}
	pqt13 = dialog{words: pqv1[2], branch1: &vct23}

	vca2 = []string{
		"F for fortunate. As in unfortunate to hear that. We are always driven to frustation and dissatisfaction by the demons within us.",
		"G for glad to hear that. L for I think you are a liar. A for And. D for a Dbag.",
		"Maps are overrated. I found asking people for directions far preferable. You don't always get where you wanted to go, but it's always a good adventure."}
	vct21 = dialog{words: vca2[0], branch1: &pqt41, branch2: &pqt42, branch3: &pqt43}
	vct22 = dialog{words: vca2[1], branch1: &pqt51, branch2: &pqt52, branch3: &pqt43}
	vct23 = dialog{words: vca2[2], branch1: &pqt41, branch2: &pqt42, branch3: &pqt43}

	pqv4  = []string{"What?", "I absolutely agree.", "Right. I'm going now."}
	pqt41 = dialog{words: pqv4[0], branch1: &vct51}
	pqt42 = dialog{words: pqv4[1], branch1: &vct52}
	pqt43 = dialog{words: pqv4[2], branch1: &vGBt3}

	vca5  = []string{"Exactly.", "Really? You're not just kissing my ass?"}
	vct51 = dialog{words: vca5[0], branch1: &ph6event}
	vct52 = dialog{words: vca5[1], branch1: &pqt511, branch2: &pqt512}

	pqv51  = []string{"No.", "Yes."}
	pqt511 = dialog{words: pqv51[0], branch1: &vct61}
	pqt512 = dialog{words: pqv51[1], branch1: &vct62}
	vca6   = []string{"That's disapointing. I bet you're good at it.", "H for honesty."}
	vct61  = dialog{words: vca6[0], branch1: &pdt2, branch2: &pdt3, branch3: &pqt43}
	vct62  = dialog{words: vca6[1], branch1: &pdt2, branch2: &pdt3, branch3: &pqt43}

	pqv5  = []string{"No. I'm actually happy.", "You caught me. I'm miserable."}
	pqt51 = dialog{words: pqv5[0], branch1: &vct91}
	pqt52 = dialog{words: pqv5[1], branch1: &vct92}

	vca9  = []string{"Why is that?"}
	vct91 = dialog{words: vca9[0], branch1: &pqt5101, branch2: &pqt5102, branch3: &pqt5103}
	vct92 = dialog{words: vca9[0], branch1: &pqt521, branch2: &pqt522, branch3: &pqt43}

	pqv510  = []string{"I like being outside.", "I like the other Crew members.", "I like you."}
	pqt5101 = dialog{words: pqv510[0], branch1: &vct5101}
	pqt5102 = dialog{words: pqv510[1], branch1: &vct5102}
	pqt5103 = dialog{words: pqv510[2], branch1: &vct5103}

	vca510 = []string{"The great outdoors! I'm always out here. We'll see if you're still a fan after nightfall.",
		"\nI haven't met them. But are you sure they are your friends?",
		"\nHehe. Of course you do."}
	vct5101 = dialog{words: vca510[0], branch1: &pqt520}
	vct5102 = dialog{words: vca510[1], branch1: &pqt531, branch2: &pqt532, branch3: &pqt533}
	vct5103 = dialog{words: vca510[2], branch1: &pdt2, branch2: &pdt3, branch3: &pqt43}

	pqv520 = []string{"What happens after the sun goes down?"}
	pqt520 = dialog{words: pqv520[0], branch1: &vct5200}
	vca520 = []string{"Things out here get interesting." +
		"\nDuring the day it is a world of light, where the shadows are hiding behind things." +
		"\nAt night it is a world of darknes, where the light is small, scattered and everything is in shadow."}
	vct5200 = dialog{words: vca520[0], branch1: &pqt5200, branch2: &pqt5201}
	pqv5200 = []string{"That doesn't make any sense. You need light to make shadows.", "Well I'm scared now."}
	pqt5200 = dialog{words: pqv5200[0], branch1: &vct5201}
	pqt5201 = dialog{words: pqv5200[1], branch1: &vct5202}
	vca5200 = []string{"That's one way to look at it.",
		"\nHehe. Poor thing. If you're scared, you can always come here to the lake. Then at least I'll be the one scaring you."}
	vct5201 = dialog{words: vca5200[0], branch1: &pdt2, branch2: &pdt3, branch3: &pqt43}
	vct5202 = dialog{words: vca5200[1], branch1: &pdt2, branch2: &pdt3, branch3: &pqt43}

	pqv530 = []string{"Of course.", "I think so. Although my best friend isn't here.", "I'm not sure."}
	pqt531 = dialog{words: pqv530[0], branch1: &vct5300}
	pqt532 = dialog{words: pqv530[1], branch1: &vct5301}
	pqt533 = dialog{words: pqv530[2], branch1: &vct5302}

	vca5300 = []string{"That's a nice thought.", "And where is he?", "I typically assume everyone is my friend. Makes things easier."}
	vct5300 = dialog{words: vca5300[0], branch1: &pdt2, branch2: &pdt3, branch3: &pqt43}
	vct5301 = dialog{words: vca5300[1], branch1: &pqt521, branch2: &pqt522}
	vct5302 = dialog{words: vca5300[2], branch1: &pqt5300, branch2: &pqt5301}

	pqv5300 = []string{"That's stupid.", "Does that mean we're friends?"}
	pqt5300 = dialog{words: pqv5300[0], branch1: &vct5311}
	pqt5301 = dialog{words: pqv5300[1], branch1: &vct5312}
	vca5301 = []string{"And you're a jerk.", "Oh we are much more than friends."}
	vct5311 = dialog{words: vca5301[0], branch1: &pdt2, branch2: &pdt3, branch3: &pqt43}
	vct5312 = dialog{words: vca5301[1], branch1: &pdt2, branch2: &pdt3, branch3: &pqt43}

	pqv52  = []string{"I don't want to talk about it.", "Maybe once we know each other better."}
	pqt521 = dialog{words: pqv52[0], branch1: &vct521}
	pqt522 = dialog{words: pqv52[1], branch1: &vct522}

	vca52 = []string{"I did not expect you to be the dark and brooding type. Secrets are treasures, why not hoard?",
		"\nYou want to get to know me? How fun. It will be your pleasure."}
	vct521 = dialog{words: vca52[0], branch1: &pdt2, branch2: &pdt3, branch3: &pqt43}
	vct522 = dialog{words: vca52[1], branch1: &pdt2, branch2: &pdt3, branch3: &pqt43}

	// what's happening?
	pqv2  = []string{"With the forest.", "With the other campers."}
	pqt21 = dialog{words: pqv2[0], branch1: &vct31}
	pqt22 = dialog{words: pqv2[1], branch1: &vct32}
	//pqt23 = dialog{words: pqv2[2], branch1: &vct33}
	vca3 = []string{"It's a special place. Heavy with age and shadows. Did you see something?",
		"I haven't met any of them. But this forest does have an impact on people. Did you notice something?"}
	vct31 = dialog{words: vca3[0], branch1: &ph3smiler}
	vct32 = dialog{words: vca3[1], branch1: &pqt301} //need something about other campers
	//for if player has seen a smiler
	vct33 = dialog{words: vca3[0], branch1: &pqt301, branch2: &pqt302}
	//for when player has not seen a smiler
	vct35 = dialog{words: vca3[0], branch1: &pqt301}

	ph3smiler = dialog{words: "Smiler Event!", branch1: &vtest1} //arg! events need stuff or they don't work
	pqv30     = []string{"No, nothing.", "I saw some creature in the forest. It was all teeth."}
	pqt301    = dialog{words: pqv30[0], branch1: &vct69} //why is this an exit?
	pqt302    = dialog{words: pqv30[1], branch1: &vct301}
	vca31     = []string{"Adorable little bastards aren't they? I call them Smilers. Nasty bite of course. They like to play devil dice."}
	vct301    = dialog{words: vca31[0], branch1: &pqt310, branch2: &pqt40} //apparantly allowed
	pqv31     = []string{"That is not a normal animal. What are they?"}
	pqt310    = dialog{words: pqv31[0], branch1: &vct320}
	vca32     = []string{"Well how do you define a normal animal? Is a tapir a normal animal?" +
		"\nAnyways, they are not animal. Or from here. They are denizen of the Shadowlands."}
	vct320 = dialog{words: vca32[0], branch1: &pqt71, branch2: &pqt72}
	pqv7   = []string{"What do you mean Shadowlands?", "What is it doing here?"}
	pqt71  = dialog{words: pqv7[0], branch1: &vct81}
	pqt72  = dialog{words: pqv7[1], branch1: &vct82}
	vca8   = []string{"You're so innocent. Not now. Ask me another time Doctor Questions.",
		"Exile? Exploration? Exploitation? Vacation? I have no idea."}
	vct81 = dialog{words: vca8[0], branch1: &pdt3, branch2: &pqt72, branch3: &pGBt2}
	vct82 = dialog{words: vca8[1], branch1: &pdt3, branch2: &pGBt2}

	vca60    = []string{"Cool beans."}
	vct69    = dialog{words: vca60[0], branch1: &ph6event}
	ph6event = dialog{words: "Defaults event!", branch1: &vtest1}
	vct70    = dialog{words: vca60[0], branch1: &pdt1, branch2: &pdt2, branch3: &pdt3}

	//can you help me?
	pqv3  = []string{"Ha. Nevermind.", "Do you know how I can get past the monsters in the forest?"} //should be an event first to allow this option
	pqt31 = dialog{words: pqv3[0], branch1: &vct69}
	pqt32 = dialog{words: pqv3[1], branch1: &vct401}

	vca40  = []string{"I wouldn't recommend petting it. Or heavy petting. It's a Shadow, it will probably play Devil Dice."}
	vct401 = dialog{words: vca40[0], branch1: &pqt40}

	pqv40 = []string{"What is Devil Dice?"}
	pqt40 = dialog{words: pqv40[0], branch1: &vct42}

	vca4 = []string{"I know a bit about it. Certainly more than you." +
		"\nIt's a game from the Shadowlands. They like to play games, no different from people." +
		"\nIt's considered a dignified alternative to kicking and biting and stabbing."}
	//vct41 = dialog{}
	vct42 = dialog{words: vca4[0], branch1: &pqt410, branch2: &pqt411}

	pqv41  = []string{"How do I play", "That's stupid"}
	pqt410 = dialog{words: pqv41[0], branch1: &vct41}
	pqt411 = dialog{words: pqv41[1], branch1: &vct69}

	vca41 = []string{"You roll dice. And you get three rolls. But they're special dice. They have shadows, moons and suns on them." +
		"\nYou're human right? So the shadows will hurt you. The suns will hurt a shadow creature. The moons give you extra rolls." +
		"\nSo simple even a boy could do it. Even an incredibly handsome but stupid boy."}
	vct41 = dialog{words: vca41[0], branch1: &pqt420, branch2: &pqt421}

	pqv42  = []string{"Special Dice? Where can I get those?", "You're stupid handsome."}
	pqt420 = dialog{words: pqv42[0], branch1: &vct420} //need event to get dice
	pqt421 = dialog{words: pqv42[1], branch1: &vct421}

	vca42 = []string{"Today is your lucky day tiger. I have a set that's just gathering dust." +
		"\nIt's a basic set. Nothing fancy but they'll let you play against someone else." +
		"\nThey are not a toy. Well they are a toy. But a toy that can actually hurt you. Like a trampoline or a matchbook.",
		"Thank you. You know, I have a set of dice that's just gathering dust." +
			"\nIt's a basic set. Nothing fancy but they'll let you play against someone else." +
			"\nThey are not a toy. Well they are a toy. But a toy that can actually hurt you. Like a trampoline or a matchbook."}
	vct420       = dialog{words: vca42[0], branch1: &pqvDiceEvent}
	vct421       = dialog{words: vca42[1], branch1: &pqvDiceEvent}
	pqvDiceEvent = dialog{words: "Dice event", branch1: &vtest1}

	vct422 = dialog{words: "There you are. Enjoy. Don't kill yourself.", branch1: &pqt430, branch2: &pqt431, branch3: &pqt432}

	pqv43  = []string{"Thanks.", "I'll be careful.", "Whatever."}
	pqt430 = dialog{words: pqv43[0], branch1: &vct430}
	pqt431 = dialog{words: pqv43[1], branch1: &vct430}
	pqt432 = dialog{words: pqv43[2], branch1: &vct430}

	vca43  = []string{"Indeed. See you around handsome."}
	vct430 = dialog{words: vca43[0], branch1: &pGBt0} //finish

	//these are not being used?
	pqv6  = []string{"Do you know how to reach the mesa?"}
	pqt61 = dialog{words: pqv6[0], branch1: &vct71}
	vca7  = []string{"I absolutely do not. And if I did, I would not tell you. Too close to the Shadowlands."}
	//it's like a portal to hell. But instead a portal to the shadowlands.
	vct71 = dialog{words: vca7[0], branch1: &pqt71, branch2: &pqt43}

	//messy goodbyes
	pvGB   = []string{"Leave me alone.", "Lates.", "Bye. See you soon."}
	pGBt0  = dialog{}
	pGBt1  = dialog{words: pvGB[0], branch1: &vGBt1}
	pGBt2  = dialog{words: pvGB[1], branch1: &vGBt2}
	pGBt3  = dialog{words: pvGB[2], branch1: &vGBt3}
	pGBray = []dialog{pGBt1, pGBt2, pGBt3}
)

//Veronica lines
var (
	vh1  = []string{"Haha. Pleasure to meet you too.", "Maybe. Let me think about it."}
	vht1 = dialog{words: vh1[0], branch1: &pht121, branch2: &pht122}
	vht2 = dialog{words: vh1[1], branch1: &pht121, branch2: &pht122, branch3: &pht123}

	vh12   = []string{"Maybe the audience is amused? I'm not funny just for you.", "Still thinking."}
	vht121 = dialog{words: vh12[0], branch1: &pht131, branch2: &pht132}
	vht122 = dialog{words: vh12[1], branch1: &pht131, branch2: &pht132}

	vh13   = []string{"Ok, I'll turn around. And I promise to peek only twice."}
	vht123 = dialog{words: vh13[0], branch1: &ph1Event}

	vh14  = []string{"My name is Veronica."}
	vh124 = dialog{words: vh14[0], branch1: &ph141, branch2: &ph142}

	vh151 = dialog{words: "Then why did you put them on?", branch1: &ph152, branch2: &ph153}
	vh152 = dialog{words: "What is it about names that start with a V? They drive the boys wild.", branch1: &ph151, branch2: &ph152, branch3: &ph153}

	vh2   = []string{"Indeed. I'm incorrigible.", "Hi Tiger."}
	vht21 = dialog{words: vh2[0], branch1: &pdt1, branch2: &pdt2, branch3: &pdt3}
	vht22 = dialog{words: vh2[1], branch1: &pdt1, branch2: &pdt2, branch3: &pdt3}

	vtest1 = dialog{words: "No idea what you're going on about", branch1: &pGBt0}

	vGB   = []string{"I'll try.", "Not funny. That Mike is a bad influence.", "Bye"}
	vGBt0 = dialog{words: "Odd test", branch1: &ph1Exit}
	vGBt1 = dialog{words: vGB[0], branch1: &ph1Exit}
	vGBt2 = dialog{words: vGB[1], branch1: &ph1Exit}
	vGBt3 = dialog{words: vGB[2], branch1: &ph1Exit}
)

//ConverserV handles conversation with Veronica
func ConverserV(cc Convo) Convo {
	vm := check.Eventcheck(5) //works but too basic
	var options []string
	var ptan dialog
	var ctan dialog
	//no function exists yet to get the player's name!
	//hellos first
	if vm == false { //check if they met
		ptan = vht01
	} else {
		ptan = vht02
	}
	options = choicemaker(ptan)
	r1 := inputs.StringarrayInput(options)
	switch r1 {
	case 1:
		ptan = *ptan.branch1
	case 2:
		ptan = *ptan.branch2
	case 3:
		ptan = *ptan.branch3
	}
	ctan = *ptan.branch1 //sets Vs response
	//hellos done, have ctan decided
	for cc.stilltalking == true {
		v1 := "\"" + ctan.words + "\"" //displays Vs response
		fmt.Println(v1)
		options = choicemaker(ctan)
		r1 = inputs.StringarrayInput(options)
		switch r1 {
		case 1:
			ptan = *ctan.branch1
		case 2:
			ptan = *ctan.branch2
		case 3:
			ptan = *ctan.branch3
		}
		ctan = *ptan.branch1 //sets Vs response
		ctan, cc = dialogchecker(ctan, cc)
	}
	cc.qa = "GB"
	return cc
}

func dialogchecker(ctan dialog, cc Convo) (dialog, Convo) {
	var options []string
	var r1 int
	var v1 string
	switch {
	case *ctan.branch1 == pGBt0: //if there is no branch in the dialog, then it defaults to this one
		//fmt.Println("Test pGBt0 event triggered")
		cc.stilltalking = false
		v1 = "\"" + ctan.words + "\"" //displays Vs response
		fmt.Println(v1)
		options = pvGB
		r1 = inputs.StringarrayInput(options)
		ctan = *pGBray[r1-1].branch1
		v1 = "\"" + ctan.words + "\""
		fmt.Println(v1)
	case *ctan.branch1 == ph1Exit: //for normal ending the conversation
		//fmt.Println("Test exit event triggered")
		v1 = "\"" + ctan.words + "\"" //displays Vs response
		fmt.Println(v1)
		cc.stilltalking = false
	case *ctan.branch1 == ph1Event: //this seems to work, needs to be the branches
		v1 := "\"" + ctan.words + "\"" //displays Vs response
		fmt.Println(v1)
		fmt.Println("She faces into the forest. You dash to the shore and rapidly put on your clothes.")
		ctan = vh124 //need a new ctan
	case *ctan.branch1 == ph3smiler: //trying to get smiler event checked
		//fmt.Println("Test smiler event triggered")
		//for when player sees smiler
		sc := check.Eventcheck(2)
		if sc == true {
			//fmt.Println("Test smiler event triggered 1")
			ctan = vct33
		} else {
			//fmt.Println("Test smiler event triggered 2")
			ctan = vct35
		}
	case *ctan.branch1 == ph4smiler: //trying to get smiler event checked
		//fmt.Println("Test smiler event triggered")
		//for when player sees smiler
		sc := check.Eventcheck(2)
		if sc == true {
			//fmt.Println("Test smiler event triggered 1")
			ctan = vct37
		} else {
			//fmt.Println("Test smiler event triggered 2")
			ctan = vct36
		}
	case *ctan.branch1 == ph6event: //return to defaults
		ctan = vct70
	case *ctan.branch1 == pqvDiceEvent: //not checking if player already has them? Not checking if backpack is full
		dd := models.ItemGetByName("devil dice")
		if dd.Loc == 20 {
			fmt.Println("\"But I see you already have a pair. Possibly my old pair. You hoarder.\"")
			ctan = vct422 // can change this later, cheat
		} else {
			v1 := "\"" + ctan.words + "\"" //displays Vs response
			fmt.Println(v1)
			fmt.Println("She hands you a set of dice.")
			dd.Loc = 20
			models.ItemUpdate(dd)
			ctan = vct422
		}
	}
	return ctan, cc
}

func choicemaker(tang dialog) []string {
	var options []string
	br1 := *tang.branch1
	options = append(options, br1.words)
	if tang.branch2 != nil {
		br2 := *tang.branch2
		options = append(options, br2.words)
		if tang.branch3 != nil {
			br3 := *tang.branch3
			options = append(options, br3.words)
		}
	}
	return options
}
