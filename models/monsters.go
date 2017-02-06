package models

//Monster is basic Monster model
type Monster struct {
	FullName  string
	ShortName string
	Health    int
	Loc       int
	Details   string
}

var monstermap = map[string]Monster{}

//Monsterset sets initial Monsters
func Monsterset() {
	md1 := "A small critter scurries around the forest. It has two little arms and two little legs." +
		"\nIt has too many little teeth in a perpetual smile. Everything is murky shadow except for that white smile."
	minion := Monster{"Smiler", "smiler", 25, 8, md1}
	MonsterUpdate(minion)
}

//MonsterGetByName grabs current Monster by short name
func MonsterGetByName(c string) Monster {
	cm := mmap()
	i := cm[c]
	return i
}

//MonsterGetByLoc grabs Monster by location
func MonsterGetByLoc(l int) ([]Monster, int) {
	cm := mmap()
	cslice := []Monster{}
	i := 0
	for _, v := range cm {
		if v.Loc == l {
			cslice = append(cslice, v)
			i++
		}
	}
	return cslice, i
}

func mmap() map[string]Monster {
	return monstermap
}

//MonsterUpdate allows updates to Monsters
func MonsterUpdate(cc Monster) Monster {
	cm := mmap()
	cm[cc.ShortName] = cc
	return cc
}
