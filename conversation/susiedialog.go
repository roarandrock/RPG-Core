package conversation

import (
	"RPG-Core/check"
	"RPG-Core/inputs"
	"RPG-Core/models"
	"fmt"
)

/*
Notes on conversations for future ones:

1. Can converser be made non-character specific and the events the only character things?

*/

var (
	sihstart = cD{"Odd Silence", []*pD{&pihs0, &pihs1}, 0}

	pihs  = []string{"Hi", "Nice view"}
	pihs0 = pD{words: pihs[0], branch: &sih0, dChange: 0}
	pihs1 = pD{words: pihs[1], branch: &sih0, dChange: 0}

	sih  = []string{"AAAAAAA!"}
	sih0 = cD{sih[0], []*pD{&psiQ0, &psiQ1}, 0}

	psiQs = []string{"Sorry! You okay?", "Haha. Don't fall off."}
	psiQ0 = pD{psiQs[0], &siA0, 0}
	psiQ1 = pD{psiQs[1], &siA1, -1}

	siAs = []string{"You scared the shit out of me. I'm fine.\nWhat are you doing here?",
		"Real funny. What are you doing here?"}
	siA0 = cD{siAs[0], []*pD{&pisQ2, &pisQ3, &pisQ4, &pisQ40}, 0}
	siA1 = cD{siAs[1], []*pD{&pisQ2, &pisQ3, &pisQ4, &pisQ40}, 0}

	pisQ2s = []string{"I need some help.", "Just wanted to meet you.", "I'm exploring."}
	pisQ2  = pD{pisQ2s[0], &siA2, 0}
	pisQ3  = pD{pisQ2s[1], &siA3, 0}
	pisQ4  = pD{pisQ2s[2], &siA4, 1}

	siA2s = []string{"I don't really know you. But I do know you are a fellow crewmember. What's up?",
		"Right. Mission accomplished. Nice to meet you. Anything else?",
		"Cool. This is a fantastic camp. It's huge. And it feels so alive."}
	siA2 = cD{siA2s[0], []*pD{&pisQ10, &pisQ11, &pisQ12}, 0}
	siA3 = cD{siA2s[1], []*pD{&pisQ2, &pisQ4, &pisQ40}, 0}
	siA4 = cD{siA2s[2], []*pD{&pisQ30, &pisQ31, &pisQ32, &pisQ40}, 0}

	//can use these in normal convo too
	pisQ10s = []string{"I need a map.", "There are monsters in the forest!",
		"I met another girl down by the lake."} //no need to check, can only get to Susie after talking to Veronica
	pisQ10 = pD{pisQ10s[0], &siA10, 0}
	pisQ11 = pD{pisQ10s[1], &siA11, 0}
	pisQ12 = pD{pisQ10s[2], &siA12, 0}

	siA10s = []string{"Honestly I don't have one anymore. I got this route memorized. But maybe there are some in the old cabin.",
		"What? You mean the bears and mosquitos?",
		"I'm not the only girl in the crew."}
	siA10 = cD{siA10s[0], []*pD{&pisQ110}, 5}
	siA11 = cD{siA10s[1], []*pD{&pisQ16, &pisQ13}, 0}
	siA12 = cD{siA10s[2], []*pD{&pisQ17, &pisQ18}, 0} //need a follow up

	//redirect to cabin
	pisQ110s = []string{"What cabin?"}
	pisQ110  = pD{pisQ110s[0], &siA51, 0}

	//monsters
	pisQ12s = []string{"No. I mean the shadows. Like the little guys with the teeth and the dice.", "Wait, there are bears here?"}
	pisQ16  = pD{pisQ12s[0], &siA16, 0}
	pisQ13  = pD{pisQ12s[1], &siA13, 0}

	siA12s = []string{"What the what? No idea what you are talking about. Is this a joke? Did Mike tell you this?",
		"Of course there are. Why do you think we have a bear bag? Technically we shouldn't be hiking alone. Terrible role models."}
	//second one is an exit
	siA16 = cD{siA12s[0], []*pD{&pisQ14, &pisQ15, &pisQ40}, 0}
	siA13 = cD{siA12s[1], []*pD{&pisQ14, &pisQ40}, 0}

	pisQ121s = []string{"No. This is something I actually saw. And I kicked its butt.", "Yep. You got me. Total prank."}
	pisQ14   = pD{pisQ121s[0], &siA14, 0}
	pisQ15   = pD{pisQ121s[1], &siA15, 1}

	siA121s = []string{"Right. I think you've played too many video games.",
		"That Mike. Probably got the idea from one of the ghost stories"}
	//later she opens up about seeing something?
	siA14 = cD{siA121s[0], []*pD{&pDefault}, 0}
	siA15 = cD{siA121s[1], []*pD{&pDefault}, 0}

	//not only girl
	pisQ17s = []string{"I know that. Do you know Veronica? I haven't seen her before.", "Wait, you're a girl?"}
	pisQ17  = pD{pisQ17s[0], &siA17, 0}
	pisQ18  = pD{pisQ17s[1], &siA18, 1}
	siA17s  = []string{"There's no Veronica in our Crew. Maybe she's from another Crew? We're not then only one at camp.",
		"Just because I can kick your butt, doesn't mean I'm not a girl."}
	siA17 = cD{siA17s[0], []*pD{&pisQ19, &pisQ20}, 0}
	siA18 = cD{siA17s[1], []*pD{&pDefault}, 0}

	pisQ19s = []string{"That makes sense.", "I suspect she's not from any Crew."}
	pisQ19  = pD{pisQ19s[0], &siA19, 1}
	pisQ20  = pD{pisQ19s[1], &siA20, 0}
	siA19s  = []string{"I am good at sense making. Anything else?", "That doesn't sound good. Be careful."}
	siA19   = cD{siA19s[0], []*pD{&pDefault}, 0}
	siA20   = cD{siA19s[1], []*pD{&pDefault}, 0}

	//exploring
	pisQ30s = []string{"It is a giant place. I keep getting lost.", "Full of sun and shadows.", "Have you seen anything cool?"}
	pisQ30  = pD{pisQ30s[0], &siA30, 0}
	pisQ31  = pD{pisQ30s[1], &siA31, 0}
	pisQ32  = pD{pisQ30s[2], &siA32, 1}

	siA30s = []string{"That's how you really learn to navigate. If you always know where you are, you never have to try.",
		"I guess. It's certainly sunny up here. In the forest, it always feels a bit darker.",
		"There's lots of stuff to see. The coolest spot is up here. I like it up here. I've seen some creepy places too."}
	siA30 = cD{siA30s[0], []*pD{&pisQ33, &pisQ34}, 0}
	siA31 = cD{siA30s[1], []*pD{&pisQ35, &pisQ36}, 0}
	siA32 = cD{siA30s[2], []*pD{&pisQ50, &pisQ51}, 3}

	pisQ33s = []string{"Wow. I get that. Cool.", "Lame. I'd rather save time and just use a map."}
	pisQ33  = pD{pisQ33s[0], &siA33, 1}
	pisQ34  = pD{pisQ33s[1], &siA34, -1}

	siA33s = []string{"Cool.", "It's your life."}
	siA33  = cD{siA33s[0], []*pD{&pDefault}, 0}
	siA34  = cD{siA33s[1], []*pD{&pDefault}, 0}

	pisQ35s = []string{"It's great up here.", "What do you mean darker?"}
	pisQ35  = pD{pisQ35s[0], &siA35, 1}
	pisQ36  = pD{pisQ35s[1], &siA36, 0}

	siA35s = []string{"That's why I'm here.",
		"It's hard to talk describe. Like the shadows are thicker, heavier. And it's quiet. Like the birds are not as noisey as they should be."}
	siA35 = cD{siA35s[0], []*pD{&pisQ38, &pisQ36}, 0}
	siA36 = cD{siA35s[1], []*pD{&pisQ11, &pisQ37}, 4}

	pisQ37s = []string{"Like it's cursed?", "Cool."}
	pisQ37  = pD{pisQ37s[0], &siA37, 0}
	pisQ38  = pD{pisQ37s[1], &siA33, 0}

	siA37s = []string{"No, not cursed. Just different. Like alien. It's just a unique place. Unlike anywhere I've been before."}
	siA37  = cD{siA37s[0], []*pD{&pDefault}, 0}

	//creepy places too
	pisQ50s = []string{"What do you like about it?", "What was the creepiest?"}
	pisQ50  = pD{pisQ50s[0], &siA50, 1}
	pisQ51  = pD{pisQ50s[1], &siA51, 0}

	siA50s = []string{"Everything. The exertion to get here. The view of the camp. The sun and the solitude.",
		"The creepiest place I've been is the abandoned cabin. Did you hear about that place?"}
	siA50 = cD{siA50s[0], []*pD{&pDefault}, 0}
	siA51 = cD{siA50s[1], []*pD{&pisQ52, &pisQ53}, 3}

	pisQ52s = []string{"No.", "Yes."}
	pisQ52  = pD{pisQ52s[0], &siA52, 0}
	pisQ53  = pD{pisQ52s[1], &siA53, 0}

	siA52s = []string{"I heard about it at the campfire, I think last night. A bunch of campers died in there." +
		"\nEveryone I think, except for one kid who's hair was turned white. Like five years back.",
		"Right. Well I went there."}
	siA52 = cD{siA52s[0], []*pD{&pisQ54, &pisQ55}, 0}
	siA53 = cD{siA52s[1], []*pD{&pisQ55}, 3}

	pisQ54s = []string{"You are not good at telling stories.", "Did you go inside?"}
	pisQ54  = pD{pisQ54s[0], &siA54, -1}
	pisQ55  = pD{pisQ54s[1], &siA59, 0}

	siA54s = []string{"You're just a bad listener. Also, you're right. It's not my strength."}
	siA54  = cD{siA54s[0], []*pD{&pisQ56, &pisQ57, &pisQ58}, 0}

	pisQ56s = []string{"Can't be great at everything. And you're great at many other things.",
		"My stories are great at putting people to sleep.",
		"We agree then. You suck."}
	pisQ56 = pD{pisQ56s[0], &siA56, 2}
	pisQ57 = pD{pisQ56s[1], &siA57, 2}
	pisQ58 = pD{pisQ56s[2], &siA58, -1}

	siA55s = []string{"Thanks.", "Haha. Well that could be useful.", "Jerk. For that, no more story."}
	siA56  = cD{siA55s[0], []*pD{&pisQ55}, 0}
	siA57  = cD{siA55s[1], []*pD{&pisQ55}, 0}
	siA58  = cD{siA55s[2], []*pD{&pDefault}, 0}

	siA59s = []string{"I did. And it was full of ghosts!"}
	siA59  = cD{siA59s[0], []*pD{&pisQ59, &pisQ60}, 4}

	pisQ59s = []string{"No!", "I don't believe you."}
	pisQ59  = pD{pisQ59s[0], &siA60, 1}
	pisQ60  = pD{pisQ59s[1], &siA61, 0}

	siA60s = []string{"Yes! I fooled you! There were no ghosts. It was really creepy though." +
		"\nIt was really dark in there. Full of dust. And old gear. Like no one had wanted to touch any of it." +
		"\nEven the animals avoided it. No spiderwebs, just dust.",
		"Damn. There were no ghosts. It was really creepy though." +
			"\nIt was really dark in there. Full of dust. And old gear. Like no one had wanted to touch any of it." +
			"\nEven the animals avoided it. No spiderwebs, just dust."}
	siA60 = cD{siA60s[0], []*pD{&pisQ62}, 0}
	siA61 = cD{siA60s[1], []*pD{&pisQ62}, 0}

	pisQ62s = []string{"Crazy. How do I get there?"}
	pisQ62  = pD{pisQ62s[0], &siA62, 0}

	siA62s = []string{"Through the forest. Let me show you on my map."} //make an event for this
	siA62  = cD{siA62s[0], []*pD{&pCabin}, 4}
	//can talk about the mesa here or later?

	//reusing the intro questions in the repeat. So they can be long.
	pisQ40 = pD{"Later. Just going to enjoy the view.", &siA40, 1}
	siA40  = cD{"Cool. See you around.", []*pD{&pExit}, 0}

	//need Normal/2nd conversation
	sstart = cD{"Odd Silence", []*pD{&psH0, &psH1}, 0}
	psHs   = []string{"Hello again.", "Hold on, I need to catch my breath."}
	psH0   = pD{psHs[0], &sH0, 0}
	psH1   = pD{psHs[1], &sH1, 0}
	sHs    = []string{"Hello.", "Take a breather."}
	sH0    = cD{sHs[0], []*pD{&pisQ2, &pisQ4, &pisQ40}, 0}
	sH1    = cD{sHs[1], []*pD{&pisQ2, &pisQ4, &pisQ40}, 0}

	//I need some help, goes back to the intro
	//And explore
	//Add one for repeat conversations?
	//need some way to bump up depth?
	//need another way to get the map, or at least explain why you need it better
	//need the map to get to the cabin?

	//Triggers
	/* can reuse:
	pExit    = pD{} //exit
	piqTest  = pD{}
	pDefault = pD{}
	*/
	//pMap   = pD{}
	pCabin = pD{}
	//Depth fail
	sDfail = cD{"Um, no. Sorry, cannot help you there.", []*pD{&pDefault}, 0}
)

//ConverserS is for conversations with Susie
//Need to update for Susie
func ConverserS(cc Convo) Convo {
	sm := check.Eventcheck(6)
	var options []string
	var npD pD
	var ncD cD
	//no function exists yet to get the player's name!
	//hellos first
	if sm == false { //check if they met
		ncD = sihstart
		jb := models.StoryblobGetByName(6) //need to make still
		fmt.Println(jb.Story)
		jb.Shown = true
		models.StoryblobUpdate(jb)
	} else {
		ncD = sstart
	}
	// need loop
	//Display options - check for valid pD or remove and attach only to events?
	//Player chooses
	//Display response -check for valid response/depth
	//check to continue

	for cc.stilltalking == true {
		options = choicemakerV2(ncD)
		r1 := inputs.StringarrayInput(options)
		npD = *ncD.branches[r1-1]
		//change depth
		if npD.dChange < 0 {
			cc.depth = depthChange(cc.depth, npD.dChange, 1) //need negative. so 1 min and 5 max
		} else {
			cc.depth = depthChange(cc.depth, npD.dChange, 5) //need negative. so 1 min and 5 max
		}
		ncD = *npD.branch
		//check if valid
		if cc.depth < ncD.dCheck {
			ncD = sDfail
		}
		v1 := "\"" + ncD.words + "\""
		fmt.Println(v1)
		//checks branches for possible events
		ncD, cc = dialogSusieEvents(ncD, cc)
	}
	cc.Character.Depth = cc.depth
	models.CharacterUpdate(cc.Character)
	return cc
}

func dialogSusieEvents(ncD cD, cc Convo) (cD, Convo) {
	switch ncD.branches[0] {
	case &pExit:
		cc.stilltalking = false
	case &pDefault:
		ncD = sH0
	case &pCabin:
		ncD = sH0
		if check.Eventcheck(7) == false {
			cblob := models.StoryblobGetByName(7)
			fmt.Println(cblob.Story)
			cblob.Shown = true
			models.StoryblobUpdate(cblob)
		} else {
			fmt.Println("\"Don't you remember? From camp go to the forest and find it from there.\"")
		}
		//implement map still
	}
	return ncD, cc
}
