package conversation

import (
	"RPG-Core/check"
	"RPG-Core/inputs"
	"RPG-Core/models"
	"fmt"
)

/*
Notes on conversations for future ones:

1. Should codify the numbering of responses - can make character response match player - like pcv420 then vct420
2. Could make a script for this. All i enter is dialog and branch numbers, it does the restrict
3. Should make two separate Dialog structs - one for Player one for Character. Without depth, the character only ever has one response

arg, it'd be good to find a way to make loops, for when the player has multiple questions
*/

type pD struct {
	words   string
	branch  *cD
	dChange int
}

type cD struct {
	words    string
	branches []*pD
	dCheck   int
}

//intro
var (
	jihstart = cD{"Odd Silence", []*pD{&pih0, &pih1}, 0}

	pih  = []string{"What's up nerd?", "Hi Josh"}
	pih0 = pD{words: pih[0], branch: &jih0, dChange: -1}
	pih1 = pD{words: pih[1], branch: &jih1, dChange: 0}

	jih  = []string{"I know you don't have any friends, doesn't mean you have to bug me.", "Hi loner. What do you want?"}
	jih0 = cD{jih[0], []*pD{&piq0, &piq1}, 0}
	jih1 = cD{jih[1], []*pD{&piq11, &piq12}, 0}

	piq  = []string{"Sorry dude.", "You're right. I'll leave."}
	piq0 = pD{piq[0], &jiq0, 1}
	piq1 = pD{piq[1], &jiq1, 0}

	jiq  = []string{"Yeh. See you around.", "Good to hear. Have fun."}
	jiq0 = cD{jiq[0], []*pD{&pExit}, 0}
	jiq1 = cD{jiq[1], []*pD{&pExit}, 0}

	piq10 = []string{"Can I borrow your flashlight?", "How are you doing?"}
	piq11 = pD{piq10[0], &jiq11, 0}
	piq12 = pD{piq10[1], &jiq12, 1}

	jiq10 = []string{"Of course. Oddly.", "I'm good. You know, nature and sun."}
	jiq11 = cD{jiq10[0], []*pD{&pExit}, 4} //not possible to get to 4 in intro
	jiq12 = cD{jiq10[1], []*pD{&piq20, &piq21}, 2}

	piqs20 = []string{"I love it.", "Better than hiding in your basement."}
	piq20  = pD{piqs20[0], &jiq20, 1}
	piq21  = pD{piqs20[1], &jiq21, -1}

	jiqs20 = []string{"Right. Well I got go put my stuff in the bear bag. See you later.", "Right. Go away, I got stuff to do."}
	jiq20  = cD{jiqs20[0], []*pD{&pBye1, &pBye2, &pBye3}, 2}
	jiq21  = cD{jiqs20[1], []*pD{&pBye1, &pBye2, &pBye3}, 0}

	//Bye
	pByes = []string{"See you.", "Lates.", "I'll leave you alone."}
	pBye1 = pD{pByes[0], &jBye1, 0}
	pBye2 = pD{pByes[1], &jBye2, 0}
	pBye3 = pD{pByes[2], &jBye3, -1}

	jByes = []string{"Maybe.", "Lame. Did you steal that from Mike?", "My prayers are answered."}
	jBye1 = cD{jByes[0], []*pD{&pExit}, 0}
	jBye2 = cD{jByes[1], []*pD{&pExit}, 0}
	jBye3 = cD{jByes[2], []*pD{&pExit}, 0}

	//Triggers
	pExit    = pD{} //exit
	piqTest  = pD{}
	pDefault = pD{}
	//pBye     = pD{}

	//Depth fail
	jDfail = cD{"No. I don't really know you and I got stuff to do. Just leave me alone.", []*pD{&pExit}, 0}

	//need Normal/2nd conversation
	jstart = cD{"Odd Silence", []*pD{&pH0, &pH1}, 0}
	pHs    = []string{"Hello again.", "Hello again. Nerd."}
	pH0    = pD{pHs[0], &jH0, 0}
	pH1    = pD{pHs[1], &jH1, -1}
	jHs    = []string{"Hi.", "Audible sigh."}
	jH0    = cD{jHs[0], []*pD{&pjQ1, &pjQ2, &pjQ3, &pjQ4}, 0}
	jH1    = cD{jHs[1], []*pD{&pjQ1, &pjQ2, &pjQ3, &pjQ4}, 0}

	pjQs = []string{"Enjoying camp?", "What's your deal?", "Can you help me?", "Nevermind."}
	pjQ1 = pD{pjQs[0], &jA1, 0}
	pjQ2 = pD{pjQs[1], &jA2, -1}
	pjQ3 = pD{pjQs[2], &jA3, 0}
	pjQ4 = pD{pjQs[3], &jA4, 0}

	jAs = []string{"Yep. It's great.", "That's direct. What is your deal?", "Sure. Pussy.", "Whatever."}
	jA1 = cD{jAs[0], []*pD{&pjQ10, &pjQ11, &pjQ12}, 0}
	jA2 = cD{jAs[1], []*pD{&pjQ20, &pjQ21, &pjQ22}, 0}
	jA3 = cD{jAs[2], []*pD{&pjQ30, &pjQ31, &pjQ32, &pjQ17}, 0}
	jA4 = cD{jAs[3], []*pD{&pExit}, 0}

	//It's great
	pjQ1s = []string{"Really?", "I concur.", "I think it's awful."}
	pjQ10 = pD{pjQ1s[0], &jA10, 0}
	pjQ11 = pD{pjQ1s[1], &jA11, 0}
	pjQ12 = pD{pjQ1s[2], &jA12, 1}

	jA10s = []string{"What? I mean it's pretty good. Nature and stuff.", "Good. Nice conversation.", "Really? Why?"}
	jA10  = cD{jA10s[0], []*pD{&pjQ15, &pjQ16}, 0}
	jA11  = cD{jA10s[1], []*pD{&pDefault}, 0}
	jA12  = cD{jA10s[2], []*pD{&pjQ18, &pjQ19}, 0}

	pjQ15s = []string{"What do you like about it specifically?", "Good ol nature."}
	pjQ15  = pD{pjQ15s[0], &jA15, 1}
	pjQ16  = pD{pjQ15s[1], &jA11, 0}

	jA15s = []string{"Ok you got me. I'm not used to this stuff at all. I'm not a big outdoors guy, like Mike."}
	jA15  = cD{jA15s[0], []*pD{&pjQ110, &pjQ111}, 4}

	pjQ110s = []string{"Sorry man. Didn't mean to interrogate you. What do you like to do?", "Hahaha. I knew it! Nerd!"}
	pjQ110  = pD{pjQ110s[0], &jA110, 0}
	pjQ111  = pD{pjQ110s[1], &jA111, -2}

	jA110s = []string{"Lots of stuff. You know? I keep busy. Watching cartoons, playing games with my friends. Lots of games actually.",
		"Fuck off dude."}
	jA110 = cD{jA110s[0], []*pD{&pjQ112, &pjQ113}, 4}
	jA111 = cD{jA110s[1], []*pD{&pBye1, &pBye2, &pBye3}, 0}

	pjQ112s = []string{"Sounds fun. So you'd rather be there, with your friends, right now?", "Hahaha. I knew it! Nerd!"}
	pjQ112  = pD{pjQ112s[0], &jA112, 0}
	pjQ113  = pD{pjQ112s[1], &jA111, -2}

	jA112s = []string{"Honestly, no. I mean, that stuff is great. But I also like it here. Hiking, among all the trees, meeting new kids." +
		"\nI just feel out of place here. A loner."}
	jA112 = cD{jA112s[0], []*pD{&pjQ114, &pjQ115}, 4}

	pjQ114s = []string{"No worries. I know the feeling. My best friend isn't here. I feel like I don't know anyone.",
		"Well that's because you are one. A loner."}
	pjQ114 = pD{pjQ114s[0], &jA114, 2}
	pjQ115 = pD{pjQ114s[1], &jA115, -1}

	jA114s = []string{"Right! I was wandering where that guy was. That sucks. You having fun without him?", "You asshole."}
	jA114  = cD{jA114s[0], []*pD{&pjQ116, &pjQ117, &pjQ118}, 4}
	jA115  = cD{jA114s[1], []*pD{&pDefault}, 0}

	pjQ116s = []string{"Yeh! It's great here.", "I'm doing okay.", "I hate this place."}
	pjQ116  = pD{pjQ116s[0], &jA116, 1}
	pjQ117  = pD{pjQ116s[1], &jA117, 1}
	pjQ118  = pD{pjQ116s[2], &jA118, 2}

	jA116s = []string{"Ha. That's enthusiasm is very impressive or very fake.", "Cool.",
		"Bummer, well I hope it's going to get better. Let me know if I can help."}
	jA116 = cD{jA116s[0], []*pD{&pDefault}, 0}
	jA117 = cD{jA116s[1], []*pD{&pDefault}, 0}
	jA118 = cD{jA116s[2], []*pD{&pDefault}, 3}

	//I think it's awful
	pjQ18s = []string{"I'd rather be with my friends.", "Too much sun.", "There are monsters in the forest."}
	pjQ18  = pD{pjQ18s[0], &jA18, 2}
	pjQ19  = pD{pjQ18s[1], &jA19, 0}
	pjQ17  = pD{pjQ18s[2], &jA17, 1}

	jA18s = []string{"Me too. What do you like to do?", "What are you a goth?", "Ha. You are a weird kid."}
	jA18  = cD{jA18s[0], []*pD{&pjQ120, &pjQ121}, 3}
	jA19  = cD{jA18s[1], []*pD{&pjQ122, &pjQ123}, 0}
	jA17  = cD{jA18s[2], []*pD{&pDefault}, 0}

	pjQ120s = []string{"Just hang out.", "Typically my best friend, he'd be here. Camping. But he couldn't join on this trip."}
	pjQ120  = pD{pjQ120s[0], &jA120, -1}
	pjQ121  = pD{pjQ120s[1], &jA114, 1}
	jA120s  = []string{"Very descriptive."}
	jA120   = cD{jA120s[0], []*pD{&pDefault}, 0}

	pjQ122s = []string{"Yes. And a vampire. So watch out.", "Labels hurt dude."}
	pjQ122  = pD{pjQ122s[0], &jA122, 1}
	pjQ123  = pD{pjQ122s[1], &jA123, 1}
	jA122s  = []string{"No, you better watch out. I'm cooking tonight and I'll add a shit-ton of garlic powder to everything.",
		"Truth."}
	jA122 = cD{jA122s[0], []*pD{&pjQ124, &pjQ125}, 2}
	jA123 = cD{jA122s[1], []*pD{&pDefault}, 0}

	pjQ124s = []string{"Ha, you're funny.", "Lame joke."}
	pjQ124  = pD{pjQ124s[0], &jA124, 1}
	pjQ125  = pD{pjQ124s[1], &jA125, -2}
	jA124s  = []string{"And you're funny looking.", "Right. Can you be somewhere else?"}
	jA124   = cD{jA124s[0], []*pD{&pDefault}, 0}
	jA125   = cD{jA124s[1], []*pD{&pDefault}, 0}

	//That's direct
	pjQ20s = []string{"Sorry. I'm not great at talking to people.", "You're moody and rude.", "I just need to borrow your flashlight."}
	pjQ20  = pD{pjQ20s[0], &jA20, 0}
	pjQ21  = pD{pjQ20s[1], &jA21, -2}
	pjQ22  = pD{pjQ20s[2], &jA22, 0}
	jA20s  = []string{"I'll give you an A+ for self pity. And a bit of advice, for free: try to get to know them. Don't just confront people out of nowhere.",
		"Who are you kid? Leave me alone.", "Really? You can have it."}
	jA20 = cD{jA20s[0], []*pD{&pDefault}, 0}
	jA21 = cD{jA20s[1], []*pD{&pDefault}, 0}
	jA22 = cD{jA20s[2], []*pD{&pFlash}, 4}

	//Flashlight event, needs to be added to event check
	pFlash = pD{}
	jFlash = cD{"Here you go.", []*pD{&pjQF1, &pjQF2}, 0}
	pjQFs  = []string{"Thanks. I'll bring it back.", "Thanks. Nerd."}
	pjQF1  = pD{pjQFs[0], &jAF1, 0}
	pjQF2  = pD{pjQFs[1], &jAF2, -1}
	jAF1   = cD{"Keep it. I have a back up.", []*pD{&pBye1, &pBye2, &pBye3}, 0}
	jAF2   = cD{"Classic loner diss.", []*pD{&pBye1, &pBye2, &pBye3}, 0}

	//Help me
	pjQ30s = []string{"I lost my flashlight. Can I borrow one from you?", "I don't know where anything is.", "I cannot carry all my cool stuff."}
	pjQ30  = pD{pjQ30s[0], &jA22, 0}
	pjQ31  = pD{pjQ30s[1], &jA30, 0}
	pjQ32  = pD{pjQ30s[2], &jA31, 0}

	jA30s = []string{"It's a big place. I'd advise finding a map.", "That sounds just awful."}
	jA30  = cD{jA30s[0], []*pD{&pjQ33, &pjQ35}, 3}
	jA31  = cD{jA30s[1], []*pD{&pjQ33, &pjQ35}, 3}

	pjQ33s = []string{"Right. And where would I find that?", "It's the worst.", "Thanks for nothing. You nerd-lord of nothing giving town."}
	pjQ33  = pD{pjQ33s[0], &jA33, 0}
	pjQ34  = pD{pjQ33s[1], &jA34, 0}
	pjQ35  = pD{pjQ33s[2], &jA35, -1}

	jA33s = []string{"You need a map to the map? I think Susie took one. She left on a hike. Wanted to see the view.",
		"Right. Well I think Mike has a larger backpack.", "Great. I banish you. Go away."}
	jA33 = cD{jA33s[0], []*pD{&pjQ36, &pjQ35}, 3}
	jA34 = cD{jA33s[1], []*pD{&pjQ36, &pjQ35}, 3}
	jA35 = cD{jA33s[2], []*pD{&pBye1, &pBye2, &pBye3}, 0}

	pjQ36s = []string{"Thanks."}
	pjQ36  = pD{pjQ36s[0], &jA36, 0}

	jA36 = cD{"Glad to help.", []*pD{&pDefault}, 3}
)

//ConverserJ is for conversations with Josh
func ConverserJ(cc Convo) Convo {
	jm := check.Eventcheck(4)
	var options []string
	var npD pD
	var ncD cD
	//no function exists yet to get the player's name!
	//hellos first
	if jm == false { //check if they met
		ncD = jihstart
		jb := models.StoryblobGetByName(4)
		fmt.Println(jb.Story)
		jb.Shown = true
		models.StoryblobUpdate(jb)
	} else {
		ncD = jstart
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
			ncD = jDfail
		}
		v1 := "\"" + ncD.words + "\""
		fmt.Println(v1, cc.depth)
		//checks branches for possible events
		ncD, cc = dialogJoshEvents(ncD, cc)
	}
	cc.Character.Depth = cc.depth
	models.CharacterUpdate(cc.Character)
	return cc
}

func choicemakerV2(tang cD) []string {
	var options []string
	n := len(tang.branches)
	for i := 0; i < n; i++ {
		br1 := *tang.branches[i]
		options = append(options, br1.words)
	}
	return options
}

func dialogJoshEvents(ncD cD, cc Convo) (cD, Convo) {
	var v1 string
	switch ncD.branches[0] {
	case &pExit:
		cc.stilltalking = false
	case &pDefault:
		ncD = jH0
	case &pFlash:
		fi := models.ItemGetByName("flashlight")
		if fi.Loc == 20 {
			fmt.Println("\"Wait a minute. You have a flashlight already. Don't get greedy.\"")
			ncD = jH0
		} else {
			fi.Loc = 20
			models.ItemUpdate(fi)
			ncD = jFlash
			v1 = "\"" + ncD.words + "\""
			fmt.Println(v1)
		}
	}
	return ncD, cc
}
