package conversation

import (
	"RPG-Core/inputs"
	"RPG-Core/models"
	"fmt"
)

//Mike, Josh, Susie and Veronica
//Need an array of responses

/* Branches=
1 - How are you?
2 - What's happening?

//qa = Hello, Player Question, Character Question, Player Answer, Character Answer, Character Tangent, Awkward Silence, Goodbye
//qa = H, PQ, CQ, PA, CA, CT, AS, GB
//depth 1-5 = Hostile, Dislike, Neutral, Like, Love

Can do a dialog tree of responses (i.e. branching tree with each stop having multiple options based on depth)
Or use depth at surface level conversation then launch Tangents? Easier in short term, less dynamic

need to add options to get items from Josh and Mike

*/

//Convo is base struct
type convo struct {
	Character    models.Character
	branch       int
	depth        int
	qa           string
	stilltalking bool
}

type cresp struct {
	ca      string //actual answer
	minD    int    //minimum Depth to get answer
	ropts   []string
	ctident string
	dD      int
	maxdD   int //max depth for changes
}

//only diff to cresp is branch-ctident. Could combine?
type presp struct {
	pq     string   //actual question
	ropts  []string //possible responses
	branch int
	dD     int //how much depth changes
	maxdD  int //min/max change
}

var (
	playerH = []string{"Hi. Fucker.", "Hello.", "Good to see you.", "Damn good to see you."}
	pq1     = presp{"How are you?", []string{"CA", "AS", "GB"}, 1, 0, 0}
	pq2     = presp{"What's happening?", []string{"CA", "AS", "GB"}, 2, 0, 0}
	pq3     = presp{"Can you help me?", []string{"CA"}, 3, 0, 0} //cheating
	pGB     = presp{"Bye", []string{"GB", "AS"}, 0, 0, 0}
	pQlist  = []presp{pq1, pq2, pq3, pGB} //list of options
	//need to start adding player answers
	playerA1 = []string{"I'm good", "Up and down", "Feeling shit", "Don't want to talk about it. Bye."}
	playerA2 = []string{"Nothing.", "A strange girl.", "Shadow monsters."}
	pA2      = []presp{
		{playerA2[0], []string{"CA21"}, 2, -1, 3}, //place depth here or on cresp? Cresp is probably better. Then remove it from presp
		{playerA2[1], []string{"CA22"}, 2, 1, 4},
		{playerA2[2], []string{"CA23"}, 2, 1, 5},
	}
	pA1 = []presp{
		{playerA1[0], []string{"CA11"}, 1, 0, 0},
		{playerA1[1], []string{"CA12"}, 1, 0, 0},
		{playerA1[2], []string{"CA13"}, 1, 0, 0},
		{playerA1[3], []string{"CA14"}, 1, 0, 0},
	}
	//these are overly specific now. Also need someone to provide general, game info
	playerA3 = []string{"I keep getting lost.", "I cannot find my flashlight. Do you have one?", "I need more space in my backpack. Cannot carry everything."}
	pA3      = []presp{
		{playerA3[0], []string{"CA31"}, 3, 0, 0}, //place depth here or on cresp? Cresp is probably better. Then remove it from presp
		{playerA3[1], []string{"CA32"}, 3, 0, 0},
		{playerA3[2], []string{"CA33"}, 3, 0, 0},
	}
)

var (
	//Mike, relaxed. One tangent a test about Josh. Not sure about the second tangent.
	mikeH  = []string{"Whoa, chill little man.", "Sup", "Good to see you too."}
	mrespH = []cresp{
		{mikeH[0], 0, []string{"PQ", "GB"}, "NA", -1, 1},
		{mikeH[1], 1, []string{"PQ", "GB"}, "NA", 0, 0},
		{mikeH[2], 2, []string{"PQ", "GB"}, "NA", 1, 4}}
	mresp1 = []cresp{{"I'm good. How are you?", 0, []string{"PA", "PQ", "GB"}, "NA", 0, 0},
		{"Honestly, not so good.", 4, []string{"CT"}, "CT1", 1, 5}} //bump this to five?
	mresp1_1 = []cresp{
		{"Good to hear", 0, []string{"PQ", "GB"}, "NA", 1, 3},
		{"Aren't we all little man?", 0, []string{"PQ", "GB"}, "NA", 1, 4},
		{"Sorry to hear. It'll turn around.", 0, []string{"PQ", "GB"}, "NA", 1, 4},
		{"It's your life. I'm here if you change your mind.", 0, []string{"GB"}, "NA", -1, 3},
	}
	mresp2 = []cresp{
		{"Just life and stuff.", 0, []string{"PQ", "GB"}, "NA", 0, 0},
		{"There are some odd things happening. What have you seen?", 2, []string{"PA", "PQ", "GB"}, "NA", 0, 0},
	}
	mresp2_2 = []cresp{ //cheating again, only one response regardless of depth
		{"Fair enough.", 0, []string{"PQ", "GB"}, "NA", 0, 0},
		{"Ha, I wouldn't call Susie strange.", 0, []string{"PQ", "GB"}, "NA", 0, 0},
		{"Whoa. Have you been binging on campfire ghost stories?", 0, []string{"PQ", "GB"}, "NA", 0, 0},
	}
	mresp3 = []cresp{
		//{"Dude. What's up?", 2, []string{"PA", "PQ", "GB"}, "NA", 1, 3}, need a low depth option here?
		{"For sure dude. What's up?", 0, []string{"PA", "PQ", "GB"}, "NA", 1, 3},
	}
	mresp3_3 = []cresp{
		{"It's a confusing place. I think Susie has a map. Last I saw she was getting ready for a big hike.", 0, []string{"PQ", "GB"}, "NA", 0, 0},
		{"It's dangerous at night without a light. Check with Josh. He's a nerd. Nerds like gear.", 0, []string{"PQ", "GB"}, "NA", 0, 0},
		{"I have a big backpack. But I cannot just give it away.", 0, []string{"PQ", "GB"}, "NA", 0, 0},
	}
	mikeGB = []string{"Lates"}
)

var (
	//need josh specific stuff
	joshH  = []string{"Hello."}
	jrespH = []cresp{
		{joshH[0], 0, []string{"PQ", "GB"}, "NA", 0, 0},
	}
	jresp3 = []cresp{
		{"Of course. Take it.", 0, []string{"PQ", "GB"}, "NA", 0, 0},
	}
)

//Converser is called by mainflow to start conversation
func Converser(cChar models.Character) {
	//defaults, starting point
	cc := convo{}
	cc.Character = cChar
	cc.stilltalking = true
	cc.depth = cc.Character.Depth
	cc.qa = "H" //starts with hello
	var cr1 cresp
	var ph1 presp
	//conversation loop starts
	for cc.stilltalking == true {
		switch cc.qa { //modify to show all options available
		case "H":
			d1 := inputs.StringarrayInput(playerH) // modify to add quotes?
			ph1 = presp{playerH[d1-1], []string{"H"}, d1 - 1, 0, 0}
			cc.depth = depthChange(cc.depth, ph1.dD, ph1.maxdD)
			cr1, cc = characterA(ph1, cc)
		case "PQ":
			options := []string{pQlist[0].pq, pQlist[1].pq, pQlist[2].pq, pQlist[3].pq} //make into a range later
			d1 := inputs.StringarrayInput(options)
			ph1 = pQlist[d1-1]
			cc.depth = depthChange(cc.depth, ph1.dD, ph1.maxdD)
			cc.qa = ph1.ropts[0]
			cr1, cc = characterA(ph1, cc)
		case "PA":
			if ph1.branch == 1 {
				d1 := inputs.StringarrayInput(playerA1)
				ph1 = pA1[d1-1]
			} else if ph1.branch == 2 {
				d1 := inputs.StringarrayInput(playerA2)
				ph1 = pA2[d1-1]
			} else if ph1.branch == 3 {
				d1 := inputs.StringarrayInput(playerA3)
				ph1 = pA3[d1-1]
			}
			cc.depth = depthChange(cc.depth, ph1.dD, ph1.maxdD)
			//need the second layer of character responses
			cc.qa = ph1.ropts[0]
			cr1, cc = characterA(ph1, cc)
		case "CT":
			cc = ctang(cc, cr1) // need to run these only once, but that will be part of character saves
			fallthrough         //goes to the default
		default: //including GB
			inputs.StringarrayInput([]string{pGB.pq})
			ph1 = pGB
			cr1, cc = characterA(ph1, cc)
			cc.stilltalking = false
		}
		fmt.Println(cr1.ca, cc.qa, cc.depth)
		cc.Character.Depth = cc.depth //sets character depth to conversation depth
		cc.Character = models.CharacterUpdate(cc.Character)
	}
}

func characterA(cpq presp, cc convo) (cresp, convo) {
	// set variables
	var curresp cresp
	var curcharM []cresp
	var hi []cresp
	var ca1 []cresp
	var ca2 []cresp
	var ca3 []cresp
	var ca1_2 []cresp
	var ca2_2 []cresp
	var ca3_3 []cresp
	var gb []string
	// set for each character
	switch cc.Character.Name {
	case "Mike":
		hi = mrespH
		ca1 = mresp1
		ca2 = mresp2
		ca3 = mresp3
		ca1_2 = mresp1_1
		ca2_2 = mresp2_2
		ca3_3 = mresp3_3
		gb = mikeGB
	case "Josh":
		hi = jrespH
		ca1 = mresp1 //placeholder
		ca2 = mresp2
		ca3 = jresp3
		ca1_2 = mresp1_1
		ca2_2 = mresp2_2
		ca3_3 = mresp3_3
		gb = mikeGB
	default:
		hi = []cresp{{"Oddly, nothing to say", 0, []string{"GB"}, "NA", 0, 0}}
		ca1 = hi
		ca2 = hi
		ca3 = hi
		ca1_2 = hi
		ca2_2 = hi
		ca3_3 = hi
		gb = []string{"Oddly bye"}
	}

	switch cc.qa {
	case "H":
		//need to make character response matrix
		crespM := [4]cresp{}
		curcharM = hi
		//should fill a matrix to 5, so for each depth there is an answer.
		//however for hello there are only 4 options
		for k := range crespM {
			for _, resp := range curcharM {
				if resp.minD <= k {
					crespM[k] = resp
				}
			}
		}
		curresp = crespM[cpq.branch]
		cc.qa = curresp.ropts[0]
		cc.depth = depthChange(cc.depth, curresp.dD, curresp.maxdD)
	case "CA":
		switch cpq.branch { //can remove this if combine characterA arrrays?
		case 1:
			crespM := [5]cresp{}
			curcharM = ca1
			for k := range crespM {
				for _, resp := range curcharM {
					if resp.minD-1 <= k {
						crespM[k] = resp
					}
				}
			}
			curresp = crespM[cc.depth-1]
			cc.qa = curresp.ropts[0]
			cc.depth = depthChange(cc.depth, curresp.dD, curresp.maxdD)
		case 2:
			crespM := [5]cresp{}
			curcharM = ca2
			for k := range crespM {
				for _, resp := range curcharM {
					if resp.minD-1 <= k {
						crespM[k] = resp
					}
				}
			}
			curresp = crespM[cc.depth-1]
			cc.qa = curresp.ropts[0]
			cc.depth = depthChange(cc.depth, curresp.dD, curresp.maxdD)
		case 3:
			crespM := [5]cresp{}
			curcharM = ca3
			for k := range crespM {
				for _, resp := range curcharM {
					if resp.minD-1 <= k {
						crespM[k] = resp
					}
				}
			}
			curresp = crespM[cc.depth-1]
			cc.qa = curresp.ropts[0]
			cc.depth = depthChange(cc.depth, curresp.dD, curresp.maxdD)
		default:
			curresp = cresp{"Oddly, nothing to say", 0, []string{"GB"}, "NA", 0, 0}
			cc.qa = "GB"
		}
	case "CA11": //cheat for second layer of conversation
		curresp = ca1_2[0]
		cc.qa = curresp.ropts[0]
		cc.depth = depthChange(cc.depth, curresp.dD, curresp.maxdD)
	case "CA12": //cheat for second layer of conversation
		curresp = ca1_2[1]
		cc.qa = curresp.ropts[0]
		cc.depth = depthChange(cc.depth, curresp.dD, curresp.maxdD)
	case "CA13": //cheat for second layer of conversation
		curresp = ca1_2[2]
		cc.qa = curresp.ropts[0]
		cc.depth = depthChange(cc.depth, curresp.dD, curresp.maxdD)
	case "CA14": //cheat for second layer of conversation
		curresp = ca1_2[3]
		cc.qa = curresp.ropts[0]
		cc.depth = depthChange(cc.depth, curresp.dD, curresp.maxdD)
	case "CA21": //cheat for second layer of conversation
		curresp = ca2_2[0]
		cc.qa = curresp.ropts[0]
		cc.depth = depthChange(cc.depth, curresp.dD, curresp.maxdD)
	case "CA22":
		curresp = ca2_2[1]
		cc.qa = curresp.ropts[0]
		cc.depth = depthChange(cc.depth, curresp.dD, curresp.maxdD)
	case "CA23":
		curresp = ca2_2[2]
		cc.qa = curresp.ropts[0]
		cc.depth = depthChange(cc.depth, curresp.dD, curresp.maxdD)
	case "CA31": //cheat for second layer of conversation
		curresp = ca3_3[0]
		cc.qa = curresp.ropts[0]
		cc.depth = depthChange(cc.depth, curresp.dD, curresp.maxdD)
	case "CA32":
		curresp = ca3_3[1]
		cc.qa = curresp.ropts[0]
		cc.depth = depthChange(cc.depth, curresp.dD, curresp.maxdD)
	case "CA33":
		curresp = ca3_3[2]
		cc.qa = curresp.ropts[0]
		cc.depth = depthChange(cc.depth, curresp.dD, curresp.maxdD)
	case "GB":
		curresp = cresp{gb[0], 0, []string{"GB"}, "NA", 0, 0}
		cc.stilltalking = false
	}
	curresp.ca = "\"" + curresp.ca + "\""
	return curresp, cc
}

func depthChange(ct int, dt int, max int) int {
	var nt int
	if dt == 0 {
		nt = ct
	} else {
		switch {
		case dt < 0:
			if ct < max {
				nt = ct
			} else {
				nt = ct + dt
				if nt < max {
					nt = max
				}
			}
		case dt > 0:
			if ct > max {
				nt = ct
			} else {
				nt = ct + dt
				if nt > max {
					nt = max
				}
			}
		}
	}
	return nt
}

type dialog struct {
	words   string
	branch1 *dialog
	branch2 *dialog
	branch3 *dialog
	//height  int
}

//testing tree
var (
	pTest1 = dialog{words: "That's surprising. What's up?", branch1: &mR1}
	mR1    = dialog{words: "Well, you know Josh?", branch1: &pR1, branch2: &pL1, branch3: &ptreeGB}
	pR1    = dialog{words: "No", branch1: &mR2}
	pL1    = dialog{words: "Yes, I met Josh", branch1: &mL2}
	mR2    = dialog{words: "Oh. He's this real awkward guy. Total nerd.", branch1: &pR3, branch2: &pR4, branch3: &ptreeGB}
	mL2    = dialog{words: "What did you think? He's a total nerd right?", branch1: &pL3, branch2: &pR4, branch3: &ptreeGB}
	pR3    = dialog{words: "Right. And what did he do to you?", branch1: &mR3}
	pL3    = dialog{words: "He seemed a bit quiet. But I thought he was ok.", branch1: &mL4}
	pR4    = dialog{words: "I hate nerds.", branch1: &mR5}
	mR3    = dialog{words: "What do you mean? He didn't do anything. He's just a nerd." +
		"\nAnd an asshole. And I want you to help me beat the shit out of him.", branch1: &pR5, branch2: &pL5, branch3: &ptreeGB}
	mL4 = dialog{words: "What, did you really talk to him? He's a total nerd." +
		"\nAnd an asshole. And I want you to help me beat the shit out of him.", branch1: &pR5, branch2: &pL5, branch3: &ptreeGB}
	mR5 = dialog{words: "Wow. Well he's a total nerd." +
		"\nAnd I want you to help me beat the shit out of him.", branch1: &pR5, branch2: &pL5, branch3: &ptreeGB}
	pR5 = dialog{words: "Of course. Let's do this.", branch1: &mR6}
	pL5 = dialog{words: "No man. I'm not doing that.", branch1: &mL6}
	mR6 = dialog{words: "Damn dude. I'm really disappointed. Seriously." +
		"\nI thought you were okay. It was like a test and you failed." +
		"\nAnd, just to clarify, it's not cool to beat other people up.", branch1: &ptreeGB1}
	mL6      = dialog{words: "Really? Why? Are you a pussy?", branch1: &pR6, branch2: &pL6}
	pR6      = dialog{words: "No I'm not a pussy. Nevermind, I'm in. Let's go punch a nerd.", branch1: &mR6}
	pL6      = dialog{words: "What? No. I just don't beat people up. What's your problem?", branch1: &mL7}
	mL7      = dialog{words: "Haha. I am a messed up dude. But you are solid. It was like a test litle man. And you passed.", branch1: &ptreeGB2}
	ptreeGB  = dialog{words: "Got to go. Later.", branch1: &mtreeGB}
	ptreeGB1 = dialog{words: "Sorry to disappoint. At least I don't set up little tests for people. Later douchebag.", branch1: &mtreeGB}
	ptreeGB2 = dialog{words: "You're solid too. But don't test me again like that.", branch1: &mtreeGB}
	mtreeGB  = dialog{words: "Lates."}
)

func ctang(cc convo, cr1 cresp) convo {
	tancont := true
	pTan := pTest1
	cTan := mR1 //could use depth check
	//pw := "\"" + pTan.words + "\""
	//fmt.Println(pw)
	for tancont == true {
		cw := "\"" + cTan.words + "\""
		fmt.Println(cw)
		options := []string{cTan.branch1.words} //causes a crash
		if cTan.branch2 != nil {
			options = append(options, cTan.branch2.words)
		}
		if cTan.branch3 != nil {
			options = append(options, cTan.branch3.words)
		}
		d1 := inputs.StringarrayInput(options)
		if d1 == 1 {
			pTan = *cTan.branch1
		} else if d1 == 2 {
			pTan = *cTan.branch2
		} else if d1 == 3 {
			pTan = *cTan.branch3
		} else {
			fmt.Println("It's over. Oddly.")
			tancont = false
		}
		//pw = "\"" + pTan.words + "\""
		//fmt.Println(pw)
		cTan = *pTan.branch1
		if cTan == mtreeGB {
			tancont = false
			if pTan == ptreeGB1 {
				cc.depth = 1 //need to make different outcomes for different characters. For now, cheating
			} else if pTan == ptreeGB2 {
				cc.depth = 5
				if cc.Character.Name == "Mike" {
					ni := models.ItemGetByName("large pack")
					fmt.Println("\"Dude, take my", ni.ShortName, "\"")
					ni.Loc = 19
					models.ItemUpdate(ni)
					oi := models.ItemGetByName("small pack")
					oi.Loc = 21
					models.ItemUpdate(oi)
				}
			}
			cw := "\"" + cTan.words + "\""
			fmt.Println(cw)
		}
	}

	cc.qa = "GB"
	return cc
}

/* old Tangent attempt

//trying to get a tangent working
mikeCT1GB = []string{"No worries. Catch up later"}
mikeCT1   = []string{
  "It's kind of tough to talk about. Sure you want to hear?",
  "Well, this is my last year here. Have to spend next summer working.",
  "The whole summer. Only way to get to college.",
  "Thanks. I feel better."}
pCT1     = []string{"Of course dude.", "Actually not really.", "About you being a pussy?", "Hold that thought, I got to go."}
pCT2     = []string{"What is up with that?"}
pCT3     = []string{"Sorry to hear"}
mikeCT1a = []string{"Haha. Exactly. Can you handle it?"}
pCT1a    = []string{"Yes"}
pCT      = [][]string{pCT1, pCT2, pCT3}

/* barely works
cT := mikeCT1
pT := pCT
pTGB := []string{pGB.pq}
cGB := mikeCT1GB
tancont := true
i := 0
for tancont == true {
  fmt.Println(cT[i])
  d1 := inputs.StringarrayInput(pT[i]) //make one that takes in player responces
  switch d1 {
  case 1:
    //nothing, only works now because of only one choice
  case 2:
    fmt.Println(cGB[0])
    tancont = false
  case 3:
    fmt.Println(mikeCT1a[0])
    pCT1a = append(pCT1a, pTGB[0])
    d2 := inputs.StringarrayInput(pCT1a)
    if d2 == 2 {
      fmt.Println(cGB[0])
      tancont = false
    }
  case 4:
    fmt.Println(cGB[0])
    tancont = false
  }
  i++
}
*/
