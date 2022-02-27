package main

import (
	"fmt"
	"math/rand"
)

//tanks damage
const (
	c4 = iota
	bazukaa
	eyp
	mine
)

//tanks health
const (
	leopar = iota
	altay
	t70
)

func findTankStat(r int) (int, int, int) {
	switch r {
	case 0:
		return leopar, 700, 50
	case 1:
		return altay, 1000, 100
	case 2:
		return t70, 800, 70
	default:
		return 0, 0, 0
	}
}

func findBombTypeAndDamage(r int) (int, int) {
	switch r {
	case 0:
		return c4, 100
	case 1:
		return bazukaa, 200
	case 2:
		return eyp, 500
	case 3:
		return mine, 600
	default:
		return 0, 0
	}
}

type Tank struct {
	name                 string
	locationX, locationY int
	isAlive              bool
	amonation            int
	healt                int
	tip                  int
}

func newTank(n string) Tank {
	var tank Tank
	tank.name = n
	randomInt := rand.Intn(3)
	tank.tip, tank.healt, _ = findTankStat(randomInt)
	tank.isAlive = true
	tank.amonation = 10
	tank.move() //tank en başta bir cordinatta başlaması için ilk adımda move metodunu çağırdık.
	return tank
}

func (t *Tank) move() {
	rx, ry := rand.Intn(10), rand.Intn(10)
	t.locationX += rx
	t.locationY += ry

}

func (t *Tank) Fire() (int, int, int) {
	frx, fry := rand.Intn(100), rand.Intn(100)
	_, _, damage := findTankStat(t.tip)
	return frx, fry, damage
}

func (t *Tank) GetingDamage(d int) bool {
	if t.healt > 0 {
		t.healt -= d
	}
	if t.healt <= 0 {
		t.isAlive = false
		//TODO: game tabledan tankı kaldır.
		return true
	}
	return false
}

type Bomb struct {
	tip                  int
	locationX, locationY int
	isAlive              bool
}

func newBomb() Bomb {
	var bomb Bomb
	tp := rand.Intn(4)
	bomb.tip = tp
	bomb.locationX, bomb.locationY = rand.Intn(200), rand.Intn(200)
	bomb.isAlive = true
	return bomb
}

//TODO: feature , tankın menzili olsun menzilde ise ateş edebilsin.Sametin.
type gameTable struct {
	tanks      []*Tank
	bombs      []*Bomb
	maxX, maxY int
}

func newGameTable() gameTable {
	var gT gameTable
	gT.maxX = 200
	gT.maxY = 200
	return gT
}

func (g gameTable) Start(turn int) {

	for i := 0; i < turn; i++ {
		//TANKLAR FİRE EDİCEK
		//Tanklar ateş ettikten sonra ammonation azalacak 0 olursa ateş edemez
		for _, t := range g.tanks {
			x, y, d := t.Fire()
			for _, dt := range g.tanks {
				if dt.locationX == x && dt.locationY == y && dt.isAlive { //is alive tanks arrayinden kaldırılırsa oda kaldırılacaktır.
					if t.name != dt.name {
						z := dt.GetingDamage(d)
						if z {
							for i, v := range g.tanks {

								if v.name == dt.name {
									// TODO  akiften sonra devam
								}
							}
						}
						fmt.Println("Tank:", t.name, "attact", dt.name, "gettingDamage:", d)
						//TODO: FAŞO eğer isAlive false çekildiyse bu tankı gameTabledan çıkartıcaksın yani silinicek.
					}
				}
			}
		}
		// tanklar ateş etmeyi bitirdi.

		//Tanklar harekete geçecek.
		//TODO tankların üst üste gelmemeli.
		for _, t := range g.tanks {
			t.move()
			for _, b := range g.bombs {
				if t.locationX == b.locationX && t.locationY == b.locationY && b.isAlive {
					_, damage := findBombTypeAndDamage(b.tip)
					t.GetingDamage(damage)
					b.isAlive = false
					fmt.Println("bomba activated for ", t.name, "bomb is :", b.tip)
				}
			}
		}

	}
}

func main() {

	//TODO randomlar çalışmıyor , bakılacak.
	tank1 := newTank("samed")
	tank2 := newTank("faho")
	tank3 := newTank("akif")

	bomb1 := newBomb()
	bomb2 := newBomb()
	bomb3 := newBomb()

	gT := newGameTable()

	gT.tanks = append(gT.tanks, &tank1, &tank2, &tank3)
	gT.bombs = append(gT.bombs, &bomb1, &bomb2, &bomb3)

	// gT.Start(10)

	for _, t := range gT.bombs {
		fmt.Println(t)
	}
	for _, t := range gT.tanks {
		fmt.Println(t)
	}
}
