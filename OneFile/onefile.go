package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//Dx returns a value between 1 and x - used for simulation of dice.
func Dx(x int) int {
	d := rand.Intn(x)
	d++
	return (d)
}

//Xdx simulates d amount of x-sided dice
func Xdx(d int, x int) int {
	totalscore := 0
	count := 0
	for count < d {
		dx := Dx(x)
		totalscore += dx
		count++
	}
	return (totalscore)
}

//DropLowest takes 4 random numbers between 1&6, removes the lowest and adds the rest.
func DropLowest() int {
	rolls := []int{Dx(6), Dx(6), Dx(6), Dx(6)}
	sort.Ints(rolls)
	droplowest := rolls[1] + rolls[2] + rolls[3]
	return droplowest
}

//rolls is a slice filled with 4*1d6. sort.Ints(rolls) orders them from lowest to highest. then you simply add the higher values.
//This isn't flexible at all (like Xdx), but it works.

var strn, dext, cons, intl, wisd, char int

var strbn, dexbn, conbn, intbn, wisbn, chabn int

var hp, hpmax, ac int

func setLowStat() {
	strn = Xdx(3, 6)
	dext = Xdx(3, 6)
	cons = Xdx(3, 6)
	intl = Xdx(3, 6)
	wisd = Xdx(3, 6)
	char = Xdx(3, 6)
}

func setHighStat() {
	strn = DropLowest()
	dext = DropLowest()
	cons = DropLowest()
	intl = DropLowest()
	wisd = DropLowest()
	char = DropLowest()
}

func setStatbonus() {
	strbn = (strn / 2) - 5
	dexbn = (dext / 2) - 5
	conbn = (cons / 2) - 5
	intbn = (intl / 2) - 5
	wisbn = (wisd / 2) - 5
	chabn = (char / 2) - 5
}

func setSubstats() {
	hpmax = Dx(6) + conbn
	hp = hpmax
	ac = 10 + dexbn

}

func printSheet() {

	fmt.Println()
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("     Strength =", strn, "(", strbn, ")")
	fmt.Println("    Dexterity =", dext, "(", dexbn, ")")
	fmt.Println(" Constitution =", cons, "(", conbn, ")")
	fmt.Println(" Intelligence =", intl, "(", intbn, ")")
	fmt.Println("       Wisdom =", wisd, "(", wisbn, ")")
	fmt.Println("     Charisma =", char, "(", chabn, ")")
	fmt.Println()
	fmt.Println(" HP=(", hp, "/", hpmax, ")")
	fmt.Println(" AC =", ac)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	//Uses fmt.Scanln as a way to select high or low.
	hilo := 0
	fmt.Println("Press 1 for normal stats (3d6) or 2 for higher stats (4d6 drop lowest)")
	fmt.Scanln(&hilo)

	if hilo == 2 {
		fmt.Println("Higher stats selected.")
		fmt.Println("Rolling dice...")
		setHighStat()
	} else {
		fmt.Println("Normal stats selected.")
		fmt.Println("Rolling dice...")
		setLowStat()
	}

	setStatbonus()
	setSubstats()

	printSheet()

	//this makes the program wait for input before closing
	fmt.Scanln(&hilo)
}
